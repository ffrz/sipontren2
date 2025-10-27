#!/bin/bash
# wait-for-db.sh: Skrip ini menunggu hingga database siap sebelum menjalankan AIR.

set -e

# Ambil variabel lingkungan yang diatur oleh docker-compose.yml
host="$DB_HOST"
user="$DB_USER"
name="$DB_NAME"
port="$DB_PORT"
password="$DB_PASSWORD"
cmd="/usr/local/bin/air"

echo "Waiting for PostgreSQL ($host:$port) to be ready..."

# Loop pengecekan koneksi menggunakan pg_isready
until PGPASSWORD=$password pg_isready -h "$host" -p "$port" -U "$user" -d "$name"; do
  >&2 echo "PostgreSQL is unavailable (DB not ready or DB '$name' not created) - sleeping..."
  sleep 1
done

>&2 echo "PostgreSQL is up and database '$name' is available. Starting AIR..."
# Eksekusi AIR setelah DB siap
exec $cmd
