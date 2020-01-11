package wsmsg

import (
	"cherry/netwebsocket"
	"protos"

	"github.com/golang/protobuf/proto"
)

var wsService *netwebsocket.Service

// Init 初始化
func Init(serv *netwebsocket.Service) {
	wsService = serv
	regHandler()
}

// regHandler 注册协议回调
func regHandler() {
	wsService.RegHandler(uint32(protos.MSGID_MSG_TEST), MsgTests{})
}

// 发送proto.message结构.
func sendToConn(connID int32, msgID uint32, pb proto.Message) {
	buf, _ := proto.Marshal(pb)
	sendBytesToConn(connID, msgID, buf)
}

// 发送bytes结构.
func sendBytesToConn(connID int32, msgID uint32, buf []byte) {
	wsService.GetClientMngr().SendToConn(connID, &netwebsocket.Msg{ID: msgID, MetaData: buf})
}

// 广播proto.message结构.
func sendToAll(msgID uint32, pb proto.Message) {
	buf, _ := proto.Marshal(pb)
	sendBytesToAll(msgID, buf)
}

// 广播bytes结构.
func sendBytesToAll(msgID uint32, buf []byte) {
	wsService.GetClientMngr().SendToAll(&netwebsocket.Msg{ID: msgID, MetaData: buf})
}

// 广播proto.message结构. (过滤发起者)
func sendWithOut(connID int32, msgID uint32, pb proto.Message) {
	buf, _ := proto.Marshal(pb)
	sendBytesWithOut(connID, msgID, buf)
}

// 广播bytes结构. (过滤发起者)
func sendBytesWithOut(connID int32, msgID uint32, buf []byte) {
	wsService.GetClientMngr().SendWithOut(&netwebsocket.Msg{ID: msgID, MetaData: buf}, connID)
}
