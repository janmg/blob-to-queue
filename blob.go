package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func blobworker(queue chan Flatevent) {
	config := configHandler()
	location := "https://" + config.Accountname + "." + config.Cloud
	fmt.Println(location)

	// keep a registry
	var registry map[string]int64
	registry = make(map[string]int64)

	//print(config)
	// List the blobs in the container
	fmt.Println("Listing the blobs in the container:")
	registry = listFiles(config.Accountname, config.Accountkey, location)
	fmt.Println(registry)

	// list all the nsg's
	// resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/
	// loop through the dates, skip the older ones and process only from the data in the registry
	/*
		resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/OCTOBER-NSG/y=2023/m=10/d=31/h=13/m=00/

		for each nsg
		resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/

	*/

	// 1. Lists all the files in the remote storage account that match the path prefix
	// 2. Filters on path_filters to only include files that match the directory and file glob (**/*.json)
	// 3. Save the listed files in a registry of known files and filesizes.
	// 4. List all the files again and compare the registry with the new filelist and put the delta in a worklist
	// 5. Process the worklist and put all events in the logstash queue.
	// 6. if there is time left, sleep to complete the interval. If processing takes more than an inteval, save the registry and continue.
	// 7. If stop signal comes, finish the current file, save the registry and quit

	// filelist = Hash.new
	// worklist = Hash.new
	// @last = start = Time.now.to_i
	// filelist.clear
	// filelist = list_blobs(false)
	// registry.store(name, { :offset => off, :length => file[:length] })

	// save_registry()
}

func listFiles(account string, key string, location string) map[string]int64 {
	var filelist map[string]int64
	filelist = make(map[string]int64)

	cred, err := azblob.NewSharedKeyCredential(config.Accountname, config.Accountkey)
	Error(err)
	client, err := azblob.NewClientWithSharedKeyCredential(location, cred, nil)
	Error(err)

	pager := client.NewListBlobsFlatPager(config.ContainerName, &azblob.ListBlobsFlatOptions{
		Include: azblob.ListBlobsInclude{Snapshots: false, Versions: true},
		// include={snapshots,metadata,uncommittedblobs,copy,deleted,tags,versions,deletedwithversions,immutabilitypolicy,legalhold,permissions}
		// showonly={deleted,files,directories}
		// prefix
		// NextMarker?
	})

	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		Error(err)

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

func fullRead(queue chan Flatevent, name string) {
	config := configHandler()
	cred, err := azblob.NewSharedKeyCredential(config.Accountname, config.Accountkey)
	Error(err)
	location := "https://" + config.Accountname + "." + config.Cloud
	client, err := azblob.NewClientWithSharedKeyCredential(location, cred, nil)
	Error(err)

	ctx := context.Background()
	/*
		// ListBlockBlob
		blobURL := fmt.Sprintf("https://%s.%s/%s/%s", config.Accountname, config.Cloud, config.ContainerName, name)
		blockBlobClient, err := blockblob.NewClientWithSharedKeyCredential(blobURL, cred, nil)
		Error(err)

		blockList, err := blockBlobClient.GetBlockList(context.Background(), blockblob.BlockListTypeAll, nil)
		// BlockListTypeCommitted
		Error(err)

			for _, blocks := range blockList.BlockList.CommittedBlocks {
				fmt.Println(*blocks.Name, *blocks.Size)
				// QTAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAw
				// WjAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAw
			}
	*/

	get, err := client.DownloadStream(ctx, config.ContainerName, name, nil)
	Error(err)

	downloadedData := bytes.Buffer{}
	retryReader := get.NewRetryReader(ctx, &azblob.RetryReaderOptions{})
	_, err = downloadedData.ReadFrom(retryReader)
	Error(err)
	//fmt.Println(downloadedData.String())

	err = retryReader.Close()
	Error(err)

	// parse the json into a flatevent struct and push it into the queue
	nsgflowlog(queue, downloadedData.Bytes(), name)
	// vnetflowlog(queue, downloadedData.Bytes(), name)
}

func partialRead(queue chan Flatevent, name string) {
	config := configHandler()
	cred, err := azblob.NewSharedKeyCredential(config.Accountname, config.Accountkey)
	Error(err)

	location := "https://" + config.Accountname + "." + config.Cloud
	client, err := azblob.NewClientWithSharedKeyCredential(location, cred, nil)
	Error(err)

	ctx := context.Background()
	/*
		httpRange := azblob.HTTPRange{
			Offset: 0,
			Count:  12,
		}
		get, err := client.DownloadStream(ctx, config.ContainerName, *blob.Name, &azblob.DownloadStreamOptions{Range: httpRange})
	*/
	get, err := client.DownloadStream(ctx, config.ContainerName, name, nil)
	Error(err)

	downloadedData := bytes.Buffer{}
	retryReader := get.NewRetryReader(ctx, &azblob.RetryReaderOptions{})
	_, err = downloadedData.ReadFrom(retryReader)
	Error(err)
	//fmt.Println(downloadedData.String())

	err = retryReader.Close()
	Error(err)

	// TODO: How to process lines?
	// parse the json into a flatevent struct and push it into the queue
	nsgflowlog(queue, downloadedData.Bytes(), name)
}
