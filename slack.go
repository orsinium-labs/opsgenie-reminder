package main

import (
	"fmt"
	"time"

	"github.com/opsgenie/opsgenie-go-sdk/alertsv2"
	"github.com/slack-go/slack"
)

func sendMessage(c Config, alert alertsv2.Alert) error {
	if c.Dry {
		return nil
	}
	api := slack.New(c.SlackToken)
	_, _, err := api.PostMessage(
		c.SlackChannel,
		slack.MsgOptionText(makeMessage(c, alert), false),
	)
	return err
}

func makeMessage(c Config, alert alertsv2.Alert) string {
	// age
	msg := "Alert is open for"
	age := int(time.Since(alert.CreatedAt).Hours())
	if age > 48 {
		msg += fmt.Sprintf(" %d days", age/24)
	} else {
		msg += fmt.Sprintf(" %d hours", age)
	}

	// priority, title, and link
	msg += "\n"
	if alert.Priority != "" {
		msg += "[" + string(alert.Priority) + "]"
	}
	url := fmt.Sprintf("%s/alert/detail/%s/details", c.OpsgenieURL, alert.ID)
	msg += fmt.Sprintf("[%s](%s)", alert.Message, url)
	return msg
}
