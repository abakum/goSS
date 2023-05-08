package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	sl "github.com/tebeka/selenium/log"
)

func s13(slide int) {
	switch deb {
	case 0, slide, -slide:
	default:
		return
	}
	wg.Add(1)
	defer wg.Done()
	var (
		params = conf.P[strconv.Itoa(slide)]
		sc     = conf.P["4"][1]
		rf     = conf.P["12"][2]
	)
	stdo.Println(params, sc, rf)
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
	ex(slide, err)
	if deb != slide {
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
	if false {
		err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
			return HasSuffix("Общая информация").nse(wd.FindElement(selenium.ByTagName, "body"))
		}, time.Minute*3)
		ex(slide, err)
	} else {
		wd.Wait(func(wd selenium.WebDriver) (bool, error) {
			return Contains("Все").nse(wd.FindElement(selenium.ByTagName, "body"))
		})
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//div[.='Все']"))
	})
	ex(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Все.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, fmt.Sprintf("//span[@title='%s']", params[1])))
	}, time.Minute*3)
	ex(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d %s.png", slide, params[1]))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//div[.='Все']"))
	})
	ex(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Все2.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, fmt.Sprintf("//span[@title='%s']", rf)))
	}, time.Minute*2)
	ex(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d %s.png", slide, rf))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//div[.='Все']"))
	})
	ex(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Все3.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, fmt.Sprintf("//span[@title='%s']", sc)))
	}, time.Minute*3)
	ex(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d %s.png", slide, sc))
	}
	wd.KeyDown(selenium.TabKey)
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weNSE(wd.FindElement(selenium.ByXPATH, "//*[text()='Пред.Неделя']"))
	}, time.Minute*3)
	ex(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Пред.Неделя.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weNil(wd.FindElement(selenium.ByXPATH, "//*[@class='circle']"))
	}, time.Minute*3)
	ex(slide, err)
	we, err := wd.FindElement(selenium.ByXPATH, "//div[@class='visualContainerHost']")
	ex(slide, err)
	// ssII(wd).crop(image.Rect(10, 10, 1880, 1006)).write(fmt.Sprintf("%02d.jpg", slide))
	ssII(we).write(fmt.Sprintf("%02d.jpg", slide))
	stdo.Printf("%02d Done", slide)
}
