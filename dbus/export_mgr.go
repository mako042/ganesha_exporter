package dbus

import (
	"log"

	"github.com/godbus/dbus"
)

type Export struct {
	ExportID  uint16
	Path      string
	Protocols map[string]bool

	Last uint64
	Time struct {
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

	var raw []interface{}
	if err := call.Store(&raw); err != nil {
		log.Panic(err)
	}

	if len(raw) < 2 {
		log.Panic("invalid dbus response: ShowExports")
	}

	header, ok := raw[0].([]interface{})
	if !ok || len(header) == 0 {
		log.Panic("invalid header format")
	}

	exportsRaw, ok := raw[1].([]interface{})
	if !ok {
		log.Panic("invalid exports format")
	}

	exports := make([]Export, 0, len(exportsRaw))

	for _, e := range exportsRaw {
		item, ok := e.([]interface{})
		if !ok || len(item) < 5 {
			continue
		}

		exp := Export{
			ExportID:  item[0].(uint16),
			Path:      item[1].(string),
			Protocols: map[string]bool{},
			Last:      item[3].(uint64),
		}

		// protocols: [(string bool), ...]
		if flags, ok := item[2].([]interface{}); ok {
			for _, f := range flags {
				pair, ok := f.([]interface{})
				if !ok || len(pair) < 2 {
					continue
				}

				name, _ := pair[0].(string)
				val, _ := pair[1].(bool)

				exp.Protocols[name] = val
			}
		}

		// time: [sec, nsec]
		if ts, ok := item[4].([]interface{}); ok && len(ts) >= 2 {
			exp.Time.Sec, _ = ts[0].(uint64)
			exp.Time.Nsec, _ = ts[1].(uint64)
		}

		exports = append(exports, exp)
	}

	return header[0].(uint64), exports
}
