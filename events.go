package main

import (
	"encoding/json"
	"fmt"
)

const (
	EventTransmitStart     = '\x02' // ASCII: Start of Text
	EventTransmitEnd       = '\x03' // ASCII: End of Text
	EventTransmitSeparator = '\x1f' // ASCII: Unit Separator
)

const (
	EventNameHostKey  = "hostKey"  // new server, never seen before
	EventNameSSHStart = "sshStart" // pipe stdin/stdout/stderr to ssh from now on
)

type EventPayloadHostKey struct {
	Host            string   `json:"h"`
	Fingerprint     string   `json:"fp"`
	HostWithSameKey []string `json:"s,omitempty"`
	OldFingerprint  *string  `json:"o,omitempty"`
}

func buildEvent(name string, payload any) ([]byte, error) {
	data := []byte{EventTransmitStart}
	data = append(data, name...)
	if payload != nil {
		data = append(data, EventTransmitSeparator)
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal payload: %w", err)
		}
		data = append(data, payloadBytes...)
	}
	data = append(data, EventTransmitEnd)
	return data, nil
}
