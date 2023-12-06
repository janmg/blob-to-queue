package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob"
)

func worker(queue chan flatevent) {
	config := configHandler()

	cred, err := azblob.NewSharedKeyCredential(config.Accountname, config.Accountkey)
	handleError(err)
	client, err := azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.%s/", config.Accountname, config.Cloud), cred, nil)
	handleError(err)

	// List the blobs in the container
	fmt.Println("Listing the blobs in the container:")

	pager := client.NewListBlobsFlatPager(config.ContainerName, &azblob.ListBlobsFlatOptions{
		Include: azblob.ListBlobsInclude{Snapshots: true, Versions: true},
	})

	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		handleError(err)

		for _, blob := range resp.Segment.BlobItems {
			fmt.Println(*blob.Name)
		}
	}

	// ListBlockBlob
	blobName := "resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/OCTOBER-NSG/y=2023/m=10/d=31/h=13/m=00/macAddress=002248A31CA3/PT1H.json"
	blobURL := fmt.Sprintf("https://%s.%s/%s/%s", config.Accountname, config.Cloud, config.ContainerName, blobName)

	blockBlobClient, err := blockblob.NewClientWithSharedKeyCredential(blobURL, cred, nil)
	handleError(err)

	blockList, err := blockBlobClient.GetBlockList(context.Background(), blockblob.BlockListTypeAll, nil)
	handleError(err)
	fmt.Println(blockList.BlockList.CommittedBlocks)

	ctx := context.Background()
	// Needs tracking of which files were read, for flowlogs should use the date/time in the directory structure, only need to remember last processed file
	// Needs implementation of partial reads, incase files grow
	// Download the blob
	get, err := client.DownloadStream(ctx, config.ContainerName, blobName, nil)
	handleError(err)

	downloadedData := bytes.Buffer{}
	retryReader := get.NewRetryReader(ctx, &azblob.RetryReaderOptions{})
	_, err = downloadedData.ReadFrom(retryReader)
	handleError(err)

	err = retryReader.Close()
	handleError(err)

	// 2023/12/02 11:19:02 connection string does not contain an EntityPath. eventHub cannot be an empty string

	// parse the json into a flatevent struct and push it into the queue
	nsgflowlog(queue, downloadedData.Bytes(), blobName)
}
