package goplatform

import (
	"bytes"
	"context"
	"encoding/json"
	"time"
)

type Project struct {
	Uuid        string         `json:"uuid"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Metadata    map[string]any `json:"metadata,omitempty"`
	Features    []any          `json:"features"`
	CreatedAt   time.Time      `json:"createdAt,omitempty"`
	UpdatedAt   time.Time      `json:"updatedAt,omitempty"`
	platformRef *Platform      `json:"-"`
}

func (p Project) GetNodes(ctx context.Context) ([]Node, error) {
	b, err := p.platformRef.fetch(ctx, httpGet, nil, "projects", p.Uuid, "nodes")
	if err != nil {
		return nil, err
	}

	var nodes response[[]Node]
	if err := json.Unmarshal(b, &nodes); err != nil {
		return nil, err
	}

	for i := 0; i < len(nodes.Data); i++ {
		nodes.Data[i].platformRef = p.platformRef
	}

	return nodes.Data, nil
}

func (p Project) GetNode(ctx context.Context, uuid string) (Node, error) {
	b, err := p.platformRef.fetch(ctx, httpGet, nil, "projects", p.Uuid, "nodes", uuid)
	if err != nil {
		var zero Node
		return zero, err
	}

	var node response[Node]
	if err := json.Unmarshal(b, &node); err != nil {
		return node.Data, err
	}

	node.Data.platformRef = p.platformRef

	return node.Data, nil
}

func (p Project) GetDevices(ctx context.Context) ([]Device, error) {
	b, err := p.platformRef.fetch(ctx, httpGet, nil, "projects", p.Uuid, "devices")
	if err != nil {
		return nil, err
	}

	var devices response[[]Device]
	if err := json.Unmarshal(b, &devices); err != nil {
		return nil, err
	}

	for i := 0; i < len(devices.Data); i++ {
		devices.Data[i].platformRef = p.platformRef
		if devices.Data[i].DeviceType != nil {
			devices.Data[i].DeviceType.platformRef = p.platformRef
		}
	}

	return devices.Data, nil
}

func (p Project) GetDevice(ctx context.Context, uuid string) (Device, error) {
	b, err := p.platformRef.fetch(ctx, httpGet, nil, "projects", p.Uuid, "devices", uuid)
	if err != nil {
		var zero Device
		return zero, err
	}

	var device response[Device]
	if err := json.Unmarshal(b, &device); err != nil {
		return device.Data, err
	}

	device.Data.platformRef = p.platformRef

	return device.Data, nil
}

func (p Project) GetDeviceTypes(ctx context.Context) ([]DeviceType, error) {
	b, err := p.platformRef.fetch(ctx, httpGet, nil, "projects", p.Uuid, "devicetypes")
	if err != nil {
		return nil, err
	}

	var devices response[[]DeviceType]
	if err := json.Unmarshal(b, &devices); err != nil {
		return nil, err
	}

	for i := 0; i < len(devices.Data); i++ {
		devices.Data[i].platformRef = p.platformRef
	}

	return devices.Data, nil
}

func (p Project) GetDeviceType(ctx context.Context, uuid string) (DeviceType, error) {
	b, err := p.platformRef.fetch(ctx, httpGet, nil, "projects", p.Uuid, "devicetypes", uuid)
	if err != nil {
		var zero DeviceType
		return zero, err
	}

	var device response[DeviceType]
	if err := json.Unmarshal(b, &device); err != nil {
		return device.Data, err
	}

	device.Data.platformRef = p.platformRef

	return device.Data, nil
}

func (p Project) CreateEvent(ctx context.Context, event Event) error {
	event.ProjectID = p.Uuid

	b, err := json.Marshal(event)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)

	_, err = p.platformRef.fetch(ctx, httpPost, body, "projects", p.Uuid, "events")
	return err
}

func (p Project) GetRules(ctx context.Context) ([]Rule, error) {
	b, err := p.platformRef.fetch(ctx, httpGet, nil, "projects", p.Uuid, "rules")
	if err != nil {
		return nil, err
	}

	var rules response[[]Rule]
	if err := json.Unmarshal(b, &rules); err != nil {
		return nil, err
	}

	for i := 0; i < len(rules.Data); i++ {
		rules.Data[i].platformRef = p.platformRef
	}

	return rules.Data, nil
}

func (p Project) GetRule(ctx context.Context, uuid string) (Rule, error) {
	b, err := p.platformRef.fetch(ctx, httpGet, nil, "projects", p.Uuid, "rules", uuid)
	if err != nil {
		var zero Rule
		return zero, err
	}

	var rule response[Rule]
	if err := json.Unmarshal(b, &rule); err != nil {
		return rule.Data, err
	}

	rule.Data.platformRef = p.platformRef

	return rule.Data, nil
}

func (p Project) SendCommand(ctx context.Context, command Command) error {
	b, err := json.Marshal(command)
	if err != nil {
		return err
	}

	_, err = p.platformRef.fetch(ctx, httpPost, bytes.NewReader(b), "projects", p.Uuid, "commands")
	return err
}
