package goplatform_test

import (
	"context"
	"os"
	"testing"

	"github.com/ApioIoT/goplatform"
	"github.com/joho/godotenv"
)

func TestUnit(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatal(err)
	}

	URI := os.Getenv("URI")
	API_KEY := os.Getenv("API_KEY")
	PROJECT_ID := os.Getenv("PROJECT_ID")
	NODE_ID := os.Getenv("NODE_ID")
	DEVICE_ID := os.Getenv("DEVICE_ID")
	DEVICE_TYPE_ID := os.Getenv("DEVICE_TYPE_ID")
	RULE_ID := os.Getenv("RULE_ID")

	t.Run("GetProjects", func(t *testing.T) {
		platform := goplatform.New(context.Background(), URI, API_KEY)

		if _, err := platform.GetProjects(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("GetProject", func(t *testing.T) {
		platform := goplatform.New(context.Background(), URI, API_KEY)

		if _, err := platform.GetProject(PROJECT_ID); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("GetNodes", func(t *testing.T) {
		platform := goplatform.New(context.Background(), URI, API_KEY)

		project, err := platform.GetProject(PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}

		nodes, err := project.GetNodes()
		if err != nil {
			t.Fatal(err)
		}

		if len(nodes) == 0 {
			t.Fail()
		}
	})

	t.Run("GetNode", func(t *testing.T) {
		platform := goplatform.New(context.Background(), URI, API_KEY)

		project, err := platform.GetProject(PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}

		if _, err := project.GetNode(NODE_ID); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("GetDevices", func(t *testing.T) {
		platform := goplatform.New(context.Background(), URI, API_KEY)

		project, err := platform.GetProject(PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}

		devices, err := project.GetDevices()
		if err != nil {
			t.Fatal(err)
		}

		if len(devices) == 0 {
			t.Fail()
		}
	})

	t.Run("GetDevice", func(t *testing.T) {
		platform := goplatform.New(context.Background(), URI, API_KEY)

		project, err := platform.GetProject(PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}

		if _, err := project.GetDevice(DEVICE_ID); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("GetDeviceTypes", func(t *testing.T) {
		platform := goplatform.New(context.Background(), URI, API_KEY)

		project, err := platform.GetProject(PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}

		devices, err := project.GetDeviceTypes()
		if err != nil {
			t.Fatal(err)
		}

		if len(devices) == 0 {
			t.Fail()
		}
	})

	t.Run("GetDeviceType", func(t *testing.T) {
		platform := goplatform.New(context.Background(), URI, API_KEY)

		project, err := platform.GetProject(PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}

		if _, err := project.GetDeviceType(DEVICE_TYPE_ID); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("CreateEvent", func(t *testing.T) {
		event := goplatform.Event{
			Description: "Evento di prova",
			Type:        "Notification",
			Source:      "test/event",
		}

		platform := goplatform.New(context.Background(), URI, API_KEY)

		project, err := platform.GetProject(PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}

		if err := project.CreateEvent(event); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("GetRules", func(t *testing.T) {
		platform := goplatform.New(context.Background(), URI, API_KEY)

		project, err := platform.GetProject(PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}

		if _, err := project.GetRules(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("GetRule", func(t *testing.T) {
		platform := goplatform.New(context.Background(), URI, API_KEY)

		project, err := platform.GetProject(PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}

		if _, err := project.GetRule(RULE_ID); err != nil {
			t.Fatal(err)
		}
	})
}
