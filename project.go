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

	var nodes Response[[]Node]
	if err := json.Unmarshal(b, &nodes); err != nil {
		return nil, err
	}

	for i := 0; i < len(nodes.Data); i++ {
		nodes.Data[i].platformRef = p.platformRef
	}

	return nodes.Data, nil
}

func (p Project) GetNode(ctx context.Context, uuid string) (Node, error) {
	var node Response[Node]

	b, err := p.platformRef.fetch(ctx, httpGet, nil, "projects", p.Uuid, "nodes", uuid)
	if err != nil {
		return node.Data, err
	}

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

	var devices Response[[]Device]
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
	var device Response[Device]

	b, err := p.platformRef.fetch(ctx, httpGet, nil, "projects", p.Uuid, "devices", uuid)
	if err != nil {
		return device.Data, err
	}

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

	var devices Response[[]DeviceType]
	if err := json.Unmarshal(b, &devices); err != nil {
		return nil, err
	}

	for i := 0; i < len(devices.Data); i++ {
		devices.Data[i].platformRef = p.platformRef
	}

	return devices.Data, nil
}

func (p Project) GetDeviceType(ctx context.Context, uuid string) (DeviceType, error) {
	var device Response[DeviceType]

	b, err := p.platformRef.fetch(ctx, httpGet, nil, "projects", p.Uuid, "devicetypes", uuid)
	if err != nil {
		return device.Data, err
	}

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

	var rules Response[[]Rule]
	if err := json.Unmarshal(b, &rules); err != nil {
		return nil, err
	}

	for i := 0; i < len(rules.Data); i++ {
		rules.Data[i].platformRef = p.platformRef
	}

	return rules.Data, nil
}

func (p Project) GetRule(ctx context.Context, uuid string) (Rule, error) {
	var rule Response[Rule]

	b, err := p.platformRef.fetch(ctx, httpGet, nil, "projects", p.Uuid, "rules", uuid)
	if err != nil {
		return rule.Data, err
	}

	if err := json.Unmarshal(b, &rule); err != nil {
		return rule.Data, err
	}

	rule.Data.platformRef = p.platformRef

	return rule.Data, nil
}
