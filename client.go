package payhook

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// ApiURL is the payhook api endpoint
	ApiURL = "https://api.payhook.org"
	// ApiVersion is the payhook api version
	ApiVersion = "v1"
	// ApiUserAgent identifies this library
	ApiUserAgent = "Go-Payhook API Agent (https://github.com/payhook-org/go-payhook)"
)

// API represents a Payhook API Client connection
type API struct {
	key    string
	client *http.Client
}

// New create a new API client
func New(apiKey string) *API {
	return &API{
		key:    apiKey,
		client: http.DefaultClient,
	}
}

// NewWithClient create a new API client with custom http client
func NewWithClient(key string, client *http.Client) *API {
	payhook := New(key)
	return payhook.WithClient(client)
}

// WithClient add custom http client into the API
func (api *API) WithClient(client *http.Client) *API {
	api.client = client
	return api
}

// CreatePayment create new payment
func (api *API) CreatePayment(params CreatePaymentParams) (*Payment, error) {
	if params.Title == "" {
		return nil, errors.New("payment title is not set")
	}
	if params.Currency == "" {
		return nil, errors.New("payment currency is not set")
	}
	if params.Amount == "" {
		return nil, errors.New("payment amount is not set")
	}

	jsonBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	resp, err := api.makeRequest("createPayment", jsonBody, &Payment{})
	if err != nil {
		return nil, err
	}

	return resp.(*Payment), nil
}

// GetPayment get payment where paymentID
func (api *API) GetPayment(paymentID uint64) (*Payment, error) {
	request := struct {
		ID uint64 `json:"id"`
	}{
		ID: paymentID,
	}

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := api.makeRequest("getPayment", jsonBody, &Payment{})
	if err != nil {
		return nil, err
	}

	return resp.(*Payment), nil
}

// DeletePayment delete payment from payhook where paymentID
func (api *API) DeletePayment(paymentID uint64) (*bool, error) {
	request := struct {
		ID uint64 `json:"id"`
	}{
		ID: paymentID,
	}

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	var deleted bool
	resp, err := api.makeRequest("deletePayment", jsonBody, &deleted)
	if err != nil {
		return nil, err
	}

	return resp.(*bool), nil
}

// makeRequest wrapper for request
func (api *API) makeRequest(
	method string,
	body []byte,
	typ interface{},
) (interface{}, error) {
	reqURL := fmt.Sprintf("%s/%s/invoke/%s", ApiURL, ApiVersion, method)

	headers := map[string]string{
		"X-Payhook-Api-Key": api.key,
		"Accept":            "application/json",
		"Content-Type":      "application/json",
	}

	return api.request(reqURL, headers, body, typ)
}

// request executes HTTP Request to the payhook and returns the result
func (api *API) request(
	reqURL string,
	headers map[string]string,
	body []byte,
	typ interface{},
) (interface{}, error) {
	req, err := http.NewRequest(http.MethodPost, reqURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", ApiUserAgent)
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := api.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var jsonData Response
	if typ != nil {
		jsonData.Result = typ
	}

	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, err
	}

	if jsonData.Message != nil {
		return nil, errors.New(*jsonData.Message)
	}

	return jsonData.Result, nil
}
