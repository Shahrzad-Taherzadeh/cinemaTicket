# Retrieve User Tickets API Documentation

## 1. Retrieve User Tickets

* **URL:** `http://localhost:8080/api/users/{id}/tickets`
* **Method:** `GET`
* **Path Param:**

  * `id` (integer) – User ID
* **Headers:**

  * `Content-Type: application/json`
* **Body:** None

---

## Request–Response Flow Summary

![Call Flow](./get_user_tickets.png)

### 1. User Request

The client sends a GET request with the user ID to retrieve all tickets owned by that user.

### 2. Service Processing & Response

The service validates user existence, fetches all tickets for the user, joins relevant screening and movie information, and returns a structured list or an appropriate error.

---

## Response (Success)

* **Status Code:** 200 OK

```json
{
  "data": [
    {
      "ticket_id": 101,
      "screening_id": 10,
      "seat_row": "A",
      "seat_number": "1",
      "movie_name": "Oppenheimer",
      "start": "2026-03-01T18:00:00Z",
      "end": "2026-03-01T21:00:00Z",
      "price": 150000
    }
  ]
}
```

If the user has no tickets:

```json
{
  "data": []
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