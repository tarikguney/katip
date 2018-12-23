package commands

import (
	"errors"
	"github.com/fsnotify/fsnotify"
	"github.com/tarikguney/katip/common"
	"log"
	"os"
	"os/exec"
)

func runWatchCommand() (success bool, err error) {
	_, fileError := os.Stat(common.KATIP_REPO)
	if os.IsExist(fileError) {
		err = errors.New("this folder is already being watched. i will continue adding to the existing history")
		success = false
	} else {
		initCommand := exec.Command(common.GIT, "init", "--separate-git-dir="+common.KATIP_REPO)
		err = initCommand.Start()
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("Event:", event)
				if event.Op == fsnotify.Write {
					log.Println("Modified file:" + event.Name)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}
	}()

	err = watcher.Add(".")
	if err != nil {
		log.Fatal(err)
	}

	<-done
	return
}
