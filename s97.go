package main

import (
	"os"
	"os/exec"
)

func s97(slide int) (ex int, err error) {
	ex = slide
	if debug != 0 {
		if debug == slide || -debug == slide {
		} else {
			return
		}
	}
	cmd := exec.Command("cmd", "/c", s2p(root, bat))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		stdo.Println()
		return
	}
	return
}
