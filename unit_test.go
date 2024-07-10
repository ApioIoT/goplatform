package goplatform_test

import (
	"log"
	"os"
	"testing"

	"github.com/ApioIoT/goplatform"
	"github.com/joho/godotenv"
)

var (
	URI        string
	API_KEY    string
	PROJECT_ID string
	DEVICE_ID  string
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	URI = os.Getenv("URI")
	API_KEY = os.Getenv("API_KEY")
	PROJECT_ID = os.Getenv("PROJECT_ID")
	DEVICE_ID = os.Getenv("DEVICE_ID")

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestGetProjects(t *testing.T) {
	platform := goplatform.New(URI, API_KEY)
	projects, err := platform.GetProjects()
	if err != nil {
		t.Fatal(err)
	}

	if len(projects) == 0 {
		t.Fail()
	}
}

func TestGetProject(t *testing.T) {
	platform := goplatform.New(URI, API_KEY)
	if _, err := platform.GetProject(PROJECT_ID); err != nil {
		t.Fatal(err)
	}
}

func TestGetDevices(t *testing.T) {
	platform := goplatform.New(URI, API_KEY)
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
}

func TestGetDevice(t *testing.T) {
	platform := goplatform.New(URI, API_KEY)
	project, err := platform.GetProject(PROJECT_ID)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := project.GetDevice(DEVICE_ID); err != nil {
		t.Fatal(err)
	}
}

func TestCreateEvent(t *testing.T) {
	platform := goplatform.New(URI, API_KEY)
	project, err := platform.GetProject(PROJECT_ID)
	if err != nil {
		t.Fatal(err)
	}

	event := goplatform.Event{
		Description: "Evento di prova",
		Type:        "Notification",
		Source:      "test/event",
	}

	if err := project.CreateEvent(event); err != nil {
		t.Fatal(err)
	}
}
