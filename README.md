# 🌍 Geo‑Track Multi‑Tenant System

A fully containerized, modular microservices-based system for **real-time location tracking** with **multi-tenancy support**, **role-based access control**, and **JWT-based authentication** via **AWS Cognito**.

## 🔧 Key Features

- 🧭 **Location Streaming**: Real-time WebSocket stream of incoming location data
- 👥 **Multi-Tenant Support**: Isolated tenant and user management
- 🔐 **Secure Authentication**: AWS Cognito login & JWT validation
- 🛡 **RBAC**: Role-based access for tenants and admins
- 📦 **Dockerized Deployment**: Easy-to-deploy using Docker Compose

---

## 📚 Documentation

For setup instructions, API reference, and deployment steps:

👉 **[View Full Documentation →](./multi-tenant-system//docs/README.md)**

---

## 💡 Technologies Used

- Go (Golang)
- PostgreSQL
- AWS Cognito
- Docker & Docker Compose
- WebSockets

---

## 🚀 Quick Start

```bash
git clone https://github.com/Mudit2297/geo-track-multi-tenant.git
cd geo-track-multi-tenant
# (Update values in env files)
docker-compose up --build
