package git

import (
	"github.com/tarikguney/katip/common"
	"os/exec"
)

func Commit(message string) error{
	cmd := exec.Command(common.GIT, "commit", "-m", message)
	return cmd.Start()
}