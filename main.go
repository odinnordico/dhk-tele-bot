package main

import (
	"fmt"
	"strconv"
	"time"

	e "github.com/odinnordico/dhk-the-bot/pkg/error"
	h "github.com/odinnordico/dhk-the-bot/pkg/handler"
	s "github.com/odinnordico/dhk-the-bot/pkg/system"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

const (
	channelSize   = 512
	pollerTimeout = 10
)

func main() {
	p := s.GetenvOrDefault("PRODUCTION", "1")
	pb := s.Must(strconv.ParseBool(p))
	log := configLogger(pb)

	t := s.GetenvOrDefault("TELEGRAM_TOKEN", "")
	if t == "" {
		s.Must(t, e.TelegramTokenNotFound{})
	}

	var poller tele.Poller

	if !pb {
		poller = &tele.LongPoller{Timeout: pollerTimeout * time.Second}
	} else {
		poller = &tele.Webhook{
			Listen: ":" + s.GetenvOrDefault("PORT", "80"),
			Endpoint: &tele.WebhookEndpoint{
				PublicURL: s.GetenvOrDefault("WEBHOOK_URL", "localhost"),
			},
		}
	}

	pref := tele.Settings{
		Token:   t,
		Poller:  poller,
		Updates: channelSize,
		Verbose: !pb,
		OnError: h.OnError,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	registerHandlers(b, log)
	log.Infow("starting bot", "poller", fmt.Sprintf("%T", poller), "production", pb)
	b.Start()
}

func configLogger(p bool) (logger *zap.SugaredLogger) {
	var l *zap.Logger
	if p {
		l = s.Must(zap.NewProduction())
	} else {
		l = s.Must(zap.NewDevelopment())
	}
	logger = l.Sugar()
	return
}

func registerHandlers(b *tele.Bot, l *zap.SugaredLogger) {
	l.Info("registering handlers")
	for _, handler := range h.GetCommands(l) {
		b.Handle(handler.GetCommand(), handler.GetHandler())
	}
	b.Handle(tele.OnText, h.OnText(l))
}
