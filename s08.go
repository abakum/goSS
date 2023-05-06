package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	sl "github.com/tebeka/selenium/log"
)

func s08(slide int) {
	wg.Add(1)
	defer wg.Done()
	switch deb {
	case 0, slide, -slide:
	default:
		return
	}
	var (
		url, user, pass = conf.R08.read()
		wes             []selenium.WebElement
		TaskClosed      = "TaskClosed.xlsx"
	)
	TaskClosed = s2p(root, doc, TaskClosed)
	sCaps := selenium.Capabilities{
		"browserName": "chrome",
	}
	cCaps := chrome.Capabilities{
		Path: chromeBin,
		Args: []string{
			`window-position="0,0"`,
		},
		Prefs: map[string]interface{}{
			"download.default_directory": s2p(root, doc),
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
			"headless=new",
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
		time.Sleep(time.Second)
	}
	wdShow(wd, slide)
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return SendKeys(user).mc(wd.FindElement(selenium.ByXPATH, "//input[@id='login_form-username']"))
	})
	er(slide, err)
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return SendKeys(pass).mc(wd.FindElement(selenium.ByXPATH, "//input[@id='login_form-password']"))
	})
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d login_form-password.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[contains(.,'Войти')]"))
	})
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Войти.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[contains(.,'По работникам и типу задачи')]"))
	})
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d По работникам и типу задачи.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[contains(.,'месяцы')]"))
	})
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d месяцы.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//ul[contains(@class,'ui-selectcheckboxmenu-multiple-container')]"))
	})
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d ui-selectcheckboxmenu-multiple-container.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//li[5]/label")) // //label[contains(.,'Обработка наряда')]
	})
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Обработка наряда.png", slide))
	}
	for i := 4; i < 9; i++ {
		wes, err = wd.FindElements(selenium.ByXPATH, "//*[contains(@class,'ui-tree-toggler')]")
		er(slide, err)
		err = wes[i].MoveTo(0, 0)
		er(slide, err)
		time.Sleep(selenium.DefaultWaitInterval * 2)
		err = wes[i].Click()
		er(slide, err)
		time.Sleep(selenium.DefaultWaitInterval * 2)
		// time.Sleep(time.Millisecond * 500)
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[contains(.,'Группа инсталляций')]"))
	})
	er(slide, err)
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[contains(.,'Группа клиентского сервиса')]"))
	})
	er(slide, err)
	ssII(wd).write(fmt.Sprintf("%02d Группа клиентского сервиса.png", slide))
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[contains(.,'ОК')]"))
	})
	er(slide, err)
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[contains(.,'Отображение фильтра')]"))
	})
	er(slide, err)
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//button[@id='report_actions_form-export_report_data']/span"))
	})
	er(slide, err)
	os.Remove(TaskClosed)
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[contains(.,'EXCEL')]"))
	})
	er(slide, err)
	ssII(wd).write(fmt.Sprintf("%02d.png", slide))
	time.Sleep(time.Second * 3) //for download
	stdo.Printf("%02d Done", slide)
}
