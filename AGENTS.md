# marketdemo

## Objetivo
Proyecto de demostración: infraestructura completa desplegada
automáticamente con Terraform, Docker y Docker Swarm.

## Stack
- Cloud: DigitalOcean
- IaC: Terraform (HCP Terraform)
- Contenedores: Docker
- Orquestación: Docker Swarm. No migrar a Kubernetes ni introducir
  otro orquestador salvo petición explícita.
- Ingress / proxy inverso: Traefik
- CI/CD: GitHub Actions
- Datos y mensajería: PostgreSQL (migraciones con Flyway), RabbitMQ, Redis
- Servicios: Celery (beat-service, scrap-service), BFF (Go), frontend.

## Comandos
Usa el Makefile como punto de entrada. Ejecuta `make help` o lee el
Makefile antes de lanzar comandos directos de terraform o docker.
Los targets que empiezan con `_` son internos: no los llames.

## Estructura
| Carpeta          | Contenido                                  |
| ---------------- | ------------------------------------------ |
| `beat-service/`  | Celery Beat: programa tareas periódicas    |
| `bff-service/`   | Backend for Frontend (API para el front)   |
| `front/`         | Aplicación web                             |
| `scrap-service/` | Scraping y recolección de datos            |
| `flyway/`        | Migraciones de base de datos               |
| `postgres/`      | Configuración de PostgreSQL                |
| `rabbitmq/`      | Configuración de RabbitMQ                  |
| `redis/`         | Configuración de Redis                     |
| `config/`        | Configuración compartida                   |
| `infra/`         | Docker Swarm, Traefik, imágenes            |
| `terraform/`     | Recursos cloud (IaC)                       |
| `scripts/`       | Scripts auxiliares                         |
| `e2e/`           | Pruebas end-to-end                         |
| `.github/`       | Workflows de CI/CD                         |

## Verificación
Tras cambiar código de un servicio, ejecuta los tests relevantes:

- `make test-unit`: unitarios del BFF (`./controller`, con `-race`)
- `make test-integration`: integración del BFF (`./integration`)
- `make test-frontend`: tests del frontend (modo no interactivo)
- `make e2e`: solo lo ejecuto yo; no lo lances (requiere TTY y navegador)
- Terraform: `terraform fmt -check` y `terraform validate`

Notas sobre los tests:
- Los tests de integración usan un Postgres de **pruebas** con datos
  sembrados (p. ej. esperan 30 productos). No cambies migraciones ni
  seeds sin avisarme: pueden romper los tests aunque el código esté bien.
- En integración la autenticación se reemplaza con un `FakeValidator`.
  Nunca uses ese patrón fuera de tests.
- Si añades o cambias un endpoint, propón el test correspondiente.
- `scrap-service` y `beat-service` no tienen tests.

## Reglas de seguridad
Prohibido sin mi aprobación explícita:
- `terraform apply`, `terraform destroy` y `terraform plan`
  (el plan pregúntamelo siempre).
- Cualquier comando que modifique un Swarm real: `docker stack deploy`,
  `docker stack rm`, `docker service update/rm`, `docker swarm ...`.
- Modificar workflows de `.github/` que manejen secretos.

Secretos:
- No leas ni modifiques `.env`, `*.tfvars` ni `~/.environments/*.env`
  (el Makefile los usa, tú no).
- No imprimas tokens ni claves. Evita comandos que puedan filtrarlos
  (p. ej. `docker inspect`, `docker compose config`, `env`).

Costos: los recursos en DigitalOcean generan gasto; evita acciones que
creen o amplíen recursos.

## Git
- Trabaja en una rama, nunca directamente en `main`.
- No hagas commit ni push sin que te lo pida.

## Convenciones
- Formatea Terraform con `terraform fmt` antes de proponer cambios.
- Respeta las versiones y dependencias definidas en cada servicio.
  Pregúntame antes de instalar dependencias nuevas o añadir servicios.
- Cambios pequeños y enfocados; explica qué cambiaste y por qué.