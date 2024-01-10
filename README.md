# opsgenie-reminder

A CLI tool that fetches from Opsgenie alerts that are open for too long and sends Slack notifications.

## Installation

Pick one of the options below.

### Download binary

1. Download a binary release from [Github Releases](https://github.com/orsinium-labs/opsgenie-reminder/releases) page.
1. Place it somewhere in your `$PATH`: `mv opsgenie-reminder ~/.local/bin/`.

### Install using Go

Install the latest version using Go:

```bash
go install github.com/orsinium-labs/opsgenie-reminder@latest
opsgenie-reminder --help
```

### Build from source

```bash
git clone --depth 1 git@github.com:orsinium-labs/opsgenie-reminder.git
cd opsgenie-reminder
go build .
./opsgenie-reminder --help
```

### Build Docker image

```bash
git clone --depth 1 git@github.com:orsinium-labs/opsgenie-reminder.git
cd opsgenie-reminder
sudo docker build --tag opsgenie-reminder .
sudo docker run opsgenie-reminder
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
