/*
 * NETCAP - Traffic Analysis Framework
 * Copyright (c) 2017 Philipp Mieden <dreadl0ck [at] protonmail [dot] ch>
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package types

import (
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

var fieldsMPLS = []string{
	"Timestamp",
	"Label",
	"TrafficClass",
	"StackBottom",
	"TTL",
}

func (a MPLS) CSVHeader() []string {
	return filter(fieldsMPLS)
}

func (a MPLS) CSVRecord() []string {
	return filter([]string{
		formatTimestamp(a.Timestamp),
		formatInt32(a.Label),              // int32
		formatInt32(a.TrafficClass),       // int32
		strconv.FormatBool(a.StackBottom), // bool
		formatInt32(a.TTL),                // int32
	})
}

func (a MPLS) NetcapTimestamp() string {
	return a.Timestamp
}

func (a MPLS) JSON() (string, error) {
	return jsonMarshaler.MarshalToString(&a)
}

var mplsMetric = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: strings.ToLower(Type_NC_MPLS.String()),
		Help: Type_NC_MPLS.String() + " audit records",
	},
	fieldsMPLS[1:],
)

func init() {
	prometheus.MustRegister(mplsMetric)
}

func (a MPLS) Inc() {
	mplsMetric.WithLabelValues(a.CSVRecord()[1:]...).Inc()
}
