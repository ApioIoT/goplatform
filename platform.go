package goplatform

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Platform struct {
	uri        string
	apiKey     string
	skipVerify bool
	ctx        context.Context
}

func New(ctx context.Context, uri, apiKey string) Platform {
	return Platform{
		uri:        uri,
		apiKey:     apiKey,
		skipVerify: false,
		ctx:        ctx,
	}
}

func (p *Platform) SetSkipVerify(value bool) {
	p.skipVerify = value
}

func (p Platform) fetch(method string, body io.Reader, path ...string) ([]byte, error) {
	u, err := url.JoinPath(p.uri, path...)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(p.ctx, strings.ToUpper(method), u, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	if p.apiKey != "" {
		req.Header.Set("Authorization", fmt.Sprintf("apiKey %s", p.apiKey))
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: p.skipVerify},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //nolint:errcheck

	if resp.StatusCode >= 400 {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("goplatform: ERR0: some error on %s %s", resp.Request.Method, resp.Request.URL)
		}

		var payload ResponseError
		if err := json.Unmarshal(b, &payload); err != nil {
			return nil, fmt.Errorf("goplatform: ERR1: some error on %s %s", resp.Request.Method, resp.Request.URL)
		}

		return nil, fmt.Errorf("goplatform: %s %s: %s: %s", resp.Request.Method, resp.Request.URL, payload.Error.Name, payload.Error.Message)
	}

	b, err := io.ReadAll(resp.Body)
	return b, err
}

func (p Platform) GetProjects() ([]Project, error) {
	b, err := p.fetch("GET", nil, "projects")
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
