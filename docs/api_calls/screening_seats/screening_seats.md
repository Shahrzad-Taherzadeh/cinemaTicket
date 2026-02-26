# View Seats for a Screening API Documentation

## 1. View Seats for a Screening

* **URL:** `http://localhost:8080/api/screenings/{id}/seats`
* **Method:** `GET`
* **Path Param:**

  * `id` (integer) – Screening ID
* **Headers:**

  * `Content-Type: application/json`
* **Body:** None

---

## Request–Response Flow Summary

![Call Flow](./get_screening_seats.png)

### 1. User Request

The client sends a GET request with the screening ID to retrieve the seat layout and current seat statuses.

### 2. Service Processing & Response

The service validates screening existence, fetches all seats associated with the screening, determines each seat's current state (`AVAILABLE`, `RESERVED`, `SOLD`), and returns the list or an appropriate error.

---

## Response (Success)

* **Status Code:** 200 OK

```json
{
  "data": [
    {
      "seat_id": 1,
      "row": "A",
      "number": "1",
      "status": "available"
    },
    {
      "seat_id": 2,
      "row": "A",
      "number": "2",
      "status": "sold"
    }
  ]
}
```

If no seats are configured for the screening:

```json
{
  "data": []
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