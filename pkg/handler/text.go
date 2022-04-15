package handler

import (
	"strings"

	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

func OnText(l *zap.SugaredLogger) func(c tele.Context) error {
	return func(c tele.Context) error {
		t := strings.TrimSpace(c.Message().Text)
		if strings.HasPrefix(t, "!") {
			l.Debugw("Message ignored", "msg", t)
			return nil
		} else if strings.HasPrefix(t, "/") {
			return OnGeneralText(l).GetHandler()(c)
		}
		c.Message()
		return c.Send("Hello!")
	}
}
