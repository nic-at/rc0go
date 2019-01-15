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

func TestZoneStatsService_Magnitude(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	wantMagnitude := &Magnitude{
		Magnitude: String("4.2"),
		Date: String("2018-3-2"),
	}

	_json, _ := json.Marshal([]interface{}{wantMagnitude})

	mux.HandleFunc(RC0ZoneStatsMagnitude, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		_, _ = fmt.Fprint(w, string(_json))
	})

	gotMagnitudes, err := client.ZoneStats.Magnitude("testzone1.at")

	if err != nil {
		t.Errorf("ZoneStats.Magnitude returned error: %v", err)
	}

	if magCount := len(gotMagnitudes); magCount != 1 {
		t.Errorf("ZoneStats.Magnitude returned %v magnitudes instead of 1", magCount)
	}

	if !reflect.DeepEqual(gotMagnitudes[0], wantMagnitude) {
		t.Errorf("ZoneStats.Magnitude returned %+v, want %+v", gotMagnitudes[0], wantMagnitude)
	}

}

func TestZoneStatsService_NXDomains(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	wantNXDomains := &NXDomain{
		Type: String("A"),
		Count: Int(2034555),
		Name: String("wwww.testzone1.at."),
	}

	_json, _ := json.Marshal([]interface{}{wantNXDomains})

	mux.HandleFunc(RC0ZoneStatsNXDomains, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		_, _ = fmt.Fprint(w, string(_json))
	})

	gotNXDomains, err := client.ZoneStats.NXDomains("testzone1.at")

	if err != nil {
		t.Errorf("ZoneStats.NXDomains returned error: %v", err)
	}

	if nxdCount := len(gotNXDomains); nxdCount != 1 {
		t.Errorf("ZoneStats.NXDomains returned %v nxdomains instead of 1", nxdCount)
	}

	if !reflect.DeepEqual(gotNXDomains[0], wantNXDomains) {
		t.Errorf("ZoneStats.NXDomains returned %+v, want %+v", gotNXDomains[0], wantNXDomains)
	}

}

func TestZoneStatsService_QNames(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	wantQName := &Query{
		Name: String("wwww.testzone1.at."),
		Type: String("A"),
		Count: Int(2034555),
	}

	_json, _ := json.Marshal([]interface{}{wantQName})

	mux.HandleFunc(RC0ZoneStatsQNames, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		_, _ = fmt.Fprint(w, string(_json))
	})

	gotQNames, err := client.ZoneStats.QNames("testzone1.at")

	if err != nil {
		t.Errorf("ZoneStats.QNames returned error: %v", err)
	}

	if qnCount := len(gotQNames); qnCount != 1 {
		t.Errorf("ZoneStats.QNames returned %v qnames instead of 1", qnCount)
	}

	if !reflect.DeepEqual(gotQNames[0], wantQName) {
		t.Errorf("ZoneStats.QNames returned %+v, want %+v", gotQNames[0], wantQName)
	}

}

func TestZoneStatsService_Queries(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	wantQueries := &PerDay{
		Date: String("2018-3-25"),
		Queries: Int(312355),
		NXDomains: Int(2132),
	}

	_json, _ := json.Marshal([]interface{}{wantQueries})

	mux.HandleFunc(RC0ZoneStatsQueries, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		_, _ = fmt.Fprint(w, string(_json))
	})

	gotQueries, err := client.ZoneStats.Queries("testzone1.at")

	if err != nil {
		t.Errorf("ZoneStats.Queries returned error: %v", err)
	}

	if qCount := len(gotQueries); qCount != 1 {
		t.Errorf("ZoneStats.Queries returned %v queries instead of 1", qCount)
	}

	if !reflect.DeepEqual(gotQueries[0], wantQueries) {
		t.Errorf("ZoneStats.Queries returned %+v, want %+v", gotQueries[0], wantQueries)
	}

}