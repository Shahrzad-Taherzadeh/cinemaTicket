# Retrieve Active Screenings for a Movie API Documentation

## 1. Retrieve Active Screenings for a Movie

* **URL:** `http://localhost:8080/api/movies/{id}/screenings`
* **Method:** `GET`
* **Path Param:**

  * `id` (integer) – Movie ID
* **Headers:**

  * `Content-Type: application/json`
* **Body:** None

---

## Request–Response Flow Summary

![Call Flow](./movie_screenings.png)

### 1. User Request

The client sends a GET request with the movie ID to retrieve all active screenings associated with that movie.

### 2. Service Processing & Response

The service validates movie existence, filters screenings by `is_active = true`, retrieves related scheduling data, and returns the list or an appropriate error response.

---

## Response (Success)

* **Status Code:** 200 OK

```json
{
  "data": [
    {
      "screening_id": 10,
      "show_time_id": 5,
      "hall_id": 2,
      "price": 150000,
      "is_active": true,
      "start": "2026-03-01T18:00:00Z",
      "end": "2026-03-01T21:00:00Z"
    }
  ]
}
```

If no active screenings exist, the response should be:

```json
{
  "data": []
}
```

---

## Response (Failure Scenarios)

### 1. Movie Not Found

* **Status Code:** 404 Not Found

```json
{
  "message": "Movie not found",
  "error": "No movie with given ID"
}
```

### 2. Invalid Movie ID Format

* **Status Code:** 400 Bad Request

```json
{
  "message": "Invalid movie ID",
  "error": "Path parameter must be a valid integer"
}
```