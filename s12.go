package main

import (
	"fmt"
	"image"
	"strconv"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	sl "github.com/tebeka/selenium/log"
)

func s12(slide int) {
	switch deb {
	case 0, slide, -slide:
	default:
		return
	}
	wg.Add(1)
	defer wg.Done()
	var (
		params = conf.P[strconv.Itoa(slide)]
		sc     = conf.P["4"][1]
	)
	stdo.Println(params, sc)
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
		err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
			return Contains(params[1]).nse(wd.FindElement(selenium.ByTagName, "body"))
		}, time.Minute*3)
		ex(slide, err)
	} else {
		wd.Wait(func(wd selenium.WebDriver) (bool, error) {
			return Contains(params[1]).nse(wd.FindElement(selenium.ByTagName, "body"))
		})
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, fmt.Sprintf("//div[@aria-label='%s']", params[1])))
	})
	ex(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d %s.png", slide, params[1]))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//*[text()='Все']"))
	})
	ex(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Все.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, fmt.Sprintf("//span[@title='%s']", params[2])))
	})
	ex(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d %s.png", slide, params[2]))
	}
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, "//*[text()='Все']"))
	}, time.Second*30)
	ex(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d Все2.png", slide))
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weMC(wd.FindElement(selenium.ByXPATH, fmt.Sprintf("//span[@title='%s']", sc)))
	})
	ex(slide, err)
	if deb == slide {
		ssII(wd).write(fmt.Sprintf("%02d %s.png", slide, sc))
	}
	wd.KeyDown(selenium.TabKey)
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		return weNil(wd.FindElement(selenium.ByXPATH, "//*[@class='circle']"))
	})
	ex(slide, err)
	we, err := wd.FindElement(selenium.ByXPATH, "//*[@id='pvExplorationHost']/div/div/exploration/div/explore-canvas/div/div[2]/div/div[2]/div[2]/visual-container-repeat/visual-container[27]/transform/div/div[3]/div/visual-modern/div/div/div[2]/div[1]/div[3]/div/div[2]")
	ex(slide, err)
	wl, err := we.Location()
	ex(slide, err)
	y := wl.Y
	we, err = wd.FindElement(selenium.ByXPATH, "//*[@id='pvExplorationHost']/div/div/exploration/div/explore-canvas/div/div[2]/div/div[2]/div[2]/visual-container-repeat/visual-container[22]/transform/div/div[3]/div/visual-modern/div/div/div[2]/div[1]/div[4]")
	ex(slide, err)
	wl, err = we.Location()
	ex(slide, err)
	ws, err := we.Size()
	ex(slide, err)
	x := wl.X + ws.Width
	we, err = wd.FindElement(selenium.ByXPATH, "//div[@class='visualContainerHost']")
	ex(slide, err)
	wl, err = we.Location()
	ex(slide, err)
	re := ssII(wd)
	if deb == slide {
		re.write(fmt.Sprintf("%02d.png", slide))
	}
	// re.crop(image.Rect(336, 0, 1580, 740)).write(fmt.Sprintf("%02d.jpg", slide))
	re.crop(image.Rect(wl.X, wl.Y, x, y)).write(fmt.Sprintf("%02d.jpg", slide))
	stdo.Printf("%02d Done", slide)
}
