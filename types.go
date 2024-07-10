package goplatform

import (
	"time"
)

type LocationPointSchema struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Device struct {
	Uuid                string               `json:"uuid,omitempty"`
	ProjectID           string               `json:"projectId"`
	DeviceTypeID        string               `json:"deviceTypeId,omitempty"`
	NodeID              string               `json:"nodeId,omitempty"`
	Encoder             string               `json:"encoder,omitempty"`
	Decoder             string               `json:"decoder,omitempty"`
	FirmwareVersion     string               `json:"firmwareVersion,omitempty"`
	PlantID             string               `json:"plantId,omitempty"`
	Location            *LocationPointSchema `json:"location,omitempty"`
	SerialNumber        string               `json:"serialNumber,omitempty"`
	Name                string               `json:"name"`
	Description         string               `json:"description,omitempty"`
	DeviceType          *DeviceType          `json:"deviceType,omitempty"`
	Metadata            interface{}          `json:"metadata"`
	State               interface{}          `json:"state"`
	StateUpdatedAt      interface{}          `json:"stateUpdatedAt"`
	LastActivityAt      string               `json:"lastActivityAt"`
	LastCommunicationAt string               `json:"lastCommunicationAt"`
	ConnectivityStatus  string               `json:"connectivityStatus"`
	LastConnectionAt    string               `json:"lastConnectionAt"`
	LastDisconnectionAt string               `json:"lastDisconnectionAt"`
	Tags                []string             `json:"tags"`
	CreatedAt           time.Time            `json:"createdAt,omitempty"`
	UpdatedAt           time.Time            `json:"updatedAt,omitempty"`
	platformRef         *Platform            `json:"-"`
}

type DeviceType struct {
	Uuid             string      `json:"uuid,omitempty"`
	ProjectID        string      `json:"projectId"`
	Visibility       string      `json:"visibility,omitempty"`
	Encoder          string      `json:"encoder,omitempty"`
	Decoder          string      `json:"decoder,omitempty"`
	FirmwareID       string      `json:"firmwareId,omitempty"`
	FirmwareVersions []string    `json:"firmwareVersions"`
	Model            string      `json:"model,omitempty"`
	Manufacturer     string      `json:"manufacturer,omitempty"`
	Category         string      `json:"category,omitempty"`
	Name             string      `json:"name"`
	Description      string      `json:"description,omitempty"`
	Metadata         interface{} `json:"metadata"`
	Commands         interface{} `json:"commands"`
	Events           interface{} `json:"events"`
	Properties       interface{} `json:"properties"`
	CreatedAt        time.Time   `json:"createdAt,omitempty"`
	UpdatedAt        time.Time   `json:"updatedAt,omitempty"`
	platformRef      *Platform   `json:"-"`
}

type Event struct {
	Uuid        string      `json:"uuid,omitempty"`
	ProjectID   string      `json:"projectId"`
	Description string      `json:"description,omitempty"`
	Type        string      `json:"type"`
	EventTime   time.Time   `json:"eventTime,omitempty"`
	Source      string      `json:"source"`
	Metadata    interface{} `json:"metadata,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

type Response[T any] struct {
	Status bool `json:"status,omitempty"`
	Data   T    `json:"data,omitempty"`
}

type ResponseError struct {
	Status bool                `json:"status,omitempty"`
	Error  ResponseErrorDetail `json:"error,omitempty"`
}

type ResponseErrorDetail struct {
	Name       string `json:"name,omitempty"`
	StatusCode int    `json:"statusCode,omitempty"`
	Message    string `json:"message,omitempty"`
	Type       string `json:"type,omitempty"`
}
