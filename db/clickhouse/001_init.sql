-- ClickHouse schema, spec.md §7.1. Applied by `make migrate-ch`.

CREATE DATABASE IF NOT EXISTS md;

CREATE TABLE IF NOT EXISTS md.trades (
    exchange LowCardinality(String),
    symbol   LowCardinality(String),
    ts       DateTime,
    price    Float64,
    qty      Float64,
    side     LowCardinality(String),
    ver      DateTime DEFAULT now64(3)
) ENGINE = ReplacingMergeTree(ver)
  ORDER BY (exchange, symbol, ts)
  TTL ts + INTERVAL 14 DAY;

CREATE TABLE IF NOT EXISTS md.funding (
    exchange   LowCardinality(String),
    symbol     LowCardinality(String),
    ts         DateTime,
    rate       Float64,
    next_ts    DateTime,
    mark_price Float64,
    ver        DateTime DEFAULT now64(3)
) ENGINE = ReplacingMergeTree(ver)
  ORDER BY (exchange, symbol, ts)
  TTL ts + INTERVAL 400 DAY;

CREATE TABLE IF NOT EXISTS md.open_interest (
    exchange LowCardinality(String),
    symbol   LowCardinality(String),
    ts       DateTime,
    oi       Float64,
    oi_value Float64,
    ver      DateTime DEFAULT now64(3)
) ENGINE = ReplacingMergeTree(ver)
  ORDER BY (exchange, symbol, ts)
  TTL ts + INTERVAL 400 DAY;

CREATE TABLE IF NOT EXISTS md.liquidations (
    exchange LowCardinality(String),
    symbol   LowCardinality(String),
    ts       DateTime,
    price    Float64,
    qty      Float64,
    side     LowCardinality(String),
    ver      DateTime DEFAULT now64(3)
) ENGINE = ReplacingMergeTree(ver)
  ORDER BY (exchange, symbol, ts)
  TTL ts + INTERVAL 400 DAY;

CREATE TABLE IF NOT EXISTS md.candles_1m (
    exchange LowCardinality(String),
    symbol   LowCardinality(String),
    ts       DateTime,
    open     AggregateFunction(argMin, Float64, DateTime),
    high     AggregateFunction(max, Float64),
    low      AggregateFunction(min, Float64),
    close    AggregateFunction(argMax, Float64, DateTime),
    vol      AggregateFunction(sum, Float64)
) ENGINE = AggregatingMergeTree
  ORDER BY (exchange, symbol, ts);

CREATE MATERIALIZED VIEW IF NOT EXISTS md.candles_1m_mv TO md.candles_1m AS
SELECT
    exchange,
    symbol,
    toStartOfMinute(ts) AS ts,
    argMinState(price, ts) AS open,
    maxState(price) AS high,
    minState(price) AS low,
    argMaxState(price, ts) AS close,
    sumState(qty) AS vol
FROM md.trades
GROUP BY exchange, symbol, ts;
