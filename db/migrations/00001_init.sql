-- +goose Up
CREATE TABLE instruments (
    exchange            text        NOT NULL,
    symbol              text        NOT NULL,
    base                text        NOT NULL,
    quote               text        NOT NULL,
    contract_multiplier numeric     NOT NULL DEFAULT 1,
    status              text        NOT NULL DEFAULT 'trading',
    updated_at          timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (exchange, symbol)
);

CREATE TABLE alerts (
    id          uuid        NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id     uuid,
    tg_chat_id  bigint,
    name        text        NOT NULL DEFAULT '',
    rule        jsonb       NOT NULL,
    channels    text[]      NOT NULL DEFAULT '{tg}',
    enabled     boolean     NOT NULL DEFAULT true,
    created_at  timestamptz NOT NULL DEFAULT now(),
    CHECK (user_id IS NOT NULL OR tg_chat_id IS NOT NULL)
);

CREATE TABLE alert_events (
    id         bigint      GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    alert_id   uuid        NOT NULL REFERENCES alerts (id) ON DELETE CASCADE,
    symbol     text        NOT NULL,
    value      jsonb       NOT NULL,
    fired_at   timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX alert_events_alert_id_fired_at ON alert_events (alert_id, fired_at DESC);

-- +goose Down
DROP TABLE alert_events;
DROP TABLE alerts;
DROP TABLE instruments;
