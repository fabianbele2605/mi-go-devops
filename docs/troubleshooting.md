# 🔧 Troubleshooting - Solución de Problemas Comunes

## 🐹 Problemas de Go

### Error: "go: cannot find main module"
```bash
# Problema: No hay go.mod en el directorio
# Solución:
go mod init github.com/usuario/proyecto
```

### Error: "package not found"
```bash
# Problema: Dependencias no descargadas
# Solución:
go mod tidy
go mod download
```

### Error: "permission denied"
```bash
# Problema: Permisos de archivo
# Solución:
chmod +x archivo
# O para el binario compilado:
chmod +x ./app
```

## 🐳 Problemas de Docker

### Error: "docker: permission denied"
```bash
# Problema: Usuario no está en grupo docker
# Solución:
sudo usermod -aG docker $USER
newgrp docker
# O reiniciar sesión
```

### Error: "no space left on device"
```bash
# Problema: Disco lleno por imágenes/contenedores
# Solución:
docker system prune -a
docker volume prune
```

### Error: "port already in use"
```bash
# Problema: Puerto ocupado
# Solución 1: Usar otro puerto
docker run -p 8081:8080 mi-app

# Solución 2: Encontrar y matar proceso
sudo lsof -i :8080
sudo kill <PID>
```

### Error: "failed to build"
```bash
# Problema: Error en Dockerfile
# Solución: Verificar sintaxis y rutas
# Común: COPY . . debe estar después de WORKDIR
```

## ⚸️ Problemas de Kubernetes

### Error: "ImagePullBackOff"
```bash
# Problema: No puede descargar imagen
# Solución 1: Imagen local (kind)
kubectl patch deployment mi-app -p '{"spec":{"template":{"spec":{"containers":[{"name":"mi-app","imagePullPolicy":"Never"}]}}}}'

# Solución 2: Cargar imagen a kind
kind load docker-image mi-app:latest --name mi-cluster

# Solución 3: Verificar nombre de imagen
kubectl describe pod <pod-name>
```

### Error: "CrashLoopBackOff"
```bash
# Problema: Pod se reinicia constantemente
# Solución: Ver logs para identificar error
kubectl logs <pod-name>
kubectl logs <pod-name> --previous

# Verificar health checks
kubectl describe pod <pod-name>
```

### Error: "Pending" pods
```bash
# Problema: Pod no puede ser programado
# Solución: Verificar recursos y nodos
kubectl describe pod <pod-name>
kubectl get nodes
kubectl top nodes

# Verificar si hay suficientes recursos
kubectl describe node <node-name>
```

### Error: "Service Unavailable"
```bash
# Problema: Service no puede alcanzar pods
# Solución: Verificar selectors y labels
kubectl get pods --show-labels
kubectl describe service <service-name>

# Verificar endpoints
kubectl get endpoints <service-name>
```

## 🏠 Problemas de kind

### Error: "cluster already exists"
```bash
# Problema: Cluster con mismo nombre existe
# Solución:
kind delete cluster --name mi-cluster
kind create cluster --name mi-cluster
```

### Error: "failed to create cluster"
```bash
# Problema: Docker no está corriendo
# Solución:
sudo systemctl start docker
# O en algunos sistemas:
sudo service docker start
```

### Error: "context not found"
```bash
# Problema: kubectl no encuentra el cluster
# Solución:
kubectl config get-contexts
kubectl config use-context kind-mi-cluster
```

## 🌐 Problemas de Ingress

### Error: "404 Not Found"
```bash
# Problema: Ingress no está funcionando
# Solución 1: Verificar controlador de Ingress
kubectl get pods -n ingress-nginx

# Solución 2: Verificar configuración
kubectl describe ingress mi-ingress

# Solución 3: Verificar /etc/hosts
cat /etc/hosts | grep mi-app.local
```

### Error: "connection refused"
```bash
# Problema: Ingress controller no está listo
# Solución: Esperar a que esté ready
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s
```

## 🔍 Comandos de Diagnóstico

### Verificación General
```bash
# Estado del cluster
kubectl cluster-info
kubectl get componentstatuses

# Estado de nodos
kubectl get nodes -o wide
kubectl describe node <node-name>

# Estado de pods
kubectl get pods -o wide
kubectl get pods --all-namespaces
```

### Logs y Eventos
```bash
# Ver eventos recientes
kubectl get events --sort-by=.metadata.creationTimestamp

# Logs de sistema
kubectl logs -n kube-system <pod-name>

# Logs de aplicación
kubectl logs -f deployment/mi-app
kubectl logs <pod-name> --previous
```

### Conectividad
```bash
# Test de conectividad interna
kubectl run test-pod --image=busybox --rm -it -- /bin/sh
# Dentro del pod:
nslookup mi-service
wget -qO- http://mi-service/health

# Port forwarding para debug
kubectl port-forward pod/<pod-name> 8080:8080
kubectl port-forward service/<service-name> 8080:80
```

## 🚨 Problemas Críticos

### Cluster No Responde
```bash
# Reiniciar cluster kind
kind delete cluster --name mi-cluster
kind create cluster --name mi-cluster

# Verificar Docker
docker ps
sudo systemctl status docker
```

### Pods en Estado Unknown
```bash
# Problema: Nodo perdió conexión
# Solución: Eliminar pods manualmente
kubectl delete pod <pod-name> --force --grace-period=0

# Verificar estado del nodo
kubectl get nodes
kubectl describe node <node-name>
```

### Out of Memory/CPU
```bash
# Problema: Recursos insuficientes
# Solución: Ajustar limits en deployment
kubectl patch deployment mi-app -p '{"spec":{"template":{"spec":{"containers":[{"name":"mi-app","resources":{"limits":{"memory":"256Mi","cpu":"200m"}}}]}}}}'
```

## 💡 Tips de Prevención

### Buenas Prácticas
1. **Siempre usar health checks** en tus aplicaciones
2. **Configurar resource limits** en Kubernetes
3. **Usar imagePullPolicy: Never** para imágenes locales
4. **Verificar logs regularmente** con `kubectl logs -f`
5. **Mantener Docker limpio** con `docker system prune`

### Comandos de Monitoreo
```bash
# Crear alias útiles
alias kgp='kubectl get pods'
alias kl='kubectl logs'
alias kd='kubectl describe'

# Script de verificación rápida
#!/bin/bash
echo "=== Estado del Cluster ==="
kubectl get nodes
echo "=== Pods ==="
kubectl get pods
echo "=== Services ==="
kubectl get services
echo "=== Ingress ==="
kubectl get ingress
```

## 📞 ¿Necesitas Más Ayuda?

Si encuentras un problema no listado aquí:

1. **Revisa los logs** con `kubectl logs`
2. **Describe el recurso** con `kubectl describe`
3. **Verifica eventos** con `kubectl get events`
4. **Busca en la documentación oficial** de Kubernetes
5. **Pregunta en la comunidad** con detalles específicos del error