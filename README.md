# 🚀 NovaERP — Enterprise Resource Planning Backend API

[![Go Version](https://img.shields.io/badge/Go-1.22%2B-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Framework](https://img.shields.io/badge/Framework-Gin-008080?style=flat&logo=gin)](https://gin-gonic.com/)
[![Database](https://img.shields.io/badge/Database-PostgreSQL-4169E1?style=flat&logo=postgresql)](https://www.postgresql.org/)
[![ORM](https://img.shields.io/badge/ORM-GORM-00599C?style=flat)](https://gorm.io/)
[![Security](https://img.shields.io/badge/Auth-RS256%20JWT%20%2B%20Refresh%20Tokens-red?style=flat&logo=jsonwebtokens)](https://jwt.io/)
[![API Docs](https://img.shields.io/badge/OpenAPI-Swagger%202.0-85EA2D?style=flat&logo=swagger)](http://localhost:8080/swagger/index.html)

NovaERP is a high-performance, modular monolithic Enterprise Resource Planning (ERP) backend built in **Go (Golang)** with **Gin** and **PostgreSQL**. Designed with domain-driven modular submodules, robust database migrations, RS256 JWT security, and interactive OpenAPI documentation.

---

## 🌟 Key Features & Architectural Highlights

* 🔐 **RS256 JWT Authentication & Token Rotation**: Asymmetric RSA 2048-bit key signing for access tokens alongside database-persisted refresh tokens with cryptographic rotation.
* 📦 **Modular Inventory System**: Submodules for **Products**, **Warehouses**, and **Stocks** with row-level locked transactional stock adjustments.
* 👥 **Comprehensive HR & Payroll Engine**:
  * **Leave Management**: Submodules for **Leave Types** (Casual, Sick, Annual) and **Leave Applications & Approval Workflows**.
  * **Attendance System**: Daily **Check-In/Check-Out** with automated work-hour calculations, late-arrival tagging, and overtime tracking.
  * **Payroll Processor**: Monthly **Payroll Batches** and automated **Payslip Generation** for active employees.
* 🛡️ **Security & CORS Middleware**: Configurable cross-origin resource sharing, authorization header extraction, and context-injected user claims.
* ⚡ **Graceful Server & DB Shutdown**: Signal-trapped OS termination handling (`SIGINT`/`SIGTERM`) with 10-second request draining and PostgreSQL pool closure.
* 📜 **Automated OpenAPI / Swagger Integration**: Interactive Swagger UI with Bearer token authentication header support.

---

## 📁 Repository Architecture

```text
novaerp/
├── cmd/
│   └── server/
│       └── main.go                  # Application entrypoint & Swagger annotations
├── docs/                            # Generated Swagger / OpenAPI documentation
├── internal/
│   ├── app/                         # Global Application context struct
│   ├── bootstrap/                   # Application bootstrapper, DB connection & AutoMigrate
│   ├── common/
│   │   ├── middleware/              # CORS, RequestID, and Auth middlewares
│   │   ├── model/                   # Base GORM models (UUID primary keys, timestamps)
│   │   ├── pagination/              # Standardized request/response pagination helpers
│   │   └── response/                # Unified API JSON response contracts
│   ├── config/                      # Environment variable loader (.env)
│   ├── database/                    # PostgreSQL GORM connection & driver config
│   ├── handler/                     # Base health check handlers
│   ├── logger/                      # Zap Structured Logger
│   ├── modules/                     # Domain Modules & Submodules
│   │   ├── auth/                    # RS256 JWT, Refresh Tokens, KeyManager & Middleware
│   │   ├── department/              # Department Management
│   │   ├── employee/                # Employee Profiles & Master Data
│   │   ├── user/                    # User Accounts & Role definitions
│   │   ├── inventory/               # Inventory Submodules
│   │   │   ├── product/             # Product Master Data & SKUs
│   │   │   ├── stock/               # Multi-warehouse Stock Adjustments
│   │   │   └── warehouse/           # Warehouse Locations
│   │   └── hr/                      # Human Resources Submodules
│   │       ├── attendance/          # Daily Check-in/out & Hours Calculation
│   │       ├── leave/               # Leave Submodules
│   │       │   ├── leaverequest/    # Employee Leave Applications & Approvals
│   │       │   └── leavetype/       # Leave Policy Categories
│   │       └── payroll/             # Monthly Payroll Batches & Payslips
│   └── routes/                      # Route aggregator & router group registration
├── .env.example                     # Environment template
├── Makefile                         # Development & build task automation
├── go.mod                           # Go module dependencies
└── README.md                        # Documentation
```

---

## 🛠️ Getting Started

### Prerequisites
* **Go**: `v1.22+`
* **PostgreSQL**: `v14+`
* **Air** *(Optional)*: Live reload CLI (`go install github.com/air-verse/air@latest`)

### 1. Clone & Set Up Environment

```bash
git clone https://github.com/tamim1715/novaerp.git
cd novaerp
cp .env.example .env
```

### 2. Configure Environment Variables (`.env`)

```env
APP_NAME=NovaERP
APP_ENV=development
PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=novaerp

GIN_MODE=debug
JWT_SECRET=super-secret-key

# RS256 RSA Keys (Optional: If empty, application auto-generates 2048-bit key pair on startup)
JWT_PRIVATE_KEY=""
JWT_PUBLIC_KEY=""

CORS_ALLOWED_ORIGINS="*"
```

---

## 🏃 Running the Application

### Development Mode (with Live Reload & Swagger auto-gen)
```bash
make dev
```

### Build Production Binary
```bash
make build
./bin/server
```

### Run All Unit Tests
```bash
make test
```

### Regenerate Swagger Documentation
```bash
make swagger
```

---

## 🔐 Authentication & Swagger UI

1. **Start the server**: `make dev`
2. **Access Swagger UI**: Open `http://localhost:8080/swagger/index.html` in your browser.
3. **Login**: Execute `POST /api/v1/auth/login` with your credentials.
4. **Authorize**:
   * Copy the returned `accessToken`.
   * Click the green **Authorize** button at the top-right of Swagger UI.
   * Enter: `Bearer <your_access_token>` and click **Authorize**.

---

## 📡 API Endpoint Overview

### 🔑 Authentication
* `POST /api/v1/auth/login` — User authentication (returns RS256 Access Token + Refresh Token)
* `POST /api/v1/auth/refresh` — Exchange refresh token for new access token
* `POST /api/v1/auth/logout` — Revoke refresh token

### 👤 User & Organization Management *(Protected)*
* `GET | POST /api/v1/users` — User Accounts & Roles
* `GET | POST /api/v1/departments` — Company Departments
* `GET | POST /api/v1/employees` — Employee Directory

### 📦 Inventory Management *(Protected)*
* `GET | POST /api/v1/inventories/products` — Product Catalog & SKUs
* `GET | POST /api/v1/inventories/warehouses` — Storage Warehouses
* `GET | POST /api/v1/inventories/stocks` — Inventory Balances
* `POST /api/v1/inventories/stocks/adjust` — Transactional Stock Adjustment

### 👥 HR & Payroll *(Protected)*
* `GET | POST /api/v1/hr/leaves/types` — Leave Categories
* `GET | POST /api/v1/hr/leaves/requests` — Leave Applications & Approvals
* `POST /api/v1/hr/attendances/check-in` — Shift Check-In
* `POST /api/v1/hr/attendances/check-out` — Shift Check-Out & Overtime Calculation
* `POST /api/v1/hr/payrolls` — Create Monthly Payroll Batch
* `POST /api/v1/hr/payrolls/:id/process` — Process Salary Slips
* `GET /api/v1/hr/payrolls/:id/payslips` — View Generated Payslips

---

[//]: # (## 📄 License)

[//]: # (This project is licensed under the MIT License.)
