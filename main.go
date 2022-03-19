package main

import (
	"auto-setup/install_manager"
	"fmt"
	"log"
	"time"
)

func main() {

	/* TODO
	- How to know the file extension in advance?
	- WTF on alcapture? : "https://advert.estsoft.com/?event=201110311523647%27,%271%27",
	*/
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
