package handler

import (
	"net/url"

	"github.com/odinnordico/dhk-the-bot/pkg/phone"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

type HandableCommand interface {
	GetCommand() string
	GetHandler() func(c tele.Context) error
}
type Command struct {
	command string
	handler func(c tele.Context) error
}

func (c *Command) GetCommand() string {
	return "/" + c.command
}

func (c *Command) GetHandler() func(c tele.Context) error {
	return c.handler
}

func OnGeneralText(l *zap.SugaredLogger) HandableCommand {
	return &Command{
		command: "",
		handler: func(c tele.Context) error {
			l.Info("text received, returning google search")
			return c.Send("https://www.google.com/search?q=" + url.QueryEscape(c.Message().Text))
		},
	}
}

func onWACommand(l *zap.SugaredLogger) HandableCommand {
	return &Command{
		command: "wa",
		handler: func(c tele.Context) error {
			l.Info("wa command received")
			numbers := c.Args()
			if len(numbers) == 0 {
				l.Errorw("no numbers provided for whatsapp easy link", "numbers", numbers)
				return nil
			}
			for _, n := range numbers {
				if number, ok := phone.ValidatePhone(n); ok {
					l.Debugw("valid number, returning chat link", "number", number)
					err := c.Send("https://wa.me/" + number)
					if err != nil {
						return err
					}
				} else {
					l.Debugw("command wa: number not valid", "number", n)
				}
			}
			return nil
		},
	}
}

func GetCommands(l *zap.SugaredLogger) []HandableCommand {
	return []HandableCommand{onWACommand(l)}
}
