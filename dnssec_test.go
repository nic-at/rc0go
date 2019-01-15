// Copyright 2019 nic.at GmbH. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rc0go

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestDNSSECService_Sign(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	want := &StatusResponse{Status: String("ok"), Message: String("Zone testzone1.at signed successfully")}

	mux.HandleFunc(RC0ZoneDNSSecSign, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	status, err := client.DNSSEC.Sign("testzone1.at")
	if err != nil {
		t.Errorf("DNSSEC.Sign returned error: %v", err)
	}

	if !reflect.DeepEqual(status, want) {
		t.Errorf("DNSSEC.Sign returned %+v, want %+v", status, want)
	}
}

func TestDNSSECService_Unsign(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	want := &StatusResponse{Status: String("ok"), Message: String("Zone testzone1.at unsigned successfully")}

	mux.HandleFunc(RC0ZoneDNSSecUnsign, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	status, err := client.DNSSEC.Unsign("testzone1.at")
	if err != nil {
		t.Errorf("DNSSEC.Unsign returned error: %v", err)
	}

	if !reflect.DeepEqual(status, want) {
		t.Errorf("DNSSEC.Ubsign returned %+v, want %+v", status, want)
	}
}

func TestDNSSECService_KeyRollover(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	want := &StatusResponse{Status: String("ok"), Message: String("Key rollover started successfully.")}

	mux.HandleFunc(RC0ZoneDNSSecKeyRollover, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	status, err := client.DNSSEC.KeyRollover("testzone1.at")
	if err != nil {
		t.Errorf("DNSSEC.KeyRollover returned error: %v", err)
	}

	if !reflect.DeepEqual(status, want) {
		t.Errorf("DNSSEC.KeyRollover returned %+v, want %+v", status, want)
	}
}

func TestDNSSECService_DSUpdate(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	want := &StatusResponse{Status: String("ok"), Message: String("Acknowledged KSK for domain 'testzone1.at'.")}

	mux.HandleFunc(RC0ZoneDNSSecDSUpdate, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	status, err := client.DNSSEC.DSUpdate("testzone1.at")
	if err != nil {
		t.Errorf("DNSSEC.DSUpdate returned error: %v", err)
	}

	if !reflect.DeepEqual(status, want) {
		t.Errorf("DNSSEC.DSUpdate returned %+v, want %+v", status, want)
	}
}

func TestDNSSECService_SimulateDSSEENEvent(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	want := &StatusResponse{Status: String("ok"), Message: String("simulate ok: Simulated DSSSEN. Had to update 1 keys for zone 'testzone1.at'")}

	mux.HandleFunc(RC0ZoneDNSSecDSSEEN, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	status, err := client.DNSSEC.SimulateDSSEENEvent("testzone1.at")
	if err != nil {
		t.Errorf("DNSSEC.SimulateDSSEENEvent returned error: %v", err)
	}

	if !reflect.DeepEqual(status, want) {
		t.Errorf("DNSSEC.SimulateDSSEENEvent returned %+v, want %+v", status, want)
	}
}

func TestDNSSECService_SimulateDSREMOVEDEvent(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	want := &StatusResponse{Status: String("ok"), Message: String("simulate ok: Simulated DSREMOVED. Had to update 1 keys for zone 'testzone1.at'")}

	mux.HandleFunc(RC0ZoneDNSSecDSREMOVED, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	status, err := client.DNSSEC.SimulateDSREMOVEDEvent("testzone1.at")
	if err != nil {
		t.Errorf("DNSSEC.SimulateDSREMOVEDEvent returned error: %v", err)
	}

	if !reflect.DeepEqual(status, want) {
		t.Errorf("DNSSEC.SimulateDSREMOVEDEvent returned %+v, want %+v", status, want)
	}
}

