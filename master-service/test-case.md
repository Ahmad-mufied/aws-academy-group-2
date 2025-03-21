# Testing for Status API

## **Test Cases**

### **1. Create New Status**
#### **Request Body (name only)**
```json
{
    "name": "Pending"
}
```
#### **Expected Response**
```json
{
    "id": "e986f0b0-acea-4ed8-a5ae-6c5addbe4a2d",
    "name": "Pending",
    "is_active": true,
    "created_by": "00000000-0000-0000-0000-000000000000",
    "created_at": "2025-03-21T10:50:50Z",
    "updated_by": "00000000-0000-0000-0000-000000000000",
    "updated_at": "2025-03-21T10:50:50Z"
}
```
---
#### **Request Body (name and is_active)**
```json
{
    "name": "Waiting",
    "is_active": false
}
```
#### **Expected Response**
```json
{
    "id": "7b3813e4-f9a3-4242-bd55-2f140d6654f0",
    "name": "Waiting",
    "is_active": false,
    "created_by": "00000000-0000-0000-0000-000000000000",
    "created_at": "2025-03-21T10:51:27Z",
    "updated_by": "00000000-0000-0000-0000-000000000000",
    "updated_at": "2025-03-21T10:51:27Z"
}
```

---

### **2. Update Existing Status**
#### **Request Body (without name)**
```json
{
    "id": "7b3813e4-f9a3-4242-bd55-2f140d6654f0",
    "is_active": true
}
```
#### **Expected Response**
```json
{
    "id": "7b3813e4-f9a3-4242-bd55-2f140d6654f0",
    "name": "Waiting",
    "is_active": true,
    "created_by": "00000000-0000-0000-0000-000000000000",
    "created_at": "2025-03-21T10:51:27Z",
    "updated_by": "00000000-0000-0000-0000-000000000000",
    "updated_at": "2025-03-21T10:52:04Z"
}
```

---
#### **Request Body (with name)**
```json
{
    "id": "e986f0b0-acea-4ed8-a5ae-6c5addbe4a2d",
    "is_active": false,
    "name": "Pending 2"
}
```
#### **Expected Response**
```json
{
    "id": "e986f0b0-acea-4ed8-a5ae-6c5addbe4a2d",
    "name": "Pending 2",
    "is_active": false,
    "created_by": "00000000-0000-0000-0000-000000000000",
    "created_at": "2025-03-21T10:50:50Z",
    "updated_by": "00000000-0000-0000-0000-000000000000",
    "updated_at": "2025-03-21T10:52:36Z"
}
```

---

### **3. Prevent Duplicate Name**
#### **Request Body**
```json
{
    "name": "Pending 2"
}
```
#### **Expected Response**
```json
{
    "error": "Failed to create status"
}
```

---

### **4. Invalid Request - Empty Name**
#### **Request Body**
```json
{
    "name": ""
}
```
#### **Expected Response**
```json
{
    "error": "Status name cannot be empty"
}
```