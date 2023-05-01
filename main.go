package main

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/tebeka/selenium"
	"github.com/xlab/closer"
)

const (
	root             = "s:"
	doc              = "doc"
	bat              = "abaku.bat"
	mov              = "abaku.mp4"
	port             = 7777
	seleniumPath     = "selenium-server.jar"
	chromeDriverPath = "chromedriver.exe "
	chromeBin        = `C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`
	userDataDir      = `Google\Chrome\User Data\Default`
)

var (
	debug int
	stdo  *log.Logger
	wg    sync.WaitGroup
	cd    string
)

func main() {
	var err error
	defer closer.Close()
	stdo = log.New(os.Stdout, "", log.Lshortfile|log.Ltime)
	cd, err = os.Getwd()
	if err != nil {
		stdo.Println(err)
		return
	}
	stdo.Println(cd)
	debug = 0
	if len(os.Args) > 1 {
		debug, err = strconv.Atoi(os.Args[1])
		if err != nil {
			debug = 0
		}
	}
	opts := []selenium.ServiceOption{
		selenium.ChromeDriver(s2p(cd, chromeDriverPath)),
	}
	service, err := selenium.NewSeleniumService(s2p(cd, seleniumPath), port, opts...)
	if err != nil {
		stdo.Println(err)
		return
	}
	defer service.Stop()
	err = loader()
	if err != nil {
		stdo.Println(err)
		saver()
		closer.Close()
	}
	go func() {
		err = s01(1)
		if err != nil {
			stdo.Println(err)
			closer.Close()
		}
	}()
	go func() {
		err = s04(4)
		if err != nil {
			stdo.Println(err)
			closer.Close()
		}
	}()
	go func() {
		err = s05(5)
		if err != nil {
			stdo.Println(err)
			closer.Close()
		}
	}()
	go func() {
		err = s08(8)
		if err != nil {
			stdo.Println(err)
			closer.Close()
		}
		err = s96(96) //imager
		if err != nil {
			stdo.Println(err)
			closer.Close()
		}
	}()
	go func() {
		err = s12(12)
		if err != nil {
			stdo.Println(err)
			closer.Close()
		}
	}()
	go func() {
		err = s13(13)
		if err != nil {
			stdo.Println(err)
			closer.Close()
		}
	}()
	time.Sleep(time.Second) //for wg.Add
	wg.Wait()
	err = s97(97) //bat
	if err != nil {
		stdo.Println(err)
		closer.Close()
	}
	go func() {
		err = s98(98) //telegram
		if err != nil {
			stdo.Println(err)
			closer.Close()
		}
	}()
	if debug == 98 {
		time.Sleep(time.Second) //for wg.Add
	}
	err = s99(99) //ss
	if err != nil {
		stdo.Println(err)
	}
	wg.Wait()
}
