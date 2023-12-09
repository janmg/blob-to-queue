# blob-to-queue
Golang application for reading from a blob storage, parsing the files and pushing the result to an event queue like kafka stream. Intended to replace the logstash module logstash-input-azure_blob_storage and move it's logic to a standalone program.

# why?
Azure/logstash-input-azureblob had limitations and I rewrote the plugin from the ground up, because I needed something to parse nsgflowlogs at scale. With Logstash 8.10 there is a Ruby dependancy problem, I can't fix. I decided to ditch Ruby and logstash dependancy and create something that could deal with files in blob storage and push it to a queue.

# golang
My problem with JAVA is the Oracle licensing requirements per JVM and they require a lot of memory reserved upfront. A compiled language will perform better and golang has more library options to code against. I can more easily glue a new feature into the logic

# kafka
nsgflowlogs are events, it would make more sense to me to have them available in an eventhub. An eventhub is an AMPQ / Kafka compatible queueing broker. This program will read from the files that are written every minute and add them as a batch to an output stream. I focus first on writing it to an eventhub, because it is available in Azure. Other output formats are planned are native kafka and amqp and maybe mqtt or any.

# flatevents
The original nsgflowlogs are nested, the nsgflowlog logic will flatten the logic to make each log entry a standalone json event.
