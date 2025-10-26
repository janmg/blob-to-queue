package common

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// TODO: Export all configs not just blob account
type Config struct {
	Accountname   string   `mapstructure:"accountName"`
	Accountkey    string   `mapstructure:"accountkey"`
	ContainerName string   `mapstructure:"containerName"`
	Cloud         string   `mapstructure:"cloud"`
	Startpolicy   string   `mapstructure:"startpolicy"`
	Resumepolicy  string   `mapstructure:"resumepolicy"`
	Timestamp     string   `mapstructure:"timestamp"`
	Registry      string   `mapstructure:"registry"`
	PathPrefix    []string `mapstructure:"path_prefix"`
	PathInclude   []string `mapstructure:"path_include"`
	PathFilter    []string `mapstructure:"path_filter"`
	Interval      int      `mapstructure:"interval"`
	Qsize         int      `mapstructure:"qsize"`
	Qwatermark    int      `mapstructure:"qwaterwark"`
	Format        string   `mapstructure:"format"`
	Type          string   `mapstructure:"type"`
	Output        []string `mapstructure:"output"`
}

/*
type output struct {
	exhaust string
	connect string
	format  string
}
*/

//type IBlob interface {
//	configPrint()
//}

func ConfigHandler() Config {
	// https://github.com/spf13/viper#watching-and-re-reading-config-files
	var conf = viper.New()

	conf.SetDefault("cloud", "blob.core.windows.net")

	conf.SetDefault("resumepolicy", "timestamp")
	// ['timestamp','registry']
	conf.SetDefault("registry", "./registry.dat")
	conf.SetDefault("timestamp", "./timestamp.json")
	conf.SetDefault("startpolicy", "start_over")
	// ['start_over','start_fresh']
	conf.SetDefault("output", "stdout")
	conf.SetDefault("interval", 60)

	// "resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/OCTOBER-NSG/y=2023/m=10/d=31/h=18/m=00"
	conf.SetDefault("path_prefix", "['**/*']")  // array of prefixes a path must start with, "resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/OCTOBER-NSG/"
	conf.SetDefault("path_include", "['**/*']") // array of strings that must occur, non-matching paths are ignored
	conf.SetDefault("path_filter", "['**/*']")  // array of strings that will be filtered out
	conf.SetDefault("type", "nsgflowlog")
	conf.SetDefault("format", "json")
	/*
		filtering down path list, only look for subdirectories and files that start with a path, then only qualify the paths that fit the filter, then exclude some that you don't want
		prefix: resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM
		path_include ** /*-NSG/** /*.json
		path_exclude ** /*y=2022/**
	*/

	conf.SetConfigFile("blob-to-queue.yaml")
	conf.SetConfigType("yaml")
	conf.AddConfigPath(".")
	err := conf.ReadInConfig()
	Error(err)

	var config Config
	conf.Unmarshal(&config)

	conf.WatchConfig()
	if conf.GetBool("fsnotify") {
		conf.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("Config file changed:", e.Name)
			conf.Unmarshal(&config)
			//lookup = append(lookup, output{"stdout", "", "Flat"})
			//lookup = append(lookup, output{"summary", "", "Flat"})
			//lookup = append(lookup, output{"azurehub", viper.GetString("eventhub.connectionString"), viper.GetString("eventhub.format")})
		})
	}
	return config
}

func configPrint(conf Config) {
	fmt.Println(conf.Accountname)
	fmt.Println(conf.Accountkey)
	fmt.Println(conf.ContainerName)
	fmt.Println(conf.Cloud)
	fmt.Println(conf.Startpolicy)
	fmt.Println(conf.Resumepolicy)
	fmt.Println(conf.Timestamp)
	fmt.Println(conf.Interval)
	fmt.Println(conf.Registry)
	fmt.Println(conf.PathPrefix)
	fmt.Println(conf.PathInclude)
	fmt.Println(conf.PathFilter)
	fmt.Println(conf.Type)
	fmt.Println(conf.Format)
	fmt.Println(conf.Output)
}
