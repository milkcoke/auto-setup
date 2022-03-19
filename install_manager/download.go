package install_manager

import (
	"fmt"
	"github.com/monaco-io/request"
	"path/filepath"
	"sync"
)

type appInstallerInfo struct {
	name      string
	uri       string
	extension string
}

var apps = [...]appInstallerInfo{
	{
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
	}, {
		name: "go-sdk-installer.msi",
		uri:  "https://go.dev/dl/go1.18.windows-amd64.msi",
	}, {
		name: "powershell-installer.msi",
		uri:  "https://github.com/PowerShell/PowerShell/releases/download/v7.2.2/PowerShell-7.2.2-win-x64.msi",
	}, {
		name: "sourctree-installer.exe",
		uri:  "https://product-downloads.atlassian.com/software/sourcetree/windows/ga/SourceTreeSetup-3.4.8.exe",
	},
}

func DownloadApps(downloadPath string) {

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
}
