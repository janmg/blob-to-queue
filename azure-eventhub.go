package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs"
)

func sendAzure(nsg flatevent) {
	fmt.Println("Azure sending")
	// "containerName"

	// "Endpoint=sb://nsgflowlogs.servicebus.windows.net/;SharedAccessKeyName=abc;SharedAccessKey=123"
	// *.servicebus.chinacloudapi.cn, *.servicebus.usgovcloudapi.net, or *.servicebus.cloudapi.de
	connectionString := "Endpoint=sb://nsgflowlogs.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=abcdefg0123456789ABCDEF="
	kfk, err := azeventhubs.NewProducerClientFromConnectionString(connectionString, "janmg", nil)
	handleError(err)

	defer kfk.Close(context.TODO())

	batch, err := kfk.NewEventDataBatch(context.TODO(), nil)
	handleError(err)
	// TODO: currently serving a batch of one, need to figure out how to suck more out of the queue?
	//err = batch.AddEventData(eventdata(nsg), nil)
	ed := eventdata(nsg)
	fmt.Println(ed)
	err = batch.AddEventData(ed, nil)
	handleError(err)

	err = kfk.SendEventDataBatch(context.Background(), batch, nil)
	handleError(err)
}

func eventdata(nsg flatevent) *azeventhubs.EventData {
	json, err := json.Marshal(nsg)
	handleError(err)
	fmt.Println(json)
	return &azeventhubs.EventData{
		Body: json,
	}
}
