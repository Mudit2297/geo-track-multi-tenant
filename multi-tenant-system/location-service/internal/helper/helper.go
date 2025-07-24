package helper

import (
	"fmt"
	"io"
	"net/http"
)

func ExecHttpRequest(method, url, token string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("unable to create http request: %w", err)
	}

	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("unable to execute http request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("unable to read http request request data: %w", err)
	}

	if res.StatusCode != 200 {
		return []byte{}, fmt.Errorf("request not authorized. status code: %v - %v", res.StatusCode, string(body))
	}

	return body, nil
}
