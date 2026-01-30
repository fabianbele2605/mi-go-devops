# ETAPA 1: Construcción (Builder)
# Usa imagen con Go completo para compilar
FROM golang:1.21-alpine AS builder

# Directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia archivos de dependencias primero (para cache de Docker)
COPY go.mod go.sum ./

# Descarga dependencias (se cachea si go.mod no cambia)
RUN go mod download

# Copia todo el código fuente
COPY . .

# Compila la aplicación (binario estático sin CGO)
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

# ETAPA 2: Runtime (Imagen final)
# Usa imagen mínima para ejecutar
FROM alpine:latest

# Crea usuario no-root por seguridad
RUN adduser -D -s /bin/sh appuser

# Directorio de trabajo
WORKDIR /app

# Copia solo el binario compilado desde la etapa builder
COPY --from=builder /app/main .

# Cambia a usuario no-root
USER appuser

# Puerto que expone la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]

