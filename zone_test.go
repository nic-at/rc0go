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

// This test is explicitly documented to understand all the others
func TestZoneManagementServiceList(t *testing.T) {

	// Prepare setup env
	client, mux, _, teardown := setup()
	defer teardown()

	// Get sample json (note: a sample zone object is within the sample paginated "response")
	dat, _ := json.Marshal(getTestDataPaginated(reflect.TypeOf(Zone{})))

	// Register route handler
	mux.HandleFunc(RC0Zones, func(w http.ResponseWriter, r *http.Request) {
		// Test HTTP method used is correct
		testMethod(t, r, "GET")

		// Send sample data as response
		_, _ = fmt.Fprint(w, string(dat))
	})

	// Here the library method should return the sample data
	zones, _, err := client.Zones.List(NewListOptions())

	if err != nil {
		t.Errorf("Zones.List returned error: %v", err)
	}

	// Here a sample zone object is constructed
	var wantZone *Zone
	_ = mapstructure.Decode(getSampleZone(), &wantZone)

	// And is compared to the one which the method returns
	if !reflect.DeepEqual(zones[0], wantZone) {
		t.Errorf("Zones.List returned %+v, want %+v", zones[0], wantZone)
	}

	// The method should have extracted the right amount of zones within the paginated response (actually 1)
	if len(zones) != 1 {
		t.Error("Zones.List returned 0 zones instead of 1")
	}

	zoneGot := reflect.TypeOf(zones[0])
	zoneType := reflect.TypeOf(&Zone{})

	// Check types are the same
	if zoneGot != zoneType {
		t.Errorf("Zones.List zone[0] type is %v and not %v", zoneGot, zoneType)
	}

}

func TestZoneManagementService_Get(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	sampleZone, _ := json.Marshal(getSampleZone())

	mux.HandleFunc(RC0Zone, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		_, _ = fmt.Fprint(w, string(sampleZone))
	})

	zone, err := client.Zones.Get("testzone1.at")

	if err != nil {
		t.Errorf("Zones.Get returned error: %v", err)
	}

	var wantZone *Zone
	_ = mapstructure.Decode(getSampleZone(), &wantZone)

	if !reflect.DeepEqual(zone, wantZone) {
		t.Errorf("Zones.Get returned %+v, want %+v", zone, wantZone)
	}

}


func TestZoneManagementService_Create(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	zoneCreate := &ZoneCreate{
		Domain: "testzone1.at",
		Type: "slave",
		Masters: []string{
			"193.0.2.2",
			"2001:db8::2",
		},
	}

	want := &StatusResponse{Status: "ok", Message: "Zone testzone1.at successfully added"}

	mux.HandleFunc(RC0Zones, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		body, err := ioutil.ReadAll(r.Body)

		var zoneCreateReceived *ZoneCreate
		_ = json.Unmarshal(body, &zoneCreateReceived)

		if err != nil {
			t.Errorf("Error reading body: %v", err)
		}

		if !reflect.DeepEqual(zoneCreateReceived, zoneCreate) {
			t.Errorf("Zone.Create returned %+v, want %+v", zoneCreateReceived, zoneCreate)
		}

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	status, err := client.Zones.Create(zoneCreate)
	if err != nil {
		t.Errorf("Zones.Create returned error: %v", err)
	}

	if !reflect.DeepEqual(status, want) {
		t.Errorf("Zones.Create returned %+v, want %+v", status, want)
	}

}

func TestZoneManagementService_Edit(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	zoneEdit := &ZoneEdit{
		Type: "slave",
		Masters: []string{
			"193.0.2.2",
			"2001:db8::2",
		},
	}

	want := &StatusResponse{Status: "ok", Message: "Zone testzone1.at successfully updated"}

	mux.HandleFunc(RC0Zone, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		body, err := ioutil.ReadAll(r.Body)

		var zoneEditReceived *ZoneEdit
		_ = json.Unmarshal(body, &zoneEditReceived)

		if err != nil {
			t.Errorf("Error reading body: %v", err)
		}

		if !reflect.DeepEqual(zoneEditReceived, zoneEdit) {
			t.Errorf("Zone.Edit returned %+v, want %+v", zoneEditReceived, zoneEdit)
		}

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	status, err := client.Zones.Edit("testzone1.at", zoneEdit)
	if err != nil {
		t.Errorf("Zones.Edit returned error: %v", err)
	}

	if !reflect.DeepEqual(status, want) {
		t.Errorf("Zones.Edit returned %+v, want %+v", status, want)
	}

}

func TestZoneManagementService_Delete(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	want := &StatusResponse{Status: "ok", Message: "Zone testzone1.at successfully removed"}

	mux.HandleFunc(RC0Zone, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	status, err := client.Zones.Delete("testzone1.at")
	if err != nil {
		t.Errorf("Zones.Delete returned error: %v", err)
	}

	if !reflect.DeepEqual(status, want) {
		t.Errorf("Zones.Delete returned %+v, want %+v", status, want)
	}

}

func TestZoneManagementService_Transfer(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	want := &StatusResponse{Status: "ok", Message: "Zonetransfer for zone testzone1.at queued"}

	mux.HandleFunc(RC0ZoneTransfer, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	status, err := client.Zones.Transfer("testzone1.at")
	if err != nil {
		t.Errorf("Zones.Transfer returned error: %v", err)
	}

	if !reflect.DeepEqual(status, want) {
		t.Errorf("Zones.Transfer returned %+v, want %+v", status, want)
	}

}