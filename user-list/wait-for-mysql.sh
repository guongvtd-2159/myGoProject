#!/bin/sh

HOST="$1"
PORT="$2"
CMD="$3"

echo "Waiting for MySQL at $HOST:$PORT..."

# Vòng lặp cho tới khi kết nối được
while ! nc -z $HOST $PORT; do
  echo "Waiting for MySQL at $HOST:$PORT..."
  sleep 2
done

echo "MySQL is up - running the app"
exec $CMD
