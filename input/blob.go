package input

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"janmg.com/blob-to-queue/common"
	"janmg.com/blob-to-queue/format"
)

// func listFiles(resumepolicy string, account string, key string, location string, last Stamp) map[string]RegistryData {
func listFiles(resumepolicy string, connection string, last time.Time) map[string]RegistryData {
	config := common.ConfigHandler() // TODO: why load config, if we already have resumepolicy and connection as parameter?
	var filelist map[string]RegistryData = make(map[string]RegistryData)
	client := getClient(config.Connection)
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
				// TODO: must use updated timestamp
				//datestamp := time.Date(last.Year, time.Month(last.Month), last.Day, last.Hour, last.Minute, 0, 0, time.UTC)
				if blob.Properties.LastModified.After(last) {
					filelist[*blob.Name] = RegistryData{
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
				filelist[*blob.Name] = RegistryData{
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
	client := getClient(config.Connection)
	// TODO: replace with getServiceClientSharedKey
	//cred, err := azblob.NewSharedKeyCredential(config.Accountname, config.Accountkey)
	//common.Error(err)
	// location := "https://" + config.Accountname + "." + config.Cloud
	//client, err := azblob.NewClientWithSharedKeyCredential(location, cred, nil)
	//common.Error(err)

	ctx := context.Background()

	var get azblob.DownloadStreamResponse
	var err error

	if oldSize > 0 {
		// Partial read - read only the new data
		httpRange := azblob.HTTPRange{
			Offset: oldSize,
			Count:  size - oldSize,
		}

		// TODO: fix IfModifiedSince: &last ... this should be registry based
		dso := &azblob.DownloadStreamOptions{
			Range: httpRange,
			AccessConditions: &blob.AccessConditions{
				ModifiedAccessConditions: &blob.ModifiedAccessConditions{
					IfModifiedSince: &last,
				},
			},
		}
		get, err = client.DownloadStream(ctx, config.ContainerName, name, dso)
	} else {
		// Full read - read entire file
		get, err = client.DownloadStream(ctx, config.ContainerName, name, nil)
	}
	common.Error(err)

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
