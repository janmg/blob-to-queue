# blob-to-queue
Golang application for reading from an azure blob storage, listing the files that match a path filter and pushing the result to an event queue like kafka stream. Intended to offer an alternative to my own logstash module logstash-input-azure_blob_storage and move it's logic to a standalone program.

# flatevents
The original nsgflowlogs are nested json structures, the nsgflowlog logic will flatten the logic to make each log entry a standalone json event which can be filtered and converted into several formats and sent to several outputs

# configure
The configurationfile is a YAML file, which is reloadable on save, because its using spf13/viper. The configuration format allows for multiple output streams.
configure multiple outputs

blob-to-queue.yaml

# configure input
accountName: "blobstoragename"
accountKey: "AMWsmPcgy/1234567890123445abcdefghijkl/1234567890123445abcdefghijklABCDEFGHI+ASt3SvXjw=="
containerName: "insights-logs-networksecuritygroupflowevent"

# configure output
eventhub, kafka, ampq, mqtt, appendfiles, stdout, logstash
and as a bonus for nsgflowlogs, create statistics by grouping the packets and calculating stats about packets and bytes in and out.

# Running
go run blob-to-queue.go flatevent.go config.go ecs.go blob.go nsgflowlogs.go format.go azure-eventhub.go kafka.go mqtt.go ampq.go zeromq.go stdout.go append.go stats.go
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