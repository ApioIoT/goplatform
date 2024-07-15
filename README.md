## GoPlatform
An SDK for [Apio Platform](https://gitlab.com/apio-private/platform/api) in Golang.

#### Install
```bash
go get github.com/ApioIoT/goplatform
```

#### Get Projects
```golang
platform := goplatform.New(URI, API_KEY)

projects, err := platform.GetProjects()
if err != nil {
  log.Fatal(err)
}
```

#### Get Project
```golang
platform := goplatform.New(URI, API_KEY)

project, err := platform.GetProject(PROJECT_ID)
if err != nil {
  log.Fatal(err)
}
```

#### Get Nodes
```golang
platform := goplatform.New(URI, API_KEY)

project, err := platform.GetProject(PROJECT_ID)
if err != nil {
  log.Fatal(err)
}

nodes, err := project.GetNodes()
if err != nil {
  log.Fatal(err)
}
```

#### Get Node
```golang
platform := goplatform.New(URI, API_KEY)

project, err := platform.GetProject(PROJECT_ID)
if err != nil {
  log.Fatal(err)
}

node, err := project.GetNode(NODE_ID)
if err != nil {
  log.Fatal(err)
}
```

#### Get Devices
```golang
platform := goplatform.New(URI, API_KEY)

project, err := platform.GetProject(PROJECT_ID)
if err != nil {
  log.Fatal(err)
}

devices, err := project.GetDevices()
if err != nil {
  log.Fatal(err)
}
```

#### Get Device
```golang
platform := goplatform.New(URI, API_KEY)

project, err := platform.GetProject(PROJECT_ID)
if err != nil {
  log.Fatal(err)
}

device, err := project.GetDevice(DEVICE_ID)
if err != nil {
  log.Fatal(err)
}
```

#### Get DeviceTypes
```golang
platform := goplatform.New(URI, API_KEY)

project, err := platform.GetProject(PROJECT_ID)
if err != nil {
  log.Fatal(err)
}

devices, err := project.GetDeviceTypes()
if err != nil {
  log.Fatal(err)
}
```

#### Get DeviceType
```golang
platform := goplatform.New(URI, API_KEY)

project, err := platform.GetProject(PROJECT_ID)
if err != nil {
  log.Fatal(err)
}

deviceType, err := project.GetDeviceType(DEVICE_TYPE_ID)
if err != nil {
  log.Fatal(err)
}
```

#### Create Event
```golang
platform := goplatform.New(URI, API_KEY)

project, err := platform.GetProject(PROJECT_ID)
if err != nil {
  log.Fatal(err)
}

event := goplatform.Event{
  Description: "Evento di prova",
  Type:        "Notification",
  Source:      "test/event",
}

if err := project.CreateEvent(event); err != nil {
  log.Fatal(err)
}
```