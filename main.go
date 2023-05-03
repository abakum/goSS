package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/tebeka/selenium"
	"github.com/xlab/closer"
)

type sWD selenium.WebDriver

const (
	root             = "s:"
	doc              = "doc"
	bat              = "abaku.bat"
	mov              = "abaku.mp4"
	port             = 7777
	seleniumPath     = "selenium-server.jar"
	chromeDriverPath = "chromedriver.exe"
	chromeBin        = `C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`
	userDataDir      = `Google\Chrome\User Data\Default`
)

var (
	debug int
	stdo  *log.Logger
	wg    sync.WaitGroup
	cd    string
	csWD  chan sWD
	exit  int
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
	csWD = make(chan sWD, 10)
	closer.Bind(func() {
		var cmd *exec.Cmd
		stdo.Println(service)
		// service.Stop()
		// kill := false
		// for len(csWD) > 0 {
		// 	wd := <-csWD
		// 	// stdo.Println(wd)
		// 	// wdQC(wd, 0)
		// 	if wd.SessionID() != "" {
		// 		kill = true
		// 	}
		// }
		// if kill {
		cmd = exec.Command("taskkill.exe", "/F", "/IM", "java.exe", "/T")
		stdo.Println(cmd.Path, strings.Join(cmd.Args[1:], " "))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		// }
		time.Sleep(time.Second)
		cmd = exec.Command("taskkill.exe", "/F", "/IM", chromeDriverPath, "/T")
		stdo.Println(cmd.Path, strings.Join(cmd.Args[1:], " "))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		stdo.Println("main Done", exit)
		os.Exit(exit)
	})
	conf, err = loader(s2p(cd, goSSjson))
	if err != nil {
		conf.saver()
		stdo.Println()
		return
	}
	go func() {
		ex(s01(1))
	}()
	go func() {
		ex(s04(4))
	}()
	go func() {
		ex(s05(5))
	}()
	go func() {
		ex(s08(8))
		ex(s09(9))
	}()
	go func() {
		ex(s12(12))
	}()
	go func() {
		ex(s13(13))
	}()
	if debug < 97 {
		time.Sleep(time.Second) //for wg.Add
	}
	wg.Wait()
	ex(s97(97)) //bat
	go func() {
		ex(s98(98)) //telegram
	}()
	if debug == 98 {
		time.Sleep(time.Second) //for wg.Add
	}
	wg.Wait()
	ex(s99(99)) //ss
}
