package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	sl "github.com/tebeka/selenium/log"
)

func s04(slide int) {
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
	stdo.Println(params)
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
		err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
			return HasSuffix("Целевые значения").nse(wd.FindElement(selenium.ByTagName, "body"))
		})
		ex(slide, err)
	} else {
		wd.Wait(func(wd selenium.WebDriver) (bool, error) {
			return Contains("Главная").nse(wd.FindElement(selenium.ByTagName, "body"))
		})
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weNSE(wd.FindElement(selenium.ByXPATH, "//div[contains(@title,'Главная')]"))
	}, time.Minute*3)
	ex(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Главная.png", slide))
	}
	err = cb(wd, "СЦ", params[1])
	ex(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d %s.png", slide, params[1]))
	} else {
		time.Sleep(time.Second)
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weNil(wd.FindElement(selenium.ByXPATH, "//*[@class='circle']"))
	}, time.Minute*3)
	ex(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d circle.png", slide))
	}
	// err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
	// 	return weNil(wd.FindElement(selenium.ByXPATH, fmt.Sprintf("//*[contains(text(),'%s')]", params[2])))
	// }, time.Minute*3)
	// ex(slide, err)
	// if deb == slide {
	// 	ssII(wd).write(fmt.Sprintf("%02d %s.png", slide, params[2]))
	// }
	// err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
	// 	return weNil(wd.FindElement(selenium.ByXPATH, "//*[contains(text(),'УРЛ')]"))
	// }, time.Minute*3)
	// ex(slide, err)
	we, err := wd.FindElement(selenium.ByXPATH, "//div[@class='visualContainerHost']")
	ex(slide, err)
	// ssII(wd).crop(image.Rect(0, 0, 1706, 812)).write(fmt.Sprintf("%02d.jpg", slide))
	ssII(we).write(fmt.Sprintf("%02d.jpg", slide))
	stdo.Printf("%02d Done", slide)
}

func cb(wd selenium.WebDriver, key, value string) (err error) {
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return wesDMEC(wd.FindElements(selenium.ByXPATH, fmt.Sprintf("//div[contains(@aria-label,'%s')]", key)))
	})
	ex(deb, err)
	if deb > 0 {
		ssII(wd).write(fmt.Sprintf("%02d cb Click div.png", deb))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return SendKeys(value).dmec(wd.FindElements(selenium.ByXPATH, "//input[contains(@placeholder,'Поиск')]"))
	})
	ex(deb, err)
	if deb > 0 {
		ssII(wd).write(fmt.Sprintf("%02d cb SendKeys.png", deb))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return wesDMEC(wd.FindElements(selenium.ByXPATH, fmt.Sprintf("//span[contains(text(),'%s')]", value)))
	})
	ex(deb, err)
	if deb > 0 {
		ssII(wd).write(fmt.Sprintf("%02d cb Click span.png", deb))
	}
	wd.KeyDown(selenium.TabKey)
	return
}
