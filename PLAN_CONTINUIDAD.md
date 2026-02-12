# 🧭 Plan de continuidad — 30-Days-Of-Go

Este plan propone cómo continuar el repositorio después de los primeros temas base, manteniendo una documentación útil para aprendizaje real y práctica profesional en Go.

## 🎯 Objetivos de continuidad

- Mantener una **cadencia sostenible** de publicación (2 entregas por semana).
- Profundizar desde fundamentos hasta temas de producción.
- Documentar cada módulo con estructura uniforme y referencias confiables.
- Convertir el repositorio en una guía viva para perfiles orientados a backend/DevOps.

## 🗓️ Cadencia recomendada

- **Martes:** teoría + ejemplos cortos + glosario del tema.
- **Jueves:** mini proyecto o kata aplicado + retrospectiva.
- **Fin de mes:** resumen mensual con aprendizajes, errores comunes y mejoras.

## 🧱 Estructura estándar por cada nuevo día (`Days/dayXX.md`)

1. **Introducción del tema**
2. **Conceptos clave** (con definiciones precisas)
3. **Ejemplos prácticos** (código ejecutable y salida esperada)
4. **Buenas prácticas**
5. **Errores comunes y cómo evitarlos**
6. **Checklist de aprendizaje**
7. **Recursos y referencias oficiales**

> Recomendación: mantener ejemplos pequeños y progresivos para evitar saturación cognitiva.

## 🛣️ Roadmap sugerido (siguientes bloques)

### Bloque A — Fundamentos intermedios

- Punteros y manejo de memoria.
- Métodos e interfaces.
- Manejo avanzado de errores.
- Paquetes y organización modular.

### Bloque B — Go orientado a backend

- JSON avanzado (`encoding/json`).
- HTTP servers con `net/http`.
- Middleware, validaciones y versionado de API.
- Conexión a base de datos (`database/sql`).

### Bloque C — Concurrencia y rendimiento

- Goroutines y channels.
- Patrones de concurrencia (worker pool, fan-in/fan-out).
- Context propagation (`context.Context`).
- Profiling y optimización.

### Bloque D — Producción y ecosistema

- Logging estructurado.
- Testing (unitario, integración y tabla-driven tests).
- CI/CD básico para Go.
- Contenerización y despliegue.

## 📝 Plantilla para documentar cada entrega

```md
# Day XX — <Tema>

## Introducción

## Conceptos clave

## Ejemplos prácticos

## Buenas prácticas

## Errores comunes

## Checklist
- [ ] Entiendo el concepto base.
- [ ] Puedo implementarlo desde cero.
- [ ] Sé identificar errores frecuentes.

## Referencias oficiales
- <enlace 1>
- <enlace 2>
```

## 📌 Criterios de calidad editorial

- Priorizar lenguaje claro, técnico y directo.
- Explicar el **porqué** de cada decisión, no solo el cómo.
- Mantener consistencia de estilo en títulos y bloques.
- Incluir al menos **2 referencias oficiales** por tema.

## 📚 Referencias oficiales base para todo el plan

- Go Documentation: https://go.dev/doc/
- A Tour of Go: https://go.dev/tour/
- Effective Go: https://go.dev/doc/effective_go
- Go Blog: https://go.dev/blog/
- Go Packages: https://pkg.go.dev/
- The Go Memory Model: https://go.dev/ref/mem
- Language Specification: https://go.dev/ref/spec

## ✅ Métrica de avance sugerida

- % de días documentados respecto al roadmap.
- Número de ejemplos ejecutables por módulo.
- Número de referencias oficiales por documento.
- Lista de “lecciones aprendidas” por semana.

Con este plan, el repositorio puede crecer con orden, profundidad técnica y una narrativa práctica para perfiles que quieren usar Go en contextos reales.


## ☁️ Línea de trabajo cloud-native (nuevo frente)

Para mantener continuidad técnica de nivel profesional, incorporar esta secuencia en paralelo:

1. Crear una API base con `net/http` y separación por capas.
2. Empaquetar con Docker multi-stage y ejecutar localmente.
3. Definir manifiestos Kubernetes (Deployment, Service, probes y recursos).
4. Añadir observabilidad mínima (logs estructurados y métricas).
5. Documentar incidentes reales y lecciones aprendidas por iteración.

Referencia adicional del repositorio:

- 📄 [Guía avanzada de arquitectura en Go](./GUIA_ARQUITECTURA_GO.md)
