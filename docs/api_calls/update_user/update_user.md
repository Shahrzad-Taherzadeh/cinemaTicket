# Update User Name API Documentation

## 1. Update User Name

* **URL:** `http://localhost:8080/api/users/{id}`
* **Method:** `PATCH`
* **Path Param:**

  * `id` (integer) – User ID
* **Headers:**

  * `Content-Type: application/json`
* **Body Type:** JSON

### Request–Response Flow Summary

![Call Flow](./update_user.png)

1. **User Request Submission**
   The client sends a PATCH request with a new `name` to update the user's full name.

2. **Service Processing & Response**
   The service validates the user ID, updates the `name` field in the database, and returns the updated user data or an error if the operation fails.

---

## Request Body Example

```json
{
  "name": "John Doe"
}
```

---

## Response (Success)

* **Status Code:** 202 Accepted

```json
{
  "message": "User name updated successfully",
  "data": {
    "id": 1,
    "name": "John Doe",
    "phone": "09129999999"
  }
}
```

---

## Response (Failure Scenarios)

### 1. User Not Found

* **Status Code:** 404 Not Found

```json
{
  "message": "User not found",
  "error": "No user with given ID"
}
```

### 2. Invalid User ID Format

* **Status Code:** 400 Bad Request

```json
{
  "message": "Invalid user ID",
  "error": "Path parameter must be a valid integer"
}
```

### 3. Invalid Request Payload

* **Status Code:** 400 Bad Request

```json
{
  "message": "Invalid request payload",
  "error": "Field 'name' is required and must be a non-empty string"
}
```