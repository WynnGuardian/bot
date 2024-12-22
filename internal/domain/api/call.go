package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type apiCall[T, U any] struct {
	endpoint string
	data     CallData
	body     *T
}

type CallData struct {
	Token string
	Host  string
}

func NewCall[T, U any](c CallData, endpoint string, body T) *apiCall[T, U] {
	endpoint = fmt.Sprintf("%s/%s", c.Host, endpoint)
	return &apiCall[T, U]{
		endpoint: endpoint,
		data:     c,
		body:     &body,
	}
}

func (c *apiCall[T, U]) encodeBody() ([]byte, error) {
	return json.Marshal(c.body)
}

func (c *apiCall[T, U]) Post() (*APIResponse[U], error) {
	encoded, err := c.encodeBody()
	if err != nil {
		return nil, err
	}
	r, err := http.NewRequest("POST", c.endpoint, bytes.NewBuffer(encoded))
	if err != nil {
		return nil, err
	}

	DefaultHeaders(r, c.data)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resp := APIResponse[U]{}
	derr := json.NewDecoder(res.Body).Decode(&resp)
	if derr != nil {
		return nil, derr
	}

	return &resp, nil
}
