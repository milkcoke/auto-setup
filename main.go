package main

import (
	"fmt"
	"github.com/monaco-io/request"
	"sync"
	"time"
)

// WINDOWS PATH
const (
	DOWNLOAD_PATH = "C:\\Users\\CafeAlle\\Downloads\\"
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

		fmt.Println(appInstaller.name)
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
				//resp.SaveToFile(DOWNLOAD_PATH + filepath.Join(app.name, app.extension))
				resp.SaveToFile(DOWNLOAD_PATH + app.name)
			}

		}(appInstaller)

	}

	waitGroup.Wait()

	endTime := time.Now().Sub(startTime).Seconds()
	fmt.Println(endTime, "Seconds")
}
