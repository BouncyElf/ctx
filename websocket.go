package ctx

import (
	"time"

	"github.com/gorilla/websocket"
)

// TODO: impl

type WebSocket struct {
}

type WebSocketConf struct {
	HandshakeTimeout time.Duration

	Subprotocols    []string
	ReadBufferSize  int
	WriteBufferSize int
	WriteBufferPool websocket.BufferPool

	Error       func(c *Context, status int, reason error)
	CheckOrigin func(c *Context) bool

	EnableCompression bool
}

func DefaultWebSocketConf() *WebSocketConf {
	return &WebSocketConf{
		HandshakeTimeout: 10 * time.Second,
	}
}

func (c *Context) ToWebSocket(confs ...*WebSocketConf) (*WebSocket, error) {
	conf := DefaultWebSocketConf()
	if len(confs) != 0 {
		conf = confs[0]
	}
	_ = conf
	return nil, nil
}
