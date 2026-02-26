# Delete Screening API Documentation

## 1. Delete Screening

* **URL:** `http://localhost:8080/api/screenings/{id}`
* **Method:** `DELETE`
* **Path Param:**

  * `id` (integer) – Screening ID
* **Headers:**

  * `Content-Type: application/json`
* **Body:** None

---

## Request–Response Flow Summary

![Call Flow](./delete_screening.png)

### 1. User Request

The client sends a DELETE request with the screening ID to deactivate a screening.

### 2. Service Processing & Response

The service validates screening existence and business constraints (e.g., sold seats), performs a soft delete (sets `is_active` to false) if allowed, and returns a success or error response.

---

## Response (Success)

* **Status Code:** 200 OK

```json
{
  "message": "Screening deactivated successfully",
  "data": {
    "screening_id": 10,
    "is_active": false
  }
}
```

---

## Response (Failure Scenarios)

### 1. Screening Not Found

* **Status Code:** 404 Not Found

```json
{
  "message": "Screening not found",
  "error": "No screening with given ID"
}
```

### 2. Invalid Screening ID Format

* **Status Code:** 400 Bad Request

```json
{
  "message": "Invalid screening ID",
  "error": "Path parameter must be a valid integer"
}
```

### 3. Sold Seats Exist (Business Rule Violation)

* **Status Code:** 409 Conflict

```json
{
  "message": "Cannot deactivate screening, some seats already sold",
  "error": "Sold seat IDs: [2,3]"
}
```

### 4. Reserved Seats Exist (Optional Business Constraint)

* **Status Code:** 409 Conflict

```json
{
  "message": "Cannot deactivate screening, seats currently reserved",
  "error": "Reserved seats must be released before deactivation"
}
```

### 5. Screening Already Inactive

* **Status Code:** 409 Conflict

```json
{
  "message": "Screening already inactive",
  "error": "No state change required"
}
```
