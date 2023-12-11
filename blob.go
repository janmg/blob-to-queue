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
	})

	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		handleError(err)

		for _, blob := range resp.Segment.BlobItems {
			fmt.Println(*blob.Name)
			blobURL := fmt.Sprintf("https://%s.%s/%s/%s", config.Accountname, config.Cloud, config.ContainerName, *blob.Name)

			blockBlobClient, err := blockblob.NewClientWithSharedKeyCredential(blobURL, cred, nil)
			handleError(err)

			// ListBlockBlob
			blockList, err := blockBlobClient.GetBlockList(context.Background(), blockblob.BlockListTypeAll, nil)
			handleError(err)
			ctx := context.Background()
			for _, blocks := range blockList.BlockList.CommittedBlocks {
				fmt.Println(*blocks.Name, *blocks.Size)
				// Needs tracking of which files were read, for flowlogs should use the date/time in the directory structure, only need to remember last processed file
				// Needs implementation of partial reads, incase files grow
				// Download the blob
				get, err := client.DownloadStream(ctx, config.ContainerName, *blob.Name, nil)
				handleError(err)

				downloadedData := bytes.Buffer{}
				retryReader := get.NewRetryReader(ctx, &azblob.RetryReaderOptions{})
				_, err = downloadedData.ReadFrom(retryReader)
				handleError(err)
				fmt.Println(string(downloadedData.Bytes()))

				err = retryReader.Close()
				handleError(err)
				// 2023/12/02 11:19:02 connection string does not contain an EntityPath. eventHub cannot be an empty string

				// parse the json into a flatevent struct and push it into the queue
				//nsgflowlog(queue, downloadedData.Bytes(), blobName)
			}
		}
	}
}
