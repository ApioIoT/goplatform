package goplatform

import (
	"time"
)

type response[T any] struct {
	Status bool `json:"status,omitempty"`
	Data   T    `json:"data,omitempty"`
}

type responseError struct {
	Status bool                `json:"status,omitempty"`
	Error  responseErrorDetail `json:"error,omitempty"`
}

type responseErrorDetail struct {
	Name       string `json:"name,omitempty"`
	StatusCode int    `json:"statusCode,omitempty"`
	Message    string `json:"message,omitempty"`
	Type       string `json:"type,omitempty"`
}

type LocationPointSchema struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type NodeProtocol struct {
	Uuid          string         `json:"uuid"`
	Name          string         `json:"name"`
	Description   string         `json:"description,omitempty"`
	Metadata      map[string]any `json:"metadata,omitempty"`
	Configuration any            `json:"configuration,omitempty"`
}

type Node struct {
	Uuid         string               `json:"uuid,omitempty"`
	ProjectID    string               `json:"projectId"`
	PlantID      string               `json:"plantId,omitempty"`
	Name         string               `json:"name"`
	Model        string               `json:"model,omitempty"`
	NodeTypeID   string               `json:"nodeTypeId,omitempty"`
	SerialNumber string               `json:"serialNumber,omitempty"`
	Location     *LocationPointSchema `json:"location,omitempty"`
	Protocols    []NodeProtocol       `json:"protocols,omitempty"`
	Metadata     map[string]any       `json:"metadata"`
	Retry        *struct {
		Enabled          bool    `json:"enabled"`
		MaxRetries       int     `json:"maxRetries"`
		BackoffBase      float64 `json:"backoffBase"`
		BackoffFactor    float64 `json:"backoffFactor"`
		BackoffTimeLimit int     `json:"backoffTimeLimit"`
	} `json:"retry,omitempty"`
	ConnectivityStatus  string    `json:"connectivityStatus"`
	LastConnectionAt    string    `json:"lastConnectionAt"`
	LastCommunicationAt string    `json:"lastCommunicationAt"`
	LastDisconnectionAt string    `json:"lastDisconnectionAt"`
	Description         string    `json:"description,omitempty"`
	Tags                []string  `json:"tags"`
	Rules               []string  `json:"rules,omitempty"`
	CreatedAt           time.Time `json:"createdAt,omitempty"`
	UpdatedAt           time.Time `json:"updatedAt,omitempty"`
	platformRef         *Platform `json:"-"`
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
	Metadata            map[string]any       `json:"metadata"`
	State               any                  `json:"state"`
	StateUpdatedAt      any                  `json:"stateUpdatedAt"`
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
	Register            uint16 `json:"register"`
	Read                bool   `json:"read"`
	ModbusFunctionRead  int    `json:"modbusFunctionRead"`
	Write               bool   `json:"write"`
	ModbusFunctionWrite int    `json:"modbusFunctionWrite,omitempty"`
	Words               byte   `json:"words"`
	BitwiseReading      bool   `json:"bitwiseReading"`
	Properties          []struct {
		Index int    `json:"index"`
		Name  string `json:"name"`
	} `json:"properties"`
	ScaleFactor *float64 `json:"scaleFactor,omitempty"`
	Type        string   `json:"type"`
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
	Metadata    map[string]any `json:"metadata"`
	Commands    any            `json:"commands"`
	Events      any            `json:"events"`
	Properties  any            `json:"properties"`
	CreatedAt   time.Time      `json:"createdAt,omitempty"`
	UpdatedAt   time.Time      `json:"updatedAt,omitempty"`
	platformRef *Platform      `json:"-"`
}

type Event struct {
	Uuid        string         `json:"uuid,omitempty"`
	ProjectID   string         `json:"projectId"`
	Description string         `json:"description,omitempty"`
	Type        string         `json:"type"`
	EventTime   time.Time      `json:"eventTime,omitempty"`
	Source      string         `json:"source"`
	Metadata    map[string]any `json:"metadata,omitempty"`
	Data        any            `json:"data,omitempty"`
}

type RuleAction struct {
	Type string `json:"type"`
	// Type wasm
	Label    *string `json:"label,omitempty"`
	Language *string `json:"language,omitempty"`
	Source   *string `json:"source,omitempty"`
	Binary   *string `json:"binary,omitempty"`
	Hash     *string `json:"hash,omitempty"`
	// Type log
	Value *string `json:"value,omitempty"`
	// Type webhook
	Uri     *string `json:"uri,omitempty"`
	Payload *string `json:"payload,omitempty"`
	// Type command
	Command  map[string]any `json:"command,omitempty"`
	DeviceId *string        `json:"deviceId,omitempty"`
	// Type delay
	Time *int64 `json:"time,omitempty"`
}

type RuleTrigger struct {
	Type     string  `json:"type"`
	DeviceId *string `json:"deviceId,omitempty"`
	Property *string `json:"property,omitempty"`
	Crontab  *string `json:"crontab,omitempty"`
	Topic    *string `json:"topic,omitempty"`
}

type RuleCondition struct {
	Devices   map[string]string `json:"devices"`
	Condition string            `json:"condition"`
	Deadline  *int              `json:"deadline,omitempty"`
}

type Rule struct {
	Uuid            string         `json:"uuid"`
	ProjectId       string         `json:"projectId"`
	Name            string         `json:"name"`
	Description     *string        `json:"description,omitempty"`
	Metadata        map[string]any `json:"metadata,omitempty"`
	Tags            []string       `json:"tags,omitempty"`
	Status          string         `json:"status"`
	AllowConcurrent bool           `json:"allowConcurrent"`
	Mode            string         `json:"mode"`
	Triggers        []RuleTrigger  `json:"triggers"`
	Actions         []RuleAction   `json:"actions"`
	ElseActions     []RuleAction   `json:"elseActions"`
	Condition       *RuleCondition `json:"condition,omitempty"`
	platformRef     *Platform      `json:"-"`
}

type Measure struct {
	ProjectId string `json:"projectId"`
	DeviceId  string `json:"deviceId"`
	Timestamp int64  `json:"timestamp"`
	Name      string `json:"name"`
	Value     any    `json:"value"`
}
