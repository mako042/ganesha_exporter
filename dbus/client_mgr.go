package dbus

import (
	"log"

	"github.com/godbus/dbus"
	"golang.org/x/sys/unix"
)

// Client represents an NFS Ganesha client
type Client struct {
	Client   string
	NFSv3    bool
	MNTv3    bool
	NLMv4    bool
	RQUOTA   bool
	NFSv40   bool
	NFSv41   bool
	NFSv42   bool
	Plan9    bool
	LastTime unix.Timespec
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

func (mgr ClientMgr) ShowClients() (unix.Timespec, []Client) {
	call := mgr.dbusObject.Call("org.ganesha.nfsd.clientmgr.ShowClients", 0)
	if call.Err != nil {
		log.Panic(call.Err)
	}

	if len(call.Body) < 2 {
		log.Panic("invalid dbus response: ShowClients - not enough body elements")
	}

	// Parse header (tt) - unix.Timespec {Sec, Nsec}
	header, ok := call.Body[0].([]interface{})
	if !ok || len(header) < 2 {
		log.Panic("invalid header format")
	}
	utime := unix.Timespec{
		Sec:  int64(header[0].(uint64)),
		Nsec: int64(header[1].(uint64)),
	}

	// Parse clients array - it's [][]interface{} (array of arrays)
	clientsRaw, ok := call.Body[1].([][]interface{})
	if !ok {
		log.Panic("invalid clients format")
	}

	clients := make([]Client, 0, len(clientsRaw))

	for _, item := range clientsRaw {
		if len(item) < 5 {
			continue
		}

		client := Client{
			Client: item[0].(string),
		}

		// Parse protocols: ((sb)(sb)(sb)(sb)(sb)(sb)(sb)(sb)(sb))
		if protos, ok := item[1].([]interface{}); ok {
			for i, p := range protos {
				if i >= 9 {
					break
				}
				pair, ok := p.([]interface{})
				if !ok || len(pair) < 2 {
					continue
				}
				enabled := pair[1].(bool)
				switch i {
				case 0:
					client.NFSv3 = enabled
				case 1:
					client.MNTv3 = enabled
				case 2:
					client.NLMv4 = enabled
				case 3:
					client.RQUOTA = enabled
				case 4:
					client.NFSv40 = enabled
				case 5:
					client.NFSv41 = enabled
				case 6:
					client.NFSv42 = enabled
				case 7:
					client.Plan9 = enabled
				}
			}
		}

		// LastTime: (tt) - two uint64 values
		if ts, ok := item[2].([]interface{}); ok && len(ts) >= 2 {
			client.LastTime = unix.Timespec{
				Sec:  int64(ts[0].(uint64)),
				Nsec: int64(ts[1].(uint64)),
			}
		}

		clients = append(clients, client)
	}

	return utime, clients
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

func (mgr ClientMgr) GetNFSv42IO(ipaddr string) BasicStats {
	out := BasicStats{}
	call := mgr.dbusObject.Call("org.ganesha.nfsd.clientstats.GetNFSv42IO", 0, ipaddr)
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

func (mgr ClientMgr) GetNFSv42Layouts(ipaddr string) PNFSOperations {
	out := PNFSOperations{}
	call := mgr.dbusObject.Call("org.ganesha.nfsd.clientstats.GetNFSv42Layouts", 0, ipaddr)
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
