FROM migrate/migrate:v4.15.2

WORKDIR /migrations

COPY ./migration /migration

ENTRYPOINT ["migrate"]

CMD ["-path", "/migrations", "-database", "mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}", "up"]
