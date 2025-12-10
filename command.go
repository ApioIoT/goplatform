package goplatform

import (
	"encoding/json"
	"maps"
	"strings"
	"time"

	"github.com/google/uuid"
)

type CommandStatus string

const (
	CMD_STATUS_PENDING   CommandStatus = "pending"
	CMD_STATUS_RECEIVED  CommandStatus = "received"
	CMD_STATUS_COMPLETED CommandStatus = "completed"
	CMD_STATUS_FAILED    CommandStatus = "failed"
)

type CommandRequest struct {
	Name           string                     `json:"name"`
	ProjectId      string                     `json:"projectId"`
	NodeId         *string                    `json:"nodeId,omitempty"`
	DeviceId       *string                    `json:"deviceId,omitempty"`
	Parameters     CommandParameters          `json:"parameters"`
	Metadata       map[string]any             `json:"metadata,omitempty"`
	DownlinkRetry  *CommandRequestRetryOption `json:"downlinkRetry,omitempty"`
	ExecutionRetry *CommandRequestRetryOption `json:"executionRetry,omitempty"`
}

func (c CommandRequest) MakeCommand(status ...CommandStatus) Command {
	now := time.Now()

	_status := CMD_STATUS_PENDING
	if len(status) > 0 {
		_status = status[0]
	}

	cmd := Command{
		Uuid:       uuid.NewString(),
		Name:       c.Name,
		ProjectId:  c.ProjectId,
		NodeId:     c.NodeId,
		DeviceId:   c.DeviceId,
		Parameters: c.Parameters,
		Metadata:   c.Metadata,
		Status:     _status,
		CreatedAt:  &now,
	}

	if c.DownlinkRetry != nil {
		cmd.DownlinkRetry = &DownlinkRetryOption{
			MaxRetries: c.DownlinkRetry.MaxRetries,
		}
	}
	if c.ExecutionRetry != nil {
		cmd.ExecutionRetry = &ExecutionRetryOption{
			MaxRetries: c.ExecutionRetry.MaxRetries,
		}
	}

	switch _status {
	case CMD_STATUS_PENDING:
		// Intentionally empty
	case CMD_STATUS_RECEIVED:
		cmd.ReceivedAt = &now
	case CMD_STATUS_COMPLETED:
		cmd.CompletedAt = &now
	case CMD_STATUS_FAILED:
		cmd.FailedAt = &now
	}

	return cmd
}

type CommandRequestRetryOption struct {
	MaxRetries *int `json:"maxRetries,omitempty"`
}

type Command struct {
	Uuid           string                `json:"uuid"`
	Name           string                `json:"name"`
	ProjectId      string                `json:"projectId"`
	NodeId         *string               `json:"nodeId,omitempty"`
	DeviceId       *string               `json:"deviceId,omitempty"`
	Parameters     CommandParameters     `json:"parameters"`
	Metadata       map[string]any        `json:"metadata,omitempty"`
	Status         CommandStatus         `json:"status"`
	DownlinkRetry  *DownlinkRetryOption  `json:"downlinkRetry,omitempty"`
	ExecutionRetry *ExecutionRetryOption `json:"executionRetry,omitempty"`
	CreatedAt      *time.Time            `json:"createdAt,omitempty"`
	UpdatedAt      *time.Time            `json:"updatedAt,omitempty"`
	ReceivedAt     *time.Time            `json:"receivedAt,omitempty"`
	CompletedAt    *time.Time            `json:"completedAt,omitempty"`
	FailedAt       *time.Time            `json:"failedAt,omitempty"`
}

type DownlinkRetryOption struct {
	MaxRetries  *int       `json:"maxRetries,omitempty"`
	RetryCount  *int       `json:"retryCount,omitempty"`
	NextRetryAt *time.Time `json:"nextRetryAt,omitempty"`
}

type ExecutionRetryOption struct {
	MaxRetries *int `json:"maxRetries,omitempty"`
	RetryCount *int `json:"retryCount,omitempty"`
}

type CommandParameter map[string]any
type CommandParameters []CommandParameter

func (c *CommandParameters) UnmarshalJSON(data []byte) error {
	var single CommandParameter
	if err := json.Unmarshal(data, &single); err != nil {
		var array []CommandParameter
		if err := json.Unmarshal(data, &array); err != nil {
			return err
		}
		*c = array
	} else {
		*c = append(*c, single)
	}

	return nil
}

func (c CommandParameters) MarshalJSON() ([]byte, error) {
	if len(c) == 0 {
		return []byte("null"), nil
	}

	if len(c) == 1 {
		return json.Marshal(c[0])
	}

	res := []string{}
	for _, el := range c {
		b, err := json.Marshal(el)
		if err != nil {
			return nil, err
		}
		res = append(res, string(b))
	}
	return []byte("[" + strings.Join(res, ",") + "]"), nil
}

type CommandAck struct {
	Uuid   string         `json:"uuid"`
	Status CommandStatus  `json:"status"`
	Extras map[string]any `json:"-"`
}

func (c *CommandAck) UnmarshalJSON(data []byte) error {
	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if uuid, ok := raw["uuid"].(string); ok {
		c.Uuid = uuid
	}
	if status, ok := raw["status"].(string); ok {
		c.Status = CommandStatus(status)
	}

	delete(raw, "uuid")
	delete(raw, "status")
	c.Extras = raw

	return nil
}

func (c CommandAck) MarshalJSON() ([]byte, error) {
	data := map[string]any{
		"uuid":   c.Uuid,
		"status": c.Status,
	}

	maps.Copy(data, c.Extras)

	return json.Marshal(data)
}
