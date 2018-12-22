package commands

import (
	"errors"
	"github.com/fsnotify/fsnotify"
	"katip/common"
	"log"
	"os"
	"os/exec"
)

func runWatchCommand() (success bool, err error) {
	_, fileError := os.Stat(common.KATIP_REPO)
	if fileError == nil {
		err = errors.New("This folder is already being watched. I will continue adding to the existing history.")
		success = false
	} else {
		initCommand := exec.Command(common.GIT, "init", "--separate-git-dir="+common.KATIP_REPO)
		err = initCommand.Start()
		success = err == nil
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func(){
		for {
			select {
				case event, ok := <- watcher.Events:
					if !ok {
						return
					}
					log.Println("Event:", event)
					if event.Op == fsnotify.Write{
						log.Println("Modified file:" + event.Name)
					}

					case err,ok := <-watcher.Errors:
						if !ok{
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

	<- done
	return 
}