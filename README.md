# Azure VNET Flow Logs
In Azure, it's possible to log network flowlogs to a blobstorage, which gives you an hourly file pt1h.json file that grows every minute. This application can list the files from the blobstorage, read them and process them as individual events and feed them to a queue for further processing. Then an interval later, if the file has grown it can read the delta and process new events. Events can be written to Elasticsearch, Kafka or Eventhub or any other protocol. I have made the output processor for AMPQ, Fluentd, Fluxdb, Keyval, Redis, MQTT or ZeroMQ.

# Logstash_azure_blob_storage
This Golang application does not need logstash. I wrote and maintained the logstash plugin logstash_azure_blob_storage in Ruby, as a fix to the original logstash plugin logstash-input-azureblob. But due to Faraday or Nogiri dependancy conflicts also my logstash_azure_blob_storage only works on some of the later logstash versions. This standalone application does not have to consider those dependancies.

# configure
The blob-to-queue.yaml configurationfile is a YAML file, uses spf13/viper, which makes it reloadable on save. The configuration format allows for multiple output streams.

# configure input
```
accountName: "blobstoragename"
accountKey: "AMWsmPcgy/1234567890123445abcdefghijkl/1234567890123445abcdefghijklABCDEFGHI+ASt3SvXjw=="
containerName: "insights-logs-networksecuritygroupflowevent"
```

# configure output
The configuration file uses an array for output to activate the output formats and then configure the outputs. For instance set "output" to write to elasticsearch and to file and then configure the destination and format you like to 
```
output: ["elasticsearch","file"]
file:
  filename: "./nsg.log"
  format: "csv"
elasticsearch:
  addresses: "https://elastic-1:9200"
  serviceToken: "AAEAAWVsYXRmEtblE"
  format: "ecs"
  index: "nsg-"
```
# Microsoft blobstorage
pt1h.json logfiles are written to blobstorage every hour. They have a header in the first block and a footer in the last block and the content grows every minute with another json fragment. These fragments contain all the events for that minute. The blob-to-queue application can read the increments every minute and process them
```
Block 0000: QTAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAw {message: [
Block 0001: Bobaba                                          {"timestamp": 98765432, "mac":"00:01:02:AA:AB:AC"}
Block 0002: Bobabb                                          ,{"timestamp": 98765433, "mac":"00:01:02:AA:AB:AD"}
Block FFFF: WjAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAw ]}
```

# resumepolicy
Can be set to	timestamp or registry. In case blob-to-queue is stopped, it can continue processing from the file that was last processed. For vnetflowlogs and nsgflowlogs, the resuming can be done based on the last processed timestamp. The registry will contain the directory that was last processed, because the timestamp is included in the filepath. "y=2023/m=10/d=31/h=14/m=00". The file may have grown with a blob since the last time the file was processed.

To make this application suitable for any other format it can be configured to scan all the files in the storage account or scan part of the storage account based on a file filter. Then a registry file can be kept to remember how many bytes of which files have been processed. The registry will then contain a list of files and the amount of bytes read since the last time, blob-to-queue can then start reading only the new fragments.

# flatevents
The original nsgflowlogs and vnetflowlogs are nested json structures, the logic will flatten each log entry as a standalone json event which can be filtered and converted into several formats and sent to several outputs

In the output, it is possible to specify ECS as a format, this is an elasticsearch format that tries to unify the JSON fields for correlating the events.

# Running
go run blob-to-queue.go

blob-to-queue v1.0-dev

Listing the blobs in the container:
```
nsgflowlog count:  127  in file resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/OCTOBER-NSG/y=2023/m=10/d=31/h=13/m=00/macAddress=002248A31CA3/PT1H.json
nsgflowlog count:  477  in file resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/OCTOBER-NSG/y=2023/m=10/d=31/h=14/m=00/macAddress=002248A31CA3/PT1H.json
```

# why?
I wrote janmg/logstash-input-azure_blob_storage based on the original Azure/logstash-input-azureblob, which had scalability issues and I rewrote the plugin from the ground up, because I couldn't fix the original and had to process nsgflowlogs at scale. With Logstash 8.10 there is a Ruby dependancy problem, I can't fix. I decided to ditch Ruby and logstash dependancy and create something that could deal with files in blob storage and push it to a queue, while now instead of using a single file plugin, break the logic in separate files for beter extensibility.

# golang
My problem with JAVA is the Oracle licensing requirements per JVM for large enterprises and how JVM's require a lot of memory reserved upfront. A compiled language will perform better and golang has more library options to code against. I can more easily glue a new feature into the logic, for instance Azure Eventhub, Kafka, AMPQ, MQTT all have libraries where I only need to setup a connection and then send them some logevent. The format of the logevent also is easier to control, because sometimes you want CSV, JSON or just a summary of the network connections or maybe a live view, although delayed by a couple of minutes, by the nsgflowlog writing to the storage account and reading by the plugin.

# kafka
nsgflowlogs are events, it would make more sense to me to have them natively available in an eventhub. An eventhub is an AMPQ / Kafka compatible queueing broker. This program will read from the files that are written every minute and add them as a batch to an output stream. I focus first on writing it to an eventhub, because it is available in Azure. Other output formats are planned are native kafka and amqp and maybe mqtt or any. Eventhubs without traffic already cost me 16 euros per months, so having a cost effective alternative is important for a single individual.

