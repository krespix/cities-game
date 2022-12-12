package message

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Processor interface {
	ProcessMessage(ctx context.Context, upd tgbotapi.Update) ([]tgbotapi.Chattable, error)
}

type processor struct {
}

func (p *processor) ProcessMessage(ctx context.Context, upd tgbotapi.Update) ([]tgbotapi.Chattable, error) {
	return []tgbotapi.Chattable{tgbotapi.NewMessage(upd.FromChat().ID, "hello")}, nil
}

func New() Processor {
	return &processor{}
}
