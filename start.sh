#!/bin/sh

set -e

echo "Postgres is up - executing migrations"

source /app/app.env
/app/migrate -dir ./migration/ -v postgres "$DB_SOURCE" up

echo "Migration complete - starting the app"

echo "start the app"
exec "$@"
