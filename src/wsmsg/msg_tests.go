package wsmsg

import (
	"cherry/base"
	"cherry/netmsgparser"
	"cherry/netwebsocket"
	"protos"
)

// MsgTests TODO
type MsgTests struct {
	req protos.ReqTest
	res protos.ResTest
}

// Handle TODO
func (owner *MsgTests) Handle(ev *netwebsocket.Event) {
	if err := netmsgparser.Unmarshal(ev.MetaData, &owner.req); err != nil {
		base.LogError("msg_tests.go - netmsgparser unmarshal: ", err)
		return
	}

	// debug.
	base.LogDebug("msg_tests.go - req info: ", ev.Connid, ev.Msgid, owner.req.String())

	owner.res.Info = owner.req.GetInfo()
	sendToConn(ev.Connid, ev.Msgid, &owner.res)
}
