# 📦 Guía de Instalación - DevOps Tools

## 🛠️ Herramientas Principales

### 1. Git
```bash
# Ubuntu/Debian
sudo apt update && sudo apt install git

# Verificar
git --version
```

### 2. Go (Golang)
```bash
# Descargar e instalar Go 1.21
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz

# Agregar al PATH
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# Verificar
go version
```

### 3. Docker
```bash
# Instalar Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Agregar usuario al grupo docker
sudo usermod -aG docker $USER
newgrp docker

# Verificar
docker version
```

### 4. kubectl (Cliente de Kubernetes)
```bash
# Instalar con snap
sudo snap install kubectl --classic

# Verificar
kubectl version --client
```

### 5. kind (Kubernetes in Docker)
```bash
# Descargar kind
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.20.0/kind-linux-amd64

# Hacer ejecutable y mover
chmod +x ./kind
sudo mv ./kind /usr/local/bin/kind

# Verificar
kind version
```

## 🔧 Herramientas de Calidad

### golangci-lint
```bash
# Instalar
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2

# Verificar
golangci-lint --version
```

### gosec
```bash
# Instalar
go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest

# Verificar
gosec --version
```

### trivy
```bash
# Instalar
sudo apt-get update
sudo apt-get install wget apt-transport-https gnupg lsb-release
wget -qO - https://aquasecurity.github.io/trivy-repo/deb/public.key | sudo apt-key add -
echo "deb https://aquasecurity.github.io/trivy-repo/deb $(lsb_release -sc) main" | sudo tee -a /etc/apt/sources.list.d/trivy.list
sudo apt-get update
sudo apt-get install trivy

# Verificar
trivy --version
```

## ✅ Verificación Completa

```bash
# Ejecutar todos los comandos de verificación
echo "=== Verificando instalaciones ==="
git --version
go version
docker version
kubectl version --client
kind version
golangci-lint --version
gosec --version
trivy --version
echo "=== ¡Todas las herramientas instaladas! ==="
```

## 🚀 Siguiente Paso

Una vez instaladas todas las herramientas, puedes seguir la guía principal en `README.md` para crear tu primer proyecto DevOps.