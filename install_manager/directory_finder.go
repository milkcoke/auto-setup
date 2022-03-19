package install_manager

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
)

var downloadDirPath string

func visitDir(path string, dir fs.DirEntry, err error) error {

	if dir.Name() != "Downloads" {
		return nil
	}
	fmt.Printf("Visited : %s\n", path)

	downloadDirPath = path
	return nil
}

func GetDownloadDir() string {
	currentUser, err := user.Current()
	homeDir := currentUser.HomeDir

	if err != nil {
		log.Fatalln(err.Error())
	}

	// Get `download` directory from current user's home directory.
	err = filepath.WalkDir(homeDir, visitDir)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return downloadDirPath
}

func OpenDownloadDir() {
	file, err := os.Open(downloadDirPath)
	if err != nil {
		log.Fatalln("Some error exists")
	}

	// Open download directory in explorer
	cmd := exec.Command(`explorer`, filepath.Join(file.Name(), "./"))

	err = cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
