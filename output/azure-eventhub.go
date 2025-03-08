package output

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs"
	"janmg.com/blob-to-queue/common"
	"janmg.com/blob-to-queue/format"
)

func SendAzure(nsg format.Flatevent) {
	fmt.Println("Azure Eventhub sending")
	// "containerName"

	// "Endpoint=sb://nsgflowlogs.servicebus.windows.net/;SharedAccessKeyName=abc;SharedAccessKey=123"
	// *.servicebus.chinacloudapi.cn, *.servicebus.usgovcloudapi.net, or *.servicebus.cloudapi.de

	// TODO: Make into a config item ...
	connectionString := "Endpoint=sb://nsgflowlogs.servicebus.windows.net/;SharedAccessKeyName=nsgflowlogs;SharedAccessKey=yq6akzSdE8YU1fsV1RoF9KVGodWvfMgx8+AEhHTKP9A="
	kfk, err := azeventhubs.NewProducerClientFromConnectionString(connectionString, "janmg", nil)
	common.Error(err)
	// after removing the eventhub, this is the error
	// 2024/01/20 13:30:13 (connlost): dial tcp: lookup nsgflowlogs.servicebus.windows.net: no such host

	defer kfk.Close(context.TODO())

	batch, err := kfk.NewEventDataBatch(context.TODO(), nil)
	common.Error(err)
	// TODO: currently serving a batch of one, need to figure out how to suck more out of the queue?
	//err = batch.AddEventData(eventdata(nsg), nil)
	ed := eventdata(nsg)
	err = batch.AddEventData(ed, nil)
	common.Error(err)

	//TODO: thread and send based on time and availability
	err = kfk.SendEventDataBatch(context.Background(), batch, nil)
	common.Error(err)
}

func eventdata(nsg format.Flatevent) *azeventhubs.EventData {
	json, err := json.Marshal(nsg)
	common.Error(err)
	fmt.Println(json)
	return &azeventhubs.EventData{
		Body: json,
	}
}
