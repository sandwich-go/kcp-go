package kcp

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sandwich-go/logbus/monitor"
)

var (
	lengthBucket = []float64{5, 10, 15, 20, 30, 50, 70, 100, 150, 200, 250, 300, 350, 400, 500}
	kcpCount     = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "kcp_active_count",
		})
	kcpUpdateTiming = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "kcp_update_time",
			Buckets: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 15, 18, 20, 25, 30, 35, 40},
		})
	sendQueueLen = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "send_queue_length",
			Buckets: lengthBucket,
		})
	recvQueueLen = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "recv_queue_length",
			Buckets: lengthBucket,
		})
	sendBufferLen = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "send_buffer_length",
			Buckets: lengthBucket,
		})
	recvBufferLen = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "recv_buffer_length",
			Buckets: lengthBucket,
		})
)

func init() {
	monitor.RegisterCollector(kcpCount)
	monitor.RegisterCollector(kcpUpdateTiming)
	monitor.RegisterCollector(sendQueueLen)
	monitor.RegisterCollector(recvQueueLen)
	monitor.RegisterCollector(sendBufferLen)
	monitor.RegisterCollector(recvBufferLen)
	monitor.RegisterCollector(newSnmpCollector())
}

func newSnmpCollector() prometheus.Collector {
	return &snmpCollector{
		BytesSentMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_bytes_sent",
			Help: "Bytes sent from upper level",
		}),
		BytesReceivedMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_bytes_received",
			Help: "Bytes received to upper level",
		}),
		MaxConnMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_max_connections",
			Help: "Maximum number of connections ever reached",
		}),
		ActiveOpensMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_active_opens",
			Help: "Accumulated active open connections",
		}),
		PassiveOpensMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_passive_opens",
			Help: "Accumulated passive open connections",
		}),
		CurrEstabMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_current_established",
			Help: "Current number of established connections",
		}),
		InErrsMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_in_errors",
			Help: "UDP read errors reported from net.PacketConn",
		}),
		InCsumErrorsMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_in_checksum_errors",
			Help: "Checksum errors from CRC32",
		}),
		KCPInErrorsMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_kcp_input_errors",
			Help: "Packet input errors reported from KCP",
		}),
		InPktsMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_in_packets",
			Help: "Incoming packets count",
		}),
		OutPktsMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_out_packets",
			Help: "Outgoing packets count",
		}),
		InSegsMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_in_segments",
			Help: "Incoming KCP segments",
		}),
		OutSegsMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_out_segments",
			Help: "Outgoing KCP segments",
		}),
		InBytesMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_in_bytes",
			Help: "UDP bytes received",
		}),
		OutBytesMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_out_bytes",
			Help: "UDP bytes sent",
		}),
		RetransSegsMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_retransmitted_segments",
			Help: "Accumulated retransmitted segments",
		}),
		FastRetransSegsMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_fast_retransmitted_segments",
			Help: "Accumulated fast retransmitted segments",
		}),
		EarlyRetransSegsMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_early_retransmitted_segments",
			Help: "Accumulated early retransmitted segments",
		}),
		LostSegsMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_lost_segments",
			Help: "Number of segments inferred as lost",
		}),
		RepeatSegsMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_duplicate_segments",
			Help: "Number of duplicated segments",
		}),
		FECRecoveredMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_fec_recovered",
			Help: "Correct packets recovered from FEC",
		}),
		FECErrsMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_fec_errors",
			Help: "Incorrect packets recovered from FEC",
		}),
		FECParityShardsMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_fec_parity_shards",
			Help: "FEC segments received",
		}),
		FECShortShardsMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "kcp_snmp_fec_short_shards",
			Help: "Number of data shards that are not enough for recovery",
		}),
	}
}

type snmpCollector struct {
	BytesSentMetric        prometheus.Gauge
	BytesReceivedMetric    prometheus.Gauge
	MaxConnMetric          prometheus.Gauge
	ActiveOpensMetric      prometheus.Gauge
	PassiveOpensMetric     prometheus.Gauge
	CurrEstabMetric        prometheus.Gauge
	InErrsMetric           prometheus.Gauge
	InCsumErrorsMetric     prometheus.Gauge
	KCPInErrorsMetric      prometheus.Gauge
	InPktsMetric           prometheus.Gauge
	OutPktsMetric          prometheus.Gauge
	InSegsMetric           prometheus.Gauge
	OutSegsMetric          prometheus.Gauge
	InBytesMetric          prometheus.Gauge
	OutBytesMetric         prometheus.Gauge
	RetransSegsMetric      prometheus.Gauge
	FastRetransSegsMetric  prometheus.Gauge
	EarlyRetransSegsMetric prometheus.Gauge
	LostSegsMetric         prometheus.Gauge
	RepeatSegsMetric       prometheus.Gauge
	FECRecoveredMetric     prometheus.Gauge
	FECErrsMetric          prometheus.Gauge
	FECParityShardsMetric  prometheus.Gauge
	FECShortShardsMetric   prometheus.Gauge
}

func (c *snmpCollector) Describe(ch chan<- *prometheus.Desc) {
	c.BytesSentMetric.Describe(ch)
	c.BytesReceivedMetric.Describe(ch)
	c.MaxConnMetric.Describe(ch)
	c.ActiveOpensMetric.Describe(ch)
	c.PassiveOpensMetric.Describe(ch)
	c.CurrEstabMetric.Describe(ch)
	c.InErrsMetric.Describe(ch)
	c.InCsumErrorsMetric.Describe(ch)
	c.KCPInErrorsMetric.Describe(ch)
	c.InPktsMetric.Describe(ch)
	c.OutPktsMetric.Describe(ch)
	c.InSegsMetric.Describe(ch)
	c.OutSegsMetric.Describe(ch)
	c.InBytesMetric.Describe(ch)
	c.OutBytesMetric.Describe(ch)
	c.RetransSegsMetric.Describe(ch)
	c.FastRetransSegsMetric.Describe(ch)
	c.EarlyRetransSegsMetric.Describe(ch)
	c.LostSegsMetric.Describe(ch)
	c.RepeatSegsMetric.Describe(ch)
	c.FECRecoveredMetric.Describe(ch)
	c.FECErrsMetric.Describe(ch)
	c.FECParityShardsMetric.Describe(ch)
	c.FECShortShardsMetric.Describe(ch)
}

func (c *snmpCollector) Collect(metrics chan<- prometheus.Metric) {
	snmp := DefaultSnmp.Copy()
	c.BytesSentMetric.Set(float64(snmp.BytesSent))
	c.BytesReceivedMetric.Set(float64(snmp.BytesReceived))
	c.MaxConnMetric.Set(float64(snmp.MaxConn))
	c.ActiveOpensMetric.Set(float64(snmp.ActiveOpens))
	c.PassiveOpensMetric.Set(float64(snmp.PassiveOpens))
	c.CurrEstabMetric.Set(float64(snmp.CurrEstab))
	c.InErrsMetric.Set(float64(snmp.InErrs))
	c.InCsumErrorsMetric.Set(float64(snmp.InCsumErrors))
	c.KCPInErrorsMetric.Set(float64(snmp.KCPInErrors))
	c.InPktsMetric.Set(float64(snmp.InPkts))
	c.OutPktsMetric.Set(float64(snmp.OutPkts))
	c.InSegsMetric.Set(float64(snmp.InSegs))
	c.OutSegsMetric.Set(float64(snmp.OutSegs))
	c.InBytesMetric.Set(float64(snmp.InBytes))
	c.OutBytesMetric.Set(float64(snmp.OutBytes))
	c.RetransSegsMetric.Set(float64(snmp.RetransSegs))
	c.FastRetransSegsMetric.Set(float64(snmp.FastRetransSegs))
	c.EarlyRetransSegsMetric.Set(float64(snmp.EarlyRetransSegs))
	c.LostSegsMetric.Set(float64(snmp.LostSegs))
	c.RepeatSegsMetric.Set(float64(snmp.RepeatSegs))
	c.FECRecoveredMetric.Set(float64(snmp.FECRecovered))
	c.FECErrsMetric.Set(float64(snmp.FECErrs))
	c.FECParityShardsMetric.Set(float64(snmp.FECParityShards))
	c.FECShortShardsMetric.Set(float64(snmp.FECShortShards))
}
