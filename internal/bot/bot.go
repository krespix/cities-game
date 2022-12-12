package bot

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/krespix/cities-game/internal/processors/message"
)

type Bot interface {
	Start(ctx context.Context)
	Stop(ctx context.Context)
}

type bot struct {
	client *tgbotapi.BotAPI

	processor message.Processor
}

func (b *bot) Start(ctx context.Context) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.client.GetUpdatesChan(u)

	for update := range updates {
		go func(upd tgbotapi.Update) {
			msg, err := b.processor.ProcessMessage(ctx, upd)
			if err != nil {
				_ = b.sendErr(err, upd.FromChat().ID)
			}
			b.sendMessages(msg)
		}(update)
	}
}

func (b *bot) sendMessages(msgList []tgbotapi.Chattable) {
	for _, msg := range msgList {
		_, _ = b.client.Send(msg)
	}
}

func (b *bot) sendErr(err error, chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, err.Error())
	_, er := b.client.Send(msg)
	return er
}

func (b *bot) Stop(_ context.Context) {
	b.client.StopReceivingUpdates()
}

func New(token string, processor message.Processor) (Bot, error) {
	c, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &bot{
		client:    c,
		processor: processor,
	}, nil
}
