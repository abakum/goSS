package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func s96(slide int) (err error) {
	if debug != 0 {
		if debug == slide || -debug == slide {
		} else {
			return
		}
	}
	err = exec.Command("cmd", "/c", "start", s2p(cd, "imager.xlsb")).Run()
	if err != nil {
		stdo.Println()
		return
	}
	for _, v := range []int{2, 7, 8, 9} {
		if _, err = os.Stat(s2p(root, fmt.Sprintf("%02d.jpg", v))); errors.Is(err, os.ErrNotExist) {
			stdo.Println()
			return
		}
	}
	return
}
