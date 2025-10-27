package goplatform_test

import (
	"encoding/json"
	"testing"

	"github.com/ApioIoT/goplatform"
)

func TestAck(t *testing.T) {
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
}
