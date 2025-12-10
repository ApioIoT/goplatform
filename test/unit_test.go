package goplatform_test

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"os"

	"github.com/ApioIoT/goplatform"
	"github.com/h2non/gock"
)

const (
	API_URI        = "http://api.example.com"
	API_KEY        = "my-api-key"
	PROJECT_ID     = "my-project-id"
	NODE_ID        = "my-node-id"
	DEVICE_ID      = "my-device-id"
	DEVICE_TYPE_ID = "my-devicetype-id"
	RULE_ID        = "my-rule-id"
)

func TestUnit(t *testing.T) {

	// --- SETUP GOCK

	defer gock.Off()

	eventsUrl, _ := url.JoinPath("/projects", PROJECT_ID, "/events")
	gock.New(API_URI).
		AddMatcher(func(req *http.Request, ereq *gock.Request) (bool, error) {
			return req.Method == "POST" && req.URL.Path == eventsUrl, nil
		}).
		Persist().
		MatchHeader("Authorization", fmt.Sprintf("apiKey %s", API_KEY)).
		Reply(200)

	projects, err := os.ReadFile("./mock/projects.json")
	if err != nil {
		t.Fatalf("errore nel leggere il file: %v", err)
	}

	gock.New(API_URI).
		AddMatcher(func(req *http.Request, ereq *gock.Request) (bool, error) {
			return req.Method == "GET" && req.URL.Path == "/projects", nil
		}).
		Persist().
		MatchHeader("Authorization", fmt.Sprintf("apiKey %s", API_KEY)).
		Reply(200).
		Type("application/json").
		BodyString(string(projects))

	project, err := os.ReadFile("./mock/project.json")
	if err != nil {
		t.Fatalf("errore nel leggere il file: %v", err)
	}

	projectUrl, _ := url.JoinPath("/projects", PROJECT_ID, "/")
	gock.New(API_URI).
		AddMatcher(func(req *http.Request, ereq *gock.Request) (bool, error) {
			return req.Method == "GET" && req.URL.Path == projectUrl, nil
		}).
		Persist().
		MatchHeader("Authorization", fmt.Sprintf("apiKey %s", API_KEY)).
		Reply(200).
		Type("application/json").
		BodyString(string(project))

	node, err := os.ReadFile("./mock/node.json")
	if err != nil {
		t.Fatalf("errore nel leggere il file: %v", err)
	}

	nodeUrl, _ := url.JoinPath("/projects", PROJECT_ID, "nodes", NODE_ID)
	gock.New(API_URI).
		AddMatcher(func(req *http.Request, ereq *gock.Request) (bool, error) {
			return req.Method == "GET" && req.URL.Path == nodeUrl, nil
		}).
		Persist().
		MatchHeader("Authorization", fmt.Sprintf("apiKey %s", API_KEY)).
		Reply(200).
		Type("application/json").
		BodyString(string(node))

	nodes, err := os.ReadFile("./mock/nodes.json")
	if err != nil {
		t.Fatalf("errore nel leggere il file: %v", err)
	}

	nodesUrl, _ := url.JoinPath("/projects", PROJECT_ID, "nodes")
	gock.New(API_URI).
		AddMatcher(func(req *http.Request, ereq *gock.Request) (bool, error) {
			return req.Method == "GET" && req.URL.Path == nodesUrl, nil
		}).
		Persist().
		MatchHeader("Authorization", fmt.Sprintf("apiKey %s", API_KEY)).
		Reply(200).
		Type("application/json").
		BodyString(string(nodes))

	rules, err := os.ReadFile("./mock/rules.json")
	if err != nil {
		t.Fatalf("errore nel leggere il file: %v", err)
	}

	rulesUrl, _ := url.JoinPath("/projects", PROJECT_ID, "rules")
	gock.New(API_URI).
		AddMatcher(func(req *http.Request, ereq *gock.Request) (bool, error) {
			return req.Method == "GET" && req.URL.Path == rulesUrl, nil
		}).
		Persist().
		MatchHeader("Authorization", fmt.Sprintf("apiKey %s", API_KEY)).
		Reply(200).
		Type("application/json").
		BodyString(string(rules))

	rule, err := os.ReadFile("./mock/rule.json")
	if err != nil {
		t.Fatalf("errore nel leggere il file: %v", err)
	}

	ruleUrl, _ := url.JoinPath("/projects", PROJECT_ID, "rules", RULE_ID)
	gock.New(API_URI).
		AddMatcher(func(req *http.Request, ereq *gock.Request) (bool, error) {
			return req.Method == "GET" && req.URL.Path == ruleUrl, nil
		}).
		Persist().
		MatchHeader("Authorization", fmt.Sprintf("apiKey %s", API_KEY)).
		Reply(200).
		Type("application/json").
		BodyString(string(rule))

	devices, err := os.ReadFile("./mock/devices.json")
	if err != nil {
		t.Fatalf("errore nel leggere il file: %v", err)
	}

	devicesUrl, _ := url.JoinPath("/projects", PROJECT_ID, "devices")
	gock.New(API_URI).
		AddMatcher(func(req *http.Request, ereq *gock.Request) (bool, error) {
			return req.Method == "GET" && req.URL.Path == devicesUrl, nil
		}).
		Persist().
		MatchHeader("Authorization", fmt.Sprintf("apiKey %s", API_KEY)).
		Reply(200).
		Type("application/json").
		BodyString(string(devices))

	device, err := os.ReadFile("./mock/device.json")
	if err != nil {
		t.Fatalf("errore nel leggere il file: %v", err)
	}

	deviceUrl, _ := url.JoinPath("/projects", PROJECT_ID, "devices", DEVICE_ID)
	gock.New(API_URI).
		AddMatcher(func(req *http.Request, ereq *gock.Request) (bool, error) {
			return req.Method == "GET" && req.URL.Path == deviceUrl, nil
		}).
		Persist().
		MatchHeader("Authorization", fmt.Sprintf("apiKey %s", API_KEY)).
		Reply(200).
		Type("application/json").
		BodyString(string(device))

	devicetypes, err := os.ReadFile("./mock/devicetypes.json")
	if err != nil {
		t.Fatalf("errore nel leggere il file: %v", err)
	}

	devicetypesUrl, _ := url.JoinPath("/projects", PROJECT_ID, "devicetypes")
	gock.New(API_URI).
		AddMatcher(func(req *http.Request, ereq *gock.Request) (bool, error) {
			return req.Method == "GET" && req.URL.Path == devicetypesUrl, nil
		}).
		Persist().
		MatchHeader("Authorization", fmt.Sprintf("apiKey %s", API_KEY)).
		Reply(200).
		Type("application/json").
		BodyString(string(devicetypes))

	devicetype, err := os.ReadFile("./mock/devicetype.json")
	if err != nil {
		t.Fatalf("errore nel leggere il file: %v", err)
	}

	devicetypeUrl, _ := url.JoinPath("/projects", PROJECT_ID, "devicetypes", DEVICE_TYPE_ID)
	gock.New(API_URI).
		AddMatcher(func(req *http.Request, ereq *gock.Request) (bool, error) {
			return req.Method == "GET" && req.URL.Path == devicetypeUrl, nil
		}).
		Persist().
		MatchHeader("Authorization", fmt.Sprintf("apiKey %s", API_KEY)).
		Reply(200).
		Type("application/json").
		BodyString(string(devicetype))

	commandUrl, _ := url.JoinPath("/projects", PROJECT_ID, "commands")
	gock.New(API_URI).
		AddMatcher(func(req *http.Request, ereq *gock.Request) (bool, error) {
			return req.Method == "POST" && req.URL.Path == commandUrl, nil
		}).
		Persist().
		MatchHeader("Authorization", fmt.Sprintf("apiKey %s", API_KEY)).
		Reply(200)

	// --- TESTS

	t.Run("Invalid API Key", func(t *testing.T) {
		platform := goplatform.New(goplatform.PlatformConfig{
			Uri:    API_URI,
			ApiKey: "invalid-api-key",
		})
		_, err := platform.GetProjects(context.Background())
		if err == nil {
			t.Fatal("expected error with invalid API key")
		}
	})

	t.Run("Empty API Key", func(t *testing.T) {
		platform := goplatform.New(goplatform.PlatformConfig{
			Uri:    API_URI,
			ApiKey: "",
		})
		_, err := platform.GetProjects(context.Background())
		if err == nil {
			t.Fatal("expected error with empty API key")
		}
	})

	t.Run("Non-existent Project", func(t *testing.T) {
		platform := goplatform.New(goplatform.PlatformConfig{
			Uri:    API_URI,
			ApiKey: API_KEY,
		})
		_, err := platform.GetProject(context.Background(), "non-existent-project")
		if err == nil {
			t.Fatal("expected error for non-existent project")
		}
	})

	t.Run("Non-existent Node", func(t *testing.T) {
		platform := goplatform.New(goplatform.PlatformConfig{
			Uri:    API_URI,
			ApiKey: API_KEY,
		})
		project, err := platform.GetProject(context.Background(), PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}
		_, err = project.GetNode(context.Background(), "non-existent-node")
		if err == nil {
			t.Fatal("expected error for non-existent node")
		}
	})

	t.Run("Non-existent Device", func(t *testing.T) {
		platform := goplatform.New(goplatform.PlatformConfig{
			Uri:    API_URI,
			ApiKey: API_KEY,
		})
		project, err := platform.GetProject(context.Background(), PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}
		_, err = project.GetDevice(context.Background(), "non-existent-device")
		if err == nil {
			t.Fatal("expected error for non-existent device")
		}
	})

	t.Run("Non-existent DeviceType", func(t *testing.T) {
		platform := goplatform.New(goplatform.PlatformConfig{
			Uri:    API_URI,
			ApiKey: API_KEY,
		})
		project, err := platform.GetProject(context.Background(), PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}
		_, err = project.GetDeviceType(context.Background(), "non-existent-devicetype")
		if err == nil {
			t.Fatal("expected error for non-existent device type")
		}
	})

	t.Run("Non-existent Rule", func(t *testing.T) {
		platform := goplatform.New(goplatform.PlatformConfig{
			Uri:    API_URI,
			ApiKey: API_KEY,
		})
		project, err := platform.GetProject(context.Background(), PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}
		_, err = project.GetRule(context.Background(), "non-existent-rule")
		if err == nil {
			t.Fatal("expected error for non-existent rule")
		}
	})

	t.Run("Context Cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		platform := goplatform.New(goplatform.PlatformConfig{
			Uri:    API_URI,
			ApiKey: API_KEY,
		})
		cancel()
		_, err := platform.GetProjects(ctx)
		if err == nil {
			t.Fatal("expected error due to cancelled context")
		}
	})

	t.Run("Context Timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
		defer cancel()
		platform := goplatform.New(goplatform.PlatformConfig{
			Uri:    API_URI,
			ApiKey: API_KEY,
		})
		_, err := platform.GetProjects(ctx)
		if err == nil {
			t.Fatal("expected error due to context timeout")
		}
	})

	t.Run("GetProjects", func(t *testing.T) {
		platform := goplatform.New(goplatform.PlatformConfig{
			Uri:    API_URI,
			ApiKey: API_KEY,
		})

		if _, err := platform.GetProjects(context.Background()); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("GetProject", func(t *testing.T) {
		platform := goplatform.New(goplatform.PlatformConfig{
			Uri:    API_URI,
			ApiKey: API_KEY,
		})
		project, err := platform.GetProject(context.Background(), PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}
		if project.Uuid != PROJECT_ID {
			t.Fatalf("expected project Uuid to be '%s', got '%s'", PROJECT_ID, project.Uuid)
		}
	})

	t.Run("GetNodes", func(t *testing.T) {
		platform := goplatform.New(goplatform.PlatformConfig{
			Uri:    API_URI,
			ApiKey: API_KEY,
		})

		project, err := platform.GetProject(context.Background(), PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}

		nodes, err := project.GetNodes(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		if len(nodes) == 0 {
			t.Fail()
		}
	})

	t.Run("GetNode", func(t *testing.T) {
		platform := goplatform.New(goplatform.PlatformConfig{
			Uri:    API_URI,
			ApiKey: API_KEY,
		})
		project, err := platform.GetProject(context.Background(), PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}
		node, err := project.GetNode(context.Background(), NODE_ID)
		if err != nil {
			t.Fatal(err)
		}
		if node.Uuid != NODE_ID {
			t.Fatalf("expected node Uuid to be '%s', got '%s'", NODE_ID, node.Uuid)
		}
	})

	t.Run("GetDevices", func(t *testing.T) {
		platform := goplatform.New(goplatform.PlatformConfig{
			Uri:    API_URI,
			ApiKey: API_KEY,
		})

		project, err := platform.GetProject(context.Background(), PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}

		devices, err := project.GetDevices(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		if len(devices) == 0 {
			t.Fail()
		}
	})

	t.Run("GetDevice", func(t *testing.T) {
		platform := goplatform.New(goplatform.PlatformConfig{
			Uri:    API_URI,
			ApiKey: API_KEY,
		})
		project, err := platform.GetProject(context.Background(), PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}
		device, err := project.GetDevice(context.Background(), DEVICE_ID)
		if err != nil {
			t.Fatal(err)
		}
		if device.Uuid != DEVICE_ID {
			t.Fatalf("expected device Uuid to be '%s', got '%s'", DEVICE_ID, device.Uuid)
		}
	})

	t.Run("GetDeviceTypes", func(t *testing.T) {
		platform := goplatform.New(goplatform.PlatformConfig{
			Uri:    API_URI,
			ApiKey: API_KEY,
		})

		project, err := platform.GetProject(context.Background(), PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}

		devices, err := project.GetDeviceTypes(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		if len(devices) == 0 {
			t.Fail()
		}
	})

	t.Run("GetDeviceType", func(t *testing.T) {
		platform := goplatform.New(goplatform.PlatformConfig{
			Uri:    API_URI,
			ApiKey: API_KEY,
		})
		project, err := platform.GetProject(context.Background(), PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}
		devicetype, err := project.GetDeviceType(context.Background(), DEVICE_TYPE_ID)
		if err != nil {
			t.Fatal(err)
		}
		if devicetype.Uuid != DEVICE_TYPE_ID {
			t.Fatalf("expected deviceType Uuid to be '%s', got '%s'", DEVICE_TYPE_ID, devicetype.Uuid)
		}
	})

	t.Run("CreateEvent", func(t *testing.T) {
		event := goplatform.Event{
			Description: "Evento di prova",
			Type:        "Notification",
			Source:      "test/event",
		}

		platform := goplatform.New(goplatform.PlatformConfig{
			Uri:    API_URI,
			ApiKey: API_KEY,
		})

		project, err := platform.GetProject(context.Background(), PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}

		if err := project.CreateEvent(context.Background(), event); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("GetRules", func(t *testing.T) {
		platform := goplatform.New(goplatform.PlatformConfig{
			Uri:    API_URI,
			ApiKey: API_KEY,
		})

		project, err := platform.GetProject(context.Background(), PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}

		if _, err := project.GetRules(context.Background()); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("GetRule", func(t *testing.T) {
		platform := goplatform.New(goplatform.PlatformConfig{
			Uri:    API_URI,
			ApiKey: API_KEY,
		})
		project, err := platform.GetProject(context.Background(), PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}
		rule, err := project.GetRule(context.Background(), RULE_ID)
		if err != nil {
			t.Fatal(err)
		}
		if rule.Uuid != RULE_ID {
			t.Fatalf("expected rule Uuid to be '%s', got '%s'", RULE_ID, rule.Uuid)
		}
	})

	t.Run("Run Command", func(t *testing.T) {
		ctx := context.Background()

		platform := goplatform.New(goplatform.PlatformConfig{
			Uri:    API_URI,
			ApiKey: API_KEY,
		})

		project, err := platform.GetProject(ctx, PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}

		deviceId := "my-device-id"
		downlinkMaxRetris := 120
		executionMaxRetris := 10

		req := goplatform.CommandRequest{
			Name:      "Sample",
			ProjectId: "my-project-id",
			DeviceId:  &deviceId,
			Parameters: goplatform.CommandParameters{
				goplatform.CommandParameter{
					"address": 1,
				},
			},
			DownlinkRetry: &goplatform.CommandRequestRetryOption{
				MaxRetries: &downlinkMaxRetris,
			},
			ExecutionRetry: &goplatform.CommandRequestRetryOption{
				MaxRetries: &executionMaxRetris,
			},
		}

		command := req.MakeCommand()

		if err := project.SendCommand(ctx, command); err != nil {
			t.Fatal(err)
		}
	})
}
