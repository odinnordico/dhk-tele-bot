package handler

import (
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

func GetDefaultHandler(l *zap.SugaredLogger) func(c tele.Context) error {
	return func(c tele.Context) error {
		l.Debugw("default command", "message", c.Message().Text, "full", c.Message())
		return nil
	}
}
