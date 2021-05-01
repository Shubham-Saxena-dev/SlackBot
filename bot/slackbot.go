package bot

import (
	"Slackbot/response"
	"fmt"
	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
)

type Bot interface {
	StartBot()
}

type botConfig struct {
	client *slack.Client
	rtm    *slack.RTM
}

func NewBotConfiguration(client *slack.Client, rtm *slack.RTM) Bot {

	return &botConfig{
		client: client,
		rtm:    rtm,
	}
}

func (bot *botConfig) StartBot() {

	go bot.rtm.ManageConnection()

IncomingEvents:

	for {

		select {

		case msg := <-bot.rtm.IncomingEvents:
			log.Info("Receiving events")

			switch event := msg.Data.(type) {
			case *slack.ConnectedEvent:
				log.Info("Connected Event details ", event.Info, event.ConnectionCount)

			case *slack.MessageEvent:
				info := bot.rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s> ", info.User.ID)

				if event.User != prefix {
					response.Response(bot.rtm, event)
				}

			case slack.RTMError:
				log.Error("Error occurred: ", event.Error())

			case *slack.InvalidAuthEvent:
				log.Error("Invalid credentials or bot token")
				break IncomingEvents

			default:

			}
		}
	}

}
