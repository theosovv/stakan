// Package connectors holds the exchange connector contract (spec.md §6.1)
// and per-exchange adapters in subpackages.
package connectors

import (
	"context"
	"time"

	"github.com/theosovv/stakan/pkg/model"
)

// Status reports connector health for /healthz.
type Status struct {
	Exchange    string    `json:"exchange"`
	Connected   bool      `json:"connected"`
	LastEventAt time.Time `json:"last_event_at,omitempty"`
	Buffered    int       `json:"buffered"`
	Dropped     uint64    `json:"dropped"`
}

// ExchangeConnector normalizes a single exchange's WS streams into
// canonical events. A new exchange is one subpackage implementing this
// interface.
type ExchangeConnector interface {
	Connect(ctx context.Context) error
	Subscribe(symbols []string) error
	Events() <-chan model.Envelope
	Health() Status
}
