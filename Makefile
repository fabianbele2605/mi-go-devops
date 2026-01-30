# Declara que estos targets no son archivos
.PHONY: build run test lint

# Construye la imagen Docker
build:
	docker build -t mi-go-devops .

# Ejecuta la aplicación localmente (sin Docker)
run:
	go run cmd/main.go

# Ejecuta la aplicación en Docker
run-docker:
	docker run -p 8080:8080 mi-go-devops

# Ejecuta todos los tests con output detallado
test:
	go test -v ./...

# Ejecuta el linter para verificar calidad de código
lint:
	golangci-lint run