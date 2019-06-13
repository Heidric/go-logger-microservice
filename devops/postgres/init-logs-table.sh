#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE TABLE public.logs
    (
        id bigserial NOT NULL,
        user_id bigint,
        severity integer NOT NULL,
        log_type character varying(255) NOT NULL,
        section character varying(255) NOT NULL,
        description character varying(1000),
        happened_at timestamp with time zone NOT NULL,
        additional_data character varying(10000),
        PRIMARY KEY (id)
    )
    WITH (
        OIDS = FALSE
    );

    ALTER TABLE public.logs
        OWNER to $POSTGRES_USER;
EOSQL
