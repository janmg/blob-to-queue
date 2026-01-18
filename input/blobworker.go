package input

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"janmg.com/blob-to-queue/common"
	"janmg.com/blob-to-queue/format"
)

type RegistryData struct {
	ContentLength int64     `json:"content_length"`
	ContentRead   int64     `json:"content_read"`
	LastModified  time.Time `json:"last_modified"`
}

var last time.Time

// Package-level state and control for graceful shutdown
var (
	mutex         sync.Mutex
	registry      map[string]RegistryData
	stopRequested bool
	stoppedCh     chan struct{}
)

func Blobworker(queue chan format.Flatevent) {
	config := common.ConfigHandler()

	// Use package-level registry/last so other packages can request final save
	mutex.Lock()
	if registry == nil {
		registry = make(map[string]RegistryData)
	}
	mutex.Unlock()

	// NSGFlowLogs grow in predictable directories and files, only need to keep a timestamp pointer to the last processed directory, other files may have arbitrary names and the registry will track which files are new and which ones have grown.
	// https://pkg.go.dev/time
	if config.Resumepolicy == "timestamp" {
		if config.Startpolicy == "resume" {
			fmt.Println("Reading timestamp from: ", config.Timestamp)
			var err error
			last, err = readTimestamp(config.Timestamp)
			if err != nil {
				last = time.Time{}
				fmt.Println("Can't read timestamp for start_fresh, will start_over instead")
			}
		}
		if config.Startpolicy == "start_fresh" {
			last = time.Now()
			fmt.Println("Starting fresh")
		}
		if config.Startpolicy == "start_over" {
			fmt.Println("Starting over, clearing the timestamp")
			last = time.Time{}
		}
		fmt.Printf("Resuming from timestamp: %v\n", last.Format(time.RFC3339))
	}

	if config.Resumepolicy == "registry" {
		// Read the registry if it exists
		if config.Startpolicy == "start_over" {
			fmt.Println("Starting over, clearing the registry")
			os.Remove(config.Registry) // Delete the registry file
			mutex.Lock()
			registry = make(map[string]RegistryData)
			mutex.Unlock()
		} else {
			var err error
			fmt.Println("Resuming from the registry")
			r, err := loadRegistry("registry.json")
			if err != nil {
				mutex.Lock()
				registry = make(map[string]RegistryData)
				mutex.Unlock()
				fmt.Println("Can't read registry for start_fresh, will start_over instead")
			} else {
				mutex.Lock()
				registry = r
				mutex.Unlock()
			}
		}
		fmt.Println("Resuming from registry")
	}

	// Do the first loop
	doLoop(config, queue, registry, last)

	stoppedCh = make(chan struct{})
	defer func() {
		close(stoppedCh)
	}()

	// Once the first loop is done, continue at every interval
	interval := time.NewTicker(time.Duration(config.Interval) * time.Second)
	defer interval.Stop()
	for range interval.C {
		if stopRequested {
			writeTimestamp(config.Timestamp, last)
			break
		}
		doLoop(config, queue, registry, last)
	}
}

func doLoop(config common.Config, queue chan format.Flatevent, registry map[string]RegistryData, last time.Time) {
	// Should use a queue to signal that it's time to stop the ingress
	for i := 0; len(queue) > config.Qwatermark; i++ {
		fmt.Printf("Hit watermark, queue is at %d pause for 10 seconds", len(queue))
		time.Sleep(10 * time.Second)
		if i%5 == 0 {
			fmt.Println("1 Minute and still above the queue watermark, maybe the output is not processing?")
		}
	}

	//location := "https://" + config.Accountname + "." + config.Cloud
	// fmt.Println(location)

	// list all the nsg's
	// resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/
	// loop through the dates, skip the older ones and process only from the data in the registry
	//	resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/OCTOBER-NSG/y=2023/m=10/d=31/h=13/m=00/
	//	for each nsg
	//	resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/

	// 1. Lists all the files in the remote storage account that match the path prefix
	var filelist map[string]RegistryData = make(map[string]RegistryData)
	filelist = listFiles(config.Resumepolicy, config.Connection, last)
	//fmt.Printf("Received filelist with %d entries\n", len(filelist))
	var fullfiles = 0
	var partialfiles = 0

	// 2. Filters on path_filters to only include files that match the directory and file glob (**/*.json)
	// TODO: filter on the timestamp so that only fresh files get processed

	// 3. Compare the list of files to the the registry with the new filelist

	// 4. Process the worklist and put all events in the logstash queue.

	// Read based on modified flags

	// Iterate filelist in deterministic order (sorted by filename)
	names := make([]string, 0, len(filelist))
	for n := range filelist {
		names = append(names, n)
	}
	sort.Strings(names)

	fmt.Println("Listing the blobs in the container:")
	for _, name := range names {
		metadata := filelist[name]
		if oldMeta, exists := registry[name]; exists && oldMeta.ContentLength != metadata.ContentLength {
			// File exists in registry and size changed - PARTIAL READ
			fmt.Printf("%s grew by %d bytes\n", name, metadata.ContentLength-oldMeta.ContentLength)
			read(queue, name, oldMeta.ContentLength, metadata.ContentLength)
			last = timeFromFile(name)
			partialfiles++
		} else if !exists {
			// File doesn't exist in registry - NEW FILE
			fmt.Printf("%s is new and has %d bytes\n", name, metadata.ContentLength)
			read(queue, name, 0, metadata.ContentLength)
			last = timeFromFile(name)
			fullfiles++
		}
		// If exists and size unchanged, skip processing
	}
	//convert to log item
	fmt.Printf("Found %d new files and %d updated files\n", fullfiles, partialfiles)
	// TODO: This doesn't work well? ... need to consider timestamps, fullfiles is 0 and updated files are 7 ... ???

	// 5. Save the registry with files and sizes to a file
	// TODO: this should only be done when the shutdown is triggered
	//if config.Resumepolicy == "timestamp" {
	//	filename := "resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/OCTOBER-NSG/y=2023/m=10/d=31/h=17/m=00/macAddress=002248A31CA3/PT1H.json"
	//	writeTimestamp(config.Timestamp, time.Now()) // TODO This should be the last processed timestamp
	//}
	if config.Resumepolicy == "registry" {
		saveRegistry(config.Registry, filelist)
	}

	// 6. if there is time left, sleep to complete the interval. If processing takes more than an inteval, save the registry and continue.
	// ... try to sync the timer to when the files are actually written to the storage account and wait an additional 5 seconds before reading.
	// ... did storage accounts implement some time of difference tracking journal?
	// 7. If stop signal comes, finish the current file, save the registry and quit
}

func timeFromFile(name string) time.Time {
	re := regexp.MustCompile(`y=(\d{4})/m=(\d{2})/d=(\d{2})/h=(\d{2})`)
	matches := re.FindStringSubmatch(name)
	if matches == nil {
		fmt.Println("Could not extract date from filename")
		return time.Time{}
	}
	year, _ := strconv.Atoi(matches[1])
	month, _ := strconv.Atoi(matches[2])
	day, _ := strconv.Atoi(matches[3])
	hour, _ := strconv.Atoi(matches[4])
	return time.Date(year, time.Month(month), day, hour, 0, 0, 0, time.UTC)
}

func loadRegistry(path string) (map[string]RegistryData, error) {
	file, err := os.Open(path)
	if err != nil {
		return make(map[string]RegistryData), err
	}
	defer file.Close()

	var registry map[string]RegistryData
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&registry)
	if err != nil {
		return make(map[string]RegistryData), err
	}
	return registry, nil
}

func saveRegistry(path string, filelist map[string]RegistryData) error {
	// resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/OCTOBER-NSG/y=2023/m=10/d=31/h=17/m=00/macAddress=002248A31CA3/PT1H.json
	// resourceId=/SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/OCTOBER-NSG/y=2023/m=10/d=31/h=18/m=00/macAddress=002248A31CA3/PT1H.json
	// y=2023/m=10/d=31/h=18/m=00/macAddress=002248A31CA3/PT1H.json
	file, err := os.Create(path)
	common.Error(err)
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(filelist)
	return err
}

func readTimestamp(path string) (time.Time, error) {
	var ts time.Time

	file, err := os.Open(path)
	// If the file doesn't exist, return empty timestamp
	if os.IsNotExist(err) {
		return ts, nil
	}
	if err != nil {
		common.Warning(err)
		return ts, err
	}
	defer file.Close()

	// Read the file contents
	data := make([]byte, 20)
	n, err := file.Read(data)
	if err != nil && err.Error() != "EOF" {
		common.Warning(err)
		return ts, err
	}

	tsStr := strings.TrimSpace(string(data[:n]))
	ts, err = time.Parse(time.RFC3339, tsStr)
	if err != nil {
		common.Warning(err)
		return time.Time{}, err
	}

	return ts, nil
}

func writeTimestamp(path string, ts time.Time) error {
	file, err := os.Create(path)
	common.Warning(err)
	defer file.Close()

	_, err = file.Write([]byte(ts.Format(time.RFC3339)))
	common.Warning(err)
	return err
}

func StopAndWait() {
	stopRequested = true
	if stoppedCh == nil {
		return
	}
	<-stoppedCh
}

func SaveState() error {
	cfg := common.ConfigHandler()
	if cfg.Resumepolicy == "timestamp" {
		return writeTimestamp(cfg.Timestamp, last)
	} else {
		mutex.Lock()
		defer mutex.Unlock()
		return saveRegistry(cfg.Registry, registry)
	}
}
