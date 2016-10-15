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
	"time"
)

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

func eventLoop() {
	//TODO
	var urlBase = "https://192.168.33.5:5665"
	var apiUser = "root"
	var apiPass = "icinga"

	eventsEndpoint := urlBase + "/v1/events"
	fmt.Println("Check if HTTP connection is intact")
	//TODO

	netClient := initHTTPClient()

	var requestBody = []byte(`{ "queue": "mine", "types": [ "CheckResult" ]}`)
	req, err := http.NewRequest("POST", eventsEndpoint, bytes.NewBuffer(requestBody))
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(apiUser, apiPass)

	resp, err := netClient.Do(req)
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
		}
		fmt.Println("Processing message")
		fmt.Println(string(line))
	}

	fmt.Println("Connection was closed by the server")
}

func cleanup() {
	fmt.Println("Exit cleanup.")
}

func main() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	for {
		//does long polling until infinity
		eventLoop()

		//before retrying immediately, wait a while
		time.Sleep(time.Second * 5)
	}
}
