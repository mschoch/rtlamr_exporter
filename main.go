//  Copyright (c) 2018 Marty Schoch
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var bindAddr = flag.String("addr", ":9415", "http listen address")

var (
	powerConsumed = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "power_consumed_kwh_total",
			Help: "Number of kilowatt hours consumed.",
		},
	)
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(powerConsumed)
}

func main() {
	flag.Parse()

	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(*bindAddr, nil)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	var last int
	for err == nil {
		var scm StandardConsumptionMsg
		jsonErr := json.Unmarshal([]byte(text), &scm)
		if jsonErr != nil {
			text, err = reader.ReadString('\n')
			continue
		}
		diff := scm.Message.Consumption - last
		if diff > 0 {
			// dont allow backwards
			powerConsumed.Add(float64(diff))
			last = scm.Message.Consumption
		}
		text, err = reader.ReadString('\n')
	}
}
