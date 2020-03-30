package consumer

import "github.com/xuzhuoxi/LegoMQ-go/message"

var (
	msgNil message.IMessageContext = nil

	msgEmpty = message.NewMessageContext("", "",
		"", nil)
	msgDefault = message.NewMessageContext("1",
		"192.168.1.1", nil, "default")
)
