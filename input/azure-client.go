package input

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"janmg.com/blob-to-queue/common"
)

// https://learn.microsoft.com/en-us/azure/storage/blobs/storage-blob-go-get-started
func getServiceClientTokenCredential(accountURL string) *azblob.Client {
	// Create a new service client with token credential
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	common.Error(err)

	client, err := azblob.NewClient(accountURL, credential, nil)
	common.Error(err)
	return client
}

func getServiceClientSAS(accountURL string, sasToken string) *azblob.Client {
	// Create a new service client with an existing SAS token

	// Append the SAS to the account URL with a "?" delimiter
	accountURLWithSAS := fmt.Sprintf("%s?%s", accountURL, sasToken)

	client, err := azblob.NewClientWithNoCredential(accountURLWithSAS, nil)
	common.Error(err)
	return client
}

func getServiceClientSharedKey(accountName string, accountKey string) *azblob.Client {
	// Create a new service client with shared key credential
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	common.Error(err)

	config := common.ConfigHandler()
	accountURL := fmt.Sprintf("https://%s.%s", accountName, config.Cloud)

	client, err := azblob.NewClientWithSharedKeyCredential(accountURL, credential, nil)
	common.Error(err)
	return client
}

func getServiceClientConnectionString(connectionString string) *azblob.Client {
	// DefaultEndpointsProtocol=https;AccountName=<accountName>;AccountKey=<accountKey>;EndpointSuffix=core.windows.net
	client, err := azblob.NewClientFromConnectionString(connectionString, nil)
	common.Error(err)
	return client
}

func getClient[T any](x T) *azblob.Client {
	// Robustly extract a connection string from the generic value and call the connection-string client.
	// Support common cases: string, []string (take first), fmt.Stringer.
	var conn string
	switch v := any(x).(type) {
	case string:
		conn = v
	case []string:
		if len(v) > 0 {
			conn = v[0]
		}
	case fmt.Stringer:
		conn = v.String()
	default:
		// Last resort: format the value as a string
		conn = fmt.Sprintf("%v", v)
	}

	return getServiceClientConnectionString(conn)
}
