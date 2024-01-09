package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type State struct {
	path  string
	state map[string]time.Time
}

func NewState(path string) State {
	return State{path: path}
}

func (s *State) Load() error {
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

func (s *State) Dump() error {
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
