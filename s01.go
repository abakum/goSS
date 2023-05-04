package main

import (
	"fmt"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	sl "github.com/tebeka/selenium/log"
)

func s01(slide int) (ex int, err error) {
	ex = slide
	wg.Add(1)
	defer wg.Done()
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
	csWD <- wd
	if debug != slide {
		wd.ResizeWindow("", 1920, 1080)
	}
	err = wd.Get(url)
	if err != nil {
		stdo.Println()
		return
	}
	wdShow(wd, slide)
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		we, err = wd.FindElement(selenium.ByXPATH, "//table[contains(@class,'weather')]")
		return weNSE(we, err)
	})
	if err != nil {
		stdo.Println()
		return
	}
	// zr.crop(we, fmt.Sprintf("%02d.jpg", slide))
	ssII(we).write(fmt.Sprintf("%02d.jpg", slide))
	return
}
