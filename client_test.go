// Copyright 2019 nic.at GmbH. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rc0go

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

const (
	baseURLPath = "/api"
)

func setup() (client *Client, _mux *mux.Router, serverURL string, teardown func()) {

	_mux = mux.NewRouter()

	apiHandler := http.NewServeMux()
	apiHandler.Handle(baseURLPath + "/" + defaultAPIVersion+"/", http.StripPrefix(baseURLPath+"/"+defaultAPIVersion, _mux))
	apiHandler.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, "Client.BaseURLPath prefix is not preserved in the dnssecRequest URL.", http.StatusInternalServerError)
	})

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(apiHandler)

	client, _ = NewClient("test123")

	_url, _ := url.Parse(server.URL + baseURLPath + "/")
	client.BaseURL = _url

	return client, _mux, server.URL, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func getTestDataPaginated(dataType reflect.Type) map[string]interface{} {

	data :=	map[string]interface{}{
				"current_page": 1,
				"from": 1,
				"last_page": 1,
				"next_page_url": "null",
				"path": "https://my.rcodezero.at/api/v1/zones",
				"per_page": 100,
				"prev_page_url": "null",
				"to": 2,
				"total": 2,
			}

	switch dataType {

	case reflect.TypeOf(Zone{}):
		data["data"] = []interface{}{getSampleZone()}
		break

	case reflect.TypeOf(RRType{}):
		data["data"] = []interface{}{getSampleRRSet()}
		break

	default:
		panic("no such type")

	}

	return data
}

func getSampleZone() map[string]interface{} {
	return map[string]interface{}{
		"domain": "testzone1.at",
		"type": "SLAVE",
		"dnssec": "yes",
		"created": "2018-04-09T09:27:31Z",
		"serial": "20180411",
		"masters": []string{
			"193.0.2.2",
			"2001:db8::2",
		},
	}
}

func getSampleRRSet() map[string]interface{} {
	return map[string]interface{}{
		"name": "www.testzone1.at.",
		"type": "A",
		"ttl": 3600,
		"records": []interface{}{
			map[string]interface{}{
				"content":  "10.10.0.2",
				"disabled": false,
			},
		},
	}
}