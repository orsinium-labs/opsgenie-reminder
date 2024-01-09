package main

import (
	"time"

	"github.com/spf13/pflag"
)

type Config struct {
	MinAge         time.Duration
	RemindEvery    time.Duration
	Filter         string
	OpsgenieToken  string
	OpsgenieAPIURL string
	OpsgenieURL    string
	StatePath      string
	SlackToken     string
	SlackChannel   string
}

func (c *Config) Flags() *pflag.FlagSet {
	fs := pflag.NewFlagSet("opsgenie-reminder", pflag.ExitOnError)
	fs.DurationVar(
		&c.MinAge, "min-age", time.Hour*24*7,
		"notify only about alerts at least that old",
	)
	fs.DurationVar(
		&c.RemindEvery, "remind-every", time.Hour*24*7,
		"re-send notification if an alert is open that long after the previous notification",
	)
	fs.StringVar(
		&c.Filter, "filter", "",
		"opsgenie query to filter alerts",
	)
	fs.StringVar(
		&c.OpsgenieToken, "opsgenie-token", "",
		"opsgenie API token to use for sending requests",
	)
	fs.StringVar(
		&c.OpsgenieAPIURL, "opsgenie-api-url", "https://api.eu.opsgenie.com",
		"opsgenie API URL to use for sending requests",
	)
	fs.StringVar(
		&c.OpsgenieURL, "opsgenie-url", "https://eu.opsgenie.com",
		"the base opsgenie URL to use for generating web links to alerts",
	)
	fs.StringVar(
		&c.StatePath, "state-path", ".opsgenie-reminder-state.json",
		"the path to the JSON file to use for storing the state between runs",
	)
	fs.StringVar(
		&c.SlackToken, "slack-token", "",
		"slack API token to use for sending messages",
	)
	fs.StringVar(
		&c.SlackToken, "slack-channel", "",
		"slack channel name to use for sending messages",
	)
	return fs
}
