package slack

import (
	"fmt"
	"o2clock/constants/appconstant"
	"o2clock/constants/errormsg"
	"o2clock/core/slackclient"
	"o2clock/utils"
	"o2clock/utils/log"

	"github.com/nlopes/slack"
)

func SlackbotReciveMsgSetup() {
	slackclient.InitSlack()
	rtm := slackclient.CreateRTM()

	//manage the connection
	go rtm.ManageConnection()
	//incoming message
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if len(ev.BotID) == 0 {
				go IncomingMessages(ev)
			}
		}
	}
}

// incoming msg
func IncomingMessages(ev *slack.MessageEvent) {
	err := SendRequestInfomation(ev)
	if err != nil {
		log.Error.Fatalln(errormsg.ERR_SLACK, err)
	}
}

func SendRequestInfomation(ev *slack.MessageEvent) error {
	var msg string
	var isRequest = true
	switch ev.Text {
	case appconstant.SRVINFO:
		msg = utils.CurrentMemStatus()
		break
	default:
		isRequest = false
	}
	fmt.Println("Channel :", ev.Channel)
	if isRequest {
		_, _, err := slackclient.GetSlackClient().PostMessage(ev.Channel, slack.MsgOptionText(msg, true))
		if err != nil {
			log.Error.Println(errormsg.ERR_SLACK, err)
			return err
		}
	}
	return nil
}

func CommonSendMsgSlack(msg, channel string) error {
	_, _, err := slackclient.GetSlackClient().PostMessage(channel, slack.MsgOptionText(msg, true))
	if err != nil {
		log.Error.Println(errormsg.ERR_SLACK, err)
		return err
	}
	return nil
}
