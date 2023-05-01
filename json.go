package main

import (
	"encoding/json"
	"os"
)

const (
	goSSjson = "goSS.json"
)

type uup struct {
	Url,
	User,
	Pass string
}
type tc struct {
	Token string
	Chat  int64
}
type config struct {
	R01 string
	R04 string
	R05 string
	R08 uup
	R12 string
	R13 string
	R98 tc
	R99 uup
	Ids []int
}

var (
	conf = config{}
)

func loader() (err error) {
	bytes, err := os.ReadFile(s2p(cd, goSSjson))
	if err != nil {
		stdo.Println()
		return
	}
	err = json.Unmarshal(bytes, &conf)
	if err != nil {
		stdo.Println()
		return
	}
	stdo.Println("loader done")
	stdo.Println(conf)
	return
}

func saver() (err error) {
	stdo.Println(conf)
	bytes, err := json.Marshal(conf)
	if err != nil {
		stdo.Println(err)
		return
	}
	out, err := os.Create(s2p(cd, goSSjson))
	if err != nil {
		stdo.Println("")
		return
	}
	defer out.Close()
	_, err = out.Write(bytes)
	if err != nil {
		stdo.Println(err)
		return
	}
	stdo.Println("saver done")
	return
}
