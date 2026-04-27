package dbus

import (
	"log"

	"github.com/godbus/dbus"
)

// Client Structure of the output of ShowClients dbus call
type Client struct {
	Name      string
	Protocols map[string]bool
	Last      uint64
	Stats     map[string]uint64
	Total     struct {
		A uint64
		B uint64
	}
}

// ClientMgr is a handle to dbus object ClientMgr
type ClientMgr struct {
	dbusObject dbus.BusObject
}

// NewClientMgr Get a new ClientMgr
func NewClientMgr() ClientMgr {
	conn, err := dbus.SystemBus()
	if err != nil {
		log.Panic(err)
	}
	return ClientMgr{
		dbusObject: conn.Object(
			"org.ganesha.nfsd",
			"/org/ganesha/nfsd/ClientMgr",
		),
	}
}

func (m ClientMgr) ShowClients() (uint64, []Client) {
	call := m.dbusObject.Call("org.ganesha.nfsd.clientmgr.ShowClients", 0)

	var raw []interface{}
	if err := call.Store(&raw); err != nil {
		log.Panic(err)
	}

	header := raw[0].([]interface{})
	clientsRaw := raw[1].([]interface{})

	var clients []Client

	for _, c := range clientsRaw {
		item := c.([]interface{})

		client := Client{
			Name:      item[0].(string),
			Protocols: map[string]bool{},
			Stats:     map[string]uint64{},
		}

		// protocols
		protos := item[1].([]interface{})
		for _, p := range protos {
			v := p.([]interface{})
			client.Protocols[v[0].(string)] = v[1].(bool)
		}

		client.Last = item[2].(uint64)

		// (ststst)
		stats := item[3].([]interface{})
		for i := 0; i < len(stats); i += 2 {
			key := stats[i].(string)
			val := stats[i+1].(uint64)
			client.Stats[key] = val
		}

		// (tt)
		total := item[4].([]interface{})
		client.Total.A = total[0].(uint64)
		client.Total.B = total[1].(uint64)

		clients = append(clients, client)
	}

	return header[0].(uint64), clients
}

func (mgr ClientMgr) GetNFSv3IO(ipaddr string) BasicStats {
	out := BasicStats{}
	call := mgr.dbusObject.Call("org.ganesha.nfsd.clientstats.GetNFSv3IO", 0, ipaddr)
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

func (mgr ClientMgr) GetNFSv40IO(ipaddr string) BasicStats {
	out := BasicStats{}
	call := mgr.dbusObject.Call("org.ganesha.nfsd.clientstats.GetNFSv40IO", 0, ipaddr)
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

func (mgr ClientMgr) GetNFSv41IO(ipaddr string) BasicStats {
	out := BasicStats{}
	call := mgr.dbusObject.Call("org.ganesha.nfsd.clientstats.GetNFSv41IO", 0, ipaddr)
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

func (mgr ClientMgr) GetNFSv41Layouts(ipaddr string) PNFSOperations {
	out := PNFSOperations{}
	call := mgr.dbusObject.Call("org.ganesha.nfsd.clientstats.GetNFSv41Layouts", 0, ipaddr)
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
