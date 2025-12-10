package goplatform

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/h2non/gock"
)

type httpMethod string

const (
	httpGet    httpMethod = "GET"
	httpPost   httpMethod = "POST"
	httpPut    httpMethod = "PUT"
	httpDelete httpMethod = "DELETE"
)

type Platform struct {
	uri    string
	apiKey string
	client *http.Client
}

type Config struct {
	Uri        string
	ApiKey     string
	SkipVerify bool
}

func New(config Config) Platform {
	client := &http.Client{}

	if config.SkipVerify {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: config.SkipVerify},
		}
	}

	gock.InterceptClient(client) // For test

	return Platform{
		uri:    config.Uri,
		apiKey: config.ApiKey,
		client: client,
	}
}

func (p Platform) fetch(ctx context.Context, method httpMethod, body io.Reader, path ...string) ([]byte, error) {
	baseUri, err := url.Parse(p.uri)
	if err != nil {
		return nil, err
	}

	fullPath, err := url.JoinPath(baseUri.Path, path...)
	if err != nil {
		return nil, err
	}
	baseUri.Path = fullPath

	req, err := http.NewRequestWithContext(ctx, string(method), baseUri.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if p.apiKey != "" {
		req.Header.Set("Authorization", fmt.Sprintf("apiKey %s", p.apiKey))
	}

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //nolint:errcheck

	if resp.StatusCode >= 400 {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("goplatform: ERR0: some error on %s %s", resp.Request.Method, resp.Request.URL)
		}

		var payload responseError
		if err := json.Unmarshal(b, &payload); err != nil {
			return nil, fmt.Errorf("goplatform: ERR1: some error on %s %s", resp.Request.Method, resp.Request.URL)
		}

		return nil, fmt.Errorf("goplatform: %s %s: %s: %s", resp.Request.Method, resp.Request.URL, payload.Error.Name, payload.Error.Message)
	}

	b, err := io.ReadAll(resp.Body)
	return b, err
}

func (p Platform) GetProjects(ctx context.Context) ([]Project, error) {
	b, err := p.fetch(ctx, httpGet, nil, "projects")
	if err != nil {
		return nil, err
	}

	var projects response[[]Project]
	if err := json.Unmarshal(b, &projects); err != nil {
		return nil, err
	}

	for i := 0; i < len(projects.Data); i++ {
		projects.Data[i].platformRef = &p
	}

	return projects.Data, nil
}

func (p Platform) GetProject(ctx context.Context, uuid string) (Project, error) {
	b, err := p.fetch(ctx, httpGet, nil, "projects", uuid, "/")
	if err != nil {
		var zero Project
		return zero, err
	}

	var project response[Project]
	if err := json.Unmarshal(b, &project); err != nil {
		var zero Project
		return zero, err
	}

	project.Data.platformRef = &p

	return project.Data, nil
}
