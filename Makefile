SHELL := /bin/bash
.SHELLFLAGS := -eu -o pipefail -c

COMPOSE := docker compose -f deploy/docker-compose.yml
GOOSE := go run github.com/pressly/goose/v3/cmd/goose@latest
GOOSE_PG := -dir db/migrations postgres "postgres://stakan:$${PG_PASSWORD:-stakan}@localhost:5432/stakan?sslmode=disable"

.PHONY: up down logs ps migrate migrate-ch test test-integration lint gen build

up: ## Start infrastructure
	$(COMPOSE) up -d redpanda clickhouse postgres redis prometheus grafana

down: ## Stop everything
	$(COMPOSE) down

logs:
	$(COMPOSE) logs -f --tail=100

ps:
	$(COMPOSE) ps

migrate: ## goose up (PostgreSQL)
	$(GOOSE) $(GOOSE_PG) up

migrate-status:
	$(GOOSE) $(GOOSE_PG) status

migrate-ch: ## Apply ClickHouse schema
	docker exec -i stakan-clickhouse clickhouse-client --user stakan --password "$${CH_PASSWORD:-stakan}" --multiquery < db/clickhouse/001_init.sql

test: ## Unit tests (Go + web)
	go test ./...
	@if [ -d web/node_modules ]; then npm --prefix web test -- --run; else echo "web: node_modules missing, skip (run: npm --prefix web install)"; fi

test-integration: ## Integration tests (testcontainers)
	go test -tags=integration -count=1 ./...

lint:
	golangci-lint run
	@if [ -d web/node_modules ]; then npm --prefix web run lint; else echo "web: node_modules missing, skip"; fi

gen: ## Generate Go server from api/openapi.yaml
	@echo "oapi-codegen: wired in at stage 1 (api/openapi.yaml is a stub)"

build: ## Build all Go binaries + web
	go build ./...
	@if [ -d web/node_modules ]; then npm --prefix web run build; else echo "web: node_modules missing, skip"; fi
