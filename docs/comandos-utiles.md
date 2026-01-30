# 🔧 Comandos Útiles - DevOps

## 🐹 Comandos de Go

### Desarrollo
```bash
# Inicializar módulo Go
go mod init github.com/usuario/proyecto

# Descargar dependencias
go mod tidy

# Ejecutar aplicación
go run cmd/main.go

# Compilar binario
go build -o app cmd/main.go

# Ejecutar tests
go test ./...

# Tests con coverage
go test -cover ./...

# Verificar código
go vet ./...
```

### Linting y Calidad
```bash
# Ejecutar golangci-lint
golangci-lint run

# Análisis de seguridad
gosec ./...

# Formatear código
go fmt ./...
```

## 🐳 Comandos de Docker

### Básicos
```bash
# Construir imagen
docker build -t nombre-imagen .

# Ejecutar contenedor
docker run -p 8080:8080 nombre-imagen

# Ejecutar en segundo plano
docker run -d -p 8080:8080 nombre-imagen

# Ver contenedores corriendo
docker ps

# Ver todas las imágenes
docker images

# Detener contenedor
docker stop <container-id>

# Eliminar contenedor
docker rm <container-id>

# Eliminar imagen
docker rmi <image-id>
```

### Limpieza
```bash
# Limpiar contenedores parados
docker container prune

# Limpiar imágenes no usadas
docker image prune

# Limpieza completa
docker system prune -a
```

## ⚸️ Comandos de Kubernetes

### Cluster Management
```bash
# Ver información del cluster
kubectl cluster-info

# Ver nodos
kubectl get nodes

# Ver todos los recursos
kubectl get all

# Ver recursos específicos
kubectl get pods
kubectl get services
kubectl get deployments
kubectl get ingress
```

### Despliegue
```bash
# Aplicar manifiestos
kubectl apply -f archivo.yaml
kubectl apply -f carpeta/

# Eliminar recursos
kubectl delete -f archivo.yaml

# Ver detalles de un recurso
kubectl describe pod nombre-pod
kubectl describe service nombre-service

# Ver logs
kubectl logs nombre-pod
kubectl logs -f deployment/nombre-deployment
```

### Debugging
```bash
# Ejecutar comando en pod
kubectl exec -it nombre-pod -- /bin/sh

# Port forwarding
kubectl port-forward pod/nombre-pod 8080:8080
kubectl port-forward service/nombre-service 8080:80

# Ver eventos
kubectl get events --sort-by=.metadata.creationTimestamp

# Ver recursos con más detalles
kubectl get pods -o wide
kubectl get all -o yaml
```

## 🏠 Comandos de kind

### Cluster Management
```bash
# Crear cluster
kind create cluster --name mi-cluster

# Crear cluster con configuración
kind create cluster --name mi-cluster --config config.yaml

# Ver clusters
kind get clusters

# Eliminar cluster
kind delete cluster --name mi-cluster

# Cargar imagen Docker al cluster
kind load docker-image mi-imagen:latest --name mi-cluster
```

## 🔍 Comandos de Troubleshooting

### Verificar Estado
```bash
# Estado general del cluster
kubectl get componentstatuses

# Recursos por namespace
kubectl get all -n kube-system

# Uso de recursos
kubectl top nodes
kubectl top pods

# Verificar conectividad
kubectl run test-pod --image=busybox --rm -it -- /bin/sh
```

### Logs y Debugging
```bash
# Logs de múltiples pods
kubectl logs -l app=mi-app

# Logs anteriores (si el pod se reinició)
kubectl logs nombre-pod --previous

# Describir recursos para ver eventos
kubectl describe deployment mi-deployment
kubectl describe pod mi-pod

# Ver configuración actual
kubectl get deployment mi-deployment -o yaml
```

## 📊 Comandos de Monitoreo

### Estado de la Aplicación
```bash
# Verificar health checks
curl http://localhost:8080/health

# Verificar endpoints
curl -i http://localhost:8080/

# Monitorear logs en tiempo real
kubectl logs -f deployment/mi-app

# Ver métricas básicas
kubectl top pods
```

## 🚀 Makefile Útil

```makefile
.PHONY: build run test lint clean deploy

# Desarrollo local
build:
	go build -o bin/app cmd/main.go

run:
	go run cmd/main.go

test:
	go test -v ./...

lint:
	golangci-lint run

# Docker
docker-build:
	docker build -t mi-app .

docker-run:
	docker run -p 8080:8080 mi-app

# Kubernetes
k8s-deploy:
	kubectl apply -f kubernetes/

k8s-delete:
	kubectl delete -f kubernetes/

k8s-logs:
	kubectl logs -f deployment/mi-app

# Limpieza
clean:
	docker system prune -f
	go clean -cache
```

## 💡 Tips Útiles

### Aliases Recomendados
```bash
# Agregar a ~/.bashrc
alias k='kubectl'
alias kgp='kubectl get pods'
alias kgs='kubectl get services'
alias kgd='kubectl get deployments'
alias kdp='kubectl describe pod'
alias kl='kubectl logs'
```

### Variables de Entorno
```bash
# Para desarrollo
export KUBECONFIG=~/.kube/config
export GO111MODULE=on
export GOPROXY=https://proxy.golang.org
```