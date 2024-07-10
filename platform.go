package goplatform

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Platform struct {
	uri    string
	apiKey string
}

func New(uri, apiKey string) Platform {
	return Platform{
		uri:    uri,
		apiKey: apiKey,
	}
}

func (p Platform) makeRequest(method string, body io.Reader, path ...string) (*http.Request, error) {
	u, err := url.JoinPath(p.uri, path...)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(strings.ToUpper(method), u, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	if p.apiKey != "" {
		req.Header.Set("Authorization", fmt.Sprintf("apiKey %s", p.apiKey))
	}

	return req, nil
}

func (p Platform) GetProjects() ([]Project, error) {
	req, err := p.makeRequest("GET", nil, "projects")
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := errorFromResponse(resp); err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var projects Response[[]Project]
	if err := json.Unmarshal(b, &projects); err != nil {
		return nil, err
	}

	for i := 0; i < len(projects.Data); i++ {
		projects.Data[i].platformRef = &p
	}

	return projects.Data, nil
}

func (p Platform) GetProject(uuid string) (Project, error) {
	// var project Response[Project]

	// req, err := p.makeRequest("GET", nil, "projects", uuid)
	// if err != nil {
	// 	return project.Data, err
	// }

	// client := &http.Client{}

	// resp, err := client.Do(req)
	// if err != nil {
	// 	return project.Data, err
	// }
	// defer resp.Body.Close()

	// if err := errorFromResponse(resp); err != nil {
	// 	return project.Data, err
	// }

	// b, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return project.Data, err
	// }

	// if err := json.Unmarshal(b, &project); err != nil {
	// 	return project.Data, err
	// }

	// project.Data.platformRef = &p

	// return project.Data, nil

	return Project{
		platformRef: &p,
		Uuid:        uuid,
	}, nil
}
