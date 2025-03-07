FROM mysql:latest

WORKDIR /seeders

COPY ./seeders /seeders
COPY ./docker/scripts/seeder_entrypoint.sh /entrypoint.sh
COPY .env.docker /.env

RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/bin/sh", "/entrypoint.sh"]