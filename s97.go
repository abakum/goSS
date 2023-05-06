package main

import (
	"os"
	"os/exec"
)

func s97(slide int) {
	switch deb {
	case 0, slide, -slide:
	default:
		return
	}
	cmd := exec.Command("cmd", "/c", s2p(root, bat))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	er(slide, cmd.Run())
	stdo.Printf("%02d Done", slide)
}
