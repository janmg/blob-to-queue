# blob-to-queue
Golang application for reading from an azure blob storage, listing the files that match a path filter and pushing the result to an event queue like kafka stream. Intended to offer an alternative to my own logstash module logstash-input-azure_blob_storage and move it's logic to a standalone program.

# why?
I wrote janmg/logstash-input-azure_blob_storage based on the original Azure/logstash-input-azureblob, which had scalability issues and I rewrote the plugin from the ground up, because I couldn't fix the original and had to process nsgflowlogs at scale. With Logstash 8.10 there is a Ruby dependancy problem, I can't fix. I decided to ditch Ruby and logstash dependancy and create something that could deal with files in blob storage and push it to a queue, while setting the possibility to break the 600 lines in one file, to a more logical separation of functions.

# golang
My problem with JAVA is the Oracle licensing requirements per JVM for large enterprises and how JVM's require a lot of memory reserved upfront. A compiled language will perform better and golang has more library options to code against. I can more easily glue a new feature into the logic, for instance Azure Eventhub, Kafka, AMPQ, MQTT all have libraries where I only need to setup a connection and then send them some logevent. The format of the logevent also is easier to control, because sometime you want CSV, JSON or just a summary of the network connections or maybe a live view, although delayed by a couple of minutes, by the nsgflowlog writing to the storage account and reading by the plugin.

# kafka
nsgflowlogs are events, it would make more sense to me to have them natively available in an eventhub. An eventhub is an AMPQ / Kafka compatible queueing broker. This program will read from the files that are written every minute and add them as a batch to an output stream. I focus first on writing it to an eventhub, because it is available in Azure. Other output formats are planned are native kafka and amqp and maybe mqtt or any. Eventhubs without traffic already cost me 16 euros per months, so having a cost effective alternative is important for a single individual.

# flatevents
The original nsgflowlogs are nested, the nsgflowlog logic will flatten the logic to make each log entry a standalone json event.
