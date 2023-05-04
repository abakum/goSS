package main

import (
	"fmt"
	"image"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	sl "github.com/tebeka/selenium/log"
)

func s13(slide int) (ex int, err error) {
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
	csWD <- wd
	if debug != slide {
		wd.ResizeWindow("", 1920, 1080)
	}
	err = getEmbed(wd, url)
	if err != nil {
		stdo.Println()
		return
	}
	wdShow(wd, slide)
	if debug == slide {
		ssII(wd).write(fmt.Sprintf("%02d get.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return WebDriver{wd}.sf(wd.FindElement(selenium.ByXPATH, "//iframe"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		ssII(wd).write(fmt.Sprintf("%02d iframe.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return HasSuffix("Общая информация").nse(wd.FindElement(selenium.ByTagName, "body"))
	}, time.Minute*3)
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		ssII(wd).write(fmt.Sprintf("%02d Общая информация.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//*[text()='Все']"))
	}, time.Second*30)
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		ssII(wd).write(fmt.Sprintf("%02d Все.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[@title='ЮГ']"))
	}, time.Minute*3)
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		ssII(wd).write(fmt.Sprintf("%02d Юг.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//*[text()='Все']"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		ssII(wd).write(fmt.Sprintf("%02d Все2.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[@title='Ростовский']"))
	}, time.Minute*2)
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		ssII(wd).write(fmt.Sprintf("%02d Ростовский.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//*[text()='Все']"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		ssII(wd).write(fmt.Sprintf("%02d Все3.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[@title='СЦ г.Миллерово']"))
	}, time.Minute*3)
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		ssII(wd).write(fmt.Sprintf("%02d Миллерово.png", slide))
	}
	wd.KeyDown(selenium.TabKey)
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weNSE(wd.FindElement(selenium.ByXPATH, "//*[text()='Пред.Неделя']"))
	}, time.Minute*3)
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		ssII(wd).write(fmt.Sprintf("%02d Пред.Неделя.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weNil(wd.FindElement(selenium.ByXPATH, "//*[@class='circle']"))
	}, time.Minute*3)
	if err != nil {
		stdo.Println()
		return
	}
	re := ssII(wd)
	if debug == slide {
		// ssII(wd).write(fmt.Sprintf("%02d.png", slide))
		re.write(fmt.Sprintf("%02d.png", slide))
	}
	// saveCropWX(wd, fmt.Sprintf("%02d.jpg", slide), image.Rect(10, 10, 1880, 1006))
	re.crop(image.Rect(10, 10, 1880, 1006)).write(fmt.Sprintf("%02d.jpg", slide))
	return
}
