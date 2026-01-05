# Pi Zero Web Project (Go + Gin + Docker + Delve)

This project is a minimal web service written in **Go** using the **Gin** framework, designed to run on a Raspberry Pi Zero 2 W (ARM64) and to be easily developed, tested and debugged from a x86_64 workstation using Docker and VSCode.

It supports:

- Native x86 development and debugging
- Cross-compilation to ARM64
- ARM64 execution inside Docker (emulating a Pi Zero 2 W environment)
- Remote debugging via Delve

---

## Architecture Overview

| Layer | Purpose |
|------|----------|
Local host (x86_64) | Development, testing and debugging |
Docker ARM64 container | Emulates Raspberry Pi Zero 2 W runtime |
Go + Gin | Web service |
Delve | Remote debugger |

---

## Requirements

- Docker Engine + Docker Compose v2
- VSCode
- VSCode Go extension
- Go >= 1.25 (for local builds)
- Delve (automatically installed inside container)

---

## Project Structure

```
pi-zero-web-project/
├── docker-compose.yaml
├── Dockerfile.arm64
├── Dockerfile.debug
├── go.mod
├── go.sum
├── main.go
├── internal/
│   └── routes/
│       └── routes.go
└── .vscode/
├── launch.json
└── tasks.json (optional / empty)

```

---

## Running Locally (x86)

```sh
go run .
```

Service will be available at:

```
http://localhost:8080
```

---

## Building for ARM64

```sh
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o app-arm64
```

---

## Running in Docker (ARM64)

```sh
docker compose build
docker compose up
```

or only the debug container:

```sh
docker compose up debug
```

---

## Debugging with VSCode

#### 1. Start the debug container

```sh
docker compose up debug
```

Wait until you see:

```
API server listening at: [::]:2345
```

#### 2. Attach VSCode debugger
Press F5 and select:

```
Attach to Pi Zero (manual)
```

#### 3. Stop debugging

Press Stop in VSCode (detaches debugger)

Stop container manually:

```sh
docker compose stop debug
```

---

### Available Endpoints

| Endpoint | Description |
| --- | --- |
| GET /ping | Liveness test | 
| GET /welcome | Welcome message |
| GET /health | Health probe |

Example:

```sh
curl http://localhost:8080/ping
```

---

### Why this approach?

- Debugging directly inside Docker is slower and less stable

- x86 native debugging is fast and reliable

- ARM64 runtime is validated in Docker only when needed

- This mirrors real embedded development workflows

---

### Notes

- This project targets Raspberry Pi Zero 2 W (ARMv8-A / Cortex-A53).

- ARM64 is fully supported and recommended.

- 32-bit builds are not used.

---

### License

MIT (or your preferred license)

