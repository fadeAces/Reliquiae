package integration

import (
	"github.com/nats-io/go-nats"
)

type Trilobita struct {
	conn *nats.Conn
} 

func NewThirdClient() {
	tri := Trilobita{}
	nc, _ := nats.Connect(nats.DefaultURL)
	tri.conn = nc
	
}