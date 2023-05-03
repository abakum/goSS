package main

import (
	"errors"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func s98(slide int) (ex int, err error) {
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
		bot         *telego.Bot
		token, chat = conf.R98.read()
		file        *os.File
		medias      []telego.InputMedia
		messages    []telego.Message
	)
	inds := []int{1, 4, 5, 8, 12, 13, 97}
	for _, v := range inds {
		if _, err = os.Stat(i2p(v)); errors.Is(err, os.ErrNotExist) {
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
	medias = []telego.InputMedia{}
	for _, v := range inds {
		file, err = os.Open(i2p(v))
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
	messages, err = bot.SendMediaGroup(tu.MediaGroup(tu.ID(chat)).WithMedia(medias...))
	if len(messages) != len(medias) {
		for _, v := range messages {
			bot.DeleteMessage(&telego.DeleteMessageParams{ChatID: tu.ID(chat), MessageID: v.MessageID})
		}
		stdo.Println()
		return
	}
	for _, v := range conf.Ids {
		if v == 0 {
			continue
		}
		bot.DeleteMessage(&telego.DeleteMessageParams{ChatID: tu.ID(chat), MessageID: v})
	}
	conf.Ids = []int{}
	for _, v := range messages {
		conf.Ids = append(conf.Ids, v.MessageID)
	}
	err = conf.saver()
	if err != nil {
		stdo.Println()
		return
	}
	return
}
