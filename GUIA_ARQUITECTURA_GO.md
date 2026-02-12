# 🏗️ Guía avanzada de arquitectura en Go

Esta guía está orientada a llevar el aprendizaje de Go a un nivel profesional de ingeniería de software: arquitectura, metodologías de trabajo, microservicios y despliegue en contenedores/Kubernetes.

## 1) Arquitectura de software en Go: principios prácticos

En Go, la arquitectura debe priorizar:

- **Simplicidad estructural** (paquetes pequeños, acoplamiento bajo).
- **Composición sobre herencia** (structs + interfaces pequeñas).
- **Dependencias explícitas** (constructor functions + interfaces en bordes).
- **Observabilidad desde el diseño** (logs, métricas y tracing como parte del flujo).

### Enfoque recomendado: Clean/Hexagonal pragmática

No necesitas “dogma”, necesitas fronteras claras:

- `domain/`: reglas de negocio puras.
- `application/`: casos de uso.
- `infrastructure/`: DB, HTTP, mensajería, archivos.
- `interfaces/`: handlers REST/gRPC, CLI, jobs.

```txt
.
├── cmd/
│   └── api/main.go
├── internal/
│   ├── domain/
│   │   └── user.go
│   ├── application/
│   │   └── create_user.go
│   ├── infrastructure/
│   │   ├── postgres/
│   │   └── logger/
│   └── interfaces/
│       └── http/
├── pkg/
│   └── shared/
└── go.mod
```

> Regla de oro: `domain` no debe importar infraestructura.

---

## 2) Metodologías para desarrollar servicios en Go

### 2.1 Trunk-based + PRs pequeñas

- Commits atómicos y revisables.
- PRs cortas con contexto técnico.
- Deploy frecuente para reducir riesgo.

### 2.2 Vertical slices

Implementar por flujo de negocio completo (endpoint + caso de uso + persistencia), en lugar de “hacer primero todas las entidades”.

### 2.3 Contract-first APIs

Definir primero contrato (OpenAPI / Proto), luego implementación.

### 2.4 Quality gates mínimos

- `go test ./...`
- `go vet ./...`
- linters (`staticcheck`, `golangci-lint`)
- escaneo de dependencias en CI

---

## 3) Ejemplo de servicio HTTP con arquitectura por capas

### Domain

```go
// internal/domain/user.go
package domain

type User struct {
    ID    string
    Email string
    Name  string
}
```

### Application (caso de uso)

```go
// internal/application/create_user.go
package application

import "context"

type CreateUserInput struct {
    Email string
    Name  string
}

type UserRepository interface {
    Save(ctx context.Context, email, name string) (string, error)
}

type CreateUser struct {
    repo UserRepository
}

func NewCreateUser(repo UserRepository) *CreateUser {
    return &CreateUser{repo: repo}
}

func (uc *CreateUser) Execute(ctx context.Context, in CreateUserInput) (string, error) {
    return uc.repo.Save(ctx, in.Email, in.Name)
}
```

### Interface HTTP

```go
// internal/interfaces/http/handler_user.go
package http

import (
    "encoding/json"
    "net/http"

    "context"

    "your/module/internal/application"
)

type CreateUserUseCase interface {
    Execute(rctx context.Context, in application.CreateUserInput) (string, error)
}

type UserHandler struct {
    create CreateUserUseCase
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Email string `json:"email"`
        Name  string `json:"name"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "invalid json", http.StatusBadRequest)
        return
    }

    id, err := h.create.Execute(r.Context(), application.CreateUserInput{Email: req.Email, Name: req.Name})
    if err != nil {
        http.Error(w, "could not create user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    _ = json.NewEncoder(w).Encode(map[string]string{"id": id})
}
```

> Nota: el ejemplo es intencionalmente compacto. En producción, añade validaciones, errores tipados y observabilidad.

---

## 4) Microservicios en Go: cuándo sí y cuándo no

### Usa microservicios cuando:

- Tienes dominios de negocio claramente separados.
- Necesitas escalar componentes de forma independiente.
- Tu equipo puede sostener complejidad operativa (CI/CD, observabilidad, seguridad).

### Evítalos cuando:

- El producto está en etapa temprana.
- No hay límites de dominio claros.
- No tienes madurez operativa para mantener múltiples servicios.

### Patrones útiles en Go para microservicios

- API Gateway + BFF.
- Outbox pattern para consistencia eventual.
- Retry con backoff + circuit breaker.
- Idempotencia en operaciones críticas.
- `context.Context` propagado en toda la cadena.

---

## 5) Empaquetado en contenedores (Docker)

### Dockerfile multi-stage recomendado

```dockerfile
# Stage 1: build
FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/service ./cmd/api

# Stage 2: runtime
FROM gcr.io/distroless/static-debian12
WORKDIR /
COPY --from=builder /bin/service /service

EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/service"]
```

### Buenas prácticas de imagen

- Imagen final mínima (distroless/alpine según necesidad).
- Ejecutar como usuario no root.
- Variables de configuración por entorno, no hardcodeadas.
- Versionar imágenes con tags inmutables (`1.3.2`, `sha-...`).

---

## 6) Despliegue en Kubernetes (Pods/Deployments)

### Deployment base

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-api
spec:
  replicas: 2
  selector:
    matchLabels:
      app: go-api
  template:
    metadata:
      labels:
        app: go-api
    spec:
      containers:
        - name: api
          image: ghcr.io/tu-org/go-api:1.0.0
          ports:
            - containerPort: 8080
          env:
            - name: APP_ENV
              value: "production"
          resources:
            requests:
              cpu: "100m"
              memory: "128Mi"
            limits:
              cpu: "500m"
              memory: "512Mi"
          readinessProbe:
            httpGet:
              path: /healthz
              port: 8080
            initialDelaySeconds: 5
            periodSeconds: 10
          livenessProbe:
            httpGet:
              path: /livez
              port: 8080
            initialDelaySeconds: 10
            periodSeconds: 15
```

### Service base

```yaml
apiVersion: v1
kind: Service
metadata:
  name: go-api
spec:
  selector:
    app: go-api
  ports:
    - port: 80
      targetPort: 8080
  type: ClusterIP
```

### Checklist K8s para Go

- Health endpoints (`/healthz`, `/livez`, `/readyz`).
- Shutdown elegante con señales (`SIGTERM`) + timeout.
- Límites/requests definidos.
- Logs en stdout/stderr (sin archivos locales).
- Configuración externa con ConfigMaps/Secrets.

---

## 7) Observabilidad y operación

Mínimos de un servicio Go en producción:

- **Logging estructurado** (JSON, niveles, correlation IDs).
- **Métricas** (latencia, errores, throughput).
- **Tracing distribuido** (OpenTelemetry).
- **SLIs/SLOs** para medir confiabilidad.

Si no puedes observarlo, no puedes operarlo.

---

## 8) Plan de evolución sugerido para este repositorio

1. Documentar primero un servicio monolítico modular en Go.
2. Añadir módulo de observabilidad y testing.
3. Containerizar y desplegar en cluster local (`kind`/`k3d`).
4. Separar un módulo como microservicio (si hay frontera de dominio real).
5. Documentar trade-offs y costo operativo de la decisión.

---

## 📚 Referencias oficiales y recomendadas

### Go

- https://go.dev/doc/
- https://go.dev/ref/spec
- https://go.dev/doc/effective_go
- https://pkg.go.dev/net/http
- https://pkg.go.dev/context

### Contenedores y Kubernetes

- https://docs.docker.com/build/building/multi-stage/
- https://kubernetes.io/docs/concepts/workloads/pods/
- https://kubernetes.io/docs/concepts/workloads/controllers/deployment/
- https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/

### Arquitectura y prácticas

- https://12factor.net/
- https://opentelemetry.io/docs/

Con esta guía puedes pasar de “aprender sintaxis” a “diseñar y operar software Go en producción” con criterio técnico y metodología.
