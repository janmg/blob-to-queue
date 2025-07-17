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

func Blobworker(queue chan format.Flatevent) {
	config := common.ConfigHandler()

	// The registry is the list of files already processed, the filelist is a list of currently seen files, than read the diffence and process the files
	var registry map[string]int64 = make(map[string]int64)

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
		// TODO: startpolicy, if it is start_over, then clear the registry and start from scratch
		if config.Startpolicy == "start_over" {
			fmt.Println("Starting over, clearing the registry")
			registry = make(map[string]int64)
		} else {
			var err error
			fmt.Println("Resuming from the registry")
			registry, err = loadRegistry("registry.json")
			if err != nil {
				registry = make(map[string]int64)
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

func doLoop(config common.Config, queue chan format.Flatevent, registry map[string]int64, last Stamp) {
	for i := 0; len(queue) > config.Qwatermark; i++ {
		fmt.Printf("Hit watermark, queue is at %d pause for 10 seconds", len(queue))
		time.Sleep(10 * time.Second)
		if i > 5 {
			fmt.Println("Giving up waiting for queue")
			return
		}
	}

	location := "https://" + config.Accountname + "." + config.Cloud
	fmt.Println(location)

	// list all the nsg's
	// resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/
	// loop through the dates, skip the older ones and process only from the data in the registry
	//	resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/OCTOBER-NSG/y=2023/m=10/d=31/h=13/m=00/
	//	for each nsg
	//	resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/

	// 1. Lists all the files in the remote storage account that match the path prefix
	var filelist map[string]int64 = make(map[string]int64)
	filelist = listFiles(config.Accountname, config.Accountkey, location)
	var fullfiles = 0
	var partialfiles = 0

	// TODO: Filter based on timestamp

	// 2. Filters on path_filters to only include files that match the directory and file glob (**/*.json)
	// TODO: filter on the timestamp so that only fresh files get processed

	// 3. Compare the list of files to the the registry with the new filelist and read the
	// 4. Process the worklist and put all events in the logstash queue.

	// Read based on modified flags

	// TODO: This lists the files based on the registry!?
	fmt.Println("Listing the blobs in the container:")
	for name, size := range filelist {
		if oldSize, exists := registry[name]; !exists || oldSize != size {
			//convert to log item
			fmt.Printf("%s grew by %d bytes\n", name, size-oldSize)
			read(queue, name, oldSize, size)
			partialfiles++
		} else {
			//convert to log item
			fmt.Printf("%s is new and has %d bytes\n", name, size)
			read(queue, name, 0, size)
			fullfiles++
		}
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

func listFiles(account string, key string, location string) map[string]int64 {
	config := common.ConfigHandler()
	var filelist map[string]int64 = make(map[string]int64)

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
		resp, err := pager.NextPage(context.TODO())
		common.Error(err)

		for _, blob := range resp.Segment.BlobItems {

			// TODO: this filter could keep track of the last file being read, but what about partial reads.
			// Needs tracking of which files were read, for flowlogs should use the date/time in the directory structure, only need to remember last processed file
			/*
				if *blob.Name > "resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/OCTOBER-NSG/y=2023/m=10/d=31/h=18/m=00" {
					fmt.Println(*blob.Name)

				}
			*/
			//registry
			//fullRead(queue, *blob.Name)
			filelist[*blob.Name] = *blob.Properties.ContentLength
		}
	}
	return filelist
}

/*
// not needed because partial read can read from 0 also??
func fullRead(queue chan format.Flatevent, name string) {
	config := common.ConfigHandler()
	cred, err := azblob.NewSharedKeyCredential(config.Accountname, config.Accountkey)
	common.Error(err)
	location := "https://" + config.Accountname + "." + config.Cloud
	client, err := azblob.NewClientWithSharedKeyCredential(location, cred, nil)
	common.Error(err)

	ctx := context.Background()
		// ListBlockBlob
		blobURL := fmt.Sprintf("https://%s.%s/%s/%s", config.Accountname, config.Cloud, config.ContainerName, name)
		blockBlobClient, err := blockblob.NewClientWithSharedKeyCredential(blobURL, cred, nil)
		common.Error(err)

		blockList, err := blockBlobClient.GetBlockList(context.Background(), blockblob.BlockListTypeAll, nil)
		// BlockListTypeCommitted
		common.Error(err)

			for _, blocks := range blockList.BlockList.CommittedBlocks {
				fmt.Println(*blocks.Name, *blocks.Size)
				// QTAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAw
				// WjAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAw
			}

	get, err := client.DownloadStream(ctx, config.ContainerName, name, nil)
	common.Error(err)

	downloadedData := bytes.Buffer{}
	retryReader := get.NewRetryReader(ctx, &azblob.RetryReaderOptions{})
	_, err = downloadedData.ReadFrom(retryReader)
	common.Error(err)
	//fmt.Println(downloadedData.String())

	err = retryReader.Close()
	common.Error(err)

}
*/

// Read the files with the httpRange
func read(queue chan format.Flatevent, name string, oldSize int64, size int64) {

	config := common.ConfigHandler()
	cred, err := azblob.NewSharedKeyCredential(config.Accountname, config.Accountkey)
	common.Error(err)

	location := "https://" + config.Accountname + "." + config.Cloud
	client, err := azblob.NewClientWithSharedKeyCredential(location, cred, nil)
	common.Error(err)

	ctx := context.Background()
	httpRange := azblob.HTTPRange{
		Offset: oldSize + 1,
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

	get, err := client.DownloadStream(ctx, config.ContainerName, name, dso)
	//get, err := client.DownloadStream(ctx, config.ContainerName, name, nil)
	common.Error(err)

	downloadedData := bytes.Buffer{}
	// TODO: account for the missing header in case of partial reads
	if oldSize > 0 {
		prefix := []byte("{\"message\": {")
		downloadedData.Write(prefix)
	}

	retryReader := get.NewRetryReader(ctx, &azblob.RetryReaderOptions{})
	_, err = downloadedData.ReadFrom(retryReader)
	common.Error(err)
	//fmt.Println(downloadedData.String())

	err = retryReader.Close()
	common.Error(err)

	// TODO: should make a distinction between log types, grok lines, json lines, json structures
	// parse the json into a flatevent struct and push it into the queue
	// nsgflowlog(queue, downloadedData.Bytes(), name)
	// vnetflowlog(queue, downloadedData.Bytes(), name)

	// parse the json into a flatevent struct and push it into the queue
	nsgflowlog(queue, downloadedData.Bytes(), name)
}

func modtime() *time.Time {
	t := time.Now().Add(-24 * time.Hour)
	return &t
}

func loadRegistry(path string) (map[string]int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return make(map[string]int64), err
	}
	defer file.Close()

	var registry map[string]int64
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&registry)
	if err != nil {
		return make(map[string]int64), err
	}
	return registry, nil
}

func saveRegistry(path string, filelist map[string]int64) {
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
