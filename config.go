package main

import (
	"errors"
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
	TeamsPath      string
	SlackToken     string
	SlackChannel   string
	MaxMessages    int
	Dry            bool
	One            bool
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
		&c.OpsgenieAPIURL, "opsgenie-api-url", "api.eu.opsgenie.com",
		"opsgenie API URL to use for sending requests",
	)
	fs.StringVar(
		&c.OpsgenieURL, "opsgenie-url", "https://eu.opsgenie.com",
		"the base opsgenie URL to use for generating web links to alerts",
	)
	fs.StringVar(
		&c.StatePath, "state-path", ".opsgenie-reminder-state.json",
		"path to the JSON file to use for storing the state between runs",
	)
	fs.StringVar(
		&c.TeamsPath, "teams-path", "",
		"path to the text file mapping opsgenie team IDs to slack groups",
	)
	fs.StringVar(
		&c.SlackToken, "slack-token", "",
		"slack API token to use for sending messages",
	)
	fs.StringVar(
		&c.SlackChannel, "slack-channel", "",
		"slack channel name to use for sending messages",
	)
	fs.IntVar(
		&c.MaxMessages, "max-messages", 10,
		"how many messages (at most) can be sent in a single run",
	)
	fs.BoolVar(
		&c.Dry, "dry", false,
		"do not send slack messages, useful for debugging",
	)
	fs.BoolVar(
		&c.One, "one", false,
		"send one slack message and exit, useful for debugging",
	)
	return fs
}

func (c Config) Validate() error {
	if c.OpsgenieToken == "" {
		return errors.New("--opsgenie-token is required")
	}
	if c.SlackToken == "" {
		return errors.New("--slack-token is required")
	}
	if c.SlackChannel == "" {
		return errors.New("--slack-channel is required")
	}
	return nil
}
