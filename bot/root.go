package bot

import (
	"fmt"
	"github.com/spf13/viper"
	tb "gopkg.in/telebot.v3"
	"time"
)

var (
	Bot *tb.Bot
)

func Start() error {
	var err error
	setting := tb.Settings{
		Token:   viper.GetString("telegram.token"),
		Updates: 100,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second, AllowedUpdates: []string{
			"message",
			"chat_member",
			"inline_query",
			"callback_query",
		}},
		OnError: func(err error, context tb.Context) {
			fmt.Printf("%+v\n", err)
		},
	}
	if viper.GetString("telegram.proxy") != "" {
		setting.URL = viper.GetString("telegram.proxy")
	}
	Bot, err = tb.NewBot(setting)
	if err != nil {
		return err
	}
	RegisterCommands()
	RegisterHandle()
	Bot.Start()
	return nil
}

func RegisterCommands() {
	_ = Bot.SetCommands([]tb.Command{
		{
			Text:        StartCmd,
			Description: "Hello🙌",
		},
		{
			Text:        AllAdCmd,
			Description: "查看所有广告",
		},
		{
			Text:        AddAdCmd,
			Description: "添加广告",
		},
		{
			Text:        DelAdCmd,
			Description: "删除广告",
		},
	})
}

func RegisterHandle() {
	Bot.Handle(StartCmd, func(c tb.Context) error {
		return c.Send("🙋你好，我是一个AI反广告机器人。二开作者是Notetoday https://www.github.com/notetoday/ai-anti-bot")
	}, PreCmdMiddleware)
	creatorOnly := Bot.Group()
	creatorOnly.Use(CreatorCmdMiddleware)
	creatorOnly.Handle(AllAdCmd, AllAd)
	creatorOnly.Handle(AddAdCmd, AddAd)
	creatorOnly.Handle(DelAdCmd, DelAd)

	groupOnly := Bot.Group()
	groupOnly.Use(PreGroupMiddleware)
	groupOnly.Handle(tb.OnText, OnTextMessage)
	groupOnly.Handle(tb.OnSticker, OnStickerMessage)
	groupOnly.Handle(tb.OnPhoto, OnPhotoMessage)

	Bot.Handle(tb.OnChatMember, OnChatMemberMessage)
}
