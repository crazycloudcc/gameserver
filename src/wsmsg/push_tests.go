package wsmsg

import (
	"cherry/netwebsocket"
)

// PushTests TODO
func PushTests(ev *netwebsocket.Event) {
	sendBytesToConn(ev.Connid, ev.Msgid, ev.MetaData)
}
