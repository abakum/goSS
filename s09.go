package main

import (
	"errors"
	"os"
	"os/exec"
)

func s09(slide int) (ex int, err error) {
	ex = slide
	if debug != 0 {
		if debug == slide || -debug == slide || debug == -8 {
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
		if _, err = os.Stat(i2p(v)); errors.Is(err, os.ErrNotExist) {
			stdo.Println()
			return
		}
	}
	return
}
