# Master Microservice

This is a microservice for maintaining master data.

## Setup Using Docker Compose
Run this command to build the image and run the container using Docker Compose:

```bash
docker compose up -d --build
```

If you see the result like this, then the app should be running:

```bash
✔ Network master-network           Created     0.1s
✔ Volume "master-db-data"          Created     0.0s
✔ Container master_service_db      Healthy     33.7s
✔ Container master-service         Started     34.1s
✔ Container migrate_master_db      Exited      34.6s
✔ Container seed_master_db         Started     34.9s
```

Make sure the app runs correctly in Docker by executing:

```bash
docker compose ps
```

Congratulations! Now your master service is running. If you access the API, it will show a welcome message:

```bash
curl --location 'http://localhost:8001/'
```

---

# API Endpoints

## Roles

### Get All Roles
Retrieve all roles stored in the database.

```http
GET /roles
```

**Response:**
```json
[
  {
    "id": "5e638abb-0333-11f0-8fa9-0242ac130002",
    "name": "Admin",
    "is_active": true,
    "created_by": "00000000-0000-0000-0000-000000000001",
    "created_at": "2025-03-17T13:26:01Z",
    "updated_by": "00000000-0000-0000-0000-000000000001",
    "updated_at": "2025-03-17T13:26:01Z"
  },
  {
    "id": "5e638ca3-0333-11f0-8fa9-0242ac130002",
    "name": "Manager",
    "is_active": true,
    "created_by": "00000000-0000-0000-0000-000000000001",
    "created_at": "2025-03-17T13:26:01Z",
    "updated_by": "00000000-0000-0000-0000-000000000001",
    "updated_at": "2025-03-17T13:26:01Z"
  },
  {
    "id": "5e638dfa-0333-11f0-8fa9-0242ac130002",
    "name": "Supervisor",
    "is_active": true,
    "created_by": "00000000-0000-0000-0000-000000000001",
    "created_at": "2025-03-17T13:26:01Z",
    "updated_by": "00000000-0000-0000-0000-000000000001",
    "updated_at": "2025-03-17T13:26:01Z"
  },
  {
    "id": "5e638e54-0333-11f0-8fa9-0242ac130002",
    "name": "Staff",
    "is_active": true,
    "created_by": "00000000-0000-0000-0000-000000000001",
    "created_at": "2025-03-17T13:26:01Z",
    "updated_by": "00000000-0000-0000-0000-000000000001",
    "updated_at": "2025-03-17T13:26:01Z"
  }
]
```

### Insert Role
Create a role.

```http
POST /roles
```

**Request Body:**
```json
{
    "name": "Developer"
}
```

**Response:**
```json
{
    "id": "ab445a56-10ed-42f6-9ac2-949766fdca28",
    "name": "Developer",
    "is_active": true,
    "created_by": "00000000-0000-0000-0000-000000000000",
    "created_at": "2025-03-17T13:29:46.996924963Z",
    "updated_by": "00000000-0000-0000-0000-000000000000",
    "updated_at": "2025-03-17T13:29:46.999Z"
}
```

### Update Role
Update a role.

```http
POST /roles
```

**Request Body:**
```json
{
    "id": "ab445a56-10ed-42f6-9ac2-949766fdca28",
    "name": "Developer",
    "is_active": false
}
```

**Response:**
```json
{
    "id": "ab445a56-10ed-42f6-9ac2-949766fdca28",
    "name": "Developer",
    "is_active": false,
    "created_by": "00000000-0000-0000-0000-000000000000",
    "created_at": "2025-03-17T13:29:47Z",
    "updated_by": "00000000-0000-0000-0000-000000000000",
    "updated_at": "2025-03-17T13:31:22Z"
}
```

## Status

### Get All Status
Retrieve all statuses stored in the database.

```http
GET /status
```

**Response:**
```json
[
    {
        "id": "5e64c58a-0333-11f0-8fa9-0242ac130002",
        "name": "All",
        "is_active": true,
        "created_by": "00000000-0000-0000-0000-000000000001",
        "created_at": "2025-03-17T13:26:01Z",
        "updated_by": "00000000-0000-0000-0000-000000000001",
        "updated_at": "2025-03-17T13:26:01Z"
    },
    {
        "id": "5e64c71f-0333-11f0-8fa9-0242ac130002",
        "name": "Active",
        "is_active": true,
        "created_by": "00000000-0000-0000-0000-000000000001",
        "created_at": "2025-03-17T13:26:01Z",
        "updated_by": "00000000-0000-0000-0000-000000000001",
        "updated_at": "2025-03-17T13:26:01Z"
    },
    {
        "id": "5e64c8a6-0333-11f0-8fa9-0242ac130002",
        "name": "Inactive",
        "is_active": true,
        "created_by": "00000000-0000-0000-0000-000000000001",
        "created_at": "2025-03-17T13:26:01Z",
        "updated_by": "00000000-0000-0000-0000-000000000001",
        "updated_at": "2025-03-17T13:26:01Z"
    }
]
```

### Insert Status
Create a status.

```http
POST /status
```

**Request Body:**
```json
{
    "name": "Pending"
}
```

**Response:**
```json
{
  "id": "f9f6be8b-0b91-457e-9530-61ae03edd431",
  "name": "Pending",
  "is_active": true,
  "created_by": "00000000-0000-0000-0000-000000000000",
  "created_at": "2025-03-17T13:33:29.227983763Z",
  "updated_by": "00000000-0000-0000-0000-000000000000",
  "updated_at": "2025-03-17T13:33:29.228Z"
}
```

### Update Status
Update a status.

```http
POST /status
```

**Request Body:**
```json
{
    "id": "f9f6be8b-0b91-457e-9530-61ae03edd431",
    "name": "Pending",
    "is_active": false
}
```

**Response:**
```json
{
    "id": "f9f6be8b-0b91-457e-9530-61ae03edd431",
    "name": "Pending",
    "is_active": false,
    "created_by": "00000000-0000-0000-0000-000000000000",
    "created_at": "0001-01-01T00:00:00Z",
    "updated_by": "00000000-0000-0000-0000-000000000000",
    "updated_at": "2025-03-17T13:34:02.436Z"
}
```