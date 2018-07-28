package bot

import (
	"gopkg.in/telegram-bot-api.v4"
	"ircop/tgsupport/cfg"
	"ircop/tgsupport/logger"
	"time"
	"runtime/debug"
	"net/http"
	"fmt"
	"strings"
)

func Init() error {
	b, err := tgbotapi.NewBotAPI(cfg.Config.Token)
	if nil != err {
		return err
	}
	b.Debug = cfg.Config.Debug

	logger.Debug("Authorizing on account %s", b.Self.UserName)
	_, err = b.SetWebhook(tgbotapi.NewWebhook(cfg.Config.WHUrl))
	if nil != err {
		return err
	}

	info, err := b.GetWebhookInfo()
	if nil != err {
		return err
	}

	if info.LastErrorDate != 0 {
		logger.Err("Last webhook error at %v: %s", info.LastErrorDate, info.LastErrorMessage)
	}

	updates := b.ListenForWebhook("/" + b.Token)

	for {
		startBot(updates, b)
		logger.Err("startBot failed! Sleeping 10 sec. and restarting.")
		time.Sleep(time.Second * 10)
	}
}

func startBot(updates tgbotapi.UpdatesChannel, api *tgbotapi.BotAPI) {
	go http.ListenAndServe(cfg.Config.ListenAddr, nil)

	for update := range updates {
		go processUpdate(update, api)
	}
}

func processUpdate(update tgbotapi.Update, api *tgbotapi.BotAPI) {
	defer func() {
		if rec := recover(); rec != nil {
			logger.Panic("Recovered from panic in processUpdate: %+v\n%s\n", rec, debug.Stack())
		}
	}()

	// we just forwarding EVERYTHING from user to support-chat and COPY support replys to user
	logger.Debug("New update: %+v", update)

	if update.Message != nil {
		logger.Debug("Message: %+v", update.Message)

		if update.Message.Chat != nil {
			logger.Debug("Chat: %+v", update.Message.Chat)
		}

		// msgs from support chat
		if update.Message.Chat != nil && update.Message.Chat.ID == cfg.Config.SupportChat {
			// support chat message
			if update.Message.ReplyToMessage != nil && update.Message.ReplyToMessage.ForwardFrom != nil {
				sendAnswer(update, api)
				return
			}
			return
		}

		// msg from user
		if strings.HasPrefix(update.Message.Text, "/") {
			logger.Debug("Returning")
			return
		}

		if update.Message.From != nil && update.Message.Chat != nil {
			forward := tgbotapi.NewForward(cfg.Config.SupportChat, update.Message.Chat.ID, update.Message.MessageID)
			_, err := api.Send(forward)
			if nil != err {
				logger.Err("Failed to forward message: %s", err.Error())
			}
		}
	}
}

// copy text ; images ; files ; etc.
func sendAnswer(update tgbotapi.Update, api *tgbotapi.BotAPI) {
	text := update.Message.Text

	if text == "" {
		return
	}

	msg := tgbotapi.NewMessage(int64(update.Message.ReplyToMessage.ForwardFrom.ID), text)
	msg.ParseMode = "Markdown"
	msg.DisableWebPagePreview = true

	_, err := api.Send(msg)
	if nil != err {
		logger.Err("Failed to send message to user: %s", err.Error())

		// try to notify
		m2 := tgbotapi.NewMessage(cfg.Config.SupportChat, fmt.Sprintf("Failed to send message: %s", err.Error()))
		api.Send(m2)
	}
}
