package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func ExecRequest(reqType, url, method string, data interface{}, token string) (*http.Response, []byte, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, []byte{}, err
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, []byte{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, []byte{}, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, []byte{}, err
	}

	return resp, respBody, nil
}
