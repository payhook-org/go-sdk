package payhook

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	// ApiURL is the payhook api endpoint
	ApiURL = "https://api-dev.payhook.org"
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

	resp, err := api.makeRequest("createPayment", jsonBody, nil, &Payment{})
	if err != nil {
		return nil, err
	}

	return resp.(*Payment), nil
}

// GetPayment get payment where paymentID
func (api *API) GetPayment(paymentID uint64) (*Payment, error) {
	urlValues := map[string]string{
		"id": strconv.FormatUint(paymentID, 10),
	}

	resp, err := api.makeRequest("getPayment", nil, urlValues, &Payment{})
	if err != nil {
		return nil, err
	}

	return resp.(*Payment), nil
}

// DeletePayment delete payment from payhook where paymentID
func (api *API) DeletePayment(paymentID uint64) (*bool, error) {
	urlValues := map[string]string{
		"id": strconv.FormatUint(paymentID, 10),
	}

	var deleted bool
	resp, err := api.makeRequest("deletePayment", nil, urlValues, &deleted)
	if err != nil {
		return nil, err
	}

	return resp.(*bool), nil
}

// makeRequest wrapper for request
func (api *API) makeRequest(
	method string,
	body []byte,
	urlValues map[string]string,
	typ interface{},
) (interface{}, error) {
	reqURL := fmt.Sprintf("%s/%s/invoke/%s", ApiURL, ApiVersion, method)

	headers := map[string]string{
		"X-Payhook-Api-Key": api.key,
		"Accept":            "application/json",
		"Content-Type":      "application/json",
	}

	return api.request(reqURL, headers, urlValues, body, typ)
}

// request executes HTTP Request to the payhook and returns the result
func (api *API) request(
	reqURL string,
	headers, urlValues map[string]string,
	body []byte,
	typ interface{},
) (interface{}, error) {
	req, err := http.NewRequest(http.MethodPost, reqURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	for key, value := range urlValues {
		req.URL.Query().Add(key, value)
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
