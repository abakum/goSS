package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	sl "github.com/tebeka/selenium/log"
)

func s99(slide int) {
	wg.Add(1)
	defer wg.Done()
	switch deb {
	case 0, slide, -slide:
	default:
		return
	}
	var (
		url, user, pass = conf.R99.read()
		we              selenium.WebElement
	)
	sCaps := selenium.Capabilities{
		"browserName": "chrome",
	}
	cCaps := chrome.Capabilities{
		Path: chromeBin,
		Args: []string{
			`window-position="0,0"`,
			fmt.Sprintf("user-data-dir=%s", s2p(os.Getenv("LOCALAPPDATA"), userDataDir)), //once
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
	err = wd.Get(url)
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d get.png", slide))
	} else {
		time.Sleep(time.Second * 2)
	}
	wdShow(wd, slide)
	currentURL, err := wd.CurrentURL()
	er(slide, err)
	if url != currentURL {
		err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
			return SendKeys(user).mc(wd.FindElement(selenium.ByXPATH, "//input[@name='ar-user-name']"))
		})
		er(slide, err)
		if deb == slide {
			ssII(wd).write(fmt.Sprintf("%02d ar-user-name.png", slide))
		}
		err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
			return weMC(wd.FindElement(selenium.ByXPATH, "//button[@type='submit']"))
		})
		er(slide, err)
		if deb == slide {
			ssII(wd).write(fmt.Sprintf("%02d submit_login.png", slide))
		}
		err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
			return SendKeys(pass).mc(wd.FindElement(selenium.ByXPATH, "//input[@name='ar-user-password']"))
		})
		er(slide, err)
		if deb == slide {
			ssII(wd).write(fmt.Sprintf("%02d ar-user-password.png", slide))
		}
		err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
			return weMC(wd.FindElement(selenium.ByXPATH, "//button[@type='submit']"))
		})
		er(slide, err)
		if deb == slide {
			ssII(wd).write(fmt.Sprintf("%02d submit_password.png", slide))
		}
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByCSSSelector, ".multiBtnInner_xbp:nth-child(1)"))
	}, time.Minute*2)
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Редактировать.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByCSSSelector, ".menu-button_J9B > svg"))
	})
	timeout := false
	if err != nil {
		timeout = strings.Contains(err.Error(), "timeout after")
		if !timeout {
			er(slide, err)
			return
		}
	}
	if !timeout {
		err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
			return weMC(wd.FindElement(selenium.ByCSSSelector, ".align-left_-232488494:nth-child(3)")) // "//button[contains(.,'Удалить')]"
		})
		er(slide, err)
		if deb == slide {
			ssII(wd).write(fmt.Sprintf("%02d Удалить.png", slide))
		}
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByCSSSelector, ".addFilesBtn_RvX"))
	})
	er(slide, err)
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//button[contains(.,'Файл')]"))
	})
	er(slide, err)
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return SendKeys(s2p(root, mov)).nse(wd.FindElement(selenium.ByXPATH, "//form/input"))
	})
	er(slide, err)
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weNSE(wd.FindElement(selenium.ByXPATH, "//*[contains(text(),'Загрузка')]"))
	})
	er(slide, err)
	// if deb == slide {
	// 	closeWindowWithTitle("Открытие")
	// }
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Загрузка.png", slide))
	}
	we, err = wd.FindElement(selenium.ByXPATH, "//*[contains(text(),'Загрузка отменена')]")
	if err != nil {
		_, err = nse(false, err)
		if err != nil {
			er(slide, err)
			return
		}
	}
	if we != nil {
		ssII(wd).write("99 Загрузка отменена.png")
		er(slide, fmt.Errorf("загрузка отменена"))
		return
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weNSE(wd.FindElement(selenium.ByXPATH, "//*[contains(text(),'Загрузка завершена')]"))
	})
	er(slide, err)
	ssII(wd).write("98.png")
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByCSSSelector, ".preset-primary_-525226473 > svg"))
	})
	er(slide, err)
	time.Sleep(time.Second * 3)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Закрыть Загрузка завершена.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByCSSSelector, ".multiBtnInner_xbp:nth-child(4)"))
	})
	er(slide, err)
	ssII(wd).write(fmt.Sprintf("%02d.png", slide))
	stdo.Printf("%02d Done\n", slide)
}
