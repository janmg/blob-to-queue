# blob-to-queue
Golang application for reading from a blob storage, parsing the files and pushing the result to an event queue like kafka stream. Intended to replace the logstash module logstash-input-azure_blob_storage and move it's logic to a standalone program.

# why?
Because I originally needed something to parse nsgflowlogs and I rewrote the logstash plugin because I couldn't fix the original Azure/logstash-input-azureblob plugin. The due to Ruby dependancy problems I couldn't fix my own plugin I decided to ditch Ruby and logstash dependancy and create something that could deal with files in blob storage and push it to a queue.

