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

func s99(slide int) (ex int, err error) {
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
	err = wd.Get(url)
	if err != nil {
		stdo.Println()
		return
	}
	wdShow(wd, slide)
	if debug == slide {
		saveWd(wd, fmt.Sprintf("%02d get.png", slide))
	}
	currentURL, err := wd.CurrentURL()
	if err != nil {
		stdo.Println()
		return
	}
	if url != currentURL {
		err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
			return SendKeys(user).mc(wd.FindElement(selenium.ByXPATH, "//input[@name='ar-user-name']"))
		})
		if err != nil {
			stdo.Println()
			return
		}
		if debug == slide {
			saveWd(wd, fmt.Sprintf("%02d ar-user-name.png", slide))
		}
		err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
			return weMC(wd.FindElement(selenium.ByXPATH, "//button[@type='submit']"))
		}, time.Second*30)
		if err != nil {
			stdo.Println()
			return
		}
		if debug == slide {
			saveWd(wd, fmt.Sprintf("%02d submit_login.png", slide))
		}
		err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
			return SendKeys(pass).mc(wd.FindElement(selenium.ByXPATH, "//input[@name='ar-user-password']"))
		})
		if err != nil {
			stdo.Println()
			return
		}
		if debug == slide {
			saveWd(wd, fmt.Sprintf("%02d ar-user-password.png", slide))
		}
		err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
			return weMC(wd.FindElement(selenium.ByXPATH, "//button[@type='submit']"))
		})
		if err != nil {
			stdo.Println()
			return
		}
		if debug == slide {
			saveWd(wd, fmt.Sprintf("%02d submit_password.png", slide))
		}
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByCSSSelector, ".multiBtnInner_xbp:nth-child(1)"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		saveWd(wd, fmt.Sprintf("%02d Редактировать.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByCSSSelector, ".menu-button_J9B > svg"))
	})
	timeout := false
	if err != nil {
		timeout = strings.Contains(err.Error(), "timeout after")
		if !timeout {
			stdo.Println()
			return
		}
	}
	if !timeout {
		err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
			return weMC(wd.FindElement(selenium.ByCSSSelector, ".align-left_-232488494:nth-child(3)")) // "//button[contains(.,'Удалить')]"
		})
		if err != nil {
			stdo.Println()
			return
		}
		if debug == slide {
			saveWd(wd, fmt.Sprintf("%02d Удалить.png", slide))
		}
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByCSSSelector, ".addFilesBtn_RvX"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//button[contains(.,'Файл')]"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return SendKeys(s2p(root, mov)).nse(wd.FindElement(selenium.ByXPATH, "//form/input"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weNSE(wd.FindElement(selenium.ByXPATH, "//*[contains(text(),'Загрузка')]"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	// exec.Command("taskkill","/fi","WINDOWTITLE eq Открытие").Run()
	err = closeWindowWithTitle("Открытие")
	if err != nil {
		stdo.Println(err)
	}
	if debug == slide {
		saveWd(wd, fmt.Sprintf("%02d Загрузка.png", slide))
	}
	we, err = wd.FindElement(selenium.ByXPATH, "//*[contains(text(),'Загрузка отменена')]")
	if err != nil {
		_, err = nse(false, err)
		if err != nil {
			stdo.Println()
			return
		}
	}
	if we != nil {
		saveWd(wd, "99 Загрузка отменена.png")
		err = fmt.Errorf("загрузка отменена")
		return
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weNSE(wd.FindElement(selenium.ByXPATH, "//*[contains(text(),'Загрузка завершена')]"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	saveWd(wd, "98.png")
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByCSSSelector, ".preset-primary_-525226473 > svg"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	time.Sleep(time.Second * 3)
	if debug == slide {
		saveWd(wd, fmt.Sprintf("%02d Закрыть Загрузка завершена.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByCSSSelector, ".multiBtnInner_xbp:nth-child(4)"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	saveWd(wd, fmt.Sprintf("%02d.png", slide))
	stdo.Printf("%02d Done\n", slide)
	return
}
