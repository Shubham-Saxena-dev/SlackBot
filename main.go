package main

import (
	"Slackbot/bot"
	"fmt"
	"github.com/nlopes/slack"
	"os"
)

func main() {

	fmt.Println("Hi, This is a slack bot demo.")

	//export SLACK_TOKEN=abcd123...
	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token)
	rtm := api.NewRTM()
	start := bot.NewBotConfiguration(api, rtm)
	start.StartBot()
}
