package common

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Output struct {
	Format            string `mapstructure:"format"`
	Filename          string `mapstructure:"filename"`
	Connectionstring  string `mapstructure:"connection"`
	Connectiondetails string `mapstructure:"connectiondetails"`
}

// OutputConfig is a unified configuration structure for outputs
type OutputConfig struct {
	Connection string   `mapstructure:"connection"`
	Token      string   `mapstructure:"token"`
	Format     string   `mapstructure:"format"`
	Details    []string `mapstructure:"details"`
	Filename   string   `mapstructure:"filename"`
}

type Config struct {
	Connection    string   `mapstructure:"connection"`
	Accountname   string   `mapstructure:"accountName"`
	Accountkey    string   `mapstructure:"accountkey"`
	SasToken      string   `mapstructure:"sastoken"`
	ContainerName string   `mapstructure:"containerName"`
	Cloud         string   `mapstructure:"cloud"`
	Type          string   `mapstructure:"type"`
	Format        string   `mapstructure:"format"`
	Interval      int      `mapstructure:"interval"`
	Startpolicy   string   `mapstructure:"startpolicy"`
	Resumepolicy  string   `mapstructure:"resumepolicy"`
	Timestamp     string   `mapstructure:"timestamp"`
	Registry      string   `mapstructure:"registry"`
	PathPrefix    []string `mapstructure:"path_prefix"`
	PathInclude   []string `mapstructure:"path_include"`
	PathFilter    []string `mapstructure:"path_filter"`
	Qsize         int      `mapstructure:"qsize"`
	Qwatermark    int      `mapstructure:"qwaterwark"`
	Filename      bool     `mapstructure:"addfilename"`
	Environment   string   `mapstructure:"environment"`
	Output        []string `mapstructure:"output"`
	// Per-output structured configs
	Stdout        OutputConfig `mapstructure:"stdout"`
	File          OutputConfig `mapstructure:"file"`
	Summary       OutputConfig `mapstructure:"summary"`
	Eventhub      OutputConfig `mapstructure:"eventhub"`
	Kafka         OutputConfig `mapstructure:"kafka"`
	Ampq          OutputConfig `mapstructure:"ampq"`
	Mqtt          OutputConfig `mapstructure:"mqtt"`
	Logstash      OutputConfig `mapstructure:"logstash"`
	Elasticsearch OutputConfig `mapstructure:"elasticsearch"`
}

func ConfigHandler() Config {
	// https://github.com/spf13/viper#watching-and-re-reading-config-files
	var conf = viper.New()

	// TODO: create logic to build AccountName and AccountKey into a Connectionstring, also for SAS token
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

	// Read primary config file
	conf.SetConfigFile("blob-to-queue.yaml")
	conf.SetConfigType("yaml")
	conf.AddConfigPath(".")
	err := conf.ReadInConfig()
	if err != nil {
		os.Exit(2)
	}

	// If a private config exists, load it and merge so private values override
	priv := viper.New()
	priv.SetConfigFile("blob-to-queue-private.yaml")
	priv.SetConfigType("yaml")
	priv.AddConfigPath(".")
	if err := priv.ReadInConfig(); err == nil {
		// Merge private settings into main config (private overrides)
		conf.MergeConfigMap(priv.AllSettings())
	} else {
		// If the private file is not present, that's fine — continue with primary config
		if _, statErr := os.Stat("blob-to-queue-private.yaml"); statErr == nil {
			fmt.Printf("private file exists but failed to read — report error: %v\n", err)
		}
	}

	var config Config
	conf.Unmarshal(&config)

	conf.WatchConfig()
	if conf.GetBool("fsnotify") {
		conf.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("Config file changed:", e.Name)
			conf.Unmarshal(&config)
			// TODO: reinitialize the changed config
			//lookup = append(lookup, output{"stdout", "", "Flat"})
			//lookup = append(lookup, output{"summary", "", "Flat"})
			//lookup = append(lookup, output{"azurehub", viper.GetString("eventhub.connectionString"), viper.GetString("eventhub.format")})
		})
	}
	return config
}

func configPrint(conf Config) {
	fmt.Println(conf.Connection)
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
