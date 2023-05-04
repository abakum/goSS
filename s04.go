package main

import (
	"fmt"
	"image"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	sl "github.com/tebeka/selenium/log"
)

func s04(slide int) (ex int, err error) {
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
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return HasSuffix("Целевые значения").nse(wd.FindElement(selenium.ByTagName, "body"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		ssII(wd).write(fmt.Sprintf("%02d Целевые значения.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weNSE(wd.FindElement(selenium.ByXPATH, "//div[contains(@title,'Главная')]"))
	}, time.Minute*3)
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		ssII(wd).write(fmt.Sprintf("%02d Главная.png", slide))
	}
	err = cb(wd, "СЦ", "СЦ г.Миллерово")
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		ssII(wd).write(fmt.Sprintf("%02d СЦ г.Миллерово.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weNil(wd.FindElement(selenium.ByXPATH, "//*[@class='circle']"))
	}, time.Minute*3)
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		ssII(wd).write(fmt.Sprintf("%02d circle.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weNil(wd.FindElement(selenium.ByXPATH, "//*[contains(text(),'ВЛГ')]"))
	}, time.Minute*3)
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		ssII(wd).write(fmt.Sprintf("%02d ВЛГ.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weNil(wd.FindElement(selenium.ByXPATH, "//*[contains(text(),'УРЛ')]"))
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
	//saveCropWd(wd, fmt.Sprintf("%02d.jpg", slide), image.Rectangle{image.Pt(0, 0), image.Pt(1706, 812)})
	// ir(image.Rect(0, 0, 1706, 812)).crop(wd, fmt.Sprintf("%02d.jpg", slide)) //x7 y7 x3 y3
	re.crop(image.Rect(0, 0, 1706, 812)).write(fmt.Sprintf("%02d.jpg", slide))
	return
}

func cb(wd selenium.WebDriver, key, value string) (err error) {
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return wesDMEC(wd.FindElements(selenium.ByXPATH, fmt.Sprintf("//div[contains(@aria-label,'%s')]", key)))
	})
	if err != nil {
		stdo.Println()
		return
	}
	if debug > 0 {
		ssII(wd).write(fmt.Sprintf("%02d cb Click div.png", debug))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return SendKeys(value).dmec(wd.FindElements(selenium.ByXPATH, "//input[contains(@placeholder,'Поиск')]"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	if debug > 0 {
		ssII(wd).write(fmt.Sprintf("%02d cb SendKeys.png", debug))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return wesDMEC(wd.FindElements(selenium.ByXPATH, fmt.Sprintf("//span[contains(text(),'%s')]", value)))
	})
	if err != nil {
		stdo.Println()
		return
	}
	if debug > 0 {
		ssII(wd).write(fmt.Sprintf("%02d cb Click span.png", debug))
	}
	wd.KeyDown(selenium.TabKey)
	return
}
