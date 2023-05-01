package main

import (
	"fmt"
	"image"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	sl "github.com/tebeka/selenium/log"
)

func s05(slide int) (err error) {
	if debug != 0 {
		if debug == slide || -debug == slide {
		} else {
			return
		}
	}
	wg.Add(1)
	defer wg.Done()
	var (
		url = conf.R05
	)
	sCaps := selenium.Capabilities{
		"browserName": "chrome",
	}
	cCaps := chrome.Capabilities{
		Path: chromeBin,
		Args: []string{
			`window-position="0,0"`,
		},
	}
	if debug == slide {
		// selenium.SetDebug(true)
		sCaps.SetLogLevel(sl.Server, sl.All) //sl "github.com/tebeka/selenium/log"
		cCaps.Args = append(cCaps.Args,
			"start-maximized",
		)
	} else {
		cCaps.Args = append(cCaps.Args,
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
	err = getEmbed(wd, url)
	if err != nil {
		stdo.Println()
		return
	}
	wdShow(wd)
	if debug == slide {
		saveWd(wd, fmt.Sprintf("%02d get.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return WebDriver{wd}.sf(wd.FindElement(selenium.ByXPATH, "//iframe"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		saveWd(wd, fmt.Sprintf("%02d iframe.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return HasSuffix("Время на все работы ЦЭ").nse(wd.FindElement(selenium.ByTagName, "body"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		saveWd(wd, fmt.Sprintf("%02d Время на все работы ЦЭ.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//div[contains(@title,'Статистика по сотрудникам')]"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		saveWd(wd, fmt.Sprintf("%02d Статистика по сотрудникам.png", slide))
	}
	err = cb(wd, "СЦ/ЦЭ", "СЦ г.Миллерово")
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		saveWd(wd, fmt.Sprintf("%02d СЦ г.Миллерово.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weNSE(wd.FindElement(selenium.ByXPATH, "//*[contains(text(),'Ср. производительность сотрудника')]"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		saveWd(wd, fmt.Sprintf("%02d Ср. производительность сотрудника.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weNil(wd.FindElement(selenium.ByXPATH, "//*[@class='circle']"))
	}, time.Minute*3)
	if err != nil {
		stdo.Println()
		return
	}
	saveWd(wd, fmt.Sprintf("%02d.png", slide))
	saveCropWd(wd, fmt.Sprintf("%02d.jpg", slide), image.Rect(10, 10, 1888, 818)) //x7 y7 x3 y3
	stdo.Printf("%02d Done\n", slide)
	return
}
