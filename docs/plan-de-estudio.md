# 📚 Plan de Estudio - Curso DevOps Completo

## 🎯 Información del Curso

**Duración:** 8 semanas (32 horas académicas)  
**Modalidad:** Teórico-Práctico (70% práctica, 30% teoría)  
**Nivel:** Principiante a Intermedio  
**Prerrequisitos:** Conocimientos básicos de programación y Linux  

## 🏆 Objetivos de Aprendizaje

Al finalizar el curso, los estudiantes serán capaces de:
- ✅ Implementar pipelines de CI/CD completos
- ✅ Containerizar aplicaciones con Docker
- ✅ Orquestar servicios con Kubernetes
- ✅ Configurar monitoreo y logging
- ✅ Aplicar mejores prácticas de seguridad DevOps
- ✅ Gestionar infraestructura como código

---

## 📅 Cronograma Semanal

### **Semana 1: Fundamentos DevOps**
**Objetivos:** Entender qué es DevOps y preparar el entorno

#### Sesión 1.1 - Introducción a DevOps (2h)
**Teoría (45 min):**
- ¿Qué es DevOps? Historia y evolución
- Cultura DevOps vs metodologías tradicionales
- Beneficios: velocidad, calidad, colaboración
- Herramientas del ecosistema DevOps

**Práctica (75 min):**
- Instalación de herramientas (Git, Go, Docker, kubectl, kind)
- Configuración del entorno de desarrollo
- Verificación de instalaciones

**Entregable:** Entorno configurado y funcionando

#### Sesión 1.2 - Git y Colaboración (2h)
**Teoría (30 min):**
- Git workflows para DevOps
- Branching strategies (GitFlow, GitHub Flow)
- Pull Requests y Code Reviews

**Práctica (90 min):**
- Crear repositorio del proyecto
- Configurar estructura de carpetas
- Práctica con branches y merges
- Configurar protección de ramas

**Entregable:** Repositorio estructurado con README

---

### **Semana 2: Desarrollo de Aplicaciones**
**Objetivos:** Crear aplicación base y entender principios de desarrollo

#### Sesión 2.1 - Aplicación Go (2h)
**Teoría (30 min):**
- Principios de aplicaciones cloud-native
- Health checks y observabilidad
- Configuración por variables de entorno

**Práctica (90 min):**
- Desarrollar aplicación Go con endpoints REST
- Implementar health checks
- Configuración flexible con variables de entorno
- Testing básico

**Entregable:** Aplicación Go funcionando localmente

#### Sesión 2.2 - Calidad de Código (2h)
**Teoría (30 min):**
- Linting y análisis estático
- Seguridad en el código
- Testing strategies

**Práctica (90 min):**
- Configurar golangci-lint
- Implementar gosec para seguridad
- Escribir tests unitarios
- Configurar Makefile

**Entregable:** Aplicación con tests y linting

---

### **Semana 3: Containerización**
**Objetivos:** Dominar Docker y containerización

#### Sesión 3.1 - Docker Fundamentals (2h)
**Teoría (45 min):**
- ¿Qué son los contenedores?
- Docker vs VMs
- Imágenes, contenedores, registries
- Dockerfile best practices

**Práctica (75 min):**
- Crear Dockerfile multi-stage
- Build y run de contenedores
- Optimización de imágenes
- Docker Compose básico

**Entregable:** Aplicación containerizada

#### Sesión 3.2 - Docker Avanzado (2h)
**Teoría (30 min):**
- Seguridad en contenedores
- Networking y volumes
- Registry management

**Práctica (90 min):**
- Implementar usuario no-root
- Configurar health checks
- Escaneo de vulnerabilidades con Trivy
- Push a registry (GHCR)

**Entregable:** Imagen Docker segura y optimizada

---

### **Semana 4: Kubernetes Básico**
**Objetivos:** Entender orquestación con Kubernetes

#### Sesión 4.1 - Kubernetes Concepts (2h)
**Teoría (60 min):**
- Arquitectura de Kubernetes
- Pods, Services, Deployments
- ConfigMaps y Secrets
- Namespaces y Labels

**Práctica (60 min):**
- Crear cluster local con kind
- Desplegar primera aplicación
- Explorar recursos con kubectl

**Entregable:** Aplicación corriendo en Kubernetes

#### Sesión 4.2 - Manifiestos y Configuración (2h)
**Teoría (30 min):**
- YAML manifests
- Resource management
- Health checks en K8s

**Práctica (90 min):**
- Crear manifiestos completos
- Configurar ConfigMaps y Secrets
- Implementar liveness/readiness probes
- Configurar resource limits

**Entregable:** Manifiestos completos y funcionando

---

### **Semana 5: Kubernetes Avanzado**
**Objetivos:** Configurar acceso externo y persistencia

#### Sesión 5.1 - Services e Ingress (2h)
**Teoría (45 min):**
- Tipos de Services
- Ingress Controllers
- Load Balancing
- SSL/TLS termination

**Práctica (75 min):**
- Configurar Ingress NGINX
- Exponer aplicación externamente
- Configurar dominios locales
- Testing de conectividad

**Entregable:** Aplicación accesible externamente

#### Sesión 5.2 - Persistencia y Escalado (2h)
**Teoría (30 min):**
- Volumes y PersistentVolumes
- StatefulSets
- Horizontal Pod Autoscaling

**Práctica (90 min):**
- Configurar almacenamiento persistente
- Implementar escalado manual
- Configurar HPA básico
- Testing de escalado

**Entregable:** Aplicación escalable con persistencia

---

### **Semana 6: CI/CD Pipelines**
**Objetivos:** Automatizar testing y despliegue

#### Sesión 6.1 - Continuous Integration (2h)
**Teoría (45 min):**
- Principios de CI
- GitHub Actions
- Quality gates
- Testing automation

**Práctica (75 min):**
- Crear workflow de CI
- Configurar tests automáticos
- Implementar linting y security scans
- Branch protection rules

**Entregable:** Pipeline de CI funcionando

#### Sesión 6.2 - Continuous Deployment (2h)
**Teoría (30 min):**
- CD vs Continuous Delivery
- Deployment strategies
- Rollback strategies

**Práctica (90 min):**
- Crear workflow de CD
- Automatizar build y push de imágenes
- Deploy automático a Kubernetes
- Configurar environments

**Entregable:** Pipeline completo de CI/CD

---

### **Semana 7: Monitoreo y Observabilidad**
**Objetivos:** Implementar monitoreo y logging

#### Sesión 7.1 - Logging (2h)
**Teoría (45 min):**
- Structured logging
- Centralized logging
- Log aggregation

**Práctica (75 min):**
- Implementar logging estructurado
- Configurar log collection
- Análisis de logs con kubectl
- Debugging de aplicaciones

**Entregable:** Aplicación con logging estructurado

#### Sesión 7.2 - Métricas y Alertas (2h)
**Teoría (45 min):**
- Prometheus y Grafana
- Métricas de aplicación
- SLIs, SLOs, SLAs

**Práctica (75 min):**
- Exponer métricas de aplicación
- Configurar Prometheus básico
- Crear dashboards simples
- Configurar alertas básicas

**Entregable:** Monitoreo básico implementado

---

### **Semana 8: Proyecto Final y Mejores Prácticas**
**Objetivos:** Integrar todo el conocimiento en proyecto completo

#### Sesión 8.1 - Seguridad DevOps (2h)
**Teoría (60 min):**
- Security scanning
- Secrets management
- Network policies
- Supply chain security

**Práctica (60 min):**
- Implementar escaneo de seguridad
- Configurar secrets management
- Hardening de contenedores
- Security policies

**Entregable:** Aplicación con seguridad implementada

#### Sesión 8.2 - Proyecto Final (2h)
**Teoría (30 min):**
- GitOps principles
- Infrastructure as Code
- Mejores prácticas DevOps

**Práctica (90 min):**
- Presentación de proyectos finales
- Code review grupal
- Optimizaciones y mejoras
- Planificación de próximos pasos

**Entregable:** Proyecto DevOps completo

---

## 📊 Sistema de Evaluación

### Evaluación Continua (60%)
- **Entregables semanales:** 40%
- **Participación en clase:** 10%
- **Ejercicios prácticos:** 10%

### Evaluaciones Parciales (25%)
- **Examen teórico Semana 4:** 10%
- **Proyecto práctico Semana 6:** 15%

### Proyecto Final (15%)
- **Presentación:** 5%
- **Código y documentación:** 10%

### Criterios de Calificación
- **90-100:** Excelente - Implementación completa con mejores prácticas
- **80-89:** Bueno - Implementación funcional con algunas mejoras
- **70-79:** Satisfactorio - Implementación básica funcionando
- **60-69:** Suficiente - Implementación con errores menores
- **<60:** Insuficiente - Requiere refuerzo

---

## 🛠️ Recursos por Semana

### Semana 1-2: Fundamentos
- **Lecturas:** Documentación oficial de Git y Go
- **Videos:** "What is DevOps?" (YouTube)
- **Herramientas:** Git, Go, VS Code

### Semana 3-4: Containerización
- **Lecturas:** Docker documentation, Kubernetes basics
- **Videos:** Docker tutorials, K8s concepts
- **Herramientas:** Docker, kind, kubectl

### Semana 5-6: Orquestación y CI/CD
- **Lecturas:** Kubernetes networking, GitHub Actions
- **Videos:** Ingress tutorials, CI/CD best practices
- **Herramientas:** Ingress NGINX, GitHub Actions

### Semana 7-8: Observabilidad y Proyecto
- **Lecturas:** Observability patterns, Security best practices
- **Videos:** Monitoring tutorials, DevOps case studies
- **Herramientas:** Prometheus, Grafana, security scanners

---

## 📝 Entregables del Curso

### Repositorio Final debe incluir:
1. **Aplicación Go** completa y funcionando
2. **Dockerfile** optimizado y seguro
3. **Manifiestos K8s** completos
4. **Pipelines CI/CD** funcionando
5. **Documentación** completa
6. **Tests** automatizados
7. **Monitoreo** básico implementado

### Documentación Requerida:
- **README.md** con instrucciones completas
- **ARCHITECTURE.md** explicando decisiones técnicas
- **DEPLOYMENT.md** con guía de despliegue
- **TROUBLESHOOTING.md** con problemas comunes

---

## 🎯 Proyectos Adicionales (Opcionales)

### Para Estudiantes Avanzados:
1. **Multi-environment setup** (dev/staging/prod)
2. **GitOps con ArgoCD**
3. **Service Mesh con Istio**
4. **Terraform para IaC**

### Para Práctica Extra:
1. **Microservicios** con múltiples aplicaciones
2. **Database integration** con PostgreSQL
3. **Message queues** con Redis/RabbitMQ
4. **API Gateway** con Kong/Ambassador

---

## 📞 Soporte y Recursos

### Durante el Curso:
- **Office Hours:** Martes y jueves 4-5 PM
- **Slack/Discord:** Canal dedicado para dudas
- **Documentación:** Todas las guías en el repositorio

### Recursos Externos:
- [Kubernetes Documentation](https://kubernetes.io/docs/)
- [Docker Documentation](https://docs.docker.com/)
- [DevOps Roadmap](https://roadmap.sh/devops)
- [CNCF Landscape](https://landscape.cncf.io/)

---

## 🏆 Certificación

### Requisitos para Certificación:
- ✅ Asistencia mínima 80%
- ✅ Calificación final ≥ 70%
- ✅ Proyecto final aprobado
- ✅ Presentación exitosa

### Certificado Incluye:
- **Competencias adquiridas**
- **Herramientas dominadas**
- **Horas de formación**
- **Proyecto destacado**

---

**¡Bienvenido al mundo DevOps!** 🚀  
*"El objetivo no es ser perfecto desde el inicio, sino mejorar continuamente"*