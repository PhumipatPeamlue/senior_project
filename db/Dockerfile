FROM mysql:8.2

LABEL authors="Besto"


COPY init.sql /docker-entrypoint-initdb.d/

ENV MYSQL_ROOT_PASSWORD=root

EXPOSE 3306