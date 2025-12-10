![Golang version](https://img.shields.io/badge/go-1.23-brightgreen)
![License](https://img.shields.io/github/license/ApioIoT/goplatform)
![CI](https://github.com/ApioIoT/goplatform/actions/workflows/ci.yml/badge.svg)

## GoPlatform
An SDK for [Apio Platform](https://gitlab.com/apio-private/platform/api) in Golang.

#### Install
```bash
go get github.com/ApioIoT/goplatform/v2
```

#### Get Projects
```golang
platform := goplatform.New(goplatform.Config{
  Uri:    "platform-uri",
  ApiKey: "my-api-key",
})

projects, err := platform.GetProjects(context.TODO())
if err != nil {
  panic(err)
}
```

#### Get Project
```golang
platform := goplatform.New(goplatform.Config{
  Uri:    "platform-uri",
  ApiKey: "my-api-key",
})

project, err := platform.GetProject(context.TODO(), "my-project-id")
if err != nil {
  panic(err)
}
```

#### Get Nodes
```golang
platform := goplatform.New(goplatform.Config{
  Uri:    "platform-uri",
  ApiKey: "my-api-key",
})

project, err := platform.GetProject(context.TODO(), "my-project-id")
if err != nil {
  panic(err)
}

nodes, err := project.GetNodes(context.TODO())
if err != nil {
  panic(err)
}
```

#### Get Node
```golang
platform := goplatform.New(goplatform.Config{
  Uri:    "platform-uri",
  ApiKey: "my-api-key",
})

project, err := platform.GetProject(context.TODO(), "my-project-id")
if err != nil {
  panic(err)
}

node, err := project.GetNode(context.TODO(), "my-node-id")
if err != nil {
  panic(err)
}
```

#### Get Devices
```golang
platform := goplatform.New(goplatform.Config{
  Uri:    "platform-uri",
  ApiKey: "my-api-key",
})

project, err := platform.GetProject(context.TODO(), "my-project-id")
if err != nil {
  panic(err)
}

devices, err := project.GetDevices(context.TODO())
if err != nil {
  panic(err)
}
```

#### Get Device
```golang
platform := goplatform.New(goplatform.Config{
  Uri:    "platform-uri",
  ApiKey: "my-api-key",
})

project, err := platform.GetProject(context.TODO(), "my-project-id")
if err != nil {
  panic(err)
}

device, err := project.GetDevice(context.TODO(), "my-device-id")
if err != nil {
  panic(err)
}
```

#### Get DeviceTypes
```golang
platform := goplatform.New(goplatform.Config{
  Uri:    "platform-uri",
  ApiKey: "my-api-key",
})

project, err := platform.GetProject(context.TODO(), "my-project-id")
if err != nil {
  panic(err)
}

devices, err := project.GetDeviceTypes(context.TODO())
if err != nil {
  panic(err)
}
```

#### Get DeviceType
```golang
platform := goplatform.New(goplatform.Config{
  Uri:    "platform-uri",
  ApiKey: "my-api-key",
})

project, err := platform.GetProject(context.TODO(), "my-project-id")
if err != nil {
  panic(err)
}

deviceType, err := project.GetDeviceType(context.TODO(), "my-device-type-id")
if err != nil {
  panic(err)
}
```

#### Create Event
```golang
platform := goplatform.New(goplatform.Config{
  Uri:    "platform-uri",
  ApiKey: "my-api-key",
})

project, err := platform.GetProject(context.TODO(), "my-project-id")
if err != nil {
  panic(err)
}

event := goplatform.Event{
  Description: "Evento di prova",
  Type:        "Notification",
  Source:      "test/event",
}

if err := project.CreateEvent(context.TODO(), event); err != nil {
  panic(err)
}
```