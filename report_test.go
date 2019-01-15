// Copyright 2019 nic.at GmbH. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rc0go

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"reflect"
	"testing"
)

func TestReportService_ProblematicZones(t *testing.T) {

	client, mux, _, teardown := setup()
	defer teardown()

	dat, _ := json.Marshal(getTestDataPaginated(reflect.TypeOf(Zone{})))

	mux.HandleFunc(RC0ReportsProblematiczones, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, _ = fmt.Fprint(w, string(dat))
	})

	probZones, _, err := client.Reports.ProblematicZones()

	if err != nil {
		t.Errorf("RRSet.List returned error: %v", err)
	}

	var wantProbZone *ProbZone
	_ = mapstructure.Decode(getSampleZone(), &wantProbZone)

	if pzCount := len(probZones); pzCount != 1 {
		t.Errorf("Reports.ProblematicZones returned %v zones instead of 1", pzCount)
	}

	if !reflect.DeepEqual(probZones[0], wantProbZone) {
		t.Errorf("Reports.ProblematicZones returned %+v, want %+v", probZones[0], wantProbZone)
	}

	pzGot := reflect.TypeOf(probZones[0])
	pzType := reflect.TypeOf(&ProbZone{})

	if pzGot != pzType {
		t.Errorf("Reports.ProblematicZones probZones[0] type is %v and not %v", pzGot, pzType)
	}

}
