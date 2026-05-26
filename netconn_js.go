package websocket

import "net"

func (nc *netConn) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (nc *netConn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }
