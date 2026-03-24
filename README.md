# Credit Decision Service

A credit application and decision engine built with Go, React, and AWS.

## Architecture

- **Application Service** — Go + Chi HTTP server for credit application intake
- **Infrastructure** — Terraform, DynamoDB, deployed on AWS
- **Frontend** — React (planned)

## Tech Stack

- Go 1.22+
- Chi (HTTP router)
- AWS DynamoDB
- Docker
- Terraform

## Getting Started

### Prerequisites

- Go 1.22+
- Docker
- AWS CLI

### Run Locally
```bash
# Start DynamoDB Local
make infra-up

# Create tables
make db-setup

# Run the application service
make run-app
```

### Available Make Commands

| Command       | Description                          |
|---------------|--------------------------------------|
| `make dev`    | Start infra, create tables, run app  |
| `make infra-up` | Start DynamoDB Local              |
| `make infra-down` | Stop DynamoDB Local             |
| `make db-setup` | Create DynamoDB tables            |
| `make run-app` | Run the application service        |
| `make test-app` | Run tests                         |

## Project Structure
```
credit_decision/
├── services/
│   └── application/     # Credit application service (Go)
├── infra/
│   └── terraform/       # AWS infrastructure
├── web/                 # React frontend (planned)
├── scripts/             # Dev tooling
└── docker-compose.yml   # Local infrastructure
```
