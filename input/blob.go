package input

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"janmg.com/blob-to-queue/common"
	"janmg.com/blob-to-queue/format"
)

type Stamp struct {
	Year   int `json:"year"`
	Month  int `json:"month"`
	Day    int `json:"day"`
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

type FileMetadata struct {
	ContentLength int64     `json:"content_length"`
	ContentRead   int64     `json:"content_read"`
	LastModified  time.Time `json:"last_modified"`
}

func Blobworker(queue chan format.Flatevent) {
	config := common.ConfigHandler()

	// The registry is the list of files already processed, the filelist is a list of currently seen files, than read the diffence and process the files
	var registry map[string]FileMetadata = make(map[string]FileMetadata)

	// NSGFlowLogs grow in predictable directories and files, only need to keep a timestamp pointer to the last processed directory, other files may have arbitrary names and the registry will track which files are new and which ones have grown.
	// https://pkg.go.dev/time
	var last Stamp
	if config.Resumepolicy == "timestamp" {
		if config.Startpolicy == "start_over" {
			fmt.Println("Starting over, clearing the timestamp")
			last = Stamp{0, 0, 0, 0, 0}
		} else {
			var err error
			last, err = readTimestamp(config.Timestamp)
			if err != nil {
				last = Stamp{0, 0, 0, 0, 0}
				fmt.Println("Can't read timestamp for start_fresh, will start_over instead")
			}
		}
		fmt.Printf("Resuming from timestamp: %v\n", last)
	} else if config.Resumepolicy == "registry" {
		// Read the registry if it exists
		if config.Startpolicy == "start_over" {
			fmt.Println("Starting over, clearing the registry")
			os.Remove(config.Registry) // Delete the registry file
			registry = make(map[string]FileMetadata)
		} else {
			var err error
			fmt.Println("Resuming from the registry")
			registry, err = loadRegistry("registry.json")
			if err != nil {
				registry = make(map[string]FileMetadata)
				fmt.Println("Can't read registry for start_fresh, will start_over instead")
			}
		}
		fmt.Println("Resuming from registry")
	}

	// Initial Sync
	doLoop(config, queue, registry, last)

	// Interval is a timestamp thing, use the last processed timestamp to process the last few files
	interval := time.NewTicker(time.Duration(config.Interval) * time.Second)
	defer interval.Stop()

	for range interval.C {
		doLoop(config, queue, registry, last)
	}
}

func doLoop(config common.Config, queue chan format.Flatevent, registry map[string]FileMetadata, last Stamp) {
	// Should use a queue to signal that it's time to stop the ingress
	for i := 0; len(queue) > config.Qwatermark; i++ {
		fmt.Printf("Hit watermark, queue is at %d pause for 10 seconds", len(queue))
		time.Sleep(10 * time.Second)
		if i%5 == 0 {
			fmt.Println("1 Minute and still above the queue watermark, maybe the output is not processing?")
		}
	}

	location := "https://" + config.Accountname + "." + config.Cloud
	// fmt.Println(location)

	// list all the nsg's
	// resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/
	// loop through the dates, skip the older ones and process only from the data in the registry
	//	resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/OCTOBER-NSG/y=2023/m=10/d=31/h=13/m=00/
	//	for each nsg
	//	resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/

	// 1. Lists all the files in the remote storage account that match the path prefix
	var filelist map[string]FileMetadata = make(map[string]FileMetadata)
	filelist = listFiles(config.Resumepolicy, config.Accountname, config.Accountkey, location, last)
	//fmt.Printf("Received filelist with %d entries\n", len(filelist))
	var fullfiles = 0
	var partialfiles = 0

	// 2. Filters on path_filters to only include files that match the directory and file glob (**/*.json)
	// TODO: filter on the timestamp so that only fresh files get processed

	// 3. Compare the list of files to the the registry with the new filelist

	// 4. Process the worklist and put all events in the logstash queue.

	// Read based on modified flags

	fmt.Println("Listing the blobs in the container:")
	for name, metadata := range filelist {
		if oldMeta, exists := registry[name]; exists && oldMeta.ContentLength != metadata.ContentLength {
			// File exists in registry and size changed - PARTIAL READ
			fmt.Printf("%s grew by %d bytes\n", name, metadata.ContentLength-oldMeta.ContentLength)
			read(queue, name, oldMeta.ContentLength, metadata.ContentLength)
			partialfiles++
		} else if !exists {
			// File doesn't exist in registry - NEW FILE
			fmt.Printf("%s is new and has %d bytes\n", name, metadata.ContentLength)
			read(queue, name, 0, metadata.ContentLength)
			fullfiles++
		}
		// If exists and size unchanged, skip processing
	}
	//convert to log item
	fmt.Printf("Found %d new files and %d updated files\n", fullfiles, partialfiles)
	// TODO: This doesn't work well? ... need to consider timestamps, fullfiles is 0 and updated files are 7 ... ???

	// 5. Save the registry with files and sizes to a file
	if config.Resumepolicy == "timestamp" {
		writeTimestamp("timestamp", time.Now())
	}
	if config.Resumepolicy == "registry" {
		saveRegistry(config.Registry, filelist)
	}

	// 6. if there is time left, sleep to complete the interval. If processing takes more than an inteval, save the registry and continue.
	// ... try to sync the timer to when the files are actually written to the storage account and wait an additional 5 seconds before reading.
	// ... did storage accounts implement some time of difference tracking journal?
	// 7. If stop signal comes, finish the current file, save the registry and quit
}

func listFiles(resumepolicy string, account string, key string, location string, last Stamp) map[string]FileMetadata {
	config := common.ConfigHandler()
	var filelist map[string]FileMetadata = make(map[string]FileMetadata)

	cred, err := azblob.NewSharedKeyCredential(config.Accountname, config.Accountkey)
	common.Error(err)
	client, err := azblob.NewClientWithSharedKeyCredential(location, cred, nil)
	common.Error(err)

	pager := client.NewListBlobsFlatPager(config.ContainerName, &azblob.ListBlobsFlatOptions{
		Include: azblob.ListBlobsInclude{Snapshots: false, Versions: true},
		// include={snapshots,metadata,uncommittedblobs,copy,deleted,tags,versions,deletedwithversions,immutabilitypolicy,legalhold,permissions}
		// showonly={deleted,files,directories}
		// prefix
		// NextMarker?
	})

	for pager.More() {
		resp, err := pager.NextPage(context.Background())
		common.Error(err)

		fmt.Printf("Processing page with %d blobs\n", len(resp.Segment.BlobItems))
		for _, blob := range resp.Segment.BlobItems {
			// generated.BlobProperties {ETag: *"0x8DE2CCA4087FF97", LastModified: *time.Time(2025-11-26T09:00:31Z){wall: 0, ext: 63899744431, loc: *(*time.Location)(0xc00005af50)}, ACL: *string nil, AccessTier: *"Hot", AccessTierChangeTime: *time.Time nil, AccessTierInferred: *true, ArchiveStatus: *github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/generated.ArchiveStatus nil, BlobSequenceNumber: *int64 nil, BlobType: *"BlockBlob", CacheControl: *"", ContentDisposition: *"", ContentEncoding: *"", ContentLanguage: *"", ContentLength: *108487, ContentMD5: []uint8 len: 0, cap: 0, nil, ContentType: *"application/octet-stream", CopyCompletionTime: *time.Time nil, CopyID: *string nil, CopyProgress: *string nil, CopySource: *string nil, CopyStatus: *github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/generated.CopyStatusType nil, CopyStatusDescription: *string nil, CreationTime: *time.Time(2025-11-26T08:09:31Z){wall: 0, ext: 63899741371, loc: *(*time.Location)(0xc00005aee0)}, CustomerProvidedKeySHA256: *string nil, DeletedTime: *time.Time nil, DestinationSnapshot: *string nil, EncryptionScope: *string nil, ExpiresOn: *time.Time nil, Group: *string nil, ImmutabilityPolicyExpiresOn: *time.Time nil, ImmutabilityPolicyMode: *github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/generated.ImmutabilityPolicyMode nil, IncrementalCopy: *bool nil, IsSealed: *bool nil, LastAccessedOn: *time.Time nil, LeaseDuration: *github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/generated.LeaseDurationType nil, LeaseState: *"available", LeaseStatus: *"unlocked", LegalHold: *bool nil, Owner: *string nil, Permissions: *string nil, RehydratePriority: *github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/generated.RehydratePriority nil, RemainingRetentionDays: *int32 nil, ResourceType: *string nil, ServerEncrypted: *true, TagCount: *int32 nil}
			// TODO: this filter could keep track of the last file being read, but what about partial reads.
			// Needs tracking of which files were read, for flowlogs should use the date/time in the directory structure, only need to remember last processed file

			// Filter on BlobType BlockBlob
			fmt.Println(*blob.Name)
			if config.Resumepolicy == "timestamp" {
				if blob.Properties.LastModified.After(time.Date(last.Year, time.Month(last.Month), last.Day, last.Hour, last.Minute, 0, 0, time.UTC)) {
					filelist[*blob.Name] = FileMetadata{
						ContentLength: *blob.Properties.ContentLength,
						ContentRead:   0,
						LastModified:  *blob.Properties.LastModified,
					}
					fmt.Printf("Added to filelist: %s (size: %d)\n", *blob.Name, *blob.Properties.ContentLength)
				} else {
					fmt.Printf("Skipped (older than lastread): %s\n", *blob.Name)
				}
			} else {
				// TODO: compare to registry
				filelist[*blob.Name] = FileMetadata{
					ContentLength: *blob.Properties.ContentLength,
					ContentRead:   0,
					LastModified:  *blob.Properties.LastModified,
				}
				// fmt.Printf("Added to filelist: %s (size: %d)\n", *blob.Name, *blob.Properties.ContentLength)
			}
		}
	}
	return filelist
}

// Read the files with the httpRange
func read(queue chan format.Flatevent, name string, oldSize int64, size int64) {

	config := common.ConfigHandler()
	cred, err := azblob.NewSharedKeyCredential(config.Accountname, config.Accountkey)
	common.Error(err)

	location := "https://" + config.Accountname + "." + config.Cloud
	client, err := azblob.NewClientWithSharedKeyCredential(location, cred, nil)
	common.Error(err)

	ctx := context.Background()

	var get azblob.DownloadStreamResponse
	var err2 error

	if oldSize > 0 {
		// Partial read - read only the new data
		httpRange := azblob.HTTPRange{
			Offset: oldSize,
			Count:  size - oldSize,
		}

		dso := &azblob.DownloadStreamOptions{
			Range: httpRange,
			AccessConditions: &blob.AccessConditions{
				ModifiedAccessConditions: &blob.ModifiedAccessConditions{
					IfModifiedSince: modtime(),
				},
			},
		}
		get, err2 = client.DownloadStream(ctx, config.ContainerName, name, dso)
	} else {
		// Full read - read entire file
		get, err2 = client.DownloadStream(ctx, config.ContainerName, name, nil)
	}
	common.Error(err2)

	downloadedData := bytes.Buffer{}
	retryReader := get.NewRetryReader(ctx, &azblob.RetryReaderOptions{})
	_, err = downloadedData.ReadFrom(retryReader)
	common.Error(err)
	//fmt.Println(downloadedData.String())

	err = retryReader.Close()
	common.Error(err)

	// TODO: should make a distinction between log formats, grok lines, json lines, json structures
	// TODO: but can only do that at the output package ... should tell the queue if the content is flatevent, line or raw
	// for flowlogs, parse the json into a flatevent struct and push it into the queue
	// nsgflowlog(queue, downloadedData.Bytes(), name)
	// vnetflowlog(queue, downloadedData.Bytes(), name)
	// maybe also flag what source this comes from?

	// parse the json into a flatevent struct and push it into the queue
	// Create a signal channel for flow control
	signal := make(chan bool, 1)
	go func() {
		for {
			if len(queue) < cap(queue)/2 {
				select {
				case signal <- true:
				default:
				}
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	if config.Type == "nsgflowlog" {
		nsgflowlog(queue, signal, downloadedData.Bytes(), name)
	}
	if config.Type == "vnetflowlog" {
		vnetflowlog(queue, signal, downloadedData.Bytes(), name)
	}
}

func loadRegistry(path string) (map[string]FileMetadata, error) {
	file, err := os.Open(path)
	if err != nil {
		return make(map[string]FileMetadata), err
	}
	defer file.Close()

	var registry map[string]FileMetadata
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&registry)
	if err != nil {
		return make(map[string]FileMetadata), err
	}
	return registry, nil
}

// TODO: why the -24? ... is this to process one day?
func modtime() *time.Time {
	t := time.Now().Add(-24 * time.Hour)
	return &t
}

func saveRegistry(path string, filelist map[string]FileMetadata) {
	// resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/OCTOBER-NSG/y=2023/m=10/d=31/h=17/m=00/macAddress=002248A31CA3/PT1H.json
	// resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/OCTOBER-NSG/y=2023/m=10/d=31/h=18/m=00/macAddress=002248A31CA3/PT1H.json
	// y=2023/m=10/d=31/h=18/m=00/macAddress=002248A31CA3/PT1H.json
	file, err := os.Create(path)
	common.Error(err)
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(filelist)
	common.Error(err)
}

// TODO Why not save this straight as a string timestamp instead of a json!?
func readTimestamp(path string) (Stamp, error) {
	var ts Stamp

	file, err := os.Open(path)
	common.Error(err)
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&ts)
	return ts, err
}

func writeTimestamp(path string, t time.Time) error {
	ts := Stamp{
		Year:   t.Year(),
		Month:  int(t.Month()),
		Day:    t.Day(),
		Hour:   t.Hour(),
		Minute: t.Minute(),
	}

	file, err := os.Create(path)
	common.Error(err)
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(ts)
}
