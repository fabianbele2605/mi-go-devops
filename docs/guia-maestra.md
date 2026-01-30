# 🚀 Guía Completa de DevOps para Principiantes

> **Tu tutor personal para dominar DevOps paso a paso**

¡Bienvenido a tu viaje en DevOps! Esta guía te llevará desde cero hasta crear un pipeline completo de CI/CD con Go, Docker y Kubernetes.

## 📋 Tabla de Contenidos

- [🎯 ¿Qué vas a lograr?](#-qué-vas-a-lograr)
- [🛠️ Preparando tu entorno](#️-preparando-tu-entorno)
- [📁 Creando la estructura del proyecto](#-creando-la-estructura-del-proyecto)
- [📝 Definiendo estándares](#-definiendo-estándares)
- [🐳 Containerización con Docker](#-containerización-con-docker)
- [⚸️ Orquestación con Kubernetes](#️-orquestación-con-kubernetes)
- [🏠 Cluster local para pruebas](#-cluster-local-para-pruebas)
- [🔄 Pipeline de Integración Continua (CI)](#-pipeline-de-integración-continua-ci)
- [🚀 Pipeline de Despliegue Continuo (CD)](#-pipeline-de-despliegue-continuo-cd)
- [🔐 Gestión de secretos](#-gestión-de-secretos)
- [✅ Validación final](#-validación-final)
- [⭐ Mejoras avanzadas](#-mejoras-avanzadas)
- [📚 Orden de implementación](#-orden-de-implementación)

---

## 🎯 ¿Qué vas a lograr?

Al finalizar esta guía tendrás un **repositorio Go profesional** que:

✅ **Construye** una imagen Docker reproducible y segura  
✅ **Pasa** tests automáticos, linting y análisis de seguridad  
✅ **Publica** la imagen en un registry (GitHub Container Registry)  
✅ **Despliega** automáticamente en Kubernetes  
✅ **Soporta** múltiples entornos (dev/staging/prod)  

**¿Qué es DevOps?** Es la práctica de automatizar y mejorar el proceso entre desarrollo y operaciones, permitiendo entregas más rápidas y confiables.

---

## 🛠️ Preparando tu entorno

### Herramientas principales que necesitas:

| Herramienta | ¿Para qué sirve? | Comando de verificación |
|-------------|------------------|------------------------|
| **Git** | Control de versiones | `git --version` |
| **Go** | Lenguaje de programación | `go version` |
| **Docker** | Containerización | `docker version` |
| **kubectl** | Cliente de Kubernetes | `kubectl version --client` |
| **kind/minikube** | Kubernetes local | `kind version` |

### Herramientas de calidad:

| Herramienta | ¿Para qué sirve? |
|-------------|------------------|
| **golangci-lint** | Análisis estático de código Go |
| **gosec** | Análisis de seguridad en Go |
| **trivy** | Escaneo de vulnerabilidades |

### 🔧 Instalación rápida:

```bash
# Verifica que todo esté instalado
go version && docker version && kubectl version --client
```

---

## 📁 Creando la estructura del proyecto

### Paso 1: Crear el repositorio

1. Ve a GitHub y crea un repo llamado `mi-go-devops`
2. Clónalo localmente:
   ```bash
   git clone https://github.com/tu-usuario/mi-go-devops.git
   cd mi-go-devops
   ```

### Paso 2: Estructura de carpetas estándar

```
mi-go-devops/
├── cmd/                    # Puntos de entrada de la aplicación
├── internal/               # Lógica interna (no exportable)
├── configs/                # Archivos de configuración
├── kubernetes/             # Manifiestos de Kubernetes
├── .github/workflows/      # Pipelines de CI/CD
├── docs/                   # Documentación
├── scripts/                # Scripts de automatización
├── README.md               # Este archivo
├── .gitignore             # Archivos a ignorar
├── Dockerfile             # Instrucciones para crear la imagen
├── Makefile               # Comandos automatizados
└── go.mod                 # Dependencias de Go
```

**¿Por qué esta estructura?** Sigue las convenciones de Go y facilita el mantenimiento y escalabilidad del proyecto.

---

## 📝 Definiendo estándares

### En tu README.md incluye:

- **Cómo ejecutar localmente** (modo desarrollo)
- **Cómo ejecutar con Docker**
- **Cómo correr tests y linting**
- **Variables de entorno** necesarias
- **Puertos y endpoints** (health checks)

### Convenciones de Git:

- **Rama principal:** `main`
- **Rama de desarrollo:** `develop`
- **Pull Request obligatorio** para merge a `main`
- **Protección de ramas** (requiere CI verde)

**¿Por qué estándares?** Facilitan la colaboración en equipo y mantienen la calidad del código.

---

## 🐳 Containerización con Docker

### ¿Qué es Docker?
Docker empaqueta tu aplicación con todas sus dependencias en un "contenedor" que puede ejecutarse en cualquier lugar de manera consistente.

### Checklist para un Dockerfile profesional:

✅ **Multi-stage build** (separar compilación de ejecución)  
✅ **Imagen base mínima** (distroless o alpine)  
✅ **Usuario no-root** (seguridad)  
✅ **Variables de entorno** configurables  
✅ **Puerto expuesto** correctamente  
✅ **Healthcheck** (opcional si usas Kubernetes)  

### Prueba local:

```bash
# Construir la imagen
docker build -t mi-app .

# Ejecutar el contenedor
docker run -p 8080:8080 mi-app

# Probar el endpoint
curl http://localhost:8080/health
```

---

## ⚸️ Orquestación con Kubernetes

### ¿Qué es Kubernetes?
Kubernetes orquesta contenedores, manejando escalado, recuperación ante fallos y balanceo de carga automáticamente.

### Manifiestos mínimos necesarios:

#### 1. **Deployment** (`kubernetes/deployment.yaml`)
- **Réplicas:** 2 (alta disponibilidad)
- **Recursos:** requests y limits de CPU/memoria
- **Probes:** liveness y readiness
- **Variables:** via ConfigMap/Secret
- **Seguridad:** SecurityContext no-root

#### 2. **Service** (`kubernetes/service.yaml`)
- **Tipo:** ClusterIP (interno)
- **Puerto:** mapeo correcto

#### 3. **Ingress** (`kubernetes/ingress.yaml`)
- **Controller:** Nginx
- **Routing:** host/path

#### 4. **ConfigMap** (`kubernetes/configmap.yaml`)
- **Variables no sensibles:** APP_ENV, LOG_LEVEL

#### 5. **Secret** (`kubernetes/secret.yaml`)
- **Variables sensibles:** DB_PASSWORD, tokens

---

## 🏠 Cluster local para pruebas

### Opción A: kind (recomendado)

```bash
# Crear cluster
kind create cluster --name devops-learning

# Verificar
kubectl get nodes
```

### Opción B: minikube

```bash
# Iniciar cluster
minikube start

# Habilitar ingress
minikube addons enable ingress
```

**¿Por qué local primero?** Te permite experimentar sin costos y aprender los conceptos básicos.

---

## 🔄 Pipeline de Integración Continua (CI)

### ¿Qué es CI?
Integración Continua ejecuta automáticamente tests y validaciones cada vez que haces cambios al código.

### Archivo: `.github/workflows/ci.yml`

#### Disparadores:
- Pull requests a `develop` y `main`
- Push a `develop`

#### Pasos del pipeline:

1. **Checkout** del código
2. **Setup Go** con cache
3. **Tests** con coverage
4. **Linting** con golangci-lint
5. **Seguridad** con gosec
6. **Build Docker** (validación)
7. **Escaneo** con trivy (opcional)

#### Regla de oro:
**Si cualquier paso falla → CI falla → PR no se puede mergear**

---

## 🚀 Pipeline de Despliegue Continuo (CD)

### ¿Qué es CD?
Despliegue Continuo automatiza la publicación y despliegue de tu aplicación cuando el código está listo.

### Archivo: `.github/workflows/cd.yml`

#### Disparadores:
- Push a `main`
- Tags `v*` para releases

#### Etapa 1: Build & Push
- **Login** a GitHub Container Registry
- **Build** con buildx
- **Tags:** `sha-<commit>`, `latest`, `vX.Y.Z`

#### Etapa 2: Deploy

**Enfoque simple:**
```bash
kubectl apply -f kubernetes/
kubectl set image deployment/mi-app container=nueva-imagen
```

**Enfoque GitOps (recomendado):**
- Actions solo publica imagen
- ArgoCD/Flux despliega automáticamente
- Más seguro y auditable

---

## 🔐 Gestión de secretos

### En GitHub (Settings → Secrets and variables → Actions):

#### Para registry:
- `GITHUB_TOKEN` (automático con permisos)
- O `DOCKERHUB_USERNAME` + `DOCKERHUB_TOKEN`

#### Para Kubernetes:
- `KUBECONFIG` (contenido del archivo)
- `KUBE_NAMESPACE`

#### Environments (opcional):
- `dev` (automático)
- `staging` (con aprobación)
- `prod` (con aprobación manual)

---

## ✅ Validación final

### Checklist de funcionamiento:

1. ✅ **CI pasa** al abrir PR
2. ✅ **Merge a main** exitoso
3. ✅ **CD construye** y publica imagen
4. ✅ **Deploy en K8s** exitoso
5. ✅ **Pods corriendo:** `kubectl get pods`
6. ✅ **Logs limpios:** `kubectl logs deployment/mi-app`
7. ✅ **Endpoint responde:** port-forward o Ingress

### Comandos de verificación:

```bash
# Estado general
kubectl get deploy,po,svc,ing

# Logs de la aplicación
kubectl logs -f deployment/mi-app

# Probar endpoint
kubectl port-forward svc/mi-app 8080:80
curl http://localhost:8080/health
```

---

## ⭐ Mejoras avanzadas

### Para un proyecto "profesional":

- 🔍 **Escaneo de imagen** con Trivy en CD
- 📋 **SBOM** (Software Bill of Materials) con Syft
- ✍️ **Firma de imagen** con Cosign
- 🎛️ **Kustomize** para múltiples entornos
- 🔄 **ArgoCD** para GitOps
- 📊 **Métricas** y observabilidad
- 🛡️ **Rate limiting** en Ingress

---

## 📚 Orden de implementación

### Para no enredarte, sigue este orden:

1. **Estructura** + README + go.mod
2. **Docker** build/run local
3. **Kubernetes** local con manifests
4. **CI** (tests/lint/security)
5. **CD** (push imagen)
6. **CD** (deploy) o GitOps
7. **Mejoras** avanzadas

---

## 🎓 ¿Listo para empezar?

**¡Perfecto!** Ahora tienes una hoja de ruta clara. Recuerda:

- 🐌 **Ve paso a paso** - no trates de hacer todo a la vez
- 🧪 **Experimenta** - rompe cosas y aprende de los errores
- 📚 **Documenta** - anota lo que aprendes
- 🤝 **Pregunta** - la comunidad DevOps es muy colaborativa

**¡Comencemos tu viaje DevOps!** 🚀