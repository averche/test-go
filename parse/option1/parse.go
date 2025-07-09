package option1

import (
	"encoding/json"
	"fmt"
	"time"
)

type Observation interface {
	Envelope() ObservationEnvelope
}

type ObservationEnvelope struct {
	ID            string      `json:"id"`
	SchemaVersion string      `json:"schema_version"`
	Timestamp     time.Time   `json:"timestamp"`
	Type          string      `json:"type"`
	NodeID        string      `json:"node_id"`
	Namespace     string      `json:"namespace"`
	PluginInfo    *PluginInfo `json:"plugin_info"`
}

type KVv2Observation struct {
	ObservationEnvelope
	Data KVv2Data
}

func (o KVv2Observation) Envelope() ObservationEnvelope {
	return o.ObservationEnvelope
}

type NamespaceObservation struct {
	ObservationEnvelope
	Data NamespaceData
}

func (o NamespaceObservation) Envelope() ObservationEnvelope {
	return o.ObservationEnvelope
}

type KVv2Data struct {
	ClientID    string `json:"client_id"`
	EntityID    string `json:"entity_id"`
	Path        string `json:"path"`
	RequestID   string `json:"request_id"`
	VersionRead int    `json:"version_read"`
}

type PluginInfo struct {
	MountClass    string `json:"mount_class"`
	MountAccessor string `json:"mount_accessor"`
	MountPath     string `json:"mount_path"`
	Plugin        string `json:"plugin"`
	Version       string `json:"version"`
}

type NamespaceData struct {
	CustomMetadata map[string]any `json:"custom_metadata"`
	ID             string         `json:"id"`
	Path           string         `json:"path"`
}

func Parse(data []byte) (Observation, error) {
	var o ObservationEnvelope
	if err := json.Unmarshal(data, &o); err != nil {
		return nil, fmt.Errorf("failed to parse: %w", err)
	}

	switch o.Type {
	case "kvv2/secret/read":
		var kvv2 KVv2Observation
		if err := json.Unmarshal(data, &kvv2); err != nil {
			return nil, fmt.Errorf("failed to parse: %w", err)
		}
		return kvv2, nil

	case "namespace/create":
		var namespace NamespaceObservation
		if err := json.Unmarshal(data, &namespace); err != nil {
			return nil, fmt.Errorf("failed to parse: %w", err)
		}
		return namespace, nil

	default:
		return nil, fmt.Errorf("unrecognized type: %s", o.Type)
	}
}
