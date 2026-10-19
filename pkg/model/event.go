// Package model defines canonical market-data events, the single
// interchange format between connectors, the bus and consumers.
// See spec.md §4.5.
package model

import (
	"encoding/json"
	"time"
)

// EventType is the canonical event kind.
type EventType string

// Canonical event kinds written to Redpanda topics md.*.
const (
	EventTypeTrade       EventType = "trade"
	EventTypeFunding     EventType = "funding"
	EventTypeOpenInteres EventType = "oi"
	EventTypeLiquidation EventType = "liq"
)

// Envelope is the canonical event wrapper written to Redpanda topics md.*.
type Envelope struct {
	V        int             `json:"v"`
	Type     EventType       `json:"type"`
	Exchange string          `json:"exchange"`
	Symbol   string          `json:"symbol"`
	TsEvent  time.Time       `json:"ts_event"`
	TsRecv   time.Time       `json:"ts_recv"`
	Seq      int64           `json:"seq"`
	Payload  json.RawMessage `json:"payload"`
}

// EnvelopeVersion is the current envelope schema version.
const EnvelopeVersion = 1

// TradePayload is a normalized trade tick.
type TradePayload struct {
	Price float64 `json:"price"`
	Qty   float64 `json:"qty"`
	Side  string  `json:"side"` // buy | sell
}

// FundingPayload is a normalized funding rate update.
type FundingPayload struct {
	Rate      float64   `json:"rate"`
	NextTs    time.Time `json:"next_ts"`
	MarkPrice float64   `json:"mark_price"`
}

// OpenInterestPayload is a normalized open interest snapshot.
type OpenInterestPayload struct {
	OI      float64 `json:"oi"`       // in contracts/base asset
	OIValue float64 `json:"oi_value"` // in quote currency (USD)
}

// LiquidationPayload is a normalized forced liquidation.
type LiquidationPayload struct {
	Price float64 `json:"price"`
	Qty   float64 `json:"qty"`
	Side  string  `json:"side"` // liquidated position side: long | short
}
