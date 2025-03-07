#!/bin/sh
set -e

if [ -f "/.env" ]; then
  . /.env
fi

echo "Starting migration with database: mysql://$DB_USER:$DB_PASSWORD@tcp($DB_HOST:$DB_PORT)/$DB_NAME"
migrate -path=/migration -database="mysql://$DB_USER:$DB_PASSWORD@tcp($DB_HOST:$DB_PORT)/$DB_NAME" up

echo "Migration completed successfully"