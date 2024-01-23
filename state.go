package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

// State stores for each alert the last time when notification for it has been sent.
//
// Closed alerts and alert that didn't have a notification sent are excluded
// from the state to save space.
type State struct {
	path  string
	state map[string]time.Time
}

func NewState(path string) State {
	return State{path: path}
}

func (s *State) Load() error {
	// if state does not exists, do nothing
	if _, err := os.Stat(s.path); errors.Is(err, os.ErrNotExist) {
		s.state = make(map[string]time.Time)
		return nil
	}
	raw, err := os.ReadFile(s.path)
	if err != nil {
		return fmt.Errorf("read file: %v", err)
	}
	err = json.Unmarshal(raw, &s.state)
	if err != nil {
		return fmt.Errorf("parse JSON: %v", err)
	}
	return nil
}

// Get the last time when the notification was sent for the alert.
func (s *State) Get(a alert.Alert) (time.Time, bool) {
	t, ok := s.state[a.Id]
	return t, ok
}

// Update stores the last notification sending time for the given alert.
func (s *State) Update(a alert.Alert) {
	s.state[a.Id] = time.Now()
}

// Sync removes old alerts from the state.
//
// It stores all current alerts that are present in the current state
// and nothing else. In other words, it is intersection of the old and new state.
func (s *State) Sync(current []alert.Alert) {
	newState := make(map[string]time.Time)
	for _, a := range current {
		updated, found := s.state[a.Id]
		if found {
			newState[a.Id] = updated
		}
	}
	s.state = newState
}

func (s State) Dump() error {
	// don't save the empty state
	if len(s.state) == 0 {
		return nil
	}
	raw, err := json.Marshal(s.state)
	if err != nil {
		return fmt.Errorf("serialize into JSON: %v", err)
	}
	err = os.WriteFile(s.path, raw, 0o600)
	if err != nil {
		return fmt.Errorf("write file: %v", err)
	}
	return nil
}
