// Copyright 2019 nic.at GmbH. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rc0go

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
)

type RRSetService service

type RRType struct {
	Name    *string   `json:"name"`
	Rtype   *string   `json:"type"`
	TTL     *int      `json:"ttl"`
	Records *[]Record `json:"records"`
}

type Record struct {
	Content  *string `json:"content"`
	Disabled *bool   `json:"disabled"`
}

type RRSetEdit struct {
	Name 		*string   `json:"name"`
	Type 		*string   `json:"type"`
	ChangeType  *string   `json:"changetype"`
	Records 	[]*Record `json:"records"`
}

const (
	ChangeTypeADD 	 = "add"
	ChangeTypeUPDATE = "update"
	ChangeTypeDELETE = "delete"
)

// List all RRSets
//
// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-zone-management-rrsets-get
func (s *RRSetService) List(zone string) ([]*RRType, *Page, error) {

	resp, err := s.client.NewRequest().
		SetPathParams(
			map[string]string{
				"zone": zone,
			}).
		Get(
			s.client.BaseURL.String() +
				s.client.APIVersion +
				RC0ZoneRRSets,
		)

	if err != nil {
		return nil, nil, err
	}

	var page *Page

	err = json.Unmarshal(resp.Body(), &page)

	if err != nil {
		return nil, nil, err
	}

	var rrset []*RRType

	err = mapstructure.WeakDecode(page.Data, &rrset)
	if err != nil {
		return nil, nil, err
	}

	return rrset, page, nil
}

func (s *RRSetService) Create(zone string, rrsetCreate []*RRSetEdit) (*StatusResponse, error) {

	resp, err := s.client.NewRequest().
		SetPathParams(
			map[string]string{
				"zone": zone,
			}).
		SetBody(rrsetCreate).
		Patch(
			s.client.BaseURL.String() +
				s.client.APIVersion +
				RC0ZoneRRSets,
		)

	if err != nil {
		return nil, err
	}

	return s.client.ResponseToRC0StatusResponse(resp)
}


func (s *RRSetService) Edit(zone string, rrsetEdit []*RRSetEdit) (*StatusResponse, error) {

	resp, err := s.client.NewRequest().
		SetPathParams(
			map[string]string{
				"zone": zone,
			}).
		SetBody(rrsetEdit).
		Patch(
			s.client.BaseURL.String() +
				s.client.APIVersion +
				RC0ZoneRRSets,
		)

	if err != nil {
		return nil, err
	}

	return s.client.ResponseToRC0StatusResponse(resp)
}

func (s *RRSetService) Delete(zone string, rrsetDelete []*RRSetEdit) (*StatusResponse, error) {

	resp, err := s.client.NewRequest().
		SetPathParams(
			map[string]string{
				"zone": zone,
			}).
		SetBody(rrsetDelete).
		Patch(
			s.client.BaseURL.String() +
				s.client.APIVersion +
				RC0ZoneRRSets,
		)

	if err != nil {
		return nil, err
	}

	return s.client.ResponseToRC0StatusResponse(resp)
}

