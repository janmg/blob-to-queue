# Blob-to-Queue Callflow Diagram

## Architecture Overview

```
┌────────────────────────────────────┐
│          MAIN APPLICATION          │
│         (blob-to-queue.go)         │
└────────────────────────────────────┘
                    │
        ┌───────────┼─────────┐
        ▼                     ▼
┌─────────────────┐   ┌──────────────┐
│  Blobworker     │   │   Flatevent  │
│  (goroutine)    │   │   (Channel)  │
└─────────────────┘   └──────────────┘
```

## Detailed Callflow

### 1. Initialization Phase

```
main()
  │
  ├─► common.ConfigHandler()
  │     └─► Load blob-to-queue.yaml
  │           ├─► accountName, accountKey, containerName
  │           ├─► resumepolicy (timestamp/registry)
  │           ├─► startpolicy (start_over/start_fresh)
  │           ├─► type (nsgflowlog/vnetflowlog)
  │           ├─► output[] (elasticsearch, kafka, etc.)
  │           └─► interval, qsize, qwatermark
  │
  ├─► signal.Notify(SIGINT, SIGTERM)
  │     └─► Graceful shutdown handler
  │
  ├─► make(chan format.Flatevent, qsize)
  │     └─► Create buffered channel
  │
  ├─► go input.Blobworker(queue)  [GOROUTINE 1]
  │
  └─► send(queue)  [MAIN LOOP]
```

### 2. Blob Worker Flow (Goroutine 1)

```
input.Blobworker(queue)
  │
  ├─► Initialize registry/timestamp
  │     ├─► If resumepolicy == "registry"
  │     │     └─► loadRegistry("registry.json")
  │     └─► If resumepolicy == "timestamp"
  │           └─► readTimestamp("timestamp.json")
  │
  ├─► doLoop() [INITIAL SYNC]
  │
  └─► time.NewTicker(interval)
        └─► for range interval.C
              └─► doLoop() [PERIODIC SYNC]


doLoop(config, queue, registry, last)
  │
  ├─► Check queue watermark
  │     └─► If len(queue) > qwatermark
  │           └─► Sleep 10s until queue drains
  │
  ├─► listFiles(resumepolicy, account, key, location, last)
  │     │
  │     ├─► NewSharedKeyCredential(accountName, accountKey)
  │     ├─► NewClientWithSharedKeyCredential(location, cred)
  │     ├─► NewListBlobsFlatPager(containerName)
  │     │
  │     └─► for pager.More()
  │           └─► NextPage(context.Background())
  │                 │
  │                 ├─► If resumepolicy == "timestamp"
  │                 │     └─► Filter: blob.Name > lastread timestamp
  │                 │
  │                 └─► Return: map[filename]size
  │
  ├─► Compare filelist vs registry
  │     │
  │     ├─► If exists && size changed → PARTIAL READ
  │     │     └─► read(queue, name, oldSize, size)
  │     │
  │     └─► If !exists → NEW FILE
  │           └─► read(queue, name, 0, size)
  │
  └─► Save state
        ├─► If resumepolicy == "timestamp"
        │     └─► writeTimestamp("timestamp", now)
        └─► If resumepolicy == "registry"
              └─► saveRegistry("registry.json", filelist)
```

### 3. Blob Read & Parse Flow

```
read(queue, name, oldSize, size)
  │
  ├─► NewSharedKeyCredential()
  ├─► NewClientWithSharedKeyCredential()
  │
  ├─► DownloadStream(ctx, containerName, name, options)
  │     └─► HTTPRange: {Offset: oldSize+1, Count: size-oldSize}
  │
  ├─► downloadedData.ReadFrom(retryReader)
  │
  └─► Parse based on config.Type
        │
        ├─► If type == "nsgflowlog"
        │     └─► nsgflowlog(queue, data, blobname)
        │
        └─► If type == "vnetflowlog"
              └─► vnetflowlog(queue, data, blobname)


nsgflowlog(queue, flowlogs, blobname)
  │
  ├─► json.Unmarshal(flowlogs, &nsgflowlogs)
  │
  └─► for each record
        └─► for each flow
              └─► for each tuple
                    │
                    ├─► addtuples(event, nsgflow)
                    │     └─► Parse: timestamp, srcIP, dstIP, srcPort, dstPort
                    │           proto, direction, action, state, packets, bytes
                    │
                    └─► queue <- event  [SEND TO CHANNEL]


vnetflowlog(queue, flowlogs, blobname)
  │
  ├─► json.Unmarshal(flowlogs, &vnetflowlogs)
  │
  └─► for each record
        └─► for each flow
              └─► for each flowGroup
                    └─► for each tuple
                          │
                          ├─► addtuples(event, vnetflow)
                          │
                          └─► queue <- event  [SEND TO CHANNEL]
```

### 4. Output Processing Flow (Main Loop)

```
send(queue)
  │
  └─► for {  [INFINITE LOOP]
        │
        ├─► event := <-queue  [RECEIVE FROM CHANNEL]
        │
        └─► for each output in config.Output[]
              │
              ├─► case "elasticsearch"
              │     └─► output.SendElasticsearch(event)
              │
              ├─► case "eventhub"
              │     └─► output.SendAzure(event)
              │
              ├─► case "kafka"
              │     └─► output.SendKafka(event)
              │           └─► kafka.DialLeader()
              │                 kafka.WriteMessages()
              │
              ├─► case "mqtt"
              │     └─► output.SendMQTT(event)
              │
              ├─► case "ampq"
              │     └─► output.SendAMPQ(event)
              │           ├─► amqp.Dial(amqpURL)
              │           ├─► conn.Channel()
              │           └─► channel.PublishWithContext()
              │
              ├─► case "zeromq"
              │     └─► output.SendZERO(event)
              │           ├─► zmq4.NewPub()
              │           └─► socket.Send()
              │
              ├─► case "keyval"
              │     └─► output.SendKeyval(event)
              │           ├─► clientv3.New() [etcd]
              │           └─► client.Put(key, value)
              │
              ├─► case "redis"
              │     └─► output.SendRedis(event)
              │
              ├─► case "fluent"
              │     └─► output.SendFluent(event)
              │
              ├─► case "fluxdb"
              │     └─► output.SendFlux(event)
              │
              ├─► case "file"
              │     └─► output.AppendFile(event)
              │
              ├─► case "stdout"
              │     └─► output.Stdout(event)
              │           └─► format.Format(config.Format, event)
              │
              └─► case "summary"
                    └─► output.Statistics(event)
```

## Data Flow Diagram

```
┌─────────────────────────────────────────────────────────────────┐
│                    Azure Blob Storage                           │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │ /SUBSCRIPTIONS/.../NETWORKSECURITYGROUPS/               │    │
│  │   ├─ y=2023/m=10/d=31/h=13/m=00/PT1H.json               │    │
│  │   └─ y=2023/m=10/d=31/h=14/m=00/PT1H.json               │    │
│  └─────────────────────────────────────────────────────────┘    │
└─────────────────────────────────────────────────────────────────┘
                            │
                            │ Azure SDK
                            │ (List & Download)
                            ▼
┌─────────────────────────────────────────────────────────────────┐
│                       Blobworker                                │
│  ┌──────────────┐     ┌──────────────┐     ┌──────────────┐     │
│  │  listFiles() │────►│   read()     │────►│nsgflowlog()  │     │
│  │              │     │              │     │vnetflowlog() │     │
│  └──────────────┘     └──────────────┘     └──────────────┘     │
│         │                    │                      │           │
│         ▼                    ▼                      ▼           │
│  ┌──────────────┐     ┌──────────────┐     ┌──────────────┐     │
│  │   Registry/  │     │  HTTP Range  │     │Parse FlowLog │     │
│  │  Timestamp   │     │   Download   │     │    JSON      │     │
│  └──────────────┘     └──────────────┘     └──────────────┘     │
└─────────────────────────────────────────────────────────────────┘
                             │
                             │ queue <- event
                             ▼
┌─────────────────────────────────────────────────────────────────┐
│                  Buffered Channel Queue                         │
│              chan format.Flatevent (qsize)                      │
└─────────────────────────────────────────────────────────────────┘
                             │
                             │ event := <-queue
                             ▼
┌─────────────────────────────────────────────────────────────────┐
│                        send() Loop                              │
│                  Multiple Output Handlers                       │
│  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────┐            │
│  │Elastic   │ │ Kafka    │ │ AMQP     │ │ ZeroMQ   │            │
│  │search    │ │          │ │          │ │          │  ...       │
│  └──────────┘ └──────────┘ └──────────┘ └──────────┘            │
└─────────────────────────────────────────────────────────────────┘
                            │
                            │ Send to external systems
                            ▼
┌─────────────────────────────────────────────────────────────────┐
│              External Systems (Destinations)                    │
│  • Elasticsearch    • Azure Event Hub   • Kafka                 │
│  • RabbitMQ (AMQP)  • ZeroMQ            • etcd (keyval)         │
│  • Redis            • Fluentd           • InfluxDB              │
│  • MQTT             • Files             • stdout                │
└─────────────────────────────────────────────────────────────────┘
```

## Concurrency Model

```
┌──────────────────────────────────────────────────────────────┐
│                     Main Goroutine                           │
│                                                              │
│  ┌─────────────────────────────────────────────────────┐     │
│  │  send() - Infinite Loop                             │     │
│  │  • Reads from channel                               │     │
│  │  • Dispatches to multiple outputs                   │     │
│  └─────────────────────────────────────────────────────┘     │
└──────────────────────────────────────────────────────────────┘
                            ▲
                            │
                    Buffered Channel
                   (format.Flatevent)
                            │
┌──────────────────────────────────────────────────────────────┐
│                  Blobworker Goroutine                        │
│                                                              │
│  ┌─────────────────────────────────────────────────────┐     │
│  │  Ticker-based Loop (interval)                       │     │
│  │  • Lists blobs from Azure Storage                   │     │
│  │  • Downloads new/updated files                      │     │
│  │  • Parses flow logs                                 │     │
│  │  • Writes events to channel                         │     │
│  └─────────────────────────────────────────────────────┘     │
└──────────────────────────────────────────────────────────────┘
                            ▲
                            │
┌──────────────────────────────────────────────────────────────┐
│                 Signal Handler Goroutine                     │
│                                                              │
│  ┌─────────────────────────────────────────────────────┐     │
│  │  Waits for SIGINT/SIGTERM                           │     │
│  │  • Monitors queue drainage                          │     │
│  │  • Graceful shutdown                                │     │
│  └─────────────────────────────────────────────────────┘     │
└──────────────────────────────────────────────────────────────┘
```

## State Management

### Resume Policies

**Timestamp Mode:**
```
timestamp.json:
{
  "year": 2023,
  "month": 10,
  "day": 31,
  "hour": 13,
  "minute": 0
}

→ Filters blobs: y=2023/m=10/d=31/h=13/m=00/
→ Only processes files newer than last timestamp
```

**Registry Mode:**
```
registry.json:
{
  "resourceId=.../PT1H.json": 45678,
  "resourceId=.../PT1H.json": 89012
}

→ Tracks file sizes
→ Detects new files (!exists)
→ Detects updated files (size changed)
→ Partial reads (oldSize to newSize)
```

## Key Features

1. **Concurrent Processing**: Producer (Blobworker) and Consumer (send) run in parallel
2. **Backpressure Handling**: Queue watermark prevents memory overflow
3. **Resume Capability**: Timestamp or registry-based state tracking
4. **Partial Reads**: HTTP Range requests for incremental file processing
5. **Multiple Outputs**: Fan-out to multiple destinations simultaneously
6. **Graceful Shutdown**: Signal handler ensures queue drainage before exit
7. **Configurable Interval**: Periodic sync based on configuration
