package commands

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/tarikguney/katip/common"
	"github.com/tarikguney/katip/git"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func runWatchCommand() (success bool, err error) {
	_, fileError := os.Stat(common.KATIP_REPO)

	writeGitIgnore()

	if !os.IsNotExist(fileError) {
		err = errors.New("this folder is already being watched. i will continue adding to the existing history")
		success = false
	} else {
		initCommand := exec.Command(common.GIT, "init", "--separate-git-dir="+common.KATIP_REPO)
		err = initCommand.Start()
	}

	watcher, err := fsnotify.NewWatcher()

	// Remove all files under .katip-repo from watcher list.
	// todo Continously ignore .katip-repo files as you commit, otherwise it enters an infinite-loop.
	_ = filepath.Walk(".katip-repo", func(path string, info os.FileInfo, err error) error {
		if err != nil{
			return err
		}
		err = watcher.Remove(path)
		return err
	})

	_ = watcher.Remove(".katip-repo")
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
				if event.Op == fsnotify.Write {
					err := git.Add()
					time.Sleep(2 * time.Second)
					if err != nil {
						log.Fatal(err)
						return
					}
					var resultMessage = fmt.Sprintf("%v is recorded in the history.", event.Name)
					err = git.Commit(resultMessage)
					if err != nil {
						log.Fatal(err)
					}
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

func writeGitIgnore() {
	var ignoreKatipRepo = ".katip-repo/*"
	_ = ioutil.WriteFile(".gitignore", []byte(ignoreKatipRepo), 0644)
}
