#!/bin/sh
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
CREATE TABLE orders
(
    id          SERIAL PRIMARY KEY,
    description text,
    created_at  timestamp not null default now()
);

CREATE TABLE clients
(
    id         SERIAL PRIMARY KEY,
    name       text,
    created_at timestamp not null default now()
);

INSERT INTO orders(description)
values ('first_order'),
       ('second_order');
INSERT INTO clients(name)
values ('Top'),
       ('Bob');
EOSQL