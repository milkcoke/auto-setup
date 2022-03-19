package main

import (
	"auto-setup/directory_finder"
	"fmt"
	"github.com/monaco-io/request"
	"log"
	"path/filepath"
	"sync"
	"time"
)

type appInstallerInfo struct {
	name      string
	uri       string
	extension string
}

func main() {

	/* TODO
	- How to know the file extension in advance?
	- WTF on alcapture? : "https://advert.estsoft.com/?event=201110311523647%27,%271%27",
	*/
	downloadPath := directory_finder.GetDownloadDir()
	if downloadPath == "" {
		log.Fatalln("Not found downloads directory")
	}

	apps := [...]appInstallerInfo{{
		name: "github_installer.exe",
		uri:  "https://github.com/git-for-windows/git/releases/download/v2.35.1.windows.2/Git-2.35.1.2-64-bit.exe",
		//extension: "exe",
	}, {
		name: "jetbrainToolBox_installer.exe",
		uri:  "https://download.jetbrains.com/toolbox/jetbrains-toolbox-1.22.10970.exe",
		//extension: "exe",
	}, {
		name: "notion_installer.msi",
		uri:  "https://www.notion.so/desktop/windows/download",
		//extension: "msi",
	}, {
		name: "nodejs_installer.msi",
		uri:  "https://nodejs.org/dist/v16.14.0/node-v16.14.0-x64.msi",
		//extension: "msi",
	},
	}

	startTime := time.Now()

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(apps))

	for _, appInstaller := range apps {

		go func(app appInstallerInfo) {
			defer waitGroup.Done()

			client := request.Client{
				URL:    app.uri,
				Method: "GET",
			}
			resp := client.Send()

			fmt.Println("=======================")
			if resp.Response().StatusCode != 200 && resp.Response().StatusCode != 304 {
				fmt.Printf("Failed to donwload : %s\n", app.name)
			} else {
				fmt.Println("Success to download : ", app.name)
				resp.SaveToFile(filepath.Join(downloadPath, app.name))
			}

		}(appInstaller)

	}

	waitGroup.Wait()

	endTime := time.Now().Sub(startTime).Seconds()
	fmt.Println(endTime, "Seconds")

	directory_finder.OpenDownloadDir()

}
