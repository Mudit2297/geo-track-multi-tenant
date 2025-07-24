# 🚀 Deployment Guide

This document describes how to deploy the **Geo‑Track Multi‑Tenant** system locally using Docker Compose.

---

## 📁 Project Structure
```bash
.
├── multi-tenant-system
│   ├── auth-service/
│   ├── db/
│   ├── docker-compose.yaml
│   ├── docs/
│   ├── location-service/
│   ├── simulator/
│   ├── streaming-service/
│   └── tenant-service/
└── README.md
```


---

## ⚙️ Prerequisites

- Docker & Docker Compose installed
- AWS Cognito User Pool and App Client (with client secret)
- Go 1.21+ installed (for local development)
- Git

---

## 🧪 Environment Setup

### 1. Clone the repository

```bash
git clone https://github.com/Mudit2297/geo-track-multi-tenant.git
cd geo-track-multi-tenant
```

### 2. Configure .env

Create a .env file at the root with the following:

```env
# PostgreSQL
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=geo_track
DB_HOST: postgres
DB_PORT: 5432

# Cognito (used by auth-service)
COGNITO_CLIENT_ID=your_cognito_app_client_id
COGNITO_CLIENT_SECRET=your_cognito_app_client_secret
COGNITO_USER_POOL_ID=your_user_pool_id
COGNITO_REGION=ap-south-1
COGNITO_JWKS_URL=https://cognito-idp.ap-south-1.amazonaws.com/<COGNITO_USER_POOL_ID>/.well-known/jwks.json
```
**Note:** Update with your actual AWS Cognito values.

---

## 🐳 Docker Deployment

### 1. Start all services

```bash
docker-compose up --build
```  

This will spin up:
- postgres (DB)
- auth-service (Port: 8084)
- tenant-service (Port: 8081)
- location-service (Port: 8082)
- streaming-service (Port: 8083)

### 2. Verify Services

- Auth: http://localhost:8084/login
- Tenants: http://localhost:8081/tenants
- Location: http://localhost:8082/location 
- Streamer: ws://localhost:8083/ws

---

## 🧱 Database Schema Initialization

The PostgreSQL container runs an init.sql script (mounted as a volume) that automatically creates all necessary tables and constraints.  

**Location:**  
- db/init.sql

---

## 🛠 Testing

If everything has been setup correctly, navigate to the simulator folder, substitute the required values in the main() function and run the following command.

```bash
cd ./multi-tenant-system/simulator
go run main.go
```

---

## 🧼 Teardown
```bash 
#This will stop and remove all containers and volumes.
docker-compose down -v
```

--- 

## 🛠 Troubleshooting

- DB connection errors: Ensure the DATABASE_URL in .env points to the Docker service name (postgres) instead of localhost.
- Cognito errors: Double-check client ID, secret, and region.
- WebSocket 400 error: Use a WebSocket client like Postman, Hoppscotch, or browser extension. Ensure the request uses ws:// and includes the Upgrade: websocket header.