# MIAFIN - Sistema Financiero Comunal

Prototipo funcional de sistema financiero con arquitectura simple.

## Stack

### Backend
- Go 1.22+
- Gin framework
- GORM
- PostgreSQL
- JWT authentication
- gofpdf para generación de PDFs

### Frontend
- Svelte (puro, no SvelteKit)
- Bun runtime
- Diseño minimalista con tipografía monospace

### Infraestructura
- Docker
- docker-compose

## Características

- Autenticación JWT
- Roles: admin y asesor
- Gestión de clientes
- Gestión de grupos comunales
- Préstamos individuales y comunales
- Registro de pagos
- Cálculo de liquidez
- Generación de reportes en PDF
- Logs de auditoría

## Instrucciones de ejecución

### Con Docker (Recomendado)

1. Clonar el repositorio:
```bash
git clone https://github.com/LuneLoops/MIAFIN.git
cd MIAFIN
git checkout prototype
```

2. Levantar todos los servicios:
```bash
docker-compose up --build
```

3. Acceder a la aplicación:
- Frontend: http://localhost:3000
- Backend API: http://localhost:8080

4. Credenciales por defecto:
- Email: admin@miafin.local
- Password: admin123

### Sin Docker (Desarrollo local)

#### Backend

1. Asegurarse de tener PostgreSQL corriendo

2. Configurar variables de entorno (opcional):
```bash
export DB_HOST=localhost
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=miafin
export DB_PORT=5432
export PORT=8080
```

3. Ejecutar backend:
```bash
cd backend
go mod download
go run ./cmd/api
```

#### Frontend

1. Instalar dependencias:
```bash
cd frontend
bun install
```

2. Ejecutar en modo desarrollo:
```bash
bun run dev
```

3. Acceder a http://localhost:5173

## API Endpoints

### Públicos
- POST /login - Autenticación

### Protegidos (requieren token)

#### Usuarios (admin)
- GET /usuarios
- POST /usuarios

#### Clientes
- GET /clientes
- POST /clientes

#### Grupos (crear solo admin)
- GET /grupos
- POST /grupos

#### Préstamos
- GET /prestamos
- POST /prestamos

#### Pagos
- POST /pagos

#### Reportes
- GET /reportes/liquidez
- GET /reportes/general/pdf
- GET /reportes/prestamo/:id/pdf
- GET /reportes/pago/:id/pdf

#### Logs (admin)
- GET /logs

## Modelo de datos

- Usuario (admin, asesor)
- Cliente
- Grupo
- Prestamo (comunal, individual)
- Pago
- LogTransaccion

## Estructura del proyecto

```
MIAFIN/
├── backend/
│   ├── cmd/
│   │   └── api/
│   │       └── main.go
│   ├── internal/
│   │   ├── auth/
│   │   ├── database/
│   │   ├── handlers/
│   │   ├── middleware/
│   │   └── services/
│   ├── pkg/
│   │   └── models/
│   ├── Dockerfile
│   └── go.mod
├── frontend/
│   ├── src/
│   │   ├── api.js
│   │   ├── main.js
│   │   ├── App.svelte
│   │   ├── Login.svelte
│   │   ├── Dashboard.svelte
│   │   ├── Clientes.svelte
│   │   ├── Grupos.svelte
│   │   ├── Prestamos.svelte
│   │   ├── Pagos.svelte
│   │   ├── Usuarios.svelte
│   │   ├── Logs.svelte
│   │   └── style.css
│   ├── Dockerfile
│   ├── nginx.conf
│   ├── package.json
│   └── vite.config.js
└── docker-compose.yml
```

## Notas

- Este es un prototipo con conceptos financieros reales pero simplificados
- No cumple estrictamente con regulaciones ASFI
- Los préstamos calculan interés simple
- La liquidez se calcula como: Total recuperado - Total saldo pendiente
- Todos los logs quedan registrados en la base de datos
- Los PDFs se generan en el backend y se descargan desde el frontend

## Detener los servicios

```bash
docker-compose down
```

Para eliminar también los volúmenes:
```bash
docker-compose down -v
```
