// Copyright 2019 nic.at GmbH. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rc0go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestAccSettingsService_Get(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	want := &GlobalSetting{
		Secondaries: &[]string{
			"10.10.1.2",
		},
		TSIGOut: String("mystigkey,hmac-sha256,BqpFrSK+zsvYDJ0oXZzfs3R6VVxabW3RL4GLTM/fm2QGQbvDIUZHWVzNXbAEYOC77EZFC+B4RfrdLE6soeQKUw=="),
	}

	mux.HandleFunc(RC0AccSettings, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	globalSetting, err := client.Settings.Get()
	if err != nil {
		t.Errorf("Settings.Get returned error: %v", err)
	}

	if !reflect.DeepEqual(globalSetting, want) {
		t.Errorf("Settings.Get returned %+v, want %+v", globalSetting, want)
	}

}

func TestAccSettingsService_SetSecondaries(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	type secondaries struct {
		Secondaries *[]string `json:"secondaries"`
	}

	_secondaries := &secondaries{
		Secondaries: &[]string{
			"10.10.1.2",
		},
	}

	want := &StatusResponse{Status: String("ok"), Message: String("Setting secondaries successfully configured")}

	mux.HandleFunc(RC0AccSecondaries, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		body, err := ioutil.ReadAll(r.Body)

		var secondariesReceived *secondaries
		_ = json.Unmarshal(body, &secondariesReceived)

		if err != nil {
			t.Errorf("Error reading body: %v", err)
		}

		if !reflect.DeepEqual(secondariesReceived, _secondaries) {
			t.Errorf("Settings.SetSecondaries returned %+v, want %+v", secondariesReceived, _secondaries)
		}

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	status, err := client.Settings.SetSecondaries(*_secondaries.Secondaries)
	if err != nil {
		t.Errorf("Settings.SetSecondaries returned error: %v", err)
	}

	if !reflect.DeepEqual(status, want) {
		t.Errorf("Settings.SetSecondaries returned %+v, want %+v", status, want)
	}

}

func TestAccSettingsService_RemoveTSIG(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	want := &StatusResponse{Status: String("ok"), Message: String("Setting tsigout successfully deleted")}

	mux.HandleFunc(RC0AccTsigout, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	statusResponse, err := client.Settings.RemoveTSIG()
	if err != nil {
		t.Errorf("Settings.RemoveTSIG returned error: %v", err)
	}

	if !reflect.DeepEqual(statusResponse, want) {
		t.Errorf("Settings.RemoveTSIG returned %+v, want %+v", statusResponse, want)
	}

}

func TestAccSettingsService_SetTSIG(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	type tsigkey struct {
		TSIGKey *string `json:"tsigkey"`
	}

	_tsigkey := &tsigkey{
		TSIGKey: String("10.10.1.2"),
	}

	want := &StatusResponse{Status: String("ok"), Message: String("Setting tsigout successfully configured")}

	mux.HandleFunc(RC0AccTsigout, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		body, err := ioutil.ReadAll(r.Body)

		var tsigkeyReceived *tsigkey
		_ = json.Unmarshal(body, &tsigkeyReceived)

		if err != nil {
			t.Errorf("Error reading body: %v", err)
		}

		if !reflect.DeepEqual(tsigkeyReceived, _tsigkey) {
			t.Errorf("Settings.SetTSIG returned %+v, want %+v", tsigkeyReceived, _tsigkey)
		}

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	status, err := client.Settings.SetTSIG(*_tsigkey.TSIGKey)
	if err != nil {
		t.Errorf("Settings.SetTSIG returned error: %v", err)
	}

	if !reflect.DeepEqual(status, want) {
		t.Errorf("Settings.SetTSIG returned %+v, want %+v", status, want)
	}

}

func TestAccSettingsService_RemoveSecondaries(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	want := &StatusResponse{Status: String("ok"), Message: String("Setting secondaries successfully deleted")}

	mux.HandleFunc(RC0AccSecondaries, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	statusResponse, err := client.Settings.RemoveSecondaries()
	if err != nil {
		t.Errorf("Settings.RemoveSecondaries returned error: %v", err)
	}

	if !reflect.DeepEqual(statusResponse, want) {
		t.Errorf("Settings.RemoveSecondaries returned %+v, want %+v", statusResponse, want)
	}

}
