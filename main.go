package main

import (
	"Slackbot/bot"
	"github.com/magiconair/properties"
	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
)

const (
	TOKEN    = "SLACK_TOKEN"
	FILENAME = "application.properties"
)

func main() {

	log.Info("Hi, This is a slack bot demo using Golang and RTM.")

	props := properties.MustLoadFile(FILENAME, properties.UTF8)
	token, ok := props.Get(TOKEN)
	if !ok {
		panic("Token not found")
	}
	api := slack.New(token)
	rtm := api.NewRTM()
	start := bot.NewBotConfiguration(api, rtm)
	start.StartBot()
}
