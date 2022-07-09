package main

import (
	"auto-setup/install_manager"
	"fmt"
	"log"
	"time"
)

func main() {

	downloadPath := install_manager.GetDownloadDir()

	if downloadPath == "" {
		log.Fatalln("Not found `Downloads` directory")
	}

	startTime := time.Now()
	install_manager.DownloadApps(downloadPath)
	endTime := time.Now().Sub(startTime).Seconds()

	fmt.Println(endTime, "Seconds")

	install_manager.OpenDownloadDir()
}
