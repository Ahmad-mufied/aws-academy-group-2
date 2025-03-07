FROM migrate/migrate:v4.15.2

WORKDIR /migration

COPY ./migration /migration
COPY ./docker/scripts/migrate_entrypoint.sh /entrypoint.sh
COPY .env.docker /.env

RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/bin/sh", "/entrypoint.sh"]