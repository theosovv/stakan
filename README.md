# Stakan

Real-time screener and alerting platform for crypto futures markets.

Aggregates funding rates, open interest, liquidations and trades from
multiple exchanges (Binance, Bybit, OKX, Bitget) over WebSocket, streams
them through a Kafka-compatible pipeline into ClickHouse, and serves a
live React dashboard with sub-second Telegram alerts.

Monorepo: Go services (connectors, writer, alert engine, gRPC/REST+WS
gateway, Telegram bot) + React/TypeScript SPA in /web.

Stack: Go · React · TypeScript · Redpanda (Kafka API) · ClickHouse ·
PostgreSQL · Redis · Docker

## Quick start

```bash
make up          # Redpanda, ClickHouse, PostgreSQL, Redis, Prometheus, Grafana
make migrate     # goose up (PostgreSQL)
make migrate-ch  # ClickHouse schema (md.*)
go run ./cmd/gateway                    # API stub → localhost:8080
npm --prefix web install && npm --prefix web run dev   # SPA → localhost:5173
```

Workflow: `spec.md` is the single source of truth; `AGENTS.md` defines
the human/agent split (human writes implementation, AI writes tests,
specs and reviews). Decisions live in `docs/adr/`.