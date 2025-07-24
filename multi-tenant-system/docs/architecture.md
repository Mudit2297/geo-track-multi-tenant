# Architecture Overview

This section describes the architecture of **Geo‑Track Multi‑Tenant**, a system for multi‑tenant real‑time location streaming.

         +---------------+          
         |   Auth        |          
         |  Service      |          
         +-------+-------+          
                 |                        
         JWT issue/validate                
                 |                        
    +------------v------------+         
    |     API Services        |         
    |-------------------------|         
    | 🎯 Tenant Service       | ←– Admin-only access 
    | 📍 Location Service     | ←– Tenant and admin users
    +------------+------------+         
                 |                        
        Stores data in PostgreSQL                
                 |                        
    +-------------v-------------+             
    |    Streaming Service      |             
    |  (Gorilla Web‑Socket)     |             
    +-------------+-------------+             
                 |                        
           WebSocket clients (e.g.,   
           browser, wscat, simulator)  



---

### Key Concepts

- **Microservices**: Isolated using Docker Compose, communicating internally via `backend` network.
- **Authentication**: AWS Cognito via a local `auth-service` wrapper.
- **RBAC**: Claims like `role=admin` or `role=tenant` enforced by middleware.
- **Streaming**: Location updates are broadcast via WebSocket (`/ws` endpoint).

---