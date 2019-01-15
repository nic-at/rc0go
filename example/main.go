// Copyright 2019 nic.at GmbH. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// These are some examples of how the client-lib can be used
// For great experience just run ():
//
//     $ RC0_API_KEY=YOUR_API_KEY go run main.go
//
package main

import (
	"log"
	"net/url"
	"os"
	"reflect"
	"strings"

	"github.com/nic-at/rc0go"
)

func main() {

	rc0client, err := rc0go.NewClient(os.Getenv("RC0_API_KEY"))

	if strings.Contains(os.Getenv("RC0_BASE_URL"), "rcodezero.at/api/") {
		rc0client.BaseURL, _ = url.Parse(os.Getenv("RC0_BASE_URL"))
	}

	if err != nil {
		log.Fatalf("failed to initialize rcodezero provider: %v", err)
	}

	// List all managed zones
	zones, _, err := rc0client.Zones.List()

	if err != nil {
		log.Fatalf("failed to list zones: %v", err)
	}

	log.Println("rc0client.Zones.List:")
	log.Println()
	log.Println("Type: \t\t\t", 		reflect.TypeOf(zones))
	log.Println("Domaincount: \t", 	len(zones))
	log.Println("First Domain:\t", 	zones[0].Domain)
	log.Println()

	// Get a single managed zone by name
	zone, err := rc0client.Zones.Get("golib.at")

	if err != nil {
		log.Fatalf("failed to get 'golib.at': %v", err)
	}

	log.Println("rc0client.Zones.Get:")
	log.Println()
	log.Println("TypeOf: \t\t", 		reflect.TypeOf(zone))
	log.Println("Domain: \t\t", 		zone.Domain)
	log.Println("Type: \t\t\t", 		zone.Type)
	log.Println("ID: \t\t\t", 		zone.ID)
	log.Println("DNSSECStatus: \t", 	zone.DNSSECStatus)
	log.Println("Serial: \t\t", 		zone.Serial)
	log.Println()

	// Add a new zone to rcode0
	zoneCreateRequest := &rc0go.ZoneCreate{
		Type: 		"master",
		Domain: 	"golib-example.at",
		Masters: 	[]string{},
	}

	statusResponse, err := rc0client.Zones.Create(zoneCreateRequest)

	if err != nil {
		log.Fatalf("failed to add 'golib-example.at' to rcodezero: %v", err)
	}

	log.Println("rc0client.Zones.Create:")
	log.Println()
	log.Println("Type: \t\t\t", 		 reflect.TypeOf(statusResponse))
	log.Println("Status: \t\t", 		 statusResponse.Status)
	log.Println("Status Message:\t",	 statusResponse.Message)
	log.Println()

	// Add few A RRs to the newly added zone
	var records []*rc0go.Record

	record1 		   := &rc0go.Record{
		Content:  "127.0.0.1",
		Disabled: false,
	}

	records = append(records, record1)

	record2 		   := &rc0go.Record{
		Content:  "127.0.0.2",
		Disabled: false,
	}

	records = append(records, record2)

	rrsetCreate := []*rc0go.RRSetEdit{{
		Type: 		"A",
		Name: 		"www.golib-example.at.",
		ChangeType: rc0go.ChangeTypeADD,
		Records:  	records,
	}}

	statusResponse, err = rc0client.RRSet.Create("golib-example.at", rrsetCreate)

	if eq := strings.Compare("ok", statusResponse.Status); eq != 0 {
		log.Fatalf("failed to add rrset: %v", statusResponse.Message)
	}

	log.Println("rc0client.RRSet.Create:")
	log.Println()
	log.Println("Type: \t\t\t", 		 reflect.TypeOf(statusResponse))
	log.Println("Status: \t\t", 		 statusResponse.Status)
	log.Println("Status Message:\t",	 statusResponse.Message)
	log.Println()

	// Remove the example zone from rcode0
	statusResponse, err = rc0client.Zones.Delete("golib-example.at")

	if err != nil {
		log.Fatalf("failed to remove 'golib-example.at' from rcodezero: %v", err)
	}

	log.Println("rc0client.Zones.Delete:")
	log.Println()
	log.Println("Type: \t\t\t", 		 reflect.TypeOf(statusResponse))
	log.Println("Status: \t\t", 		 statusResponse.Status)
	log.Println("Status Message:\t",	 statusResponse.Message)

}
