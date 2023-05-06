package main

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	sl "github.com/tebeka/selenium/log"
)

func s13(slide int) {
	wg.Add(1)
	defer wg.Done()
	switch deb {
	case 0, slide, -slide:
	default:
		return
	}
	var (
		url = conf.R13
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
	if false && deb != slide {
		err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
			return HasSuffix("Общая информация").nse(wd.FindElement(selenium.ByTagName, "body"))
		}, time.Minute*3)
		er(slide, err)
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//*[text()='Все']"))
	}, time.Minute*3)
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Все.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[@title='ЮГ']"))
	}, time.Minute*3)
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Юг.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//*[text()='Все']"))
	})
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Все2.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[@title='Ростовский']"))
	}, time.Minute*2)
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Ростовский.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//*[text()='Все']"))
	})
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Все3.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[@title='СЦ г.Миллерово']"))
	}, time.Minute*3)
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Миллерово.png", slide))
	}
	wd.KeyDown(selenium.TabKey)
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weNSE(wd.FindElement(selenium.ByXPATH, "//*[text()='Пред.Неделя']"))
	}, time.Minute*3)
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Пред.Неделя.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weNil(wd.FindElement(selenium.ByXPATH, "//*[@class='circle']"))
	}, time.Minute*3)
	er(slide, err)
	we, err := wd.FindElement(selenium.ByXPATH, "//div[@class='visualContainerHost']")
	er(slide, err)
	// ssII(wd).crop(image.Rect(10, 10, 1880, 1006)).write(fmt.Sprintf("%02d.jpg", slide))
	ssII(we).write(fmt.Sprintf("%02d.jpg", slide))
	stdo.Printf("%02d Done", slide)
}
