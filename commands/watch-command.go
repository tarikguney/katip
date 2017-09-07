package commands

import (
	"errors"
	"katip/common"
	"os"
	"os/exec"
)

func runWatchCommand() (success bool, err error) {
	_, fileError := os.Stat(common.KATIP_REPO)
	if fileError == nil {
		err = errors.New("This folder is already being watched. I will continue adding to the existing history.")
		success = false
		return
	}

	initCommand := exec.Command(common.GIT, "init", "--separate-git-dir="+common.KATIP_REPO)
	err = initCommand.Start()
	success = err == nil
	return
}
