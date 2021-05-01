package response

import (
	"github.com/nlopes/slack"
	"strings"
	"time"
)

func Response(rtm *slack.RTM, message *slack.MessageEvent) {

	responses := map[string]string{
		"hi":             "Hey " + rtm.GetInfo().Team.Name,
		"hello":          "Hi " + rtm.GetInfo().Team.Name,
		"what's up":      "Up are the clouds",
		"how are you?":   "I'm doing good. What about you?",
		"how's it going": "you know, as usual. You tell how's your job working for you",
		"day ?":          "Hmm...I think its " + time.Now().Weekday().String(),
		"date":           time.Now().String(),
		"all good":       "seems so..",
	}

	if value, ok := responses[strings.ToLower(message.Text)]; ok {
		rtm.SendMessage(rtm.NewOutgoingMessage(value, message.Channel))
	} else {
		rtm.SendMessage(rtm.NewOutgoingMessage("I don't know this.", message.Channel))
	}
}
