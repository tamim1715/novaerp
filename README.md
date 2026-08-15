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
* 💰 **Financial & Accounting Management Engine**:
  * **Chart of Accounts (COA)**: Hierarchical account classification (Assets, Liabilities, Equity, Revenue, Expense) with parent-child tree view and automatic GAAP/IFRS seeding.
  * **Fiscal Years & Accounting Periods**: 12-month automated financial sub-periods with period-closing and reopening controls.
  * **Double-Entry General Ledger**: Balanced debit/credit journal transactions with strict immutability, atomic posting, and auto-generated reversal entries.
  * **Financial Statements Engine**: Real-time **General Ledger**, **Trial Balance**, **Profit & Loss (Income Statement)**, and **Balance Sheet** reports.
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
│   │   ├── accounting/              # Financial & Accounting Submodules
│   │   │   ├── account/             # Chart of Accounts (COA) & Account Hierarchy Tree
│   │   │   ├── journal/             # Double-Entry General Ledger & Atomic Posting
│   │   │   ├── period/              # Fiscal Years & 12-Month Financial Sub-Periods
│   │   │   ├── report/              # Financial Statements (GL, Trial Balance, P&L, Balance Sheet)
│   │   │   └── routes.go            # Accounting Route Aggregator
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

### 🏥 System & Health *(Public)*
* `GET /` — Service status check / welcome endpoint
* `GET /health` — Service & PostgreSQL database health check
* `GET /swagger/*any` — Interactive Swagger UI / OpenAPI documentation

### 🔑 Authentication *(Public)*
* `POST /api/v1/auth/login` — User authentication (returns RS256 Access Token + Refresh Token)
* `POST /api/v1/auth/refresh` — Exchange refresh token for new access token
* `POST /api/v1/auth/logout` — Revoke refresh token

### 👤 User Management
* `GET | POST /api/v1/users` — List paginated users / Register a new user account
* `GET | PUT | DELETE /api/v1/users/:id` — Retrieve user profile, update account & roles, delete user

### 🏢 Department Management *(Protected)*
* `GET | POST /api/v1/departments` — List paginated departments / Create a department
* `GET | PUT | DELETE /api/v1/departments/:id` — Retrieve department details, update department, delete department

### 👨‍💼 Employee Directory *(Protected)*
* `GET | POST /api/v1/employees` — List paginated employees / Onboard a new employee
* `GET | PUT | DELETE /api/v1/employees/:id` — Retrieve employee profile, update employee info & status, delete employee

### 📦 Inventory Management *(Protected)*
* `GET | POST /api/v1/inventories/products` — List paginated products / Create product & SKU
* `GET | PUT | DELETE /api/v1/inventories/products/:id` — Retrieve product details, update product, delete product
* `GET | POST /api/v1/inventories/warehouses` — List paginated warehouses / Create storage warehouse
* `GET | PUT | DELETE /api/v1/inventories/warehouses/:id` — Retrieve warehouse details, update warehouse, delete warehouse
* `GET /api/v1/inventories/stocks` — List paginated inventory stock balances across warehouses
* `POST /api/v1/inventories/stocks/adjust` — Transactional stock adjustment (`IN`, `OUT`, `ADJUST` with row locking)

### 👥 HR & Payroll *(Protected)*
#### 🌴 Leave Management
* `GET | POST /api/v1/hr/leaves/types` — List leave type policies / Create leave category
* `GET | PUT | DELETE /api/v1/hr/leaves/types/:id` — Retrieve leave type details, update policy, delete leave category
* `GET | POST /api/v1/hr/leaves/requests` — List leave applications / Submit employee leave request
* `GET /api/v1/hr/leaves/requests/:id` — Retrieve leave application details by ID
* `PUT /api/v1/hr/leaves/requests/:id/status` — Approve or reject leave application (`APPROVED`, `REJECTED`)

#### ⏱️ Attendance System
* `POST /api/v1/hr/attendances/check-in` — Employee shift check-in (auto timestamp & status)
* `POST /api/v1/hr/attendances/check-out` — Employee shift check-out & automated work/overtime calculation
* `GET | POST /api/v1/hr/attendances` — List paginated attendance records / Record manual attendance entry
* `GET /api/v1/hr/attendances/:id` — Retrieve attendance record by ID

#### 💵 Payroll Processing
* `GET | POST /api/v1/hr/payrolls` — List paginated payroll batches / Create monthly payroll period
* `GET /api/v1/hr/payrolls/:id` — Retrieve payroll batch details by ID
* `POST /api/v1/hr/payrolls/:id/process` — Process and calculate salary slips for active employees
* `GET /api/v1/hr/payrolls/:id/payslips` — View generated employee payslips
* `POST /api/v1/hr/payrolls/:id/pay` — Mark payroll batch as paid and finalize status

### 💰 Financial & Accounting Core *(Protected)*
#### 📊 Chart of Accounts (COA)
* `GET | POST /api/v1/accounting/accounts` — Chart of Accounts (COA) list & create account
* `GET /api/v1/accounting/accounts/tree` — Hierarchical parent-child account tree structure
* `POST /api/v1/accounting/accounts/seed` — Seed standard GAAP/IFRS Chart of Accounts
* `GET | PUT | DELETE /api/v1/accounting/accounts/:id` — Retrieve account details, update account, safe delete account

#### 📅 Fiscal Years & Accounting Periods
* `GET | POST /api/v1/accounting/fiscal-years` — List fiscal years & sub-periods / Create fiscal year (with optional 12 monthly sub-periods)
* `GET /api/v1/accounting/fiscal-years/:id` — Retrieve fiscal year details & associated monthly accounting periods
* `POST /api/v1/accounting/fiscal-years/:id/close` — Close fiscal year & lock all sub-periods
* `GET /api/v1/accounting/periods/:id` — Retrieve monthly accounting period details by ID
* `POST /api/v1/accounting/periods/:id/close` — Close monthly accounting period
* `POST /api/v1/accounting/periods/:id/reopen` — Reopen closed accounting period

#### 📖 General Ledger & Journal Entries
* `GET | POST /api/v1/accounting/journals` — List paginated journal entries / Create balanced double-entry transaction
* `GET /api/v1/accounting/journals/:id` — Retrieve journal entry details and line items (debits/credits)
* `POST /api/v1/accounting/journals/:id/post` — Post & freeze journal entry to General Ledger
* `POST /api/v1/accounting/journals/:id/void` — Void posted journal entry (generates automated reversal entry)

#### 📈 Financial Statements & Reports
* `GET /api/v1/accounting/reports/general-ledger` — Running General Ledger statement with running balances
* `GET /api/v1/accounting/reports/trial-balance` — Trial Balance debit/credit equality verification report
* `GET /api/v1/accounting/reports/profit-loss` — Profit & Loss / Income Statement (Revenue, COGS, Gross Profit, Expenses, Net Income)
* `GET /api/v1/accounting/reports/balance-sheet` — Balance Sheet statement (Assets = Liabilities + Equity)

---

## 📄 License

This project is licensed under the MIT License.
