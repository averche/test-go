package parse

import (
	"encoding/json"
	"fmt"
	"time"
)

type Observation2 interface {
	Envelope() ObservationEnvelope2
}

type ObservationEnvelope2 struct {
	ID            string    `json:"id"`
	SchemaVersion string    `json:"schema_version"`
	Timestamp     time.Time `json:"timestamp"`
	Type          string    `json:"type"`
}

type PluginSpecificData struct {
	NodeID     string      `json:"node_id"`
	Namespace  string      `json:"namespace"`
	PluginInfo PluginInfo1 `json:"plugin_info"`
}

type GenericObservation struct {
	ObservationEnvelope2
	Internal json.RawMessage `json:"data"`
}

type KVv2Observation2 struct {
	ObservationEnvelope2
	PluginData PluginSpecificData
	Data       KVv2Data2
}

func (o KVv2Observation2) Envelope() ObservationEnvelope2 {
	return o.ObservationEnvelope2
}

type NamespaceObservation2 struct {
	ObservationEnvelope2
	Data NamespaceData2
}

func (o NamespaceObservation2) Envelope() ObservationEnvelope2 {
	return o.ObservationEnvelope2
}

type KVv2Data2 struct {
	NodeID      string      `json:"node_id"`
	Namespace   string      `json:"namespace"`
	ClientID    string      `json:"client_id"`
	EntityID    string      `json:"entity_id"`
	Path        string      `json:"path"`
	RequestID   string      `json:"request_id"`
	VersionRead int         `json:"version_read"`
	PluginInfo  PluginInfo1 `json:"plugin_info"`
}

type PluginInfo2 struct {
	MountClass    string `json:"mount_class"`
	MountAccessor string `json:"mount_accessor"`
	MountPath     string `json:"mount_path"`
	Plugin        string `json:"plugin"`
	Version       string `json:"version"`
}

type NamespaceData2 struct {
	NodeID         string         `json:"node_id"`
	Namespace      string         `json:"namespace"`
	CustomMetadata map[string]any `json:"custom_metadata"`
	ID             string         `json:"id"`
	Path           string         `json:"path"`
}

func ParseOption2(data []byte) (Observation2, error) {
	var o GenericObservation
	if err := json.Unmarshal(data, &o); err != nil {
		return nil, fmt.Errorf("failed to parse: %w", err)
	}

	switch o.Type {
	case "kvv2/secret/read":
		var kvv2 KVv2Data2
		if err := json.Unmarshal(o.Internal, &kvv2); err != nil {
			return nil, fmt.Errorf("failed to parse: %w", err)
		}
		return KVv2Observation2{
			ObservationEnvelope2: o.ObservationEnvelope2,
			Data:                 kvv2,
		}, nil

	case "namespace/create":
		var namespace NamespaceData2
		if err := json.Unmarshal(data, &namespace); err != nil {
			return nil, fmt.Errorf("failed to parse: %w", err)
		}
		return NamespaceObservation2{
			ObservationEnvelope2: o.ObservationEnvelope2,
			Data:                 namespace,
		}, nil

	default:
		return nil, fmt.Errorf("unrecognized type: %s", o.Type)
	}
}
