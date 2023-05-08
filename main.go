package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/tebeka/selenium"
	"github.com/xlab/closer"
)

const (
	doc              = "doc"
	bat              = "abaku.bat"
	mov              = "abaku.mp4"
	port             = 7777
	seleniumPath     = "selenium-server.jar"
	taskKill         = "taskkill.exe"
	java             = "java.exe"
	chromeDriverPath = "chromedriver.exe"
	chromeBin        = `C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`
	userDataDir      = `Google\Chrome\User Data\Default`
)

var (
	deb  int
	stdo *log.Logger
	wg   sync.WaitGroup
	cd   string // s:\bin
	root string // s:
	exit int
)

func main() {
	var err error
	defer closer.Close()
	stdo = log.New(os.Stdout, "", log.Lshortfile|log.Ltime)
	cd, err = os.Getwd()
	ex(2, err)
	stdo.Println(cd)
	root = filepath.Dir(cd)
	slides := []int{}
	for _, s := range os.Args[1:] {
		i, err := strconv.Atoi(s)
		if err == nil {
			slides = append(slides, i)
		}
	}
	if len(slides) == 0 {
		slides = append(slides, 0)
	}
	opts := []selenium.ServiceOption{
		selenium.ChromeDriver(filepath.Join(cd, chromeDriverPath)),
	}
	service, err := selenium.NewSeleniumService(filepath.Join(cd, seleniumPath), port, opts...)
	ex(2, err)
	closer.Bind(func() {
		deb = 2 //exit
		var cmd *exec.Cmd
		stdo.Println(service)
		cmd = exec.Command(taskKill, "/F", "/IM", java, "/T")
		stdo.Println(cmd.Path, strings.Join(cmd.Args[1:], " "))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		time.Sleep(time.Second)
		cmd = exec.Command(taskKill, "/F", "/IM", chromeDriverPath, "/T")
		stdo.Println(cmd.Path, strings.Join(cmd.Args[1:], " "))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		stdo.Println("main Done", exit)
		os.Exit(exit)
	})
	conf, err = loader(filepath.Join(cd, goSSjson))
	if err != nil {
		conf.P = map[string][]string{}
		conf.Ids = []int{}
		// conf.saver()
		ex(2, err)
		return
	}
	for _, de := range slides {
		deb = de
		go func() {
			s01(1)
		}()
		go func() {
			s04(4)
		}()
		go func() {
			s05(5)
		}()
		go func() {
			s08(8)
			s09(9)
		}()
		go func() {
			s12(12)
		}()
		go func() {
			s13(13)
		}()
		if deb < 97 {
			time.Sleep(time.Second) //for wg.Add
		}
		wg.Wait()
		s97(97) //bat
		go func() {
			s98(98) //telegram
		}()
		if deb == 98 {
			time.Sleep(time.Second) //for wg.Add
		}
		wg.Wait()
		s99(99) //ss
	}
}
