package goplatform

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Project struct {
	Uuid          string         `json:"uuid"`
	ProjectId     string         `json:"projectId"`
	Name          string         `json:"name"`
	Metadata      map[string]any `json:"metadata,omitempty"`
	Configuration any            `json:"configuration,omitempty"`
	CreatedAt     time.Time      `json:"createdAt,omitempty"`
	UpdatedAt     time.Time      `json:"updatedAt,omitempty"`
	platformRef   *Platform      `json:"-"`
}

func (p Project) GetNodes() ([]Node, error) {
	req, err := p.platformRef.makeRequest("GET", nil, "projects", p.Uuid, "nodes")
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := errorFromResponse(resp); err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
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

func (p Project) GetNode(uuid string) (Node, error) {
	var node Response[Node]

	req, err := p.platformRef.makeRequest("GET", nil, "projects", p.Uuid, "nodes", uuid)
	if err != nil {
		return node.Data, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return node.Data, err
	}
	defer resp.Body.Close()

	if err := errorFromResponse(resp); err != nil {
		return node.Data, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return node.Data, err
	}

	if err := json.Unmarshal(b, &node); err != nil {
		return node.Data, err
	}

	node.Data.platformRef = p.platformRef

	return node.Data, nil
}

func (p Project) GetDevices() ([]Device, error) {
	req, err := p.platformRef.makeRequest("GET", nil, "projects", p.Uuid, "devices")
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := errorFromResponse(resp); err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
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

func (p Project) GetDevice(uuid string) (Device, error) {
	var device Response[Device]

	req, err := p.platformRef.makeRequest("GET", nil, "projects", p.Uuid, "devices", uuid)
	if err != nil {
		return device.Data, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return device.Data, err
	}
	defer resp.Body.Close()

	if err := errorFromResponse(resp); err != nil {
		return device.Data, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return device.Data, err
	}

	if err := json.Unmarshal(b, &device); err != nil {
		return device.Data, err
	}

	device.Data.platformRef = p.platformRef

	return device.Data, nil
}

func (p Project) GetDeviceTypes() ([]DeviceType, error) {
	req, err := p.platformRef.makeRequest("GET", nil, "projects", p.Uuid, "devicetypes")
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := errorFromResponse(resp); err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
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

func (p Project) GetDeviceType(uuid string) (DeviceType, error) {
	var device Response[DeviceType]

	req, err := p.platformRef.makeRequest("GET", nil, "projects", p.Uuid, "devicetypes", uuid)
	if err != nil {
		return device.Data, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return device.Data, err
	}
	defer resp.Body.Close()

	if err := errorFromResponse(resp); err != nil {
		return device.Data, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return device.Data, err
	}

	if err := json.Unmarshal(b, &device); err != nil {
		return device.Data, err
	}

	device.Data.platformRef = p.platformRef

	return device.Data, nil
}

func (p Project) CreateEvent(event Event) error {
	event.ProjectID = p.ProjectId

	b, err := json.Marshal(event)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)

	req, err := p.platformRef.makeRequest("POST", body, "projects", p.Uuid, "events")
	if err != nil {
		return err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return errorFromResponse(resp)
}

func (p Project) GetRules() ([]Rule, error) {
	req, err := p.platformRef.makeRequest("GET", nil, "projects", p.Uuid, "rules")
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := errorFromResponse(resp); err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
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

func (p Project) GetRule(uuid string) (Rule, error) {
	var rule Response[Rule]

	req, err := p.platformRef.makeRequest("GET", nil, "projects", p.Uuid, "rules", uuid)
	if err != nil {
		return rule.Data, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return rule.Data, err
	}
	defer resp.Body.Close()

	if err := errorFromResponse(resp); err != nil {
		return rule.Data, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return rule.Data, err
	}

	if err := json.Unmarshal(b, &rule); err != nil {
		return rule.Data, err
	}

	rule.Data.platformRef = p.platformRef

	return rule.Data, nil
}
