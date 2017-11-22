package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

// TcpeekPrometheusMetric is metric for tcpeek
type TcpeekPrometheusMetric struct {
	Success TcpeekPrometheusSuccess
	Failure TcpeekPrometheusFailure
}

// TcpeekPrometheusSuccess is metric for tcpeek
type TcpeekPrometheusSuccess struct {
	Total     prometheus.Gauge
	DupSyn    prometheus.Gauge
	DupSynAck prometheus.Gauge
}

// TcpeekPrometheusFailure is metric for tcpeek
type TcpeekPrometheusFailure struct {
	Total   prometheus.Gauge
	Timeout prometheus.Gauge
	Reject  prometheus.Gauge
	Unreach prometheus.Gauge
}

// TcpeekMetric is metric for tcpeek
type TcpeekMetric struct {
	Success TcpeekSuccess `json:"success"`
	Failure TcpeekFailure `json:"failure"`
}

// TcpeekSuccess is statistics for tcpeek
type TcpeekSuccess struct {
	Total     int64 `json:"total"`
	DupSyn    int64 `json:"dupsyn"`
	DupSynAck int64 `json:"dupsynack"`
}

// TcpeekFailure is statistics for tcpeek
type TcpeekFailure struct {
	Total   int64 `json:"total"`
	Timeout int64 `json:"timeout"`
	Reject  int64 `json:"reject"`
	Unreach int64 `json:"unreach"`
}

// TcpeekFailure is statistics for tcpeek pcap
type TcpeekPcapStats struct {
	Recv   int64 `json:"recv"`
	Drop   int64 `json:"drop"`
	Ifdrop int64 `json:"ifdrop"`
}

// DescribeTcpeekPrometheusMetric return TcpeekPrometheusMetric
func DescribeTcpeekPrometheusMetric(key string) TcpeekPrometheusMetric {
	return TcpeekPrometheusMetric{
		Success: TcpeekPrometheusSuccess{
			Total: prometheus.NewGauge(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      fmt.Sprintf("%s_success_total", key),
				Help:      fmt.Sprintf("%s_success_total", key),
			}),
			DupSyn: prometheus.NewGauge(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      fmt.Sprintf("%s_success_dupsyn", key),
				Help:      fmt.Sprintf("%s_success_dupsyn", key),
			}),
			DupSynAck: prometheus.NewGauge(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      fmt.Sprintf("%s_success_dupsynack", key),
				Help:      fmt.Sprintf("%s_success_dupsynack", key),
			}),
		},
		Failure: TcpeekPrometheusFailure{
			Total: prometheus.NewGauge(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      fmt.Sprintf("%s_failure_total", key),
				Help:      fmt.Sprintf("%s_failure_total", key),
			}),
			Timeout: prometheus.NewGauge(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      fmt.Sprintf("%s_failure_timeout", key),
				Help:      fmt.Sprintf("%s_failure_timeout", key),
			}),
			Reject: prometheus.NewGauge(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      fmt.Sprintf("%s_failure_reject", key),
				Help:      fmt.Sprintf("%s_failure_reject", key),
			}),
			Unreach: prometheus.NewGauge(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      fmt.Sprintf("%s_failure_unreach", key),
				Help:      fmt.Sprintf("%s_failure_unreach", key),
			}),
		},
	}
}
