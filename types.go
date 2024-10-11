package goplatform

import (
	"time"
)

type LocationPointSchema struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type NodeProtocol struct {
	Uuid          string      `json:"uuid"`
	Name          string      `json:"name"`
	Description   string      `json:"description,omitempty"`
	Metadata      interface{} `json:"metadata,omitempty"`
	Configuration interface{} `json:"configuration,omitempty"`
	// Mapping       map[string]struct {
	// 	Address    int      `json:"address"`
	// 	Properties []string `json:"properties,omitempty"`
	// } `json:"mapping"`
}

type Node struct {
	Uuid                string               `json:"uuid,omitempty"`
	ProjectID           string               `json:"projectId"`
	PlantID             string               `json:"plantId,omitempty"`
	Name                string               `json:"name"`
	Model               string               `json:"model,omitempty"`
	NodeTypeID          string               `json:"nodeTypeId,omitempty"`
	SerialNumber        string               `json:"serialNumber,omitempty"`
	Location            *LocationPointSchema `json:"location,omitempty"`
	Protocols           []NodeProtocol       `json:"protocols,omitempty"`
	Metadata            interface{}          `json:"metadata"`
	ConnectivityStatus  string               `json:"connectivityStatus"`
	LastConnectionAt    string               `json:"lastConnectionAt"`
	LastCommunicationAt string               `json:"lastCommunicationAt"`
	LastDisconnectionAt string               `json:"lastDisconnectionAt"`
	Description         string               `json:"description,omitempty"`
	Tags                []string             `json:"tags"`
	CreatedAt           time.Time            `json:"createdAt,omitempty"`
	UpdatedAt           time.Time            `json:"updatedAt,omitempty"`
	platformRef         *Platform            `json:"-"`
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

type DeviceTypeModbusProtocol struct {
	Endianness string                     `json:"endianness"`
	Registers  []DeviceTypeModbusRegister `json:"registers"`
}

type DeviceTypeModbusRegister struct {
	Register       uint16 `json:"register"`
	ModbusFunction int    `json:"modbusFunction"`
	Words          byte   `json:"words"`
	BitwiseReading bool   `json:"bitwiseReading"`
	Properties     []struct {
		Index int    `json:"index"`
		Name  string `json:"name"`
	} `json:"properties"`
	Type string `json:"type"`
}

type DeviceTypeKnxProtocol struct {
	Properties map[string]struct {
		Address    string `json:"address"`
		SendDPT    string `json:"sendDPT"`
		ReceiveDPT string `json:"receiveDPT"`
	} `json:"properties"`
}

type DeviceType struct {
	Uuid             string   `json:"uuid,omitempty"`
	ProjectID        string   `json:"projectId"`
	Visibility       string   `json:"visibility,omitempty"`
	Encoder          string   `json:"encoder,omitempty"`
	Decoder          string   `json:"decoder,omitempty"`
	FirmwareID       string   `json:"firmwareId,omitempty"`
	FirmwareVersions []string `json:"firmwareVersions"`
	Model            string   `json:"model,omitempty"`
	Manufacturer     string   `json:"manufacturer,omitempty"`
	Category         string   `json:"category,omitempty"`
	Name             string   `json:"name"`
	Description      string   `json:"description,omitempty"`
	Protocols        *struct {
		Modbus *DeviceTypeModbusProtocol `json:"modbus,omitempty"`
		Knx    *DeviceTypeKnxProtocol    `json:"knx,omitempty"`
	} `json:"protocols,omitempty"`
	Metadata    interface{} `json:"metadata"`
	Commands    interface{} `json:"commands"`
	Events      interface{} `json:"events"`
	Properties  interface{} `json:"properties"`
	CreatedAt   time.Time   `json:"createdAt,omitempty"`
	UpdatedAt   time.Time   `json:"updatedAt,omitempty"`
	platformRef *Platform   `json:"-"`
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
