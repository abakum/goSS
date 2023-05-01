package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	sl "github.com/tebeka/selenium/log"
)

func s08(slide int) (err error) {
	if debug != 0 {
		if debug == slide || -debug == slide {
		} else {
			return
		}
	}
	wg.Add(1)
	defer wg.Done()
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
	if debug == slide {
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
	if err != nil {
		stdo.Println()
		return
	}
	defer close(wd)
	if debug != slide {
		wd.ResizeWindow("", 1920, 1080)
	}
	err = wd.Get(url)
	if err != nil {
		stdo.Println()
		return
	}
	wdShow(wd)
	if debug == slide {
		saveWd(wd, fmt.Sprintf("%02d get.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return SendKeys(user).mc(wd.FindElement(selenium.ByXPATH, "//input[@id='login_form-username']"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return SendKeys(pass).mc(wd.FindElement(selenium.ByXPATH, "//input[@id='login_form-password']"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		saveWd(wd, fmt.Sprintf("%02d login_form-password.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[contains(.,'Войти')]"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		saveWd(wd, fmt.Sprintf("%02d Войти.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[contains(.,'По работникам и типу задачи')]"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		saveWd(wd, fmt.Sprintf("%02d По работникам и типу задачи.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[contains(.,'месяцы')]"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		saveWd(wd, fmt.Sprintf("%02d месяцы.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//ul[contains(@class,'ui-selectcheckboxmenu-multiple-container')]"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		saveWd(wd, fmt.Sprintf("%02d ui-selectcheckboxmenu-multiple-container.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//li[5]/label")) // //label[contains(.,'Обработка наряда')]
	})
	if err != nil {
		stdo.Println()
		return
	}
	if debug == slide {
		saveWd(wd, fmt.Sprintf("%02d Обработка наряда.png", slide))
	}
	for i := 4; i < 9; i++ {
		wes, err = wd.FindElements(selenium.ByXPATH, "//*[contains(@class,'ui-tree-toggler')]")
		if err != nil {
			stdo.Println()
			return
		}
		err = wes[i].MoveTo(0, 0)
		if err != nil {
			stdo.Println()
			return
		}
		time.Sleep(selenium.DefaultWaitInterval)
		err = wes[i].Click()
		if err != nil {
			stdo.Println()
			return
		}
		time.Sleep(selenium.DefaultWaitInterval * 2)
		// time.Sleep(time.Millisecond * 500)
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[contains(.,'Группа инсталляций')]"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[contains(.,'Группа клиентского сервиса')]"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	saveWd(wd, fmt.Sprintf("%02d Группа клиентского сервиса.png", slide))
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[contains(.,'ОК')]"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[contains(.,'Отображение фильтра')]"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//button[@id='report_actions_form-export_report_data']/span"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	os.Remove(TaskClosed)
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[contains(.,'EXCEL')]"))
	})
	if err != nil {
		stdo.Println()
		return
	}
	saveWd(wd, fmt.Sprintf("%02d.png", slide))
	time.Sleep(time.Second * 3)
	stdo.Printf("%02d Done\n", slide)
	return
}
