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
		url:  "https://github.com/git-for-windows/git/releases/download/v2.39.0.windows.2/Git-2.39.0.2-64-bit.exe",
	}, {
		name: "jetbrainToolBox_installer.exe",
		url:  "https://download.jetbrains.com/toolbox/jetbrains-toolbox-1.27.2.13801.exe",
	}, {
		name: "notion_installer.exe",
		url:  "https://www.notion.so/desktop/windows/download",
	}, {
		name: "nodejs_installer.msi",
		url:  "https://nodejs.org/dist/v18.13.0/node-v18.13.0-x64.msi",
	}, {
		name: "go-sdk-installer.msi",
		url:  "https://go.dev/dl/go1.19.5.windows-amd64.msi",
	}, {
		name: "adobe-reader-installer.exe",
		url:  "https://admdownload.adobe.com/rdcm/installers/live/readerdc64.exe",
	}, {
		name: "powershell-installer.msi",
		url:  "https://github.com/PowerShell/PowerShell/releases/download/v7.3.1/PowerShell-7.3.1-win-x64.msi",
	}, {
		name: "sourctree-installer.exe",
		url:  "https://product-downloads.atlassian.com/software/sourcetree/windows/ga/SourceTreeSetup-3.4.11.exe",
	}, {
		name: "docker-installer.exe",
		url:  "https://desktop.docker.com/win/main/amd64/Docker%20Desktop%20Installer.exe",
	}, {
		name: "alcpature-installer.exe",
		url:  "https://aldn.altools.co.kr/setup/ALCapture304.exe",
	}, {
		name: "SCoreDream_fonts.zip",
		url:  "https://s-core.co.kr/wp-content/uploads/2020/03/S-Core_Dream_OTF.zip",
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
