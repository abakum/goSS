package main

import (
	"fmt"
	"image"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	sl "github.com/tebeka/selenium/log"
)

func s12(slide int) {
	wg.Add(1)
	defer wg.Done()
	switch deb {
	case 0, slide, -slide:
	default:
		return
	}
	var (
		url = conf.R12
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
		selenium.SetDebug(true)
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
		err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
			return Contains("Юг").nse(wd.FindElement(selenium.ByTagName, "body"))
		}, time.Minute*3)
		er(slide, err)
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//div[@aria-label='Юг']"))
	})
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Юг.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//*[text()='Все']"))
	})
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Все.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[@title='Ростовский']"))
	})
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Ростовский.png", slide))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//*[text()='Все']"))
	}, time.Second*30)
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Все2.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//span[@title='СЦ г.Миллерово']"))
	})
	er(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Миллерово.png", slide))
	}
	wd.KeyDown(selenium.TabKey)
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weNil(wd.FindElement(selenium.ByXPATH, "//*[@class='circle']"))
	})
	er(slide, err)
	we, err := wd.FindElement(selenium.ByXPATH, "//*[@id='pvExplorationHost']/div/div/exploration/div/explore-canvas/div/div[2]/div/div[2]/div[2]/visual-container-repeat/visual-container[27]/transform/div/div[3]/div/visual-modern/div/div/div[2]/div[1]/div[3]/div/div[2]")
	er(slide, err)
	wl, err := we.Location()
	er(slide, err)
	y := wl.Y
	we, err = wd.FindElement(selenium.ByXPATH, "//*[@id='pvExplorationHost']/div/div/exploration/div/explore-canvas/div/div[2]/div/div[2]/div[2]/visual-container-repeat/visual-container[22]/transform/div/div[3]/div/visual-modern/div/div/div[2]/div[1]/div[4]")
	er(slide, err)
	wl, err = we.Location()
	er(slide, err)
	ws, err := we.Size()
	er(slide, err)
	x := wl.X + ws.Width
	we, err = wd.FindElement(selenium.ByXPATH, "//div[@class='visualContainerHost']")
	er(slide, err)
	wl, err = we.Location()
	er(slide, err)
	re := ssII(wd)
	if deb == slide {
		re.write(fmt.Sprintf("%02d.png", slide))
	}
	// re.crop(image.Rect(336, 0, 1580, 740)).write(fmt.Sprintf("%02d.jpg", slide))
	re.crop(image.Rect(wl.X, wl.Y, x, y)).write(fmt.Sprintf("%02d.jpg", slide))
	stdo.Printf("%02d Done", slide)
}
