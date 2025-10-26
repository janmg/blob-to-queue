# blob-to-queue
Golang application for reading from an azure blob storage, listing the files that match a path filter and pushing the result to an event queue like kafka stream. Intended to offer an alternative to my own logstash module logstash-input-azure_blob_storage and move it's logic to a standalone program.

# configure
The configurationfile is a YAML file, which is reloadable on save, because its using spf13/viper. The configuration format allows for multiple output streams.
configure multiple outputs

blob-to-queue.yaml

output: ["stdout","file"]

# configure input
accountName: "blobstoragename"
accountKey: "AMWsmPcgy/1234567890123445abcdefghijkl/1234567890123445abcdefghijklABCDEFGHI+ASt3SvXjw=="
containerName: "insights-logs-networksecuritygroupflowevent"

# Microsoft blobstorage
Logfiles written to blobstorage in the json format have a header in the first block and a footer in the last block
Block 0000: QTAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAw {message: [
Block 0001: Bobaba                                          {"timestamp": 98765432, "mac":"00:01:02:AA:AB:AC"},
Block 0002: Bobae                                           {"timestamp": 98765433, "mac":"00:01:02:AA:AB:AD"}
Block FFFF: WjAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAw ]}

# resumepolicy
In case blob-to-queue is stopped, it can continue processing from the file that was last processed. flowlogs have the timestamp in the directory and they can be read in sequence. The registry will then contain the directory that was last processed, because the time stamp is included in the filepath. "y=2023/m=10/d=31/h=14/m=00". The file may have grown with a blob since the last time the file was processed. Because the registry also contains the amount of bytes read since the last time, blob-to-queue will read the new parts.

# flatevents
The original nsgflowlogs and vnetflowlogs are nested json structures, the logic will flatten each log entry as a standalone json event which can be filtered and converted into several formats and sent to several outputs

In the output, it is possible to specify ECS as a format, this is an elasticsearch format that tries to unify the JSON

# configure output
eventhub, kafka, ampq, mqtt, appendfiles, stdout, logstash
and as a bonus for nsgflowlogs, create statistics by grouping the packets and calculating stats about packets and bytes in and out.

# Running
go run blob-to-queue.go

blob-to-queue v1.0-dev

Listing the blobs in the container:
nsgflowlog count:  127  in file resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/OCTOBER-NSG/y=2023/m=10/d=31/h=13/m=00/macAddress=002248A31CA3/PT1H.json
nsgflowlog count:  477  in file resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/OCTOBER-NSG/y=2023/m=10/d=31/h=14/m=00/macAddress=002248A31CA3/PT1H.json

# why?
I wrote janmg/logstash-input-azure_blob_storage based on the original Azure/logstash-input-azureblob, which had scalability issues and I rewrote the plugin from the ground up, because I couldn't fix the original and had to process nsgflowlogs at scale. With Logstash 8.10 there is a Ruby dependancy problem, I can't fix. I decided to ditch Ruby and logstash dependancy and create something that could deal with files in blob storage and push it to a queue, while now instead of using a single file plugin, break the logic in separate files for beter extensibility.

# golang
My problem with JAVA is the Oracle licensing requirements per JVM for large enterprises and how JVM's require a lot of memory reserved upfront. A compiled language will perform better and golang has more library options to code against. I can more easily glue a new feature into the logic, for instance Azure Eventhub, Kafka, AMPQ, MQTT all have libraries where I only need to setup a connection and then send them some logevent. The format of the logevent also is easier to control, because sometimes you want CSV, JSON or just a summary of the network connections or maybe a live view, although delayed by a couple of minutes, by the nsgflowlog writing to the storage account and reading by the plugin.

# kafka
nsgflowlogs are events, it would make more sense to me to have them natively available in an eventhub. An eventhub is an AMPQ / Kafka compatible queueing broker. This program will read from the files that are written every minute and add them as a batch to an output stream. I focus first on writing it to an eventhub, because it is available in Azure. Other output formats are planned are native kafka and amqp and maybe mqtt or any. Eventhubs without traffic already cost me 16 euros per months, so having a cost effective alternative is important for a single individual.

Other output queues are implemented as golang modules exist for ampq, mqtt, keyval, etcetera

# What ChatGPT thinks of my code
✅ Prevent Deadlocks: Use a buffered channel and select for non-blocking writes.

✅ Faster JSON Parsing: Use json.Decoder for efficiency.

✅ Parallel Processing: Implement worker Goroutines for concurrency.

✅ Efficient Elasticsearch Writes: Use the Bulk API to reduce network calls.

✅ Better Logging & Error Handling: Implement structured logging with Zap or Logrus.