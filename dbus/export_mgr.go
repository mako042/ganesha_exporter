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

	// DEBUG: Print raw body
	log.Printf("DEBUG: ShowExports call.Body length: %d", len(call.Body))
	for i, v := range call.Body {
		log.Printf("DEBUG: call.Body[%d] type=%T value=%v", i, v, v)
	}

	var raw []interface{}
	if err := call.Store(&raw); err != nil {
		log.Printf("DEBUG: Store error: %v", err)
		log.Panic(err)
	}

	log.Printf("DEBUG: raw length: %d", len(raw))
	for i, v := range raw {
		log.Printf("DEBUG: raw[%d] type=%T value=%v", i, v, v)
		if slice, ok := v.([]interface{}); ok {
			for j, sv := range slice {
				log.Printf("DEBUG:   raw[%d][%d] type=%T value=%v", i, j, sv, sv)
				if innerSlice, ok := sv.([]interface{}); ok {
					for k, isv := range innerSlice {
						log.Printf("DEBUG:     raw[%d][%d][%d] type=%T value=%v", i, j, k, isv, isv)
					}
				}
			}
		}
	}

	if len(raw) < 2 {
		log.Panic("invalid dbus response: ShowExports")
	}

	// Parse header (tt) - two uint64 values
	header, ok := raw[0].([]interface{})
	if !ok || len(header) < 2 {
		log.Panic("invalid header format")
	}
	timestamp := header[0].(uint64)

	// Parse exports array
	exportsRaw, ok := raw[1].([]interface{})
	if !ok {
		log.Panic("invalid exports format")
	}

	exports := make([]Export, 0, len(exportsRaw))

	for idx, e := range exportsRaw {
		log.Printf("DEBUG: processing export[%d] type=%T value=%v", idx, e, e)
		item, ok := e.([]interface{})
		if !ok || len(item) < 5 {
			log.Printf("DEBUG: export[%d] has %d elements, skipping", idx, len(item))
			continue
		}

		log.Printf("DEBUG: export[%d] ExportID=%v Path=%v", idx, item[0], item[1])
		exp := Export{
			ExportID: item[0].(uint16),
			Path:     item[1].(string),
			Last:     item[3].(uint64),
		}

		// Parse protocols: ((sb)(sb)(sb)(sb)(sb)(sb)(sb))
		if protos, ok := item[2].([]interface{}); ok {
			log.Printf("DEBUG: export[%d] protocols count=%d", idx, len(protos))
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
