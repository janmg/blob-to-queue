package input

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"janmg.com/blob-to-queue/common"
	"janmg.com/blob-to-queue/format"
)

func Blobworker(queue chan format.Flatevent) {
	config := common.ConfigHandler()
	location := "https://" + config.Accountname + "." + config.Cloud
	fmt.Println(location)

	// TODO: NSGFlowLogs grow in predictable directories and files, only need to keep a pointer to the last processed directory.
	// TODO: Would need to way to manually change that pointer, incase old files need to be reprocessed
	// TODO: The worker would need a timer to check for new files or grown files?

	// keep a registry, filelist and a worklist
	// The registry is the list of files already processed, the filelist is a list of currently seen files, the worklist contains the new files or the files that grew
	var registry map[string]int64 = make(map[string]int64)
	var filelist map[string]int64 = make(map[string]int64)
	var worklist map[string]int64 = make(map[string]int64)

	interval := time.NewTimer(60 * time.Second)
	//interval := time.NewTicker(60 * time.Second)
	defer interval.Stop()

	go func() {
		<-interval.C
		fmt.Println("Interval timer fired")
		// 1. Lists all the files in the remote storage account that match the path prefix
		filelist = listFiles(config.Accountname, config.Accountkey, location)
		// 2. Filters on path_filters to only include files that match the directory and file glob (**/*.json)
		worklist = registry - filelist
		// 3. Compare the list of files to the the registry with the new filelist and put the delta in a worklist

		// 4. Process the worklist and put all events in the logstash queue.

		// 5. Save the registry with files and sizes to a file

		// 6. if there is time left, sleep to complete the interval. If processing takes more than an inteval, save the registry and continue.
		// ... try to sync the timer to when the files are actually written to the storage account and wait an additional 5 seconds before reading.
		// ... did storage accounts implement some time of difference tracking journal?
		// 7. If stop signal comes, finish the current file, save the registry and quit

	}()

	//print(config)
	// List the blobs in the container
	fmt.Println("Listing the blobs in the container:")
	fmt.Println(registry)

	/*
		       ChatGPT, suggests to drop events when the queue is full. Obviously would never do that and better to signal a monitor channel to signal to pause reading more data and check regularly if the queue is blocked or not.

			   select {
			   case queue <- event: // Non-blocking send
			   default:
			       fmt.Println("Warning: Dropped event due to full queue")
			   }
	*/

	// list all the nsg's
	// resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/
	// loop through the dates, skip the older ones and process only from the data in the registry
	/*
		resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/OCTOBER-NSG/y=2023/m=10/d=31/h=13/m=00/

		for each nsg
		resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/

	*/

	// filelist = Hash.new
	// worklist = Hash.new
	// @last = start = Time.now.to_i
	// filelist.clear
	// filelist = list_blobs(false)
	// registry.store(name, { :offset => off, :length => file[:length] })

	// save_registry()
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

func fullRead(queue chan format.Flatevent, name string) {
	config := common.ConfigHandler()
	cred, err := azblob.NewSharedKeyCredential(config.Accountname, config.Accountkey)
	common.Error(err)
	location := "https://" + config.Accountname + "." + config.Cloud
	client, err := azblob.NewClientWithSharedKeyCredential(location, cred, nil)
	common.Error(err)

	ctx := context.Background()
	/*
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
	*/

	get, err := client.DownloadStream(ctx, config.ContainerName, name, nil)
	common.Error(err)

	downloadedData := bytes.Buffer{}
	retryReader := get.NewRetryReader(ctx, &azblob.RetryReaderOptions{})
	_, err = downloadedData.ReadFrom(retryReader)
	common.Error(err)
	//fmt.Println(downloadedData.String())

	err = retryReader.Close()
	common.Error(err)

	// parse the json into a flatevent struct and push it into the queue
	nsgflowlog(queue, downloadedData.Bytes(), name)
	// vnetflowlog(queue, downloadedData.Bytes(), name)
}

func partialRead(queue chan format.Flatevent, name string) {
	config := common.ConfigHandler()
	cred, err := azblob.NewSharedKeyCredential(config.Accountname, config.Accountkey)
	common.Error(err)

	location := "https://" + config.Accountname + "." + config.Cloud
	client, err := azblob.NewClientWithSharedKeyCredential(location, cred, nil)
	common.Error(err)

	ctx := context.Background()
	/*
		httpRange := azblob.HTTPRange{
			Offset: 0,
			Count:  12,
		}
		get, err := client.DownloadStream(ctx, config.ContainerName, *blob.Name, &azblob.DownloadStreamOptions{Range: httpRange})
	*/
	get, err := client.DownloadStream(ctx, config.ContainerName, name, nil)
	common.Error(err)

	downloadedData := bytes.Buffer{}
	retryReader := get.NewRetryReader(ctx, &azblob.RetryReaderOptions{})
	_, err = downloadedData.ReadFrom(retryReader)
	common.Error(err)
	//fmt.Println(downloadedData.String())

	err = retryReader.Close()
	common.Error(err)

	// TODO: How to process lines?
	// parse the json into a flatevent struct and push it into the queue
	nsgflowlog(queue, downloadedData.Bytes(), name)
}
