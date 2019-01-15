// Copyright 2019 nic.at GmbH. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rc0go

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestRRSetService_List(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	dat, _ := json.Marshal(getTestDataPaginated(reflect.TypeOf(RRType{})))

	mux.HandleFunc(RC0ZoneRRSets, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, _ = fmt.Fprint(w, string(dat))
	})

	rrset, _, err := client.RRSet.List("testzone1.at")

	if err != nil {
		t.Errorf("RRSet.List returned error: %v", err)
	}

	var wantRRSet *RRType
	_ = mapstructure.WeakDecode(getSampleRRSet(), &wantRRSet)

	if rrsCount := len(rrset); rrsCount != 1 {
		t.Errorf("RRSet.List returned %v zones instead of 1", rrsCount)
	}

	if !reflect.DeepEqual(rrset[0], wantRRSet) {
		t.Errorf("RRSet.List returned %+v, want %+v", rrset[0], wantRRSet)
	}

	rrsGot := reflect.TypeOf(rrset[0])
	rrsType := reflect.TypeOf(&RRType{})

	if rrsGot != rrsType {
		t.Errorf("RRSet.List rrset[0] type is %v and not %v", rrsGot, rrsType)
	}
}

func TestRRSetService_Create(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	rrsetCreate := []*RRSetEdit{
			{
				Name: "www.testzone1.at.",
				Type: "A",
				ChangeType: ChangeTypeADD,
				Records: []*Record{
					{
						Content:"10.10.0.1",
					},
				},
			},
		}

	want := &StatusResponse{Status: "ok", Message: "RRsets updated"}

	mux.HandleFunc(RC0ZoneRRSets, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")

		body, err := ioutil.ReadAll(r.Body)

		var rrsetCreateReceived []*RRSetEdit
		_ = json.Unmarshal(body, &rrsetCreateReceived)

		if err != nil {
			t.Errorf("Error reading body: %v", err)
		}

		if !reflect.DeepEqual(rrsetCreateReceived, rrsetCreate) {
			t.Errorf("RRSet.Create returned %+v, want %+v", rrsetCreateReceived, rrsetCreate)
		}

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	status, err := client.RRSet.Create("testzone1.at", rrsetCreate)
	if err != nil {
		t.Errorf("RRSet.Create returned error: %v", err)
	}

	if !reflect.DeepEqual(status, want) {
		t.Errorf("RRSet.Create returned %+v, want %+v", status, want)
	}

}

func TestRRSetService_Edit(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	rrsetUpdate := []*RRSetEdit{
		{
			Name: "www.testzone1.at.",
			Type: "A",
			ChangeType: ChangeTypeUPDATE,
			Records: []*Record{
				{
					Content:"10.10.0.10",
				},
				{
					Content: "10.10.0.20",
					Disabled: true,
				},
			},
		},
	}

	want := &StatusResponse{Status: "ok", Message: "RRsets updated"}

	mux.HandleFunc(RC0ZoneRRSets, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")

		body, err := ioutil.ReadAll(r.Body)

		var rrsetUpdateReceived []*RRSetEdit
		_ = json.Unmarshal(body, &rrsetUpdateReceived)

		if err != nil {
			t.Errorf("Error reading body: %v", err)
		}

		if !reflect.DeepEqual(rrsetUpdateReceived, rrsetUpdate) {
			t.Errorf("RRSet.Edit returned %+v, want %+v", rrsetUpdateReceived, rrsetUpdate)
		}

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	status, err := client.RRSet.Edit("testzone1.at", rrsetUpdate)
	if err != nil {
		t.Errorf("RRSet.Edit returned error: %v", err)
	}

	if !reflect.DeepEqual(status, want) {
		t.Errorf("RRSet.Edit returned %+v, want %+v", status, want)
	}

}

func TestRRSetService_Delete(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	rrsetDelete := []*RRSetEdit{
		{
			Name: "www.testzone1.at.",
			Type: "A",
			ChangeType: ChangeTypeDELETE,
		},
	}

	want := &StatusResponse{Status: "ok", Message: "RRsets updated"}

	mux.HandleFunc(RC0ZoneRRSets, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")

		body, err := ioutil.ReadAll(r.Body)

		var rrsetDeleteReceived []*RRSetEdit
		_ = json.Unmarshal(body, &rrsetDeleteReceived)

		if err != nil {
			t.Errorf("Error reading body: %v", err)
		}

		if !reflect.DeepEqual(rrsetDeleteReceived, rrsetDelete) {
			t.Errorf("RRSet.Delete returned %+v, want %+v", rrsetDeleteReceived, rrsetDelete)
		}

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	status, err := client.RRSet.Delete("testzone1.at", rrsetDelete)
	if err != nil {
		t.Errorf("RRSet.Delete returned error: %v", err)
	}

	if !reflect.DeepEqual(status, want) {
		t.Errorf("RRSet.Delete returned %+v, want %+v", status, want)
	}

}