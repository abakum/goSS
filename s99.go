package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	sl "github.com/tebeka/selenium/log"
)

func s99(slide int) {
	switch deb {
	case 0, slide, -slide:
	default:
		return
	}
	wg.Add(1)
	defer wg.Done()
	var (
		we     selenium.WebElement
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
			fmt.Sprintf("user-data-dir=%s", filepath.Join(os.Getenv("LOCALAPPDATA"), userDataDir)), //once
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
	err = wd.Get(params[0])
	ex(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d get.png", slide))
	} else {
		time.Sleep(time.Second * 2)
	}
	wdShow(wd, slide)
	currentURL, err := wd.CurrentURL()
	ex(slide, err)
	if params[0] != currentURL {
		err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
			return SendKeys(params[1]).mc(wd.FindElement(selenium.ByXPATH, "//input[@name='ar-user-name']"))
		})
		ex(slide, err)
		if deb == slide {
			ssII(wd).write(fmt.Sprintf("%02d ar-user-name.png", slide))
		}
		err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
			return weMC(wd.FindElement(selenium.ByXPATH, "//button[@type='submit']"))
		})
		ex(slide, err)
		if deb == slide {
			ssII(wd).write(fmt.Sprintf("%02d submit_login.png", slide))
		}
		err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
			return SendKeys(params[2]).mc(wd.FindElement(selenium.ByXPATH, "//input[@name='ar-user-password']"))
		})
		ex(slide, err)
		if deb == slide {
			ssII(wd).write(fmt.Sprintf("%02d ar-user-password.png", slide))
		}
		err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
			return weMC(wd.FindElement(selenium.ByXPATH, "//button[@type='submit']"))
		})
		ex(slide, err)
		if deb == slide {
			ssII(wd).write(fmt.Sprintf("%02d submit_password.png", slide))
		}
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByCSSSelector, ".multiBtnInner_xbp:nth-child(1)"))
	}, time.Minute*2)
	ex(slide, err)
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
			ex(slide, err)
			return
		}
	}
	if !timeout {
		err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
			return weMC(wd.FindElement(selenium.ByCSSSelector, ".align-left_-232488494:nth-child(3)")) // "//button[contains(.,'Удалить')]"
		})
		ex(slide, err)
		if deb == slide {
			ssII(wd).write(fmt.Sprintf("%02d Удалить.png", slide))
		}
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByCSSSelector, ".addFilesBtn_RvX"))
	})
	ex(slide, err)
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//button[contains(.,'Файл')]"))
	})
	ex(slide, err)
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return SendKeys(filepath.Join(root, mov)).nse(wd.FindElement(selenium.ByXPATH, "//form/input"))
	})
	ex(slide, err)
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weNSE(wd.FindElement(selenium.ByXPATH, "//*[contains(text(),'Загрузка')]"))
	})
	ex(slide, err)
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
			ex(slide, err)
			return
		}
	}
	if we != nil {
		ssII(wd).write("99 Загрузка отменена.png")
		ex(slide, fmt.Errorf("загрузка отменена"))
		return
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weNSE(wd.FindElement(selenium.ByXPATH, "//*[contains(text(),'Загрузка завершена')]"))
	})
	ex(slide, err)
	ssII(wd).write("98.png")
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByCSSSelector, ".preset-primary_-525226473 > svg"))
	})
	ex(slide, err)
	time.Sleep(time.Second * 3)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Закрыть Загрузка завершена.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByCSSSelector, ".multiBtnInner_xbp:nth-child(4)"))
	})
	ex(slide, err)
	ssII(wd).write(fmt.Sprintf("%02d.png", slide))
	stdo.Printf("%02d Done\n", slide)
}
