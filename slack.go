package main

import (
	"fmt"
	"strings"
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
	addField := func(title, value string) {
		field := slack.AttachmentField{Title: title, Value: value, Short: true}
		fields = append(fields, field)
	}

	age := int(time.Since(alert.CreatedAt).Hours())
	addField("Age", humanizeHours(age))
	if alert.Priority != "" {
		addField("Priority", string(alert.Priority))
	}
	if alert.Owner != "" {
		addField("Owner", alert.Owner)
	}
	if alert.Report.AcknowledgedBy != "" {
		addField("Acknowledged by", alert.Report.AcknowledgedBy)
	}
	if alert.Integration.Name != "" {
		addField("Integration", alert.Integration.Name)
	}
	issues := getIssues(alert)
	if len(issues) > 0 {
		addField("Issues", strings.Join(issues, ", "))
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

func getIssues(alert alertsv2.Alert) []string {
	issues := make([]string, 0)
	if !alert.IsSeen {
		issues = append(issues, "not-seen")
	}
	if !alert.Acknowledged {
		issues = append(issues, "not-acked")
	}
	if alert.Snoozed {
		issues = append(issues, "snoozed")
	}
	if alert.Owner == "" {
		issues = append(issues, "no-owner")
	}
	return issues
}

func humanizeHours(age int) string {
	if age > 48 {
		return fmt.Sprintf(" %d days", age/24)
	}
	return fmt.Sprintf(" %d hours", age)
}
