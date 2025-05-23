#!/bin/sh

HOST="$1"
PORT="$2"
shift 2

echo "Waiting for MySQL at $HOST:$PORT..."

# Vòng lặp chờ MySQL
while ! nc -z $HOST $PORT; do
  echo "Waiting for MySQL at $HOST:$PORT..."
  sleep 2
done

echo "MySQL is up - running the app"
exec "$@"
