package install_manager

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestFind(t *testing.T) {
	downloadDirPath = GetDownloadDir()
	if "C:\\Users\\CafeAlle\\Downloads" != downloadDirPath {
		t.Error("Not Found Download directory from HOME directory")
	} else {
		file, err := os.Open(downloadDirPath)
		if err != nil {
			t.Error("Some error exists")
		}

		// Open download directory in explorer
		cmd := exec.Command(`explorer`, filepath.Join(file.Name(), "./"))
		fmt.Println(cmd)
		err = cmd.Run()
		if err != nil {
			t.Error(err)
		}
	}
}
