package main

import (
	"strconv"

	"github.com/alecthomas/kingpin/v2"
	"github.com/mako042/ganesha_exporter/dbus"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	nfsV3RequestedDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v3_requested_bytes_total",
		"Number of requested bytes for NFSv3 operations",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV3TransferedDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v3_transfered_bytes_total",
		"Number of transfered bytes for NFSv3 operations",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV3OperationsDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v3_operations_total",
		"Number of operations for NFSv3",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV3ErrorsDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v3_operations_errors_total",
		"Number of operations in error for NFSv3",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV3LatencyDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v3_operations_latency_seconds_total",
		"Cumulative time consumed by operations for NFSv3",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV3QueueWaitDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v3_operations_queue_wait_seconds_total",
		"Cumulative time spent in rpc wait queue for NFSv3",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV40RequestedDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v40_requested_bytes_total",
		"Number of requested bytes for NFSv4.0 operations",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV40TransferedDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v40_transfered_bytes_total",
		"Number of transfered bytes for NFSv4.0 operations",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV40OperationsDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v40_operations_total",
		"Number of operations for NFSv4.0",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV40ErrorsDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v40_operations_errors_total",
		"Number of operations in error for NFSv4.0",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV40LatencyDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v40_operations_latency_seconds_total",
		"Cumulative time consumed by operations for NFSv4.0",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV40QueueWaitDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v40_operations_queue_wait_seconds_total",
		"Cumulative time spent in rpc wait queue for NFSv4.0",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV41RequestedDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v41_requested_bytes_total",
		"Number of requested bytes for NFSv4.1 operations",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV41TransferedDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v41_transfered_bytes_total",
		"Number of transfered bytes for NFSv4.1 operations",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV41OperationsDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v41_operations_total",
		"Number of operations for NFSv4.1",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV41ErrorsDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v41_operations_errors_total",
		"Number of operations in error for NFSv4.1",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV41LatencyDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v41_operations_latency_seconds_total",
		"Cumulative time consumed by operations for NFSv4.1",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV41QueueWaitDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v41_operations_queue_wait_seconds_total",
		"Cumulative time spent in rpc wait queue for NFSv4.1",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV42RequestedDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v42_requested_bytes_total",
		"Number of requested bytes for NFSv4.2 operations",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV42TransferedDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v42_transfered_bytes_total",
		"Number of transfered bytes for NFSv4.2 operations",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV42OperationsDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v42_operations_total",
		"Number of operations for NFSv4.2",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV42ErrorsDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v42_operations_errors_total",
		"Number of operations in error for NFSv4.2",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV42LatencyDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v42_operations_latency_seconds_total",
		"Cumulative time consumed by operations for NFSv4.2",
		[]string{"direction", "exportid", "path"}, nil,
	)
	nfsV42QueueWaitDesc = prometheus.NewDesc(
		"ganesha_exports_nfs_v42_operations_queue_wait_seconds_total",
		"Cumulative time spent in rpc wait queue for NFSv4.2",
		[]string{"direction", "exportid", "path"}, nil,
	)
	pnfsLayoutOperationsDesc = prometheus.NewDesc(
		"ganesha_exports_pnfs_v41_layout_operations_total",
		"Numer of layout operations for pNFSv4.1",
		[]string{"type", "exportid", "path"}, nil,
	)
	pnfsLayoutErrorsDesc = prometheus.NewDesc(
		"ganesha_exports_pnfs_v41_layout_operations_errors_total",
		"Numer of layout operations in error for pNFSv4.1",
		[]string{"type", "exportid", "path"}, nil,
	)
	pnfsLayoutDelayDesc = prometheus.NewDesc(
		"ganesha_exports_pnfs_v41_layout_delay_seconds_total",
		"Cumulative delay time for pNFSv4.1",
		[]string{"direction", "exportid", "path"}, nil,
	)
	pnfsV42LayoutOperationsDesc = prometheus.NewDesc(
		"ganesha_exports_pnfs_v42_layout_operations_total",
		"Numer of layout operations for pNFSv4.2",
		[]string{"type", "exportid", "path"}, nil,
	)
	pnfsV42LayoutErrorsDesc = prometheus.NewDesc(
		"ganesha_exports_pnfs_v42_layout_operations_errors_total",
		"Numer of layout operations in error for pNFSv4.2",
		[]string{"type", "exportid", "path"}, nil,
	)
	pnfsV42LayoutDelayDesc = prometheus.NewDesc(
		"ganesha_exports_pnfs_v42_layout_delay_seconds_total",
		"Cumulative delay time for pNFSv4.2",
		[]string{"direction", "exportid", "path"}, nil,
	)
)

// ExportsCollector Collector for ganesha exports
type ExportsCollector struct {
	exportMgr                      dbus.ExportMgr
	nfsv3, nfsv40, nfsv41, pnfsv41 *bool
	nfsv42, pnfsv42                *bool
}

// NewExportsCollector creates a new collector
func NewExportsCollector() ExportsCollector {
	return ExportsCollector{
		exportMgr: dbus.NewExportMgr(),
		nfsv3:     kingpin.Flag("collector.exports.nfsv3", "Activate NFSv3 stats").Default("true").Bool(),
		nfsv40:    kingpin.Flag("collector.exports.nfsv40", "Activate NFSv4.0 stats").Default("true").Bool(),
		nfsv41:    kingpin.Flag("collector.exports.nfsv41", "Activate NFSv4.1 stats").Default("true").Bool(),
		pnfsv41:   kingpin.Flag("collector.exports.pnfsv41", "Activate pNFSv4.1 stats").Default("true").Bool(),
		nfsv42:    kingpin.Flag("collector.exports.nfsv42", "Activate NFSv4.2 stats").Default("true").Bool(),
		pnfsv42:   kingpin.Flag("collector.exports.pnfsv42", "Activate pNFSv4.2 stats").Default("true").Bool(),
	}
}

// Describe prometheus description
func (ic ExportsCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(ic, ch)
}

// Collect do the actual job
func (ic ExportsCollector) Collect(ch chan<- prometheus.Metric) {
	var ()
	_, exports := ic.exportMgr.ShowExports()
	for _, export := range exports {
		exportid := strconv.FormatUint(uint64(export.ExportID), 10)
		path := export.Path
		if *ic.nfsv3 {
			var stats dbus.BasicStats
			if export.Protocols[0].Enabled { // NFSv3
				stats = ic.exportMgr.GetNFSv3IO(export.ExportID)
			}
			ch <- prometheus.MustNewConstMetric(
				nfsV3RequestedDesc,
				prometheus.CounterValue,
				float64(stats.Read.Requested),
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV3TransferedDesc,
				prometheus.CounterValue,
				float64(stats.Read.Transfered),
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV3OperationsDesc,
				prometheus.CounterValue,
				float64(stats.Read.Total),
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV3ErrorsDesc,
				prometheus.CounterValue,
				float64(stats.Read.Errors),
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV3LatencyDesc,
				prometheus.CounterValue,
				float64(stats.Read.Latency)/1e9,
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV3QueueWaitDesc,
				prometheus.CounterValue,
				float64(stats.Read.QueueWait)/1e9,
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV3RequestedDesc,
				prometheus.CounterValue,
				float64(stats.Write.Requested),
				"write", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV3TransferedDesc,
				prometheus.CounterValue,
				float64(stats.Write.Transfered),
				"write", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV3OperationsDesc,
				prometheus.CounterValue,
				float64(stats.Write.Total),
				"write", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV3ErrorsDesc,
				prometheus.CounterValue,
				float64(stats.Write.Errors),
				"write", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV3LatencyDesc,
				prometheus.CounterValue,
				float64(stats.Write.Latency)/1e9,
				"write", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV3QueueWaitDesc,
				prometheus.CounterValue,
				float64(stats.Write.QueueWait)/1e9,
				"write", exportid, path)
		}
		if *ic.nfsv40 {
			stats := dbus.BasicStats{}
			if export.Protocols[4].Enabled { // NFSv4.0
				stats = ic.exportMgr.GetNFSv40IO(export.ExportID)
			}
			ch <- prometheus.MustNewConstMetric(
				nfsV40RequestedDesc,
				prometheus.CounterValue,
				float64(stats.Read.Requested),
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV40TransferedDesc,
				prometheus.CounterValue,
				float64(stats.Read.Transfered),
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV40OperationsDesc,
				prometheus.CounterValue,
				float64(stats.Read.Total),
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV40ErrorsDesc,
				prometheus.CounterValue,
				float64(stats.Read.Errors),
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV40LatencyDesc,
				prometheus.CounterValue,
				float64(stats.Read.Latency)/1e9,
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV40QueueWaitDesc,
				prometheus.CounterValue,
				float64(stats.Read.QueueWait)/1e9,
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV40RequestedDesc,
				prometheus.CounterValue,
				float64(stats.Write.Requested),
				"write", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV40TransferedDesc,
				prometheus.CounterValue,
				float64(stats.Write.Transfered),
				"write", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV40OperationsDesc,
				prometheus.CounterValue,
				float64(stats.Write.Total),
				"write", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV40ErrorsDesc,
				prometheus.CounterValue,
				float64(stats.Write.Errors),
				"write", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV40LatencyDesc,
				prometheus.CounterValue,
				float64(stats.Write.Latency)/1e9,
				"write", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV40QueueWaitDesc,
				prometheus.CounterValue,
				float64(stats.Write.QueueWait)/1e9,
				"write", exportid, path)
		}
		if *ic.nfsv41 {
			stats := dbus.BasicStats{}
			if export.Protocols[5].Enabled { // NFSv4.1
				stats = ic.exportMgr.GetNFSv41IO(export.ExportID)
			}
			ch <- prometheus.MustNewConstMetric(
				nfsV41RequestedDesc,
				prometheus.CounterValue,
				float64(stats.Read.Requested),
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV41TransferedDesc,
				prometheus.CounterValue,
				float64(stats.Read.Transfered),
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV41OperationsDesc,
				prometheus.CounterValue,
				float64(stats.Read.Total),
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV41ErrorsDesc,
				prometheus.CounterValue,
				float64(stats.Read.Errors),
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV41LatencyDesc,
				prometheus.CounterValue,
				float64(stats.Read.Latency)/1e9,
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV41QueueWaitDesc,
				prometheus.CounterValue,
				float64(stats.Read.QueueWait)/1e9,
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV41RequestedDesc,
				prometheus.CounterValue,
				float64(stats.Write.Requested),
				"write", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV41TransferedDesc,
				prometheus.CounterValue,
				float64(stats.Write.Transfered),
				"write", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV41OperationsDesc,
				prometheus.CounterValue,
				float64(stats.Write.Total),
				"write", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV41ErrorsDesc,
				prometheus.CounterValue,
				float64(stats.Write.Errors),
				"write", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV41LatencyDesc,
				prometheus.CounterValue,
				float64(stats.Write.Latency)/1e9,
				"write", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV41QueueWaitDesc,
				prometheus.CounterValue,
				float64(stats.Write.QueueWait)/1e9,
				"write", exportid, path)
		}
		if *ic.pnfsv41 {
			stats := dbus.PNFSOperations{}
			if export.Protocols[5].Enabled { // NFSv4.1
				stats = ic.exportMgr.GetNFSv41Layouts(export.ExportID)
			}
			ch <- prometheus.MustNewConstMetric(
				pnfsLayoutOperationsDesc,
				prometheus.CounterValue,
				float64(stats.Getdevinfo.Total),
				"getdevinfo", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsLayoutErrorsDesc,
				prometheus.CounterValue,
				float64(stats.Getdevinfo.Errors),
				"getdevinfo", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsLayoutDelayDesc,
				prometheus.CounterValue,
				float64(stats.Getdevinfo.Delays)/1e9,
				"getdevinfo", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsLayoutOperationsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutGet.Total),
				"get", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsLayoutErrorsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutGet.Errors),
				"get", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsLayoutDelayDesc,
				prometheus.CounterValue,
				float64(stats.LayoutGet.Delays)/1e9,
				"get", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsLayoutOperationsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutCommit.Total),
				"commit", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsLayoutErrorsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutCommit.Errors),
				"commit", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsLayoutDelayDesc,
				prometheus.CounterValue,
				float64(stats.LayoutCommit.Delays)/1e9,
				"commit", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsLayoutOperationsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutReturn.Total),
				"return", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsLayoutErrorsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutReturn.Errors),
				"return", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsLayoutDelayDesc,
				prometheus.CounterValue,
				float64(stats.LayoutReturn.Delays)/1e9,
				"return", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsLayoutOperationsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutRecall.Total),
				"recall", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsLayoutErrorsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutRecall.Errors),
				"recall", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsLayoutDelayDesc,
				prometheus.CounterValue,
				float64(stats.LayoutRecall.Delays)/1e9,
				"recall", exportid, path)
		}
		if *ic.nfsv42 {
			stats := dbus.BasicStats{}
			if export.Protocols[6].Enabled { // NFSv4.2
				stats = ic.exportMgr.GetNFSv42IO(export.ExportID)
			}
			ch <- prometheus.MustNewConstMetric(
				nfsV42RequestedDesc,
				prometheus.CounterValue,
				float64(stats.Read.Requested),
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV42TransferedDesc,
				prometheus.CounterValue,
				float64(stats.Read.Transfered),
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV42OperationsDesc,
				prometheus.CounterValue,
				float64(stats.Read.Total),
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV42ErrorsDesc,
				prometheus.CounterValue,
				float64(stats.Read.Errors),
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV42LatencyDesc,
				prometheus.CounterValue,
				float64(stats.Read.Latency)/1e9,
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV42QueueWaitDesc,
				prometheus.CounterValue,
				float64(stats.Read.QueueWait)/1e9,
				"read", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV42RequestedDesc,
				prometheus.CounterValue,
				float64(stats.Write.Requested),
				"write", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV42TransferedDesc,
				prometheus.CounterValue,
				float64(stats.Write.Transfered),
				"write", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV42OperationsDesc,
				prometheus.CounterValue,
				float64(stats.Write.Total),
				"write", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV42ErrorsDesc,
				prometheus.CounterValue,
				float64(stats.Write.Errors),
				"write", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV42LatencyDesc,
				prometheus.CounterValue,
				float64(stats.Write.Latency)/1e9,
				"write", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				nfsV42QueueWaitDesc,
				prometheus.CounterValue,
				float64(stats.Write.QueueWait)/1e9,
				"write", exportid, path)
		}
		if *ic.pnfsv42 {
			stats := dbus.PNFSOperations{}
			if export.Protocols[6].Enabled { // NFSv4.2
				stats = ic.exportMgr.GetNFSv42Layouts(export.ExportID)
			}
			ch <- prometheus.MustNewConstMetric(
				pnfsV42LayoutOperationsDesc,
				prometheus.CounterValue,
				float64(stats.Getdevinfo.Total),
				"getdevinfo", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsV42LayoutErrorsDesc,
				prometheus.CounterValue,
				float64(stats.Getdevinfo.Errors),
				"getdevinfo", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsV42LayoutDelayDesc,
				prometheus.CounterValue,
				float64(stats.Getdevinfo.Delays)/1e9,
				"getdevinfo", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsV42LayoutOperationsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutGet.Total),
				"get", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsV42LayoutErrorsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutGet.Errors),
				"get", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsV42LayoutDelayDesc,
				prometheus.CounterValue,
				float64(stats.LayoutGet.Delays)/1e9,
				"get", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsV42LayoutOperationsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutCommit.Total),
				"commit", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsV42LayoutErrorsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutCommit.Errors),
				"commit", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsV42LayoutDelayDesc,
				prometheus.CounterValue,
				float64(stats.LayoutCommit.Delays)/1e9,
				"commit", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsV42LayoutOperationsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutReturn.Total),
				"return", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsV42LayoutErrorsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutReturn.Errors),
				"return", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsV42LayoutDelayDesc,
				prometheus.CounterValue,
				float64(stats.LayoutReturn.Delays)/1e9,
				"return", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsV42LayoutOperationsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutRecall.Total),
				"recall", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsV42LayoutErrorsDesc,
				prometheus.CounterValue,
				float64(stats.LayoutRecall.Errors),
				"recall", exportid, path)
			ch <- prometheus.MustNewConstMetric(
				pnfsV42LayoutDelayDesc,
				prometheus.CounterValue,
				float64(stats.LayoutRecall.Delays)/1e9,
				"recall", exportid, path)
		}
	}
}
