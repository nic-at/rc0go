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

func TestMessageService_GetLatest(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	want := &Message{
		ID: 56007,
		Domain: "testzone2.at",
		Date: "2018-04-09T09:31:14Z",
		Type: "DSSEEN",
		Comment: "Simulate that the DS record has been seen in the parent zone.",
	}

	mux.HandleFunc(RC0Messages, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	message, err := client.Messages.GetLatest()
	if err != nil {
		t.Errorf("Messages.GetLatest returned error: %v", err)
	}

	if !reflect.DeepEqual(message, want) {
		t.Errorf("Messages.GetLatest returned %+v, want %+v", message, want)
	}

}

func TestMessageService_AckAndDelete(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	want := &StatusResponse{Status: "ok", Message: "Acknowledged notification '56007'"}

	mux.HandleFunc(RC0AckMessage, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	statusResponse, err := client.Messages.AckAndDelete(56007)
	if err != nil {
		t.Errorf("Messages.AckAndDelete returned error: %v", err)
	}

	if !reflect.DeepEqual(statusResponse, want) {
		t.Errorf("Messages.AckAndDelete returned %+v, want %+v", statusResponse, want)
	}

}