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

    CREATE INDEX logs_user_ids ON public.logs(user_id);
    CREATE INDEX logs_user_ids_nulls_last ON public.logs(user_id DESC NULLS LAST);
    CREATE INDEX logs_severity ON public.logs(severity);
    CREATE INDEX logs_log_type ON public.logs(log_type);
    CREATE INDEX logs_section ON public.logs(section);
    CREATE INDEX logs_description ON public.logs(description);
    CREATE INDEX logs_happened_at ON public.logs(happened_at);
    CREATE INDEX logs_happened_at_desc ON public.logs(happened_at DESC);
EOSQL
