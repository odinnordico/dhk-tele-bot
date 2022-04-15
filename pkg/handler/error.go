package handler

import (
	"fmt"

	e "github.com/odinnordico/dhk-the-bot/pkg/error"
	tele "gopkg.in/telebot.v3"
)

var OnError = func(err error, c tele.Context) {
	e.NoLoggerError(c, fmt.Errorf("error while handling event"))
}
