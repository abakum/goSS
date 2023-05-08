package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func s97(slide int) {
	switch deb {
	case 0, slide, -slide:
	default:
		return
	}
	cmd := exec.Command("cmd", "/c", filepath.Join(root, bat))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	ex(slide, cmd.Run())
	stdo.Printf("%02d Done", slide)
}
