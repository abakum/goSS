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

func (u uup) read() (string, string, string) {
	return u.Url, u.User, u.Pass
}

type tc struct {
	Token string
	Chat  int64
}

func (t tc) read() (string, int64) {
	return t.Token, t.Chat
}

type config struct {
	fn  string
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
	conf *config
)

func loader(fn string) (conf *config, err error) {
	obj := config{fn: fn}
	conf = &obj
	bytes, err := os.ReadFile(fn)
	if err != nil {
		stdo.Println("loader")
		return
	}
	err = json.Unmarshal(bytes, conf)
	if err != nil {
		stdo.Println("loader")
		return
	}
	stdo.Println("loader done")
	stdo.Println(obj.Ids)
	return
}

func (conf *config) saver() (err error) {
	stdo.Println(conf.Ids)
	bytes, err := json.Marshal(conf)
	if err != nil {
		stdo.Println("saver")
		return
	}
	err = os.WriteFile(conf.fn, bytes, 0644)
	if err != nil {
		stdo.Println("saver")
		return
	}
	stdo.Println("saver done")
	return
}
