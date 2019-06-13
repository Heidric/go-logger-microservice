FROM postgres

RUN POSTGRES_USER=postgres
RUN POSTGRES_PASSWORD=postgres
RUN POSTGRES_DB=logger

COPY ./init-logs-table.sh /docker-entrypoint-initdb.d/init-logs-table.sh

RUN chmod +x /docker-entrypoint-initdb.d/init-logs-table.sh
