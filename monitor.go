package kcp

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sandwich-go/logbus/monitor"
)

var (
	kcpCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "kcp_active_count",
		})
	kcpUpdateTiming = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "kcp_update_time",
			Buckets: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 15, 18, 20, 25, 30, 35, 40},
		})
)

func init() {
	monitor.RegisterCollector(kcpCount)
	monitor.RegisterCollector(kcpUpdateTiming)
}
