package communication

import (
	"github.com/nats-io/nats.go"
	"time"
)

type Manager struct {
	*nats.Conn
}

func NewManager(conn string) (m *Manager, err error) {
	m = &Manager{}
	m.Conn, err = nats.Connect(
		conn,
		nats.Compression(true),
		nats.NoEcho(),
		nats.ReconnectWait(time.Second),
		nats.MaxPingsOutstanding(10),
		nats.MaxReconnects(10),
	)

	return
}
