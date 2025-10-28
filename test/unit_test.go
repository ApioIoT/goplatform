package goplatform_test

import (
	"context"
	"os"
	"testing"
	"time"

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

	// Test di errori di autenticazione
	t.Run("Authentication Errors", func(t *testing.T) {
		t.Run("Invalid API Key", func(t *testing.T) {
			platform := goplatform.New(context.Background(), URI, "invalid-api-key")
			_, err := platform.GetProjects()
			if err == nil {
				t.Fatal("expected error with invalid API key")
			}
		})

		t.Run("Empty API Key", func(t *testing.T) {
			platform := goplatform.New(context.Background(), URI, "")
			_, err := platform.GetProjects()
			if err == nil {
				t.Fatal("expected error with empty API key")
			}
		})
	})

	// Test di gestione errori per progetti non esistenti
	t.Run("Non-existent Resources", func(t *testing.T) {
		platform := goplatform.New(context.Background(), URI, API_KEY)

		t.Run("Non-existent Project", func(t *testing.T) {
			_, err := platform.GetProject("non-existent-project")
			if err == nil {
				t.Fatal("expected error for non-existent project")
			}
		})

		project, err := platform.GetProject(PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}

		t.Run("Non-existent Node", func(t *testing.T) {
			_, err := project.GetNode("non-existent-node")
			if err == nil {
				t.Fatal("expected error for non-existent node")
			}
		})

		t.Run("Non-existent Device", func(t *testing.T) {
			_, err := project.GetDevice("non-existent-device")
			if err == nil {
				t.Fatal("expected error for non-existent device")
			}
		})

		t.Run("Non-existent DeviceType", func(t *testing.T) {
			_, err := project.GetDeviceType("non-existent-devicetype")
			if err == nil {
				t.Fatal("expected error for non-existent device type")
			}
		})

		t.Run("Non-existent Rule", func(t *testing.T) {
			_, err := project.GetRule("non-existent-rule")
			if err == nil {
				t.Fatal("expected error for non-existent rule")
			}
		})
	})

	// Test di validazione dei parametri
	t.Run("Parameter Validation", func(t *testing.T) {
		platform := goplatform.New(context.Background(), URI, API_KEY)
		project, err := platform.GetProject(PROJECT_ID)
		if err != nil {
			t.Fatal(err)
		}

		t.Run("Invalid Event Creation", func(t *testing.T) {
			invalidEvent := goplatform.Event{
				// Mancano i campi obbligatori Type e Source
				Description: "Test event",
			}
			err := project.CreateEvent(invalidEvent)
			if err == nil {
				t.Fatal("expected error for invalid event")
			}
		})
	})

	// Test di timeout e gestione del contesto
	t.Run("Context and Timeout Handling", func(t *testing.T) {
		t.Run("Context Cancellation", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			platform := goplatform.New(ctx, URI, API_KEY)

			// Annulla il contesto immediatamente
			cancel()

			_, err := platform.GetProjects()
			if err == nil {
				t.Fatal("expected error due to cancelled context")
			}
		})

		t.Run("Context Timeout", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
			defer cancel()

			platform := goplatform.New(ctx, URI, API_KEY)
			_, err := platform.GetProjects()
			if err == nil {
				t.Fatal("expected error due to context timeout")
			}
		})
	})

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
