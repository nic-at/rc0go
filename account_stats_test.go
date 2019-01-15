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
	"strconv"
	"strings"
	"testing"
)

func TestAccountStatsService_TopMagnitude(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	want := []*TopMagnitude{
		{
			Domain:    String("testzone1.at"),
			Magnitude: Float32(4.2),
			ID:        Int(324234324),
		},
	}

	daysToTest := 30

	mux.HandleFunc(RC0AccStatsTopDNSMagnitude, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		days, ok := r.URL.Query()["days"]

		if !ok || len(days) != 1 {
			t.Error("Url Param 'days' is missing")
		}

		d := days[0]

		if eq := strings.Compare(d, strconv.Itoa(daysToTest)); eq != 0 {
			t.Errorf("Query param days is %v and not %v", d, daysToTest)
		}

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	topMagnitudes, err := client.AccStats.TopMagnitude(daysToTest)
	if err != nil {
		t.Errorf("AccStats.TopMagnitude returned error: %v", err)
	}

	if !reflect.DeepEqual(topMagnitudes, want) {
		t.Errorf("AccStats.TopMagnitude returned %+v, want %+v", topMagnitudes, want)
	}

}

func TestAccountStatsService_TopNXDomains(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	want := []*TopNXDomain{
		{
			ID: Int(213123),
			Domain: String("testzone1.at"),
			NXDomain: NXDomain{
				Type:String("A"),
				Name: String("nosuchlabel.testzone1.at"),
				Count: Int(2034),
			},
		},
	}

	daysToTest := 30

	mux.HandleFunc(RC0AccStatsTopNXDomains, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		days, ok := r.URL.Query()["days"]

		if !ok || len(days) != 1 {
			t.Error("Url Param 'days' is missing")
		}

		d := days[0]

		if eq := strings.Compare(d, strconv.Itoa(daysToTest)); eq != 0 {
			t.Errorf("Query param days is %v and not %v", d, daysToTest)
		}

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	topNXDomains, err := client.AccStats.TopNXDomains(daysToTest)
	if err != nil {
		t.Errorf("AccStats.TopNXDomains returned error: %v", err)
	}

	if !reflect.DeepEqual(topNXDomains, want) {
		t.Errorf("AccStats.TopNXDomains returned %+v, want %+v", topNXDomains, want)
	}

}

func TestAccountStatsService_TopQNames(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	want := []*TopQuery{
		{
			ID: Int(213123),
			Domain: String("testzone.at"),
			Query: Query{
				Name: String("www.testzone1.at."),
				Type: String("A"),
				Count: Int(2034555),
			},
		},
	}

	daysToTest := 30

	mux.HandleFunc(RC0AccStatsTopQNames, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		days, ok := r.URL.Query()["days"]

		if !ok || len(days) != 1 {
			t.Error("Url Param 'days' is missing")
		}

		d := days[0]

		if eq := strings.Compare(d, strconv.Itoa(daysToTest)); eq != 0 {
			t.Errorf("Query param days is %v and not %v", d, daysToTest)
		}

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	topQNames, err := client.AccStats.TopQNames(daysToTest)
	if err != nil {
		t.Errorf("AccStats.TopQNames returned error: %v", err)
	}

	if !reflect.DeepEqual(topQNames, want) {
		t.Errorf("AccStats.TopQNames returned %+v, want %+v", topQNames, want)
	}

}

func TestAccountStatsService_TopZones(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	want := []*TopZone{
		{
			ID: Int(324234324),
			Domain: String("testzone1.at"),
			Count: Int(2034),
		},
	}

	daysToTest := 30

	mux.HandleFunc(RC0AccStatsTopZones, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		days, ok := r.URL.Query()["days"]

		if !ok || len(days) != 1 {
			t.Error("Url Param 'days' is missing")
		}

		d := days[0]

		if eq := strings.Compare(d, strconv.Itoa(daysToTest)); eq != 0 {
			t.Errorf("Query param days is %v and not %v", d, daysToTest)
		}

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	topZones, err := client.AccStats.TopZones(daysToTest)
	if err != nil {
		t.Errorf("AccStats.TopZones returned error: %v", err)
	}

	if !reflect.DeepEqual(topZones, want) {
		t.Errorf("AccStats.TopZones returned %+v, want %+v", topZones, want)
	}

}

func TestAccountStatsService_TotalQueryCount(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	want := []*QueryCount{
		{
			Date: String("2018-02-24"),
			Count: Int(3213123),
			NXCount: Int(76642),
		},
	}

	daysToTest := 30

	mux.HandleFunc(RC0AccStatsQueries, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		days, ok := r.URL.Query()["days"]

		if !ok || len(days) != 1 {
			t.Error("Url Param 'days' is missing")
		}

		d := days[0]

		if eq := strings.Compare(d, strconv.Itoa(daysToTest)); eq != 0 {
			t.Errorf("Query param days is %v and not %v", d, daysToTest)
		}

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	queryCount, err := client.AccStats.TotalQueryCount(daysToTest)
	if err != nil {
		t.Errorf("AccStats.TotalQueryCount returned error: %v", err)
	}

	if !reflect.DeepEqual(queryCount, want) {
		t.Errorf("AccStats.TotalQueryCount returned %+v, want %+v", queryCount, want)
	}

}

func TestAccountStatsService_TotalQueryCountPerCountry(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	want := []*CountryQueryCount{
		{
			CountryCode: String("AT"),
			Country: String("Austria"),
			Region: String("Europe"),
			Subregion: String("Western Europe"),
			QueryCount: Int(10353087),
		},
	}

	daysToTest := 30

	mux.HandleFunc(RC0AccStatsCountries, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		days, ok := r.URL.Query()["days"]

		if !ok || len(days) != 1 {
			t.Error("Url Param 'days' is missing")
		}

		d := days[0]

		if eq := strings.Compare(d, strconv.Itoa(daysToTest)); eq != 0 {
			t.Errorf("Query param days is %v and not %v", d, daysToTest)
		}

		_json, _ := json.Marshal(want)
		_, _ = fmt.Fprint(w, string(_json))
	})

	countryQueryCount, err := client.AccStats.TotalQueryCountPerCountry(daysToTest)
	if err != nil {
		t.Errorf("AccStats.TotalQueryCountPerCountry returned error: %v", err)
	}

	if !reflect.DeepEqual(countryQueryCount, want) {
		t.Errorf("AccStats.TotalQueryCountPerCountry returned %+v, want %+v", countryQueryCount, want)
	}

}
