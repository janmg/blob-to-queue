package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob"
)

func blobworker(queue chan flatevent) {
	config := configHandler()

	cred, err := azblob.NewSharedKeyCredential(config.Accountname, config.Accountkey)
	handleError(err)
	client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.%s/", config.Accountname, config.Cloud), cred, nil)
	handleError(err)

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
		handleError(err)

		for _, blob := range resp.Segment.BlobItems {
			if *blob.Name > "resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/OCTOBER-NSG/y=2023/m=10/d=31/h=18/m=00" {
				fmt.Println(*blob.Name)
				blobURL := fmt.Sprintf("https://%s.%s/%s/%s", config.Accountname, config.Cloud, config.ContainerName, *blob.Name)
				blockBlobClient, err := blockblob.NewClientWithSharedKeyCredential(blobURL, cred, nil)
				handleError(err)

				// ListBlockBlob
				blockList, err := blockBlobClient.GetBlockList(context.Background(), blockblob.BlockListTypeAll, nil)
				// BlockListTypeCommitted
				handleError(err)
				ctx := context.Background()
				for _, blocks := range blockList.BlockList.CommittedBlocks {
					fmt.Println(*blocks.Name, *blocks.Size)
					// QTAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAw
					// WjAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAw
				}
				// Needs tracking of which files were read, for flowlogs should use the date/time in the directory structure, only need to remember last processed file
				// Needs implementation of partial reads, incase files grow
				// Download the blob
				/*
					httpRange := azblob.HTTPRange{
						Offset: 0,
						Count:  12,
					}
					get, err := client.DownloadStream(ctx, config.ContainerName, *blob.Name, &azblob.DownloadStreamOptions{Range: httpRange})
				*/
				get, err := client.DownloadStream(ctx, config.ContainerName, *blob.Name, nil)
				handleError(err)

				downloadedData := bytes.Buffer{}
				retryReader := get.NewRetryReader(ctx, &azblob.RetryReaderOptions{})
				_, err = downloadedData.ReadFrom(retryReader)
				handleError(err)
				//fmt.Println(downloadedData.String())

				err = retryReader.Close()
				handleError(err)

				// TODO: How to process lines?
				// parse the json into a flatevent struct and push it into the queue
				nsgflowlog(queue, downloadedData.Bytes(), *blob.Name)
			}

		}
	}
}
