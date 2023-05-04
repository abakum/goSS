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
	"reflect"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/tebeka/selenium"
	sl "github.com/tebeka/selenium/log"
	"github.com/xlab/closer"
)

func wdQC(wd selenium.WebDriver, slide int) {
	id := wd.SessionID()
	if id == "" {
		return
	}
	stdo.Printf("%02d wdQC %q\n", slide, id)
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

func sErr(s string, err error) string {
	if err != nil {
		return err.Error()
	}
	return s
}
func wdShow(wd selenium.WebDriver, slide int) {
	stdo.Printf("%02d %q\n", slide, sErr(wd.Title()))
	stdo.Printf("%02d %q\n", slide, sErr(url.QueryUnescape(sErr(wd.CurrentURL()))))
}

// like path.Join but better
func s2p(s ...string) string {
	ss := []string{}
	for _, v := range s {
		ss = append(ss, filepath.SplitList(strings.Trim(v, `\/`))...)
	}
	return filepath.FromSlash(path.Join(ss...))
}

func i2p(v int) (fn string) {
	fn = fmt.Sprintf("%02d.jpg", v)
	if v == 97 {
		fn = mov
	}
	fn = s2p(root, fn)
	return
}

func GetUnexportedField(field reflect.Value) interface{} {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()
}
func SetUnexportedField(field reflect.Value, value interface{}) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).
		Elem().
		Set(reflect.ValueOf(value))
}

func ex(slide int, err error) {
	if err != nil {
		exit = slide
		stdo.Printf("%02d %q", slide, err)
		closer.Close()
	} else {
		stdo.Printf("%02d Done", slide)
	}
}

type ii struct {
	img image.Image
	err error
}

func ssII(wd any) (r *ii) {
	var pngBytes []byte
	r = &ii{}
	if wd == nil {
		stdo.Println("ssII")
		r.err = fmt.Errorf("empty img")
		return
	}
	time.Sleep(time.Second)
	switch wx := wd.(type) {
	case selenium.WebDriver:
		pngBytes, r.err = wx.Screenshot()
	case selenium.WebElement:
		pngBytes, r.err = wx.Screenshot(true)
	default:
		r.err = fmt.Errorf("not selenium.WebX")
	}
	if r.err != nil {
		stdo.Println("ssII")
		return
	}
	r.img, r.err = png.Decode(bytes.NewReader(pngBytes))
	if r.err != nil {
		stdo.Println("ssII")
	}
	return
}

func (i *ii) crop(crop image.Rectangle) (r *ii) {
	r = i
	if i.err != nil {
		return
	}
	// r = &ii{}
	// r.img, r.err = cropImage(i.img, crop)
	type subImager interface {
		SubImage(ir image.Rectangle) image.Image
	}
	// img is an Image interface. This checks if the underlying value has a
	// method called SubImage. If it does, then we can use SubImage to crop the
	// image.
	sImg, ok := i.img.(subImager)
	if !ok {
		stdo.Println("crop")
		r.err = fmt.Errorf("image does not support cropping")
		return
	}
	r.img = sImg.SubImage(crop)
	return
}

func (i *ii) write(fileName string) (err error) {
	err = i.err
	if err != nil {
		return
	}
	fullName := s2p(root, doc, fileName)
	if strings.HasSuffix(fileName, ".jpg") {
		fullName = filepath.Join(root, fileName)
	}
	out, err := os.Create(fullName)
	if err != nil {
		stdo.Println("write")
		return
	}
	defer out.Close()
	if strings.HasSuffix(fileName, ".jpg") {
		err = jpeg.Encode(out, i.img, &jpeg.Options{Quality: 100})
	} else {
		err = png.Encode(out, i.img)
	}
	if err != nil {
		stdo.Println("write")
		return
	}
	// err = exec.Command("rundll32", "url.dll,FileProtocolHandler", pJPG).Run()
	// err = exec.Command("powershell", "Start-Process", "chrome", "-argumentlist", pJPG).Run()
	// err = exec.Command("cmd", "/c", "start", "chrome", pJPG).Run()
	err = exec.Command(chromeBin, fullName).Run()
	return
}
