package main

import (
	"fmt"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	sl "github.com/tebeka/selenium/log"
)

func s01(slide int) (err error) {
	if debug != 0 {
		if debug == slide || -debug == slide {
		} else {
			return
		}
	}
	var (
		url = conf.R01
		we  selenium.WebElement
	)
	wg.Add(1)
	defer wg.Done()
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
	if debug == slide {
		// selenium.SetDebug(true)
		sCaps.SetLogLevel(sl.Server, sl.All) //sl "github.com/tebeka/selenium/log"
		// sCaps.SetLogLevel(sl.Performance, sl.All)
		// sCaps.SetLogLevel(sl.Browser, sl.Info) //not for chrome
		// sCaps.SetLogLevel(sl.Client, sl.Info)  //not for chrome
		// sCaps.SetLogLevel(sl.Driver, sl.Info)  //not for chrome
		cCaps.Args = append(cCaps.Args,
			"start-maximized",
		)
	} else {
		cCaps.Args = append(cCaps.Args,
			// "start-fullscreen",
			"kiosk",
			"headless",
		)
	}
	sCaps.AddChrome(cCaps)
	wd, err := selenium.NewRemote(sCaps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		stdo.Println()
		return
	}
	defer close(wd)
	if debug != slide {
		wd.ResizeWindow("", 1920, 1080)
	}
	err = wd.Get(url)
	if err != nil {
		stdo.Println()
		return
	}
	wdShow(wd)
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		we, err = wd.FindElement(selenium.ByXPATH, "//table[contains(@class,'weather')]")
		return weNSE(we, err)
	})
	if err != nil {
		stdo.Println()
		return
	}
	saveWe(we, fmt.Sprintf("%02d.jpg", slide))
	stdo.Printf("0%d Done\n", slide)
	return
}
