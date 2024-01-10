# opsgenie-reminder

A CLI tool that fetches from Opsgenie alerts that are open for too long and sends Slack notifications.

## Installation

```bash
go get github.com/orsinium-labs/opsgenie-reminder
```

## Usage

```bash
opsgenie-reminder \
    --opsgenie-token    ${OPSGENIE_TOKEN} \
    --opsgenie-api-url  https://api.eu.opsgenie.com \
    --opsgenie-url      https://${MY_ORG_NAME}.app.eu.opsgenie.com/ \
    --slack-token       ${SLACK_TOKEN} \
    --slack-channel     opsgenie-reminders
```

Run `opsgenie-reminder --help` to see the full list of options.

The tool stores the state between runs in a local file (`.opsgenie-reminder-state.json` by default). Make sure to preserve this file between runs, or you can get notified about the same alert.

You can also provide (using `--teams-path`) a mapping of opsgenie team IDs to Slack group names. Then the bot will also ping the team responsible for the alert. The mapping is a text file that looks like this:

```text
183c81ae-1904-41f0-aede-7c53ef6b16e8 devops-team
2fcf93e8-b32c-42ea-8828-61278c80bb25 backenders
```
