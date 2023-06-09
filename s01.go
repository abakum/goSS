package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	sl "github.com/tebeka/selenium/log"
)

func s01(slide int) {
	switch deb {
	case 0, slide, -slide:
	default:
		return
	}
	wg.Add(1)
	defer wg.Done()
	var (
		we     selenium.WebElement
		params = conf.P[strconv.Itoa(slide)]
	)
	stdo.Println(params)
	sCaps := selenium.Capabilities{
		"browserName":      "chrome",
		"pageLoadStrategy": "eager",
	}
	// detach := false
	cCaps := chrome.Capabilities{
		Path: chromeBin,
		// Detach: &detach,
		Args: []string{
			`window-position="0,0"`,
		},
		Prefs: map[string]interface{}{},
	}
	if deb == slide {
		// selenium.SetDebug(true)
		sCaps.SetLogLevel(sl.Server, sl.All) //sl "github.com/tebeka/selenium/log"
		// sCaps.SetLogLevel(sl.Performance, sl.All)
		// sCaps.SetLogLevel(sl.Browser, sl.Info) //not for chrome
		// sCaps.SetLogLevel(sl.Client, sl.Info)  //not for chrome
		// sCaps.SetLogLevel(sl.Driver, sl.Info)  //not for chrome
		// cCaps.Args = append(cCaps.Args,
		// 	"start-maximized",
		// )
	} else {
		cCaps.Args = append(cCaps.Args,
			// "kiosk",
			"headless=new",
		)
	}
	sCaps.AddChrome(cCaps)
	wd, err := selenium.NewRemote(sCaps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	ex(slide, err)
	if deb == slide {
		wd.MaximizeWindow("")
	} else {
		wd.ResizeWindow("", 1920, 1080)
	}
	err = wd.Get(params[0])
	ex(slide, err)
	time.Sleep(time.Second)
	wdShow(wd, slide)
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		we, err = wd.FindElement(selenium.ByXPATH, "//table[contains(@class,'weather')]")
		return weNSE(we, err)
	})
	ex(slide, err)
	ssII(we).write(fmt.Sprintf("%02d.jpg", slide))
	stdo.Printf("%02d Done", slide)
}
