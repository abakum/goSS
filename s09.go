package main

import (
	"errors"
	"os"
	"os/exec"
)

func s09(slide int) {
	switch deb {
	case 0, 8, -8, slide, -slide:
	default:
		return
	}
	err := exec.Command("cmd", "/c", "start", s2p(cd, "imager.xlsb")).Run()
	er(slide, err)
	for _, v := range []int{2, 7, 8, 9} {
		_, err = os.Stat(i2p(v))
		if errors.Is(err, os.ErrNotExist) {
			er(slide, err)
			return
		}
	}
	stdo.Printf("%02d Done", slide)
}
