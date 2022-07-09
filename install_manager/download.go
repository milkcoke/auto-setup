package install_manager

import (
	"fmt"
	"github.com/monaco-io/request"
	"path/filepath"
	"sync"
)

type appInstallerInfo struct {
	name string
	url  string
}

var apps = [...]appInstallerInfo{
	{
		name: "github_installer.exe",
		url:  "https://github.com/git-for-windows/git/releases/download/v2.37.0.windows.1/Git-2.37.0-64-bit.exe",
	}, {
		name: "jetbrainToolBox_installer.exe",
		url:  "https://download.jetbrains.com/toolbox/jetbrains-toolbox-1.25.12424.exe",
	}, {
		name: "notion_installer.msi",
		url:  "https://www.notion.so/desktop/windows/download",
	}, {
		name: "nodejs_installer.msi",
		url:  "https://nodejs.org/dist/v16.16.0/node-v16.16.0-x64.msi",
	}, {
		name: "go-sdk-installer.msi",
		url:  "https://go.dev/dl/go1.18.3.windows-amd64.msi",
	}, {
		name: "powershell-installer.msi",
		url:  "https://github.com/PowerShell/PowerShell/releases/download/v7.2.5/PowerShell-7.2.5-win-x64.msi",
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
				URL:    app.url,
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
