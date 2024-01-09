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
	// https://github.com/slack-go/slack/blob/master/examples/messages/messages.go
	// https://api.slack.com/methods/chat.postMessage
	_, _, err := api.PostMessage(
		c.SlackChannel,
		slack.MsgOptionAttachments(makeAttachment(c, alert)),
	)
	return err
}

// https://api.slack.com/reference/surfaces/formatting
// https://www.bacancytechnology.com/blog/develop-slack-bot-using-golang
func makeAttachment(c Config, alert alertsv2.Alert) slack.Attachment {
	fields := make([]slack.AttachmentField, 0)

	// age
	age := int(time.Since(alert.CreatedAt).Hours())
	var ageStr string
	if age > 48 {
		ageStr = fmt.Sprintf(" %d days", age/24)
	} else {
		ageStr = fmt.Sprintf(" %d hours", age)
	}
	field := slack.AttachmentField{Title: "Age", Value: ageStr, Short: true}
	fields = append(fields, field)

	if alert.Priority != "" {
		field = slack.AttachmentField{Title: "Priority", Value: string(alert.Priority), Short: true}
		fields = append(fields, field)
	}
	if alert.Owner != "" {
		field = slack.AttachmentField{Title: "Owner", Value: alert.Owner}
		fields = append(fields, field)
	}
	if alert.Report.AcknowledgedBy != "" {
		field = slack.AttachmentField{Title: "Acknowledged by", Value: alert.Report.AcknowledgedBy}
		fields = append(fields, field)
	}

	url := fmt.Sprintf("%s/alert/detail/%s/details", c.OpsgenieURL, alert.ID)
	attachment := slack.Attachment{
		Title:     alert.Message,
		TitleLink: url,
		Pretext:   "Alert is open for too long",
		Fields:    fields,
	}
	return attachment
}
