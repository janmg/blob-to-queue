package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func blobworker(queue chan Flatevent) {
	config := configHandler()
	//print(config)
	cred, err := azblob.NewSharedKeyCredential(config.Accountname, config.Accountkey)
	Error(err)
	location := "https://" + config.Accountname + "." + config.Cloud
	fmt.Println(location)
	client, err := azblob.NewClientWithSharedKeyCredential(location, cred, nil)
	Error(err)

	// List the blobs in the container
	fmt.Println("Listing the blobs in the container:")

	pager := client.NewListBlobsFlatPager(config.ContainerName, &azblob.ListBlobsFlatOptions{
		Include: azblob.ListBlobsInclude{Snapshots: false, Versions: true},
		// include={snapshots,metadata,uncommittedblobs,copy,deleted,tags,versions,deletedwithversions,immutabilitypolicy,legalhold,permissions}
		// showonly={deleted,files,directories}
		// prefix
		// NextMarker?
	})

	// list all the nsg's
	// resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/
	// loop through the dates, skip the older ones and process only from the data in the registry
	/*
	   	resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/OCTOBER-NSG/y=2023/m=10/d=31/h=13/m=00/

	   	for each nsg
	   	resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/

	       y=2023
	   	m=10
	   	d=31
	   	h=13 from 0 to 23
	   	m=00
	*/
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
			fullRead(queue, *blob.Name)
		}
	}
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

	// TODO:
	//    THIS IS WHERE BLOBS DONT GET CHOPPED IN LINES
	// TODO: How to process lines?
	// parse the json into a flatevent struct and push it into the queue
	nsgflowlog(queue, downloadedData.Bytes(), name)
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
