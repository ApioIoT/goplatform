## GoPlatform
An SDK for [Apio Platform](https://gitlab.com/apio-private/platform/api) in Golang.

#### Install
```bash
go get github.com/ApioIoT/goplatform
```

#### Get Projects
```golang
platform := goplatform.New(context.TODO(), URI, API_KEY)

projects, err := platform.GetProjects()
if err != nil {
  panic(err)
}
```

#### Get Project
```golang
platform := goplatform.New(context.TODO(), URI, API_KEY)

project, err := platform.GetProject(PROJECT_ID)
if err != nil {
  panic(err)
}
```

#### Get Nodes
```golang
platform := goplatform.New(context.TODO(), URI, API_KEY)

project, err := platform.GetProject(PROJECT_ID)
if err != nil {
  panic(err)
}

nodes, err := project.GetNodes()
if err != nil {
  panic(err)
}
```

#### Get Node
```golang
platform := goplatform.New(context.TODO(), URI, API_KEY)

project, err := platform.GetProject(PROJECT_ID)
if err != nil {
  panic(err)
}

node, err := project.GetNode(NODE_ID)
if err != nil {
  panic(err)
}
```

#### Get Devices
```golang
platform := goplatform.New(context.TODO(), URI, API_KEY)

project, err := platform.GetProject(PROJECT_ID)
if err != nil {
  panic(err)
}

devices, err := project.GetDevices()
if err != nil {
  panic(err)
}
```

#### Get Device
```golang
platform := goplatform.New(context.TODO(), URI, API_KEY)

project, err := platform.GetProject(PROJECT_ID)
if err != nil {
  panic(err)
}

device, err := project.GetDevice(DEVICE_ID)
if err != nil {
  panic(err)
}
```

#### Get DeviceTypes
```golang
platform := goplatform.New(context.TODO(), URI, API_KEY)

project, err := platform.GetProject(PROJECT_ID)
if err != nil {
  panic(err)
}

devices, err := project.GetDeviceTypes()
if err != nil {
  panic(err)
}
```

#### Get DeviceType
```golang
platform := goplatform.New(context.TODO(), URI, API_KEY)

project, err := platform.GetProject(PROJECT_ID)
if err != nil {
  panic(err)
}

deviceType, err := project.GetDeviceType(DEVICE_TYPE_ID)
if err != nil {
  panic(err)
}
```

#### Create Event
```golang
platform := goplatform.New(context.TODO(), URI, API_KEY)

project, err := platform.GetProject(PROJECT_ID)
if err != nil {
  panic(err)
}

event := goplatform.Event{
  Description: "Evento di prova",
  Type:        "Notification",
  Source:      "test/event",
}

if err := project.CreateEvent(event); err != nil {
  panic(err)
}
```