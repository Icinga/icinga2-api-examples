/******************************************************************************
 * Icinga 2 API Example: Event Streams with Golang                                                               *
 * Copyright (C) 2016 Icinga Development Team (https://www.icinga.org)        *
 *                                                                            *
 * This program is free software; you can redistribute it and/or              *
 * modify it under the terms of the GNU General Public License                *
 * as published by the Free Software Foundation; either version 2             *
 * of the License, or (at your option) any later version.                     *
 *                                                                            *
 * This program is distributed in the hope that it will be useful,            *
 * but WITHOUT ANY WARRANTY; without even the implied warranty of             *
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the              *
 * GNU General Public License for more details.                               *
 *                                                                            *
 * You should have received a copy of the GNU General Public License          *
 * along with this program; if not, write to the Free Software Foundation     *
 * Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301, USA.             *
 ******************************************************************************/

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"net"
	"net/http"
	"crypto/tls"
	"bytes"
	"bufio"
	"encoding/json"
	"time"
)

/*****************************************************************************/

//Golang has static types and does not like dynamic json
//https://github.com/xert/go-icinga2/blob/master/icinga/hosts.go
//we need to decode the following event messages which may differ based on their selected types
//type: CheckResult
//{"check_result":{"active":true,"check_source":"icinga2","command":["/usr/lib64/nagios/plugins/check_ssh","54.218.59.62"],"execution_end":1476468766.2535429001,"execution_start":1476468756.2439880371,"exit_status":2.0,"output":"CRITICAL - Socket timeout after 10 seconds","performance_data":[],"schedule_end":1476468766.2536571026,"schedule_start":1476468756.2430608273,"state":2.0,"type":"CheckResult","vars_after":{"attempt":1.0,"reachable":false,"state":2.0,"state_type":1.0},"vars_before":{"attempt":1.0,"reachable":false,"state":2.0,"state_type":1.0}},"host":"i-a78837bf","service":"ssh","timestamp":1476468766.2537300587,"type":"CheckResult"}

type EventTypeCheckResult struct {
	CheckResult	CheckResultAttrs `json:"check_result"`
	Host		string		`json:"host,omitempty"`
	Service		string		`json:"service,omitempty"`
	Timestamp	float64		`json:"timestamp,omitempty"`
	Type		string		`json:"type,omitempty"`
}

//performance data requires a different handling - strings or parsed data as PerfdataValue
//http://stackoverflow.com/questions/13364181/how-to-unmarshall-an-array-of-different-types-correctly
//{"check_result":{"active":true,"check_source":"icinga2","command":null,"execution_end":1476472102.1215960979,"execution_start":1476472102.1214408875,"exit_status":0.0,"output":"Hello from icinga2","performance_data":[{"counter":false,"crit":null,"label":"time","max":null,"min":null,"type":"PerfdataValue","unit":"","value":1476472102.12155509,"warn":null}],"schedule_end":1476472102.1215960979,"schedule_start":1476472102.1199998856,"state":3.0,"type":"CheckResult","vars_after":{"attempt":1.0,"reachable":true,"state":3.0,"state_type":0.0},"vars_before":{"attempt":1.0,"reachable":true,"state":0.0,"state_type":1.0}},"host":"icinga2","service":"random-002","timestamp":1476472102.121901989,"type":"CheckResult"}

type PerformanceDataValue struct {
	Counter		bool		`json:"counter,omitempty"`
	Crit		json.Number	`json:"crit,omitempty"`
	Label		string		`json:"label,omitempty"`
	Max		json.Number	`json:"max,omitempty"`
	Min		json.Number	`json:"min,omitempty"`
	Type		string		`json:"type,omitempty"`
	Unit		string		`json:"unit,omitempty"`
	Value		json.Number	`json:"value,omitempty"`
	Warn		json.Number	`json:"warn,omitempty"`
}

type CheckResultAttrs struct {
	Active		bool		`json:"active,omitempty"`
	CheckSource	string		`json:"check_source,omitempty"`
	Command		[]string	`json:"command,omitempty"`
	ExecutionEnd	float64		`json:"execution_end,omitempty"`
	ExecutionStart	float64		`json:"execution_start,omitempty"`
	ExitStatus	json.Number	`json:"exit_status,omitempty"`
	Output		string		`json:"output,omitempty"`
	PerformanceDataParsed []PerformanceDataValue `json:"performance_data,omitempty"`
	PerformanceDataStr []string	`json:"performance_data,omitempty"`
	ScheduleEnd	float64		`json:"schedule_end,omitempty"`
	ScheduleStart	float64		`json:"schedule_start,omitempty"`
	State		json.Number	`json:"state,omitempty"`
	Type		string		`json:"type,omitempty"`
	VarsAfter	*struct {
		Attempt		json.Number	`json:"attempt,omitempty"`
		Reachable	bool		`json:"reachable,omitempty"`
		State		json.Number	`json:"state,omitempty"`
		StateType	json.Number	`json:"state_type,omitempty"`
	} `json:"vars_after,omitempty"`
	VarsBefore	*struct {
		Attempt		json.Number	`json:"attempt,omitempty"`
		Reachable	bool		`json:"reachable,omitempty"`
		State		json.Number	`json:"state,omitempty"`
		StateType	json.Number	`json:"state_type,omitempty"`
	} `json:"vars_before,omitempty"`
}

/*****************************************************************************/

var StartTime int64
var CheckResultCount int64
var CheckResultCountObject map[string]int64

var StateChangesObject map[string][]int64

/*****************************************************************************/

func initHTTPClient() *http.Client {
	sslSkipVerify := true //TODO: config option

	//https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
	httpTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: time.Second * 5,
			KeepAlive: time.Second * 30,
		}).Dial,
		TLSHandshakeTimeout: time.Second * 5,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: sslSkipVerify,
		},
	}
	//We don't need a client timeout for http long polling!
	httpClient := &http.Client{
		Transport: httpTransport,
	}

	return httpClient
}

func handleEventTypes(response string) {
	var CheckResult EventTypeCheckResult
	err := json.Unmarshal([]byte(response), &CheckResult)
	if err == nil {
//		fmt.Printf("%+v\n", CheckResult)
		//TODO: Add fall through for different types
	} else {
//		fmt.Println(err)
//		fmt.Printf("%+v\n", CheckResult)
	}

	if CheckResult.Type == "CheckResult" {
		//fmt.Println("Processing type 'CheckResult'")

		CheckResultCount++

		var ObjectName = CheckResult.Host
		if CheckResult.Service != "" {
			ObjectName = ObjectName + "!" + CheckResult.Service
		}

		CheckResultCountObject[ObjectName] = CheckResultCountObject[ObjectName] + 1

		//TODO keep array size static
		state, _ := json.Number.Int64(CheckResult.CheckResult.State)

		StateChangesObject[ObjectName] = append(StateChangesObject[ObjectName], state)

		//periodic statistics
		if time.Now().Unix() % 2 == 0 {
			var CheckResultRate = float64(CheckResultCount) / float64(time.Now().Unix() - StartTime)
			fmt.Println("Global check result rate/second:", CheckResultRate)

			var CheckResultObjectRate = float64(CheckResultCountObject[ObjectName]) / float64(time.Now().Unix() - StartTime)
			fmt.Printf("Check result rate for object '%s': %.02f\n", ObjectName, CheckResultObjectRate)

			fmt.Println(" ")
		}
	}
}


func eventLoop() {
	//TODO
	var urlBase = "https://192.168.33.5:5665"
	var apiUser = "root"
	var apiPass = "icinga"

	eventsEndpoint := urlBase + "/v1/events"
	fmt.Println("Check if HTTP connection is intact")

	httpClient := initHTTPClient()

	var requestBody = []byte(`{ "queue": "mine", "types": [ "CheckResult" ]}`)
	//TODO: Depending on the types the user chooses we need to select our response handling

	req, err := http.NewRequest("POST", eventsEndpoint, bytes.NewBuffer(requestBody))
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(apiUser, apiPass)

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Server error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	//http://stackoverflow.com/questions/22108519/how-do-i-read-a-streaming-response-body-using-golangs-net-http-package
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')

		if err != nil {
			fmt.Println("Error reading stream", err)
			return
		}
		//fmt.Println("Processing message")
		//fmt.Println(string(line))

		//call handler for json decoding and message processing
		handleEventTypes(string(line))
	}

	fmt.Println("Connection was closed by the server")
}

func cleanup() {
	fmt.Println("Exit cleanup.")
}

/*****************************************************************************/

func main() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	StartTime = time.Now().Unix()

	CheckResultCount = 0
	CheckResultCountObject = make(map[string]int64)
	StateChangesObject = make(map[string][]int64)

	for {
		//does long polling until infinity or the server closes the connection
		eventLoop()

		//before retrying immediately, wait a while
		time.Sleep(time.Second * 5)
	}
}
