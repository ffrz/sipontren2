# STAGE 1: BUILDER (Hanya untuk instalasi AIR)
# Gunakan image Go resmi (1.21-alpine) 
FROM golang:1.21-alpine AS builder

# Set working directory di dalam container
WORKDIR /app

# Install air versi v1.45.0 (Kompatibel dengan Go 1.21)
# Ini diperlukan agar Go compiler bisa menjalankan AIR di stage 2
RUN go install github.com/cosmtrek/air@v1.45.0

# STAGE 2: DEVELOPMENT RUNTIME ENVIRONMENT
# Gunakan image Go yang sama agar toolchain 'go' tersedia untuk AIR
FROM golang:1.21-alpine

# Tambahkan bash dan client PostgreSQL (untuk pg_isready di wait-for-db.sh)
RUN apk update && apk add --no-cache bash postgresql-client

# Setel WORKDIR
WORKDIR /app

# Copy binary air dari stage 'builder'
COPY --from=builder /go/bin/air /usr/local/bin/air

# Copy script wait-for-db dan berikan izin eksekusi
COPY ./wait-for-db.sh /usr/local/bin/wait-for-db.sh
RUN chmod +x /usr/local/bin/wait-for-db.sh

# Volume mount dari docker-compose.yml akan menyediakan semua kode sumber di /app
# Kita tidak perlu COPY kode, hanya menjalankan entrypoint.

EXPOSE 8080

# Command yang dieksekusi saat container dijalankan adalah script wait-for-db
# Skrip ini akan menunggu DB siap, lalu menjalankan AIR
ENTRYPOINT ["/usr/local/bin/wait-for-db.sh"]
