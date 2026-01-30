# 🎓 Guía para Estudiantes - DevOps Paso a Paso

## 👋 ¡Bienvenido al Mundo DevOps!

Esta guía te llevará desde cero hasta crear tu primer pipeline completo de DevOps. No te preocupes si eres principiante, ¡todos empezamos así!

## 📚 ¿Qué Aprenderás?

Al completar esta guía sabrás:
- ✅ Crear aplicaciones en Go
- ✅ Containerizar con Docker
- ✅ Orquestar con Kubernetes
- ✅ Configurar Ingress para acceso externo
- ✅ Implementar health checks
- ✅ Gestionar configuraciones y secretos

## 🎯 Objetivos de Aprendizaje

### Nivel Principiante
- Entender qué es DevOps
- Crear una aplicación simple
- Usar Docker básico
- Desplegar en Kubernetes local

### Nivel Intermedio
- Implementar CI/CD pipelines
- Configurar monitoreo
- Gestionar múltiples entornos
- Aplicar mejores prácticas de seguridad

## 📋 Prerrequisitos

### Conocimientos Básicos
- Línea de comandos de Linux
- Conceptos básicos de programación
- Git básico

### Herramientas Necesarias
Sigue la guía de instalación en `docs/instalacion.md`

## 🚀 Paso a Paso - Tu Primer Proyecto DevOps

### Fase 1: Preparación (15 minutos)

#### 1.1 Crear Repositorio
```bash
# En GitHub, crear repo: mi-go-devops
git clone https://github.com/tu-usuario/mi-go-devops.git
cd mi-go-devops
```

#### 1.2 Crear Estructura
```bash
mkdir -p cmd internal configs kubernetes .github/workflows docs scripts
```

#### 1.3 Inicializar Go
```bash
go mod init github.com/tu-usuario/mi-go-devops
```

**🎯 Objetivo:** Tener la estructura base del proyecto

### Fase 2: Aplicación Go (30 minutos)

#### 2.1 Crear la Aplicación
Crear `cmd/main.go` con:
- Servidor HTTP básico
- Endpoint `/` para página principal
- Endpoint `/health` para health checks
- Puerto configurable via variable de entorno

#### 2.2 Probar Localmente
```bash
go run cmd/main.go
curl http://localhost:8080/
curl http://localhost:8080/health
```

**🎯 Objetivo:** Aplicación funcionando localmente

### Fase 3: Docker (45 minutos)

#### 3.1 Crear Dockerfile
- Multi-stage build (builder + runtime)
- Imagen base mínima (Alpine)
- Usuario no-root
- Variables de entorno

#### 3.2 Crear Makefile
```makefile
build:
    docker build -t mi-go-devops .
run-docker:
    docker run -p 8080:8080 mi-go-devops
```

#### 3.3 Probar Docker
```bash
make build
make run-docker
curl http://localhost:8080/health
```

**🎯 Objetivo:** Aplicación corriendo en contenedor

### Fase 4: Kubernetes Local (60 minutos)

#### 4.1 Crear Cluster
```bash
kind create cluster --name devops-learning
```

#### 4.2 Crear Manifiestos
En carpeta `kubernetes/`:
- `deployment.yaml` - Despliegue con 2 réplicas
- `service.yaml` - Servicio interno
- `configmap.yaml` - Variables de configuración
- `secret.yaml` - Variables sensibles
- `ingress.yaml` - Acceso externo

#### 4.3 Desplegar
```bash
kind load docker-image mi-go-devops:latest --name devops-learning
kubectl apply -f kubernetes/
```

**🎯 Objetivo:** Aplicación corriendo en Kubernetes

### Fase 5: Verificación (15 minutos)

#### 5.1 Verificar Estado
```bash
kubectl get all,ingress
kubectl logs -f deployment/mi-go-devops
```

#### 5.2 Probar Aplicación
```bash
echo "127.0.0.1 mi-go-devops.local" | sudo tee -a /etc/hosts
curl http://mi-go-devops.local/
curl http://mi-go-devops.local/health
```

**🎯 Objetivo:** Aplicación accesible externamente

## 🧠 Conceptos Clave que Debes Entender

### DevOps
- **Qué es:** Cultura y prácticas que combinan desarrollo y operaciones
- **Por qué:** Entregas más rápidas, confiables y frecuentes
- **Cómo:** Automatización, colaboración, monitoreo continuo

### Containerización
- **Qué es:** Empaquetar aplicación con sus dependencias
- **Por qué:** Consistencia entre entornos
- **Cómo:** Docker crea contenedores portables

### Orquestación
- **Qué es:** Gestión automática de contenedores
- **Por qué:** Escalado, recuperación, balanceo de carga
- **Cómo:** Kubernetes maneja múltiples contenedores

### Infrastructure as Code
- **Qué es:** Infraestructura definida en archivos
- **Por qué:** Versionado, reproducible, auditable
- **Cómo:** Manifiestos YAML de Kubernetes

## 🔍 Ejercicios Prácticos

### Ejercicio 1: Modificar la Aplicación
1. Agregar endpoint `/version` que retorne la versión
2. Reconstruir imagen Docker
3. Redesplegar en Kubernetes
4. Verificar que funciona

### Ejercicio 2: Configuración
1. Agregar nueva variable en ConfigMap
2. Usar la variable en la aplicación
3. Aplicar cambios sin reconstruir imagen
4. Verificar que se actualiza

### Ejercicio 3: Escalado
1. Cambiar réplicas de 2 a 3 en deployment
2. Aplicar cambios
3. Verificar que hay 3 pods corriendo
4. Probar que el balanceo funciona

### Ejercicio 4: Health Checks
1. Hacer que `/health` falle intencionalmente
2. Observar cómo Kubernetes reinicia el pod
3. Arreglar el health check
4. Verificar que vuelve a funcionar

## 📊 Criterios de Evaluación

### Básico (Aprobado)
- ✅ Aplicación Go funcionando
- ✅ Docker build exitoso
- ✅ Despliegue en Kubernetes
- ✅ Endpoints respondiendo

### Intermedio (Bien)
- ✅ Todo lo básico +
- ✅ Health checks funcionando
- ✅ ConfigMap y Secrets configurados
- ✅ Ingress accesible externamente

### Avanzado (Excelente)
- ✅ Todo lo intermedio +
- ✅ Ejercicios completados
- ✅ Documentación propia
- ✅ Troubleshooting resuelto independientemente

## 🚨 Errores Comunes y Soluciones

### "ImagePullBackOff"
```bash
# Problema: Kubernetes no encuentra la imagen
# Solución:
kubectl patch deployment mi-go-devops -p '{"spec":{"template":{"spec":{"containers":[{"name":"mi-go-devops","imagePullPolicy":"Never"}]}}}}'
```

### "Port already in use"
```bash
# Problema: Puerto ocupado
# Solución:
sudo lsof -i :8080
sudo kill <PID>
```

### "Connection refused"
```bash
# Problema: Servicio no accesible
# Solución:
kubectl port-forward service/mi-go-devops-service 8080:80
```

## 📝 Entregables

### Repositorio Git con:
1. **Código fuente** completo y funcionando
2. **Dockerfile** optimizado
3. **Manifiestos** de Kubernetes
4. **README.md** con instrucciones
5. **Makefile** con comandos útiles

### Demostración:
1. **Clonar** tu repositorio
2. **Ejecutar** `make build && make run-docker`
3. **Desplegar** en Kubernetes
4. **Mostrar** aplicación funcionando
5. **Explicar** cada componente

## 🎓 Próximos Pasos

Una vez completado este proyecto, puedes avanzar a:

### CI/CD Pipelines
- GitHub Actions
- Automatización de tests
- Deploy automático

### Monitoreo
- Prometheus y Grafana
- Logs centralizados
- Alertas

### Seguridad
- Escaneo de vulnerabilidades
- Secrets management
- Network policies

### Múltiples Entornos
- Dev, Staging, Production
- GitOps con ArgoCD
- Helm charts

## 💡 Tips para el Éxito

1. **Lee los errores** - Los mensajes de error contienen la solución
2. **Usa los logs** - `kubectl logs` es tu mejor amigo
3. **Practica regularmente** - DevOps se aprende haciendo
4. **Documenta todo** - Tu yo del futuro te lo agradecerá
5. **No tengas miedo** - Romper cosas es parte del aprendizaje

## 🤝 Recursos Adicionales

### Documentación Oficial
- [Kubernetes Docs](https://kubernetes.io/docs/)
- [Docker Docs](https://docs.docker.com/)
- [Go Docs](https://golang.org/doc/)

### Comunidades
- [Kubernetes Slack](https://kubernetes.slack.com/)
- [Docker Community](https://www.docker.com/community/)
- [Go Community](https://golang.org/help/)

### Cursos Recomendados
- Kubernetes Basics
- Docker Fundamentals
- Go Programming

---

**¡Recuerda:** El objetivo no es memorizar comandos, sino entender los conceptos y saber dónde buscar ayuda cuando la necesites.

**¡Éxito en tu viaje DevOps!** 🚀