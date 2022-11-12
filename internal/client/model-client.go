package client

import (
	gate "crypto/internal/gate"
	stex "crypto/internal/stex"
)

type Clients struct {
	GateClient *gate.GateClient
	StexClient *stex.StexClient
}
