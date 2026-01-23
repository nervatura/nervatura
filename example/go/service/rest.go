package service

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"slices"

	cu "github.com/nervatura/component/pkg/util"
)

type HttpClient struct {
	Config cu.SM
}

func (rest *HttpClient) request(method, path, token string, options cu.IM) (any, error) {
	service_url := "http://localhost:" + rest.Config["NT_HTTP_PORT"] + "/api/v6/"
	data, err := cu.ConvertToByte(options)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, service_url+path, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if token != "" {
		request.Header.Set("Authorization", "Bearer "+token)
	} else {
		request.Header.Set("X-API-Key", rest.Config["NT_API_KEY"])
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.Header.Get("Content-Type") == "application/pdf" {
		responseStr := ""
		responseData, err := io.ReadAll(response.Body)
		if err == nil {
			responseStr = string(responseData)
		}
		return responseStr, err
	}

	var result interface{}
	if response.StatusCode == 200 {
		err = cu.ConvertFromReader(response.Body, &result)
		if err != nil {
			return nil, errors.New(response.Status)
		}
	}
	if !slices.Contains([]int64{200, 201, 204}, int64(response.StatusCode)) {
		return nil, errors.New(response.Status)
	}
	return result, err
}

func (rest *HttpClient) Get(token, path string, query cu.IM) (any, error) {
	queryParams := url.Values{}
	for key, value := range query {
		queryParams.Add(key, cu.ToString(value, ""))
	}
	if len(queryParams) > 0 {
		path += "?" + queryParams.Encode()
	}
	return rest.request("GET", path, token, cu.IM{})
}

func (rest *HttpClient) Post(token, path string, data cu.IM) (any, error) {
	return rest.request("POST", path, token, data)
}

func (rest *HttpClient) Put(token, path string, data cu.IM) (any, error) {
	return rest.request("PUT", path, token, data)
}

func (rest *HttpClient) Delete(token, path string) (any, error) {
	return rest.request("DELETE", path, token, cu.IM{})
}
