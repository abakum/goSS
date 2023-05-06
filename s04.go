package main

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	sl "github.com/tebeka/selenium/log"
)

func s04(slide int) {
	wg.Add(1)
	defer wg.Done()
	switch deb {
	case 0, slide, -slide:
	default:
		return
	}
	var (
		url = conf.R04
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
		err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
			return HasSuffix("Целевые значения").nse(wd.FindElement(selenium.ByTagName, "body"))
		})
		er(slide, err)
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weNSE(wd.FindElement(selenium.ByXPATH, "//div[contains(@title,'Главная')]"))
	}, time.Minute*3)
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Главная.png", slide))
	}
	err = cb(wd, "СЦ", "СЦ г.Миллерово")
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d СЦ г.Миллерово.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weNil(wd.FindElement(selenium.ByXPATH, "//*[@class='circle']"))
	}, time.Minute*3)
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d circle.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weNil(wd.FindElement(selenium.ByXPATH, "//*[contains(text(),'ВЛГ')]"))
	}, time.Minute*3)
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d ВЛГ.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weNil(wd.FindElement(selenium.ByXPATH, "//*[contains(text(),'УРЛ')]"))
	}, time.Minute*3)
	er(slide, err)
	we, err := wd.FindElement(selenium.ByXPATH, "//div[@class='visualContainerHost']")
	er(slide, err)
	// ssII(wd).crop(image.Rect(0, 0, 1706, 812)).write(fmt.Sprintf("%02d.jpg", slide))
	ssII(we).write(fmt.Sprintf("%02d.jpg", slide))
	stdo.Printf("%02d Done", slide)
}

func cb(wd selenium.WebDriver, key, value string) (err error) {
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return wesDMEC(wd.FindElements(selenium.ByXPATH, fmt.Sprintf("//div[contains(@aria-label,'%s')]", key)))
	})
	er(deb, err)
	if deb > 0 {
		ssII(wd).write(fmt.Sprintf("%02d cb Click div.png", deb))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return SendKeys(value).dmec(wd.FindElements(selenium.ByXPATH, "//input[contains(@placeholder,'Поиск')]"))
	})
	er(deb, err)
	if deb > 0 {
		ssII(wd).write(fmt.Sprintf("%02d cb SendKeys.png", deb))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return wesDMEC(wd.FindElements(selenium.ByXPATH, fmt.Sprintf("//span[contains(text(),'%s')]", value)))
	})
	er(deb, err)
	if deb > 0 {
		ssII(wd).write(fmt.Sprintf("%02d cb Click span.png", deb))
	}
	wd.KeyDown(selenium.TabKey)
	return
}
