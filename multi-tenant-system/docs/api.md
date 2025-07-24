# API Reference

This document describes the REST API endpoints for all services in the Geo‑Track Multi‑Tenant system.

---

## 🔐 Auth Service (`http://localhost:8084`)

### POST `/login`

Authenticate a user using AWS Cognito via local proxy.

**Request:**
```json
{
  "username": "user@example.com",
  "password": "yourpassword"
}
```

**Response:**
```json
{
  "access_token": "eyJraWQiOiJ...",
  "id_token": "eyJhbGciOiJSUzI1NiIs...",
  "refresh_token": "eyJjdHkiOiJKV1Q...",
  "expires_in": 3600,
  "token_type": "Bearer"
}
```

### GET `/validate`

Validates a JWT token and returns decoded claims.

**HEADERS:**
Authorization: Bearer <access_token>

**Response:**
```json
{
  "sub": "user-uuid",
  "email": "user@example.com",
  "custom:role": "tenant",
  "custom:tenant_id": "tenant123"
}
```

---

## 🧩 Tenant Service (`http://localhost:8081`)

### POST `/tenants` — Admin only

Create a new tenant. Only accessible to users with role: admin.

**HEADERS:**
Authorization: Bearer <access_token>

**Request:**
```json
{
  "id": "tenant123",
  "name": "Tenant Alpha"
}
```

**Response:**  

NO RESPONSE, 201 CREATED


### GET `/tenants` — Admin only

List all registered tenants. Only accessible to users with role: admin.

**HEADERS:**
Authorization: Bearer <access_token>

**Response:**  
```
[
  {
    "id": "tenant123",
    "name": "Tenant Alpha"
  }
]
```

### GET `/tenant/:id` — Admin only

Get a particular tenant by ID. Only accessible to users with role: admin.

**HEADERS:**
Authorization: Bearer <access_token>

**Response:**
```json
{
  "id": "tenant123",
  "name": "Tenant Alpha"
}
```

---

## 📍 Location Service (`http://localhost:8081`)

### POST `/location`

Send real-time location data. Authenticated users must belong to a tenant and have role: tenant or admin.

**HEADERS:**
Authorization: Bearer <access_token>

**Request:**
```json
{
  "latitude": 12.9716,
  "longitude": 77.5946
}
```

**Response:**
```json
{
  "message": "location submitted"
}
```

---

## 📡 Streaming Service (`http://localhost:8083`)

### WebSocket `/ws`

Real-time location stream using WebSockets. Clients subscribe to receive broadcast location updates.

**Usage Example:**  
In browser extension or wscat:
```bash
ws://localhost:8083/ws
```

**Response:**
```json
{
  "tenant_id": "tenant123",
  "latitude": 12.9716,
  "longitude": 77.5946
}
```

---

## 🧪 Test Flow

1. Login via /login
2. Use token to:
    - Call /validate to see claims
    - Access tenant or location APIs based on role
3. Observe WebSocket stream from /ws

---

## 🛠 Notes

1. All protected endpoints expect a JWT in the Authorization header
2. Role-based access enforced by middleware
3. Location and tenant data is stored in PostgreSQL
4. Streamed data is emitted via broadcasting from backend services