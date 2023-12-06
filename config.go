package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Blob struct {
	Accountname   string `mapstructure:"accountName"`
	Accountkey    string `mapstructure:"accountkey"`
	ContainerName string `mapstructure:"containerName"`
	Cloud         string `mapstructure:"cloud"`
}

type output struct {
	exhaust string
	connect string
	format  string
}

func configHandler() Blob {
	// https://github.com/spf13/viper#watching-and-re-reading-config-files
	var conf = viper.New()
	conf.SetConfigFile("blob-to-kafka.yaml")
	conf.SetConfigType("yaml")
	conf.AddConfigPath(".")
	err := conf.ReadInConfig()
	handleError(err)
	var blob Blob
	conf.Unmarshal(&blob)
	handleError(err)

	blob.Cloud = "blob.core.windows.net"

	viper.WatchConfig()
	if viper.GetBool("fsnotify") {
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("Config file changed:", e.Name)
			blob.Accountname = viper.GetString("accountName")
			blob.Accountkey = viper.GetString("accountKey")
			blob.ContainerName = viper.GetString("containerName")
			blob.Cloud = viper.GetString("cloud")
			lookup = nil
			lookup = append(lookup, output{"stdout", "", "Flat"})
			lookup = append(lookup, output{"summary", "", "Flat"})
			lookup = append(lookup, output{"azurehub", viper.GetString("eventhub.connectionString"), viper.GetString("eventhub.format")})
		})
	}

	return blob
}
