package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// TODO: Export all configs not just blob account
type Blob struct {
	Accountname   string `mapstructure:"accountName"`
	Accountkey    string `mapstructure:"accountkey"`
	ContainerName string `mapstructure:"containerName"`
	Cloud         string `mapstructure:"cloud"`
}

/*
type output struct {
	exhaust string
	connect string
	format  string
}
*/

var blob Blob

type IBlob interface {
	Print()
}

func configHandler() Blob {
	// https://github.com/spf13/viper#watching-and-re-reading-config-files
	var conf = viper.New()
	conf.SetConfigFile("blob-to-queue.yaml")
	// TODO add default config file and one that contains my private secrets
	conf.SetConfigType("yaml")
	conf.AddConfigPath(".")
	err := conf.ReadInConfig()
	Error(err)

	conf.Unmarshal(&blob)
	blob.Cloud = "blob.core.windows.net"

	viper.WatchConfig()
	if viper.GetBool("fsnotify") {
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("Config file changed:", e.Name)
			conf.Unmarshal(&blob)
			//lookup = append(lookup, output{"stdout", "", "Flat"})
			//lookup = append(lookup, output{"summary", "", "Flat"})
			//lookup = append(lookup, output{"azurehub", viper.GetString("eventhub.connectionString"), viper.GetString("eventhub.format")})
		})
	}
	//blob.print()
	return blob
}

func (blob Blob) print() {
	fmt.Println(blob.Accountname)
	fmt.Println(blob.Accountkey)
	fmt.Println(blob.ContainerName)
	fmt.Println(blob.Cloud)
}
