package directory_finder

import (
	"fmt"
	"io/fs"
	"log"
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
