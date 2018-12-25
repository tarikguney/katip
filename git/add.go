package git

import (
	"github.com/tarikguney/katip/common"
	"os/exec"
)

func Add() error{
	add := exec.Command(common.GIT, "add", ".")
	return add.Start()
}
