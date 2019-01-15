// Copyright 2019 nic.at GmbH. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rc0go

import (
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL    = "https://my.rcodezero.at/api/"
	defaultAPIVersion = "v1"
	userAgent         = "rc0go"

	headerRateLimit     = "X-RateLimit-Limit"
	headerRateRemaining = "X-RateLimit-Remaining"
)

type Client struct {

	// Base URL for API requests. Defaults to the rcode0 dev API, but can be
	// set to a production or test domain. BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL

	// Version of Rcode0 API
	APIVersion string

	// API Token
	Token string

	// User agent used when communicating with the rcode0 API.
	UserAgent string

	// HTTP client used to communicate with the API.
	client *http.Client

	// Reuse a single struct instead of allocating one for each service on the heap.
	common service

	Zones  	  *ZoneManagementService
	RRSet  	  *RRSetService
	DNSSEC    *DNSSECService
	ZoneStats *ZoneStatsService
	AccStats  *AccountStatsService
	Reports   *ReportService
	Messages  *MessageService
	Settings  *AccSettingsService
}

type service struct {
	client *Client
}

type StatusResponse struct {
	Status  string `json:"status, omitempty"`
	Message string `json:"message, omitempty"`
}

type Page struct {
	Data        []interface{} `json:"data"`
	CurrentPage int           `json:"current_page, omitempty"`
	From        int           `json:"from, omitempty"`
	LastPage    int           `json:"last_page, omitempty"`
	NextPageURL string        `json:"next_page_url, omitempty"`
	Path        string        `json:"path, omitempty"`
	PerPage     int           `json:"per_page, omitempty"`
	PrevPageURL string        `json:"prev_page_url, omitempty"`
	To          int           `json:"to, omitempty"`
	Total       int           `json:"total, omitempty"`
}

// NewClient returns a new rcode0 API client.
func NewClient(token string) (*Client, error) {

	if strings.Compare(token, "") == 0 {
		return nil, fmt.Errorf("rcodezero API token is not provided")
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		BaseURL:    baseURL,
		APIVersion: defaultAPIVersion,
		Token:      token,
		UserAgent:  userAgent,
		client:     http.DefaultClient,
	}

	c.common.client = c
	c.Zones   		= (*ZoneManagementService)(&c.common)
	c.RRSet   		= (*RRSetService)(&c.common)
	c.DNSSEC    	= (*DNSSECService)(&c.common)
	c.ZoneStats 	= (*ZoneStatsService)(&c.common)
	c.AccStats 		= (*AccountStatsService)(&c.common)
	c.Settings		= (*AccSettingsService)(&c.common)
	c.Reports 		= (*ReportService)(&c.common)
	c.Messages 		= (*MessageService)(&c.common)

	return c, nil
}

// @todo
func (c *Client) NewRequest() *resty.Request {

	return resty.R().
		SetAuthToken(c.Token)
}

// @todo
func (c *Client) ResponseToRC0StatusResponse(response *resty.Response) (*StatusResponse, error) {
	var statusResponse *StatusResponse

	err := json.Unmarshal(response.Body(), &statusResponse)

	if err != nil {
		return nil, err
	}

	return statusResponse, nil
}