package goplatform

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func errorFromResponse(resp *http.Response) error {
	if resp.StatusCode >= 400 {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("goplatform: ERR0: some error on %s %s", resp.Request.Method, resp.Request.URL)
		}

		var payload ResponseError
		if err := json.Unmarshal(b, &payload); err != nil {
			return fmt.Errorf("goplatform: ERR1: some error on %s %s", resp.Request.Method, resp.Request.URL)
		}

		return fmt.Errorf("goplatform: %s %s: %s: %s", resp.Request.Method, resp.Request.URL, payload.Error.Name, payload.Error.Message)
	}

	return nil
}
