package dbus

import (
	"log"

	"github.com/godbus/dbus"
)

// Export Structure of the output of ShowExports dbus call
type Export struct {
	ExportID uint16
	Path     string

	Protocols map[string]bool

	Last uint64
	Time struct {
		Sec  uint64
		Nsec uint64
	}
}

// ExportMgr is a handle to dbus object ExportMgr
type ExportMgr struct {
	dbusObject dbus.BusObject
}

// NewExportMgr Get a new ExportMgr
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

	var raw []interface{}
	if err := call.Store(&raw); err != nil {
		log.Panic(err)
	}

	// raw = [ (tt), array ]
	header := raw[0].([]interface{})
	exportsRaw := raw[1].([]interface{})

	var exports []Export

	for _, e := range exportsRaw {
		item := e.([]interface{})

		exp := Export{
			ExportID:  item[0].(uint16),
			Path:      item[1].(string),
			Protocols: map[string]bool{},
		}

		// ((sb)(sb)...)
		flags := item[2].([]interface{})
		for _, f := range flags {
			pair := f.([]interface{})
			name := pair[0].(string)
			val := pair[1].(bool)
			exp.Protocols[name] = val
		}

		exp.Last = item[3].(uint64)

		ts := item[4].([]interface{})
		exp.Time.Sec = ts[0].(uint64)
		exp.Time.Nsec = ts[1].(uint64)

		exports = append(exports, exp)
	}

	return header[0].(uint64), exports
}

func (mgr ExportMgr) GetNFSv3IO(exportID uint16) BasicStats {
	out := BasicStats{}
	call := mgr.dbusObject.Call("org.ganesha.nfsd.exportstats.GetNFSv3IO", 0, exportID)
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

func (mgr ExportMgr) GetNFSv40IO(exportID uint16) BasicStats {
	out := BasicStats{}
	call := mgr.dbusObject.Call("org.ganesha.nfsd.exportstats.GetNFSv40IO", 0, exportID)
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

func (mgr ExportMgr) GetNFSv41IO(exportID uint16) BasicStats {
	out := BasicStats{}
	call := mgr.dbusObject.Call("org.ganesha.nfsd.exportstats.GetNFSv41IO", 0, exportID)
	if call.Err != nil {
		log.Panic(call.Err)
	}
	if !call.Body[0].(bool) {
		if err := call.Store(&out.Status, &out.Error, &out.Time); err != nil {
			log.Panic(err)
		}
		return out
	}
	if Gandi {
		if err := call.Store(
			&out.Status, &out.Error, &out.Time,
			&out.Read, &out.Write,
			&out.Open, &out.Close, &out.Getattr, &out.Lock,
		); err != nil {
			log.Panic(err)
		}
	} else {
		if err := call.Store(
			&out.Status, &out.Error, &out.Time,
			&out.Read, &out.Write,
		); err != nil {
			log.Panic(err)
		}
	}
	return out
}

func (mgr ExportMgr) GetNFSv41Layouts(exportID uint16) PNFSOperations {
	out := PNFSOperations{}
	call := mgr.dbusObject.Call("org.ganesha.nfsd.exportstats.GetNFSv41Layouts", 0, exportID)
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
