package dbus

import (
	"log"

	"github.com/godbus/dbus"
)

// ProtocolEntry represents a protocol and its enabled state
type ProtocolEntry struct {
	Name    string
	Enabled bool
}

// Export represents an NFS Ganesha export
type Export struct {
	ExportID  uint16
	Path      string
	Protocols [7]ProtocolEntry // 7 protocols: NFSv3, MNTv3, NLMv4, RQUOTA, NFSv40, NFSv41, NFSv42
	Last      uint64
	Time      struct {
		Sec  uint64
		Nsec uint64
	}
}

type ExportMgr struct {
	dbusObject dbus.BusObject
}

func NewExportMgr() ExportMgr {
	conn, err := dbus.SystemBus()
	if err != nil {
		log.Panic(err)
	}

	return ExportMgr{
		dbusObject: conn.Object(
			"org.ganesha.nfsd",
			"/org/ganesha/nfsd/ExportMgr",
		),
	}
}

func (m ExportMgr) ShowExports() (uint64, []Export) {
	call := m.dbusObject.Call("org.ganesha.nfsd.exportmgr.ShowExports", 0)
	if call.Err != nil {
		log.Panic(call.Err)
	}

	var header [2]uint64
	var exports []Export
	if err := call.Store(&header, &exports); err != nil {
		log.Panic(err)
	}

	return header[0], exports
}

// GetNFSv3IO retrieves NFSv3 IO statistics for a specific export
func (m ExportMgr) GetNFSv3IO(exportID uint16) BasicStats {
	out := BasicStats{}
	call := m.dbusObject.Call("org.ganesha.nfsd.exportstats.GetNFSv3IO", 0, exportID)
	if call.Err != nil {
		log.Panic(call.Err)
	}
	if !call.Body[0].(bool) {
		if err := call.Store(&out.Status, &out.Error, &out.Time); err != nil {
			log.Panic(err)
		}
		return out
	}
	if err := call.Store(
		&out.Status, &out.Error, &out.Time,
		&out.Read, &out.Write,
	); err != nil {
		log.Panic(err)
	}
	return out
}

// GetNFSv40IO retrieves NFSv4.0 IO statistics for a specific export
func (m ExportMgr) GetNFSv40IO(exportID uint16) BasicStats {
	out := BasicStats{}
	call := m.dbusObject.Call("org.ganesha.nfsd.exportstats.GetNFSv40IO", 0, exportID)
	if call.Err != nil {
		log.Panic(call.Err)
	}
	if !call.Body[0].(bool) {
		if err := call.Store(&out.Status, &out.Error, &out.Time); err != nil {
			log.Panic(err)
		}
		return out
	}
	if err := call.Store(
		&out.Status, &out.Error, &out.Time,
		&out.Read, &out.Write,
	); err != nil {
		log.Panic(err)
	}
	return out
}

// GetNFSv41IO retrieves NFSv4.1 IO statistics for a specific export
func (m ExportMgr) GetNFSv41IO(exportID uint16) BasicStats {
	out := BasicStats{}
	call := m.dbusObject.Call("org.ganesha.nfsd.exportstats.GetNFSv41IO", 0, exportID)
	if call.Err != nil {
		log.Panic(call.Err)
	}
	if !call.Body[0].(bool) {
		if err := call.Store(&out.Status, &out.Error, &out.Time); err != nil {
			log.Panic(err)
		}
		return out
	}
	if err := call.Store(
		&out.Status, &out.Error, &out.Time,
		&out.Read, &out.Write,
	); err != nil {
		log.Panic(err)
	}
	return out
}

// GetNFSv41Layouts retrieves pNFSv4.1 layout statistics for a specific export
func (m ExportMgr) GetNFSv41Layouts(exportID uint16) PNFSOperations {
	out := PNFSOperations{}
	call := m.dbusObject.Call("org.ganesha.nfsd.exportstats.GetNFSv41Layouts", 0, exportID)
	if call.Err != nil {
		log.Panic(call.Err)
	}
	if !call.Body[0].(bool) {
		if err := call.Store(&out.Status, &out.Error, &out.Time); err != nil {
			log.Panic(err)
		}
		return out
	}
	if err := call.Store(
		&out.Status, &out.Error, &out.Time,
		&out.Getdevinfo, &out.LayoutGet, &out.LayoutCommit, &out.LayoutReturn, &out.LayoutRecall,
	); err != nil {
		log.Panic(err)
	}
	return out
}
