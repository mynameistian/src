package process2

import (
	"chatroom.com/common/message"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}
