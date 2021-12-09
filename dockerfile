FROM mysql:8.0.23
COPY ./init.sql /docker-entrypoint-initdb.d/