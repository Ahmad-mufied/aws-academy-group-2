#!/bin/sh
set -e

if [ -f "/.env" ]; then
  . /.env
fi

echo "Running database seeders..."
mysql -h$DB_HOST -u$DB_USER -p$DB_PASSWORD $DB_NAME < /seeders/roles_status_seeder.sql

echo "Seeding completed successfully"