package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func run() error {
	// read CLI args
	c := Config{}
	fs := c.Flags()
	_ = fs.Parse(os.Args[1:])
	err := c.Validate()
	if err != nil {
		return err
	}

	// load state
	state := NewState(c.StatePath)
	err = state.Load()
	if err != nil {
		return fmt.Errorf("load state: %v", err)
	}
	defer func() {
		err := state.Dump()
		if err != nil {
			log.Println(err)
		}
	}()
	teams, err := ReadTeams(c.TeamsPath)
	if err != nil {
		return fmt.Errorf("read teams: %v", err)
	}

	log.Println("fetching alerts...")
	alerts, err := getNewAlerts(c)
	if err != nil {
		return fmt.Errorf("fetch alerts: %v", err)
	}

	log.Println("sending slack messages...")
	sent := 0
	for _, alert := range alerts {
		lastNotif, notified := state.Get(alert)

		// send notification if never notified about the alert before
		// or if the last notification was sent a long time ago
		notify := !notified || time.Since(lastNotif) >= c.RemindEvery
		if notify {
			log.Printf("\033[32m%s\033[0m %s", alert.ID, alert.Message)
			err := sendMessage(c, alert, teams)
			if err != nil {
				return fmt.Errorf("send slack message: %v", err)
			}
			state.Update(alert)
			sent += 1
			if c.One || sent >= c.MaxMessages {
				break
			}
		}
	}

	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
