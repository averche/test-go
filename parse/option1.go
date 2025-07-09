package parse

import (
	"encoding/json"
	"fmt"
	"time"
)

type Observation1 interface {
	Envelope() ObservationEnvelope1
}

type ObservationEnvelope1 struct {
	ID            string       `json:"id"`
	SchemaVersion string       `json:"schema_version"`
	Timestamp     time.Time    `json:"timestamp"`
	Type          string       `json:"type"`
	NodeID        string       `json:"node_id"`
	Namespace     string       `json:"namespace"`
	PluginInfo    *PluginInfo1 `json:"plugin_info"`
}

type KVv2Observation1 struct {
	ObservationEnvelope1
	Data KVv2Data1
}

func (o KVv2Observation1) Envelope() ObservationEnvelope1 {
	return o.ObservationEnvelope1
}

type NamespaceObservation1 struct {
	ObservationEnvelope1
	Data NamespaceData1
}

func (o NamespaceObservation1) Envelope() ObservationEnvelope1 {
	return o.ObservationEnvelope1
}

type KVv2Data1 struct {
	ClientID    string `json:"client_id"`
	EntityID    string `json:"entity_id"`
	Path        string `json:"path"`
	RequestID   string `json:"request_id"`
	VersionRead int    `json:"version_read"`
}

type PluginInfo1 struct {
	MountClass    string `json:"mount_class"`
	MountAccessor string `json:"mount_accessor"`
	MountPath     string `json:"mount_path"`
	Plugin        string `json:"plugin"`
	Version       string `json:"version"`
}

type NamespaceData1 struct {
	CustomMetadata map[string]any `json:"custom_metadata"`
	ID             string         `json:"id"`
	Path           string         `json:"path"`
}

func ParseOption1(data []byte) (Observation1, error) {
	var o ObservationEnvelope1
	if err := json.Unmarshal(data, &o); err != nil {
		return nil, fmt.Errorf("failed to parse: %w", err)
	}

	switch o.Type {
	case "kvv2/secret/read":
		var kvv2 KVv2Observation1
		if err := json.Unmarshal(data, &kvv2); err != nil {
			return nil, fmt.Errorf("failed to parse: %w", err)
		}
		return kvv2, nil

	case "namespace/create":
		var namespace NamespaceObservation1
		if err := json.Unmarshal(data, &namespace); err != nil {
			return nil, fmt.Errorf("failed to parse: %w", err)
		}
		return namespace, nil

	default:
		return nil, fmt.Errorf("unrecognized type: %s", o.Type)
	}
}
