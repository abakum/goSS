package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func s98(slide int) (err error) {
	if debug != 0 {
		if debug == slide || -debug == slide {
		} else {
			return
		}
	}
	wg.Add(1)
	defer wg.Done()
	var (
		bot      *telego.Bot
		me       *telego.User
		token    = conf.R98.Token
		chat     = conf.R98.Chat
		file     *os.File
		fn       string
		medias   []telego.InputMedia
		messages []telego.Message
	)
	inds := []int{1, 4, 5, 8, 12, 13, 97}
	for _, v := range inds {
		fn = fmt.Sprintf("%02d.jpg", v)
		if v == 97 {
			fn = mov
		}
		if _, err = os.Stat(s2p(root, fn)); errors.Is(err, os.ErrNotExist) {
			stdo.Println()
			return
		}
	}
	bot, err = telego.NewBot(token, telego.WithDefaultDebugLogger())
	if err != nil {
		stdo.Println()
		return
	}
	defer bot.Close()
	me, err = bot.GetMe()
	if err != nil {
		stdo.Println()
		return
	}
	stdo.Println(me)
	medias = []telego.InputMedia{}
	for _, v := range inds {
		fn = fmt.Sprintf("%02d.jpg", v)
		if v == 97 {
			fn = mov
		}
		file, err = os.Open(s2p(root, fn))
		if err != nil {
			stdo.Println()
			return
		}
		defer file.Close()
		switch v {
		case 1:
			medias = append(medias, tu.MediaPhoto(tu.File(file)).WithCaption("⚡#умныеЭкраны"))
		case 97:
			medias = append(medias, tu.MediaVideo(tu.File(file)))
		default:
			medias = append(medias, tu.MediaPhoto(tu.File(file)))
		}
	}
	for _, v := range conf.Ids {
		if v == 0 {
			continue
		}
		bot.DeleteMessage(&telego.DeleteMessageParams{ChatID: tu.ID(chat), MessageID: v})
	}
	messages, err = bot.SendMediaGroup(tu.MediaGroup(tu.ID(chat)).WithMedia(medias...))
	if len(messages) != len(medias) {
		for _, v := range messages {
			bot.DeleteMessage(&telego.DeleteMessageParams{ChatID: tu.ID(chat), MessageID: v.MessageID})
		}
		stdo.Println()
		return
	}
	conf.Ids = []int{}
	for _, v := range messages {
		conf.Ids = append(conf.Ids, v.MessageID)
	}
	err = saver()
	if err != nil {
		stdo.Println()
		return
	}
	return
}
