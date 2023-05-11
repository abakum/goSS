package main

import (
	"fmt"
	"image"
	"strconv"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	sl "github.com/tebeka/selenium/log"
)

func s05(slide int) {
	switch deb {
	case 0, slide, -slide:
	default:
		return
	}
	wg.Add(1)
	defer wg.Done()
	var (
		params = conf.P[strconv.Itoa(slide)]
	)
	stdo.Println(params, sc)
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
	} else {
		cCaps.Args = append(cCaps.Args,
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
	err = getEmbed(wd, params[0])
	ex(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d get.png", slide))
	} else {
		time.Sleep(time.Second)
	}
	wdShow(wd, slide)
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return WebDriver{wd}.sf(wd.FindElement(selenium.ByXPATH, "//iframe"))
	})
	ex(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d iframe.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//div[contains(@title,'Статистика по сотрудникам')]"))
	}, time.Minute*3)
	ex(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Статистика по сотрудникам.png", slide))
	}
	err = cb(wd, "СЦ/ЦЭ", sc)
	ex(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d %s.png", slide, sc))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weNSE(wd.FindElement(selenium.ByXPATH, "//*[contains(text(),'Ср. производительность сотрудника')]"))
	}, time.Minute*3)
	ex(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Ср. производительность сотрудника.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weNil(wd.FindElement(selenium.ByXPATH, "//*[@class='circle']"))
	}, time.Minute*3)
	ex(slide, err)
	we, err := wd.FindElement(selenium.ByXPATH, "//div[@class='innerContainer']")
	ex(slide, err)
	wl, err := we.Location()
	ex(slide, err)
	ws, err := we.Size()
	ex(slide, err)
	re := ssII(wd)
	if deb == slide {
		// ssII(wd).write(fmt.Sprintf("%02d.png", slide))
		re.write(fmt.Sprintf("%02d.png", slide))
	}
	// re.crop(image.Rect(10, 10, 1888, 818)).write(fmt.Sprintf("%02d.jpg", slide))
	re.crop(image.Rect(10, 10, 53+wl.X+ws.Width, 17+wl.Y+ws.Height)).write(fmt.Sprintf("%02d.jpg", slide))
	stdo.Printf("%02d Done", slide)
}
