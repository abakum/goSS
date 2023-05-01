package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/url"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/tebeka/selenium"
	sl "github.com/tebeka/selenium/log"
)

func close(wd selenium.WebDriver) {
	if debug > 0 {
		for _, logType := range []sl.Type{sl.Server, sl.Performance} {
			mess, err := wd.Log(sl.Type(logType))
			if err == nil {
				for _, mes := range mess {
					left := len(mes.Message)
					if logType == sl.Performance {
						left = 64
					}
					stdo.Println(logType, mes.Timestamp, mes.Message[:left])
				}
			}
		}
	}
	wd.Quit()
	wd.Close() //for kill chromeDriver.exe
}

func saveWe(we selenium.WebElement, fileName string) (err error) {
	if we == nil {
		stdo.Println()
		return
	}
	time.Sleep(time.Second)
	pngBytes, err := we.Screenshot(true)
	if err != nil {
		stdo.Println()
		return
	}
	fullName := s2p(root, doc, fileName)
	if strings.HasSuffix(fileName, ".jpg") {
		fullName = filepath.Join(root, fileName)
	}
	out, err := os.Create(fullName)
	if err != nil {
		stdo.Println()
		return
	}
	defer out.Close()
	if strings.HasSuffix(fileName, ".jpg") {
		var img image.Image
		img, err = png.Decode(bytes.NewReader(pngBytes))
		if err != nil {
			stdo.Println()
			return
		}
		err = jpeg.Encode(out, img, &jpeg.Options{Quality: 100})
		if err != nil {
			stdo.Println()
			return
		}
	} else {
		_, err = out.Write(pngBytes)
		if err != nil {
			stdo.Println()
			return
		}
	}
	// err = exec.Command("rundll32", "url.dll,FileProtocolHandler", pJPG).Run()
	// err = exec.Command("powershell", "Start-Process", "chrome", "-argumentlist", pJPG).Run()
	// err = exec.Command("cmd", "/c", "start", "chrome", pJPG).Run()
	err = exec.Command(chromeBin, fullName).Run()
	return
}
func saveWd(wd selenium.WebDriver, fileName string) (err error) {
	if wd == nil {
		stdo.Println()
		return
	}
	time.Sleep(time.Second)
	pngBytes, err := wd.Screenshot()
	if err != nil {
		stdo.Println()
		return
	}
	fullName := s2p(root, doc, fileName)
	if strings.HasSuffix(fileName, ".jpg") {
		fullName = filepath.Join(root, fileName)
	}
	out, err := os.Create(fullName)
	if err != nil {
		stdo.Println()
		return
	}
	defer out.Close()
	if strings.HasSuffix(fileName, ".jpg") {
		var img image.Image
		img, err = png.Decode(bytes.NewReader(pngBytes))
		if err != nil {
			stdo.Println()
			return
		}
		err = jpeg.Encode(out, img, &jpeg.Options{Quality: 100})
		if err != nil {
			stdo.Println()
			return
		}
	} else {
		_, err = out.Write(pngBytes)
		if err != nil {
			stdo.Println()
			return
		}
	}
	// err = exec.Command("rundll32", "url.dll,FileProtocolHandler", pJPG).Run()
	// err = exec.Command("powershell", "Start-Process", "chrome", "-argumentlist", pJPG).Run()
	// err = exec.Command("cmd", "/c", "start", "chrome", pJPG).Run()
	err = exec.Command(chromeBin, fullName).Run()
	return
}
func cropImage(img image.Image, crop image.Rectangle) (image.Image, error) {
	type subImager interface {
		SubImage(r image.Rectangle) image.Image
	}

	// img is an Image interface. This checks if the underlying value has a
	// method called SubImage. If it does, then we can use SubImage to crop the
	// image.
	simg, ok := img.(subImager)
	if !ok {
		return nil, fmt.Errorf("image does not support cropping")
	}

	return simg.SubImage(crop), nil
}
func saveCropWd(wd selenium.WebDriver, fileName string, crop image.Rectangle) (err error) {
	if wd == nil {
		stdo.Println()
		return
	}
	time.Sleep(time.Second)
	pngBytes, err := wd.Screenshot()
	if err != nil {
		stdo.Println()
		return
	}
	fullName := s2p(root, doc, fileName)
	if strings.HasSuffix(fileName, ".jpg") {
		fullName = filepath.Join(root, fileName)
	}
	out, err := os.Create(fullName)
	if err != nil {
		stdo.Println()
		return
	}
	defer out.Close()
	var oImg, img image.Image
	oImg, err = png.Decode(bytes.NewReader(pngBytes))
	if err != nil {
		stdo.Println()
		return
	}
	img, err = cropImage(oImg, crop)
	if err != nil {
		stdo.Println()
		return
	}
	if strings.HasSuffix(fileName, ".jpg") {
		err = jpeg.Encode(out, img, &jpeg.Options{Quality: 100})
	} else {
		err = png.Encode(out, img)
	}
	if err != nil {
		stdo.Println()
		return
	}
	// err = exec.Command("rundll32", "url.dll,FileProtocolHandler", pJPG).Run()
	// err = exec.Command("powershell", "Start-Process", "chrome", "-argumentlist", pJPG).Run()
	// err = exec.Command("cmd", "/c", "start", "chrome", pJPG).Run()
	err = exec.Command(chromeBin, fullName).Run()
	return
}

func weShow(we selenium.WebElement, err error) {
	if err != nil || we == nil || debug < 1 {
		return
	}
	// stdo.Println(we.Location())
	// stdo.Println(we.Size())
	oh, oerr := we.GetAttribute("outerHTML")
	if oerr != nil {
		stdo.Println(we)
		return
	}
	pwe, perr := we.FindElement(selenium.ByXPATH, "..")
	if perr != nil {
		stdo.Println(oh)
		return
	}
	poh, _ := pwe.GetAttribute("outerHTML")
	stdo.Println(strings.Replace(poh, oh, "▶"+oh+"◀", 1))
}
func getEmbed(wd selenium.WebDriver, url string) (err error) {
	err = wd.Get(url)
	if err != nil {
		return
	}
	err = wd.Get(url + "?rs:Embed=true")
	// _, err = wd.ExecuteScript(fmt.Sprintf("window.open(%q,%q)", url2, "_self"), nil)
	return
}

var (
	procSendMessage *syscall.Proc
)

func closeWindowWithTitle(title string) (err error) {
	const WM_CLOSE = 16
	user32, err = syscall.LoadDLL("user32.dll")
	if err != nil {
		return
	}
	defer user32.Release()
	procEnumWindows, err = user32.FindProc("EnumWindows")
	if err != nil {
		return
	}
	procGetWindowTextW, err = user32.FindProc("GetWindowTextW")
	if err != nil {
		return
	}
	procSendMessage, err = user32.FindProc("SendMessageW")
	if err != nil {
		return
	}
	h, err := FindWindow(title)
	if err != nil {
		return
	}
	stdo.Println(h, err)
	err = SendMessage(h, WM_CLOSE, 0, 0)
	return
}
func SendMessage(hWnd syscall.Handle, msg uint32, wParam, lParam uintptr) (err error) {
	r1, _, e1 := procSendMessage.Call(uintptr(hWnd), uintptr(msg), wParam, lParam)
	if r1 != 0 {
		err = e1
	}
	return
}
func wdShow(wd selenium.WebDriver) {
	stdo.Println(wd.Title())
	cu, err := wd.CurrentURL()
	if err != nil {
		return
	}
	stdo.Println(url.QueryUnescape(cu))
}

// like path.Join but better
func s2p(s ...string) string {
	ss := []string{}
	for _, v := range s {
		ss = append(ss, filepath.SplitList(strings.Trim(v, `\/`))...)
	}
	return filepath.FromSlash(path.Join(ss...))
}
