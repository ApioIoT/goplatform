package goplatform_test

import (
	"encoding/json"
	"testing"

	"github.com/ApioIoT/goplatform"
)

func TestCommand(t *testing.T) {
	t.Run("Ack", func(t *testing.T) {
		ack := goplatform.CommandAck{
			Uuid:   "1234",
			Status: goplatform.CMD_STATUS_COMPLETED,
			Extras: map[string]any{
				"extraField": "hello",
				"count":      42,
			},
		}

		b, err := json.Marshal(ack)
		if err != nil {
			t.Fatal(err)
		}

		var postAck goplatform.CommandAck
		if err := json.Unmarshal(b, &postAck); err != nil {
			t.Fatal(err)
		}

		if postAck.Uuid != "1234" {
			t.Fatalf("expected uuid to be '1234', got '%s'", postAck.Uuid)
		}
		if postAck.Status != goplatform.CMD_STATUS_COMPLETED {
			t.Fatalf("expected status to be 'completed', got '%s'", postAck.Status)
		}
		if postAck.Extras["extraField"] != "hello" {
			t.Fatalf("expected extraField to be 'hello', got '%v'", postAck.Extras["extraField"])
		}
		if postAck.Extras["count"] != float64(42) {
			t.Fatalf("expected count to be 42, got '%v'", postAck.Extras["count"])
		}
	})

	t.Run("MakeCommand", func(t *testing.T) {
		deviceId := "my-device-id"
		downlinkMaxRetris := 120
		executionMaxRetris := 10

		req := goplatform.CommandRequest{
			Name:      "Sample",
			ProjectId: "my-project-id",
			DeviceId:  &deviceId,
			Parameters: goplatform.CommandParameters{
				goplatform.CommandParameter{
					"address": 1,
				},
			},
			DownlinkRetry: &goplatform.CommandRequestRetryOption{
				MaxRetries: &downlinkMaxRetris,
			},
			ExecutionRetry: &goplatform.CommandRequestRetryOption{
				MaxRetries: &executionMaxRetris,
			},
		}

		cmd := req.MakeCommand(goplatform.CMD_STATUS_RECEIVED)

		// Verifica che l'UUID sia stato generato
		if cmd.Uuid == "" {
			t.Fatal("expected uuid to be generated")
		}

		// Verifica i campi copiati dalla request
		if cmd.Name != req.Name {
			t.Fatalf("expected name to be '%s', got '%s'", req.Name, cmd.Name)
		}
		if cmd.ProjectId != req.ProjectId {
			t.Fatalf("expected projectId to be '%s', got '%s'", req.ProjectId, cmd.ProjectId)
		}
		if *cmd.DeviceId != *req.DeviceId {
			t.Fatalf("expected deviceId to be '%s', got '%s'", *req.DeviceId, *cmd.DeviceId)
		}

		// Verifica i parametri
		if len(cmd.Parameters) != 1 {
			t.Fatalf("expected 1 parameter, got %d", len(cmd.Parameters))
		}
		if cmd.Parameters[0]["address"] != 1 {
			t.Fatalf("expected address to be 1, got %v", cmd.Parameters[0]["address"])
		}

		// Verifica le opzioni di retry
		if *cmd.DownlinkRetry.MaxRetries != 120 {
			t.Fatalf("expected downlink maxRetries to be 120, got %d", *cmd.DownlinkRetry.MaxRetries)
		}
		if *cmd.ExecutionRetry.MaxRetries != 10 {
			t.Fatalf("expected execution maxRetries to be 10, got %d", *cmd.ExecutionRetry.MaxRetries)
		}

		// Verifica lo stato e i timestamp
		if cmd.Status != goplatform.CMD_STATUS_RECEIVED {
			t.Fatalf("expected status to be '%s', got '%s'", goplatform.CMD_STATUS_RECEIVED, cmd.Status)
		}
		if cmd.CreatedAt == nil {
			t.Fatal("expected createdAt to be set")
		}
		if cmd.ReceivedAt == nil {
			t.Fatal("expected receivedAt to be set")
		}
		if cmd.CompletedAt != nil {
			t.Fatal("expected completedAt to be nil")
		}
		if cmd.FailedAt != nil {
			t.Fatal("expected failedAt to be nil")
		}
	})

	t.Run("Command with pending status", func(t *testing.T) {
		req := goplatform.CommandRequest{
			Name:      "PendingCmd",
			ProjectId: "project-1",
			Parameters: goplatform.CommandParameters{
				goplatform.CommandParameter{"action": "start"},
			},
		}

		cmd := req.MakeCommand(goplatform.CMD_STATUS_PENDING)

		if cmd.Status != goplatform.CMD_STATUS_PENDING {
			t.Fatalf("expected status to be '%s', got '%s'", goplatform.CMD_STATUS_PENDING, cmd.Status)
		}
		if cmd.CreatedAt == nil {
			t.Fatal("expected createdAt to be set")
		}
		if cmd.ReceivedAt != nil {
			t.Fatal("expected receivedAt to be nil for pending status")
		}
		if cmd.CompletedAt != nil {
			t.Fatal("expected completedAt to be nil")
		}
		if cmd.FailedAt != nil {
			t.Fatal("expected failedAt to be nil")
		}
	})

	t.Run("Command with completed status", func(t *testing.T) {
		req := goplatform.CommandRequest{
			Name:      "CompleteCmd",
			ProjectId: "project-1",
			Parameters: goplatform.CommandParameters{
				goplatform.CommandParameter{"action": "stop"},
			},
		}

		cmd := req.MakeCommand(goplatform.CMD_STATUS_COMPLETED)

		if cmd.Status != goplatform.CMD_STATUS_COMPLETED {
			t.Fatalf("expected status to be '%s', got '%s'", goplatform.CMD_STATUS_COMPLETED, cmd.Status)
		}
		if cmd.CreatedAt == nil {
			t.Fatal("expected createdAt to be set")
		}
		if cmd.CompletedAt == nil {
			t.Fatal("expected completedAt to be set for completed status")
		}
		if cmd.FailedAt != nil {
			t.Fatal("expected failedAt to be nil")
		}
	})

	t.Run("Command with failed status", func(t *testing.T) {
		req := goplatform.CommandRequest{
			Name:      "FailedCmd",
			ProjectId: "project-1",
			Parameters: goplatform.CommandParameters{
				goplatform.CommandParameter{"action": "reset"},
			},
		}

		cmd := req.MakeCommand(goplatform.CMD_STATUS_FAILED)

		if cmd.Status != goplatform.CMD_STATUS_FAILED {
			t.Fatalf("expected status to be '%s', got '%s'", goplatform.CMD_STATUS_FAILED, cmd.Status)
		}
		if cmd.CreatedAt == nil {
			t.Fatal("expected createdAt to be set")
		}
		if cmd.FailedAt == nil {
			t.Fatal("expected failedAt to be set for failed status")
		}
		if cmd.CompletedAt != nil {
			t.Fatal("expected completedAt to be nil")
		}
	})

	t.Run("Command with metadata", func(t *testing.T) {
		metadata := map[string]any{
			"source":  "test",
			"version": 1,
			"tags":    []string{"test", "automation"},
		}

		req := goplatform.CommandRequest{
			Name:      "MetadataCmd",
			ProjectId: "project-1",
			Parameters: goplatform.CommandParameters{
				goplatform.CommandParameter{"action": "test"},
			},
			Metadata: metadata,
		}

		cmd := req.MakeCommand(goplatform.CMD_STATUS_RECEIVED)

		if cmd.Metadata == nil {
			t.Fatal("expected metadata to be set")
		}
		if cmd.Metadata["source"] != "test" {
			t.Fatalf("expected metadata source to be 'test', got '%v'", cmd.Metadata["source"])
		}
		if cmd.Metadata["version"] != 1 {
			t.Fatalf("expected metadata version to be 1, got '%v'", cmd.Metadata["version"])
		}
		if tags, ok := cmd.Metadata["tags"].([]string); !ok || len(tags) != 2 || tags[0] != "test" || tags[1] != "automation" {
			t.Fatalf("expected metadata tags to be ['test', 'automation'], got '%v'", cmd.Metadata["tags"])
		}
	})

	t.Run("Command with multiple parameters", func(t *testing.T) {
		req := goplatform.CommandRequest{
			Name:      "MultiParamCmd",
			ProjectId: "project-1",
			Parameters: goplatform.CommandParameters{
				goplatform.CommandParameter{
					"address": 1,
					"value":   42,
				},
				goplatform.CommandParameter{
					"address": 2,
					"value":   84,
				},
			},
		}

		cmd := req.MakeCommand(goplatform.CMD_STATUS_RECEIVED)

		if len(cmd.Parameters) != 2 {
			t.Fatalf("expected 2 parameters, got %d", len(cmd.Parameters))
		}

		if cmd.Parameters[0]["address"] != 1 || cmd.Parameters[0]["value"] != 42 {
			t.Fatalf("expected first parameter to be {address: 1, value: 42}, got %v", cmd.Parameters[0])
		}

		if cmd.Parameters[1]["address"] != 2 || cmd.Parameters[1]["value"] != 84 {
			t.Fatalf("expected second parameter to be {address: 2, value: 84}, got %v", cmd.Parameters[1])
		}
	})

	t.Run("CommandAck Marshal/Unmarshal with different extras types", func(t *testing.T) {
		ack := goplatform.CommandAck{
			Uuid:   "test-uuid",
			Status: goplatform.CMD_STATUS_COMPLETED,
			Extras: map[string]any{
				"string":  "value",
				"number":  42,
				"boolean": true,
				"array":   []int{1, 2, 3},
				"object": map[string]any{
					"nested": "value",
					"count":  10,
				},
			},
		}

		b, err := json.Marshal(ack)
		if err != nil {
			t.Fatal(err)
		}

		var decoded goplatform.CommandAck
		if err := json.Unmarshal(b, &decoded); err != nil {
			t.Fatal(err)
		}

		// Verifica i campi base
		if decoded.Uuid != "test-uuid" {
			t.Fatalf("expected uuid to be 'test-uuid', got '%s'", decoded.Uuid)
		}
		if decoded.Status != goplatform.CMD_STATUS_COMPLETED {
			t.Fatalf("expected status to be 'completed', got '%s'", decoded.Status)
		}

		// Verifica gli extras
		if decoded.Extras["string"] != "value" {
			t.Fatalf("expected string extra to be 'value', got '%v'", decoded.Extras["string"])
		}
		if decoded.Extras["number"] != float64(42) {
			t.Fatalf("expected number extra to be 42, got '%v'", decoded.Extras["number"])
		}
		if decoded.Extras["boolean"] != true {
			t.Fatalf("expected boolean extra to be true, got '%v'", decoded.Extras["boolean"])
		}

		// Verifica array
		if arr, ok := decoded.Extras["array"].([]interface{}); !ok {
			t.Fatal("expected array to be []interface{}")
		} else {
			if len(arr) != 3 {
				t.Fatalf("expected array length to be 3, got %d", len(arr))
			}
			if arr[0] != float64(1) || arr[1] != float64(2) || arr[2] != float64(3) {
				t.Fatalf("expected array to be [1,2,3], got %v", arr)
			}
		}

		// Verifica oggetto annidato
		if obj, ok := decoded.Extras["object"].(map[string]interface{}); !ok {
			t.Fatal("expected object to be map[string]interface{}")
		} else {
			if obj["nested"] != "value" {
				t.Fatalf("expected nested value to be 'value', got '%v'", obj["nested"])
			}
			if obj["count"] != float64(10) {
				t.Fatalf("expected count to be 10, got '%v'", obj["count"])
			}
		}
	})

	t.Run("CommandAck Marshal/Unmarshal without extras", func(t *testing.T) {
		ack := goplatform.CommandAck{
			Uuid:   "simple-uuid",
			Status: goplatform.CMD_STATUS_RECEIVED,
		}

		b, err := json.Marshal(ack)
		if err != nil {
			t.Fatal(err)
		}

		var decoded goplatform.CommandAck
		if err := json.Unmarshal(b, &decoded); err != nil {
			t.Fatal(err)
		}

		if decoded.Uuid != "simple-uuid" {
			t.Fatalf("expected uuid to be 'simple-uuid', got '%s'", decoded.Uuid)
		}
		if decoded.Status != goplatform.CMD_STATUS_RECEIVED {
			t.Fatalf("expected status to be 'received', got '%s'", decoded.Status)
		}
		if len(decoded.Extras) != 0 {
			t.Fatalf("expected no extras, got %v", decoded.Extras)
		}
	})

	t.Run("CommandAck with invalid JSON", func(t *testing.T) {
		invalidJSON := []byte(`{"uuid": "test", "status": "completed", "invalid": }`)
		var ack goplatform.CommandAck

		err := json.Unmarshal(invalidJSON, &ack)
		if err == nil {
			t.Fatal("expected error when unmarshaling invalid JSON")
		}
	})

	t.Run("CommandParameters Marshal single parameter", func(t *testing.T) {
		params := goplatform.CommandParameters{
			goplatform.CommandParameter{
				"address": 1,
				"value":   42,
			},
		}
		b, err := json.Marshal(params)
		if err != nil {
			t.Fatal(err)
		}
		// Deve essere un oggetto JSON, non un array
		if string(b) != `{"address":1,"value":42}` {
			t.Fatalf("expected JSON object, got %s", string(b))
		}
	})

	t.Run("CommandParameters Marshal multiple parameters", func(t *testing.T) {
		params := goplatform.CommandParameters{
			goplatform.CommandParameter{"address": 1, "value": 42},
			goplatform.CommandParameter{"address": 2, "value": 84},
		}
		b, err := json.Marshal(params)
		if err != nil {
			t.Fatal(err)
		}
		// Deve essere un array JSON
		if string(b) != `[{"address":1,"value":42},{"address":2,"value":84}]` {
			t.Fatalf("expected JSON array, got %s", string(b))
		}
	})

	t.Run("CommandParameters Marshal empty", func(t *testing.T) {
		params := goplatform.CommandParameters{}
		b, err := json.Marshal(params)
		if err != nil {
			t.Fatal(err)
		}
		if string(b) != "null" {
			t.Fatalf("expected null, got %s", string(b))
		}
	})
}
