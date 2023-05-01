package main

import (
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

type SendKeys string

func nse(done bool, err error) (bool, error) {
	const noSuchElement = "no such element"
	if err == nil {
		return done, nil
	}
	if strings.Contains(err.Error(), noSuchElement) {
		return done, nil
	}
	// stdo.Println(err)
	return false, err
}
func weNil(we selenium.WebElement, err error) (bool, error) {
	if err != nil {
		return nse(true, err)
	}
	return we == nil, err
}
func weMC(we selenium.WebElement, err error) (bool, error) {
	if err != nil {
		return nse(false, err)
	}
	if we == nil {
		return false, nil
	}
	err = we.MoveTo(0, 0)
	if err != nil {
		return false, err
	}
	time.Sleep(selenium.DefaultWaitInterval)
	err = we.Click() // _, err := wd.ExecuteScript("arguments[0].click()", []interface{}{we})
	if err != nil {
		return false, nil
	}
	weShow(we, err)
	return true, err
}
func (s SendKeys) mc(we selenium.WebElement, err error) (bool, error) {
	if err != nil {
		return nse(false, err)
	}
	if we == nil {
		return false, nil
	}
	err = we.MoveTo(0, 0)
	if err != nil {
		return false, err
	}
	time.Sleep(selenium.DefaultWaitInterval)
	err = we.Click() // _, err := wd.ExecuteScript("arguments[0].click()", []interface{}{we})
	if err != nil {
		return false, nil
	}
	err = we.SendKeys(string(s))
	if err != nil {
		return false, err
	}
	return true, err
}

func weNSE(we selenium.WebElement, err error) (bool, error) {
	if err != nil {
		return nse(false, err)
	}
	return we != nil, err
}

type WebDriver []selenium.WebDriver

func (w WebDriver) sf(we selenium.WebElement, err error) (bool, error) {
	if err != nil {
		return nse(false, err)
	}
	if we == nil {
		return false, nil
	}
	err = w[0].SwitchFrame(we)
	if err != nil {
		return false, err
	}
	return true, err
}

func (s SendKeys) nse(we selenium.WebElement, err error) (bool, error) {
	if err != nil {
		return nse(false, err)
	}
	if we == nil {
		return false, nil
	}
	err = we.SendKeys(string(s))
	if err != nil {
		return false, err
	}
	return true, err
}

func wesDMEC(wes []selenium.WebElement, err error) (bool, error) {
	if err != nil {
		return nse(false, err)
	}
	for _, v := range wes {
		ok, err := v.IsDisplayed()
		if !ok || err != nil {
			continue
		}
		err = v.MoveTo(0, 0)
		if err != nil {
			continue
		}
		time.Sleep(selenium.DefaultWaitInterval)
		ok, err = v.IsEnabled()
		if !ok || err != nil {
			continue
		}
		err = v.Click()
		if err != nil {
			continue
		}
		weShow(v, err)
		return true, err
	}
	return false, err
}
func (s SendKeys) dmec(wes []selenium.WebElement, err error) (bool, error) {
	if err != nil {
		return nse(false, err)
	}
	for _, v := range wes {
		ok, err := v.IsDisplayed()
		if !ok || err != nil {
			continue
		}
		err = v.MoveTo(0, 0)
		if err != nil {
			continue
		}
		time.Sleep(selenium.DefaultWaitInterval)
		ok, err = v.IsEnabled()
		if !ok || err != nil {
			continue
		}
		err = v.Click()
		if err != nil {
			continue
		}
		err = v.SendKeys(string(s))
		if err != nil {
			continue
		}
		weShow(v, err)
		return true, err
	}
	return false, err
}
func (s SendKeys) dme(wes []selenium.WebElement, err error) (bool, error) {
	if err != nil {
		return nse(false, err)
	}
	for _, v := range wes {
		ok, err := v.IsDisplayed()
		if !ok || err != nil {
			continue
		}
		err = v.MoveTo(0, 0)
		if err != nil {
			continue
		}
		time.Sleep(selenium.DefaultWaitInterval)
		ok, err = v.IsEnabled()
		if !ok || err != nil {
			continue
		}
		err = v.SendKeys(string(s))
		if err != nil {
			continue
		}
		return true, err
	}
	return false, err
}

type HasSuffix string

func (s HasSuffix) nse(we selenium.WebElement, err error) (bool, error) {
	if err != nil {
		return nse(false, err)
	}
	if we == nil {
		return false, nil
	}
	Text, _ := we.Text()
	return strings.HasSuffix(Text, string(s)), err
}

type Contains string

func (s Contains) nse(we selenium.WebElement, err error) (bool, error) {
	if err != nil {
		return nse(false, err)
	}
	if we == nil {
		return false, nil
	}
	Text, _ := we.Text()
	return strings.Contains(Text, string(s)), err
}
