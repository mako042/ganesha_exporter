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

	// Use call.Body directly instead of Store
	if len(call.Body) < 2 {
		log.Panic("invalid dbus response: ShowExports - not enough body elements")
	}

	// Parse header (tt) - two uint64 values
	header, ok := call.Body[0].([]interface{})
	if !ok || len(header) < 2 {
		log.Panic("invalid header format")
	}
	timestamp := header[0].(uint64)

	// Parse exports array - it's [][]interface{} (array of arrays)
	exportsRaw, ok := call.Body[1].([][]interface{})
	if !ok {
		log.Printf("DEBUG: exportsRaw type assertion failed, got %T", call.Body[1])
		log.Panic("invalid exports format")
	}

	exports := make([]Export, 0, len(exportsRaw))

	for _, item := range exportsRaw {
		if len(item) < 5 {
			continue
		}

		exp := Export{
			ExportID: item[0].(uint16),
			Path:     item[1].(string),
			Last:     item[3].(uint64),
		}

		// Parse protocols: ((sb)(sb)(sb)(sb)(sb)(sb)(sb))
		if protos, ok := item[2].([]interface{}); ok {
			for i, p := range protos {
				if i >= 7 {
					break
				}
				pair, ok := p.([]interface{})
				if !ok || len(pair) < 2 {
					continue
				}
				if i < len(exp.Protocols) {
					exp.Protocols[i].Name = pair[0].(string)
					exp.Protocols[i].Enabled = pair[1].(bool)
				}
			}
		}

		// Parse time: (tt) - two uint64 values
		if ts, ok := item[4].([]interface{}); ok && len(ts) >= 2 {
			exp.Time.Sec = ts[0].(uint64)
			exp.Time.Nsec = ts[1].(uint64)
		}

		exports = append(exports, exp)
	}

	return timestamp, exports
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
