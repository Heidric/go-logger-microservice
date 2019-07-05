FROM postgres

COPY ./init-logs-table.sh /docker-entrypoint-initdb.d/init-logs-table.sh

RUN chmod +x /docker-entrypoint-initdb.d/init-logs-table.sh
