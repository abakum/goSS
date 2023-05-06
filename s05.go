package main

import (
	"fmt"
	"image"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	sl "github.com/tebeka/selenium/log"
)

func s05(slide int) {
	wg.Add(1)
	defer wg.Done()
	switch deb {
	case 0, slide, -slide:
	default:
		return
	}
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
	if deb == slide {
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
	er(slide, err)
	if deb != slide {
		wd.ResizeWindow("", 1920, 1080)
	}
	err = getEmbed(wd, url)
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d get.png", slide))
	} else {
		time.Sleep(time.Second)
	}
	wdShow(wd, slide)
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return WebDriver{wd}.sf(wd.FindElement(selenium.ByXPATH, "//iframe"))
	})
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d iframe.png", slide))
	}
	if deb != slide {
		err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
			return HasSuffix("Время на все работы ЦЭ").nse(wd.FindElement(selenium.ByTagName, "body"))
		}, time.Minute*3)
		er(slide, err)
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//div[contains(@title,'Статистика по сотрудникам')]"))
	})
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Статистика по сотрудникам.png", slide))
	}
	err = cb(wd, "СЦ/ЦЭ", "СЦ г.Миллерово")
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d СЦ г.Миллерово.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weNSE(wd.FindElement(selenium.ByXPATH, "//*[contains(text(),'Ср. производительность сотрудника')]"))
	}, time.Minute*3)
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Ср. производительность сотрудника.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weNil(wd.FindElement(selenium.ByXPATH, "//*[@class='circle']"))
	}, time.Minute*3)
	er(slide, err)
	we, err := wd.FindElement(selenium.ByXPATH, "//div[@class='innerContainer']")
	er(slide, err)
	wl, err := we.Location()
	er(slide, err)
	ws, err := we.Size()
	er(slide, err)
	re := ssII(wd)
	if deb == slide {
		// ssII(wd).write(fmt.Sprintf("%02d.png", slide))
		re.write(fmt.Sprintf("%02d.png", slide))
	}
	// re.crop(image.Rect(10, 10, 1888, 818)).write(fmt.Sprintf("%02d.jpg", slide))
	re.crop(image.Rect(10, 10, 53+wl.X+ws.Width, 17+wl.Y+ws.Height)).write(fmt.Sprintf("%02d.jpg", slide))
	stdo.Printf("%02d Done", slide)
}
