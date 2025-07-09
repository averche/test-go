package parse

import (
	"testing"
)

var kvv2Data1 = []byte(`{
  "id": "7dabe4fe-b6af-436d-8b1a-b042c81d9102",
  "schema_version": "v1",
  "timestamp": "2025-06-25T22:36:37.092298929Z",
  "type": "kvv2/secret/read",
  "node_id": "data-plane-vault-0",
  "data": {
    "client_id": "d6c5fbd7-d064-6687-d446-f389dd07dde8",
    "entity_id": "d6c5fbd7-d064-6687-d446-f389dd07dde8",
    "path": "data/secret100",
    "request_id": "aee4a874-5882-b7a9-231f-d85b64288a96",
    "version_read": 1
  },
  "namespace": "my-namespace/",
  "plugin_info": {
    "mount_class": "secret",
    "mount_accessor": "kv_403f6a62",
    "mount_path": "kv-v2/",
    "plugin": "kv",
    "version": "2"
  }
}`)

var kvv2Data2 = []byte(`{
  "id": "7dabe4fe-b6af-436d-8b1a-b042c81d9102",
  "schema_version": "v1",
  "timestamp": "2025-06-25T22:36:37.092298929Z",
  "type": "kvv2/secret/read",
  "data": {
  	"node_id": "data-plane-vault-0",
  	"namespace": "my-namespace/",
    "client_id": "d6c5fbd7-d064-6687-d446-f389dd07dde8",
    "entity_id": "d6c5fbd7-d064-6687-d446-f389dd07dde8",
    "path": "data/secret100",
    "request_id": "aee4a874-5882-b7a9-231f-d85b64288a96",
    "version_read": 1,
	"plugin_info": {
		"mount_class": "secret",
		"mount_accessor": "kv_403f6a62",
		"mount_path": "kv-v2/",
		"plugin": "kv",
		"version": "2"
	}
  }
}`)

var namespaceData1 = []byte(`{
  "id": "f6a9c9c2-9428-4104-abde-ff61dae3a685",
  "schema_version": "v1",
  "timestamp": "2025-06-25T22:36:11.172451916Z",
  "type": "namespace/create",
  "node_id": "data-plane-vault-0",
  "data": {
    "custom_metadata": {},
    "id": "4sExn",
    "path": "my-namespace/"
  },
  "namespace": "my-namespace/",
  "plugin_info": null
}`)

var namespaceData2 = []byte(`{
  "id": "f6a9c9c2-9428-4104-abde-ff61dae3a685",
  "schema_version": "v1",
  "timestamp": "2025-06-25T22:36:11.172451916Z",
  "type": "namespace/create",
  "data": {
    "namespace": "my-namespace/",
    "node_id": "data-plane-vault-0",
    "custom_metadata": {},
    "id": "4sExn",
    "path": "my-namespace/"
  }
}`)

func BenchmarkParseOption1(b *testing.B) {
	for b.Loop() {
		_, err := ParseOption1(kvv2Data1)
		if err != nil {
			b.Fatalf("ParseOption1 failed: %v", err)
		}
		_, err = ParseOption1(namespaceData1)
		if err != nil {
			b.Fatalf("ParseOption1 failed: %v", err)
		}
	}
}

func BenchmarkParseOption2(b *testing.B) {
	for b.Loop() {
		_, err := ParseOption2(kvv2Data2)
		if err != nil {
			b.Fatalf("ParseOption2 failed: %v", err)
		}
		_, err = ParseOption2(namespaceData2)
		if err != nil {
			b.Fatalf("ParseOption2 failed: %v", err)
		}
	}
}
