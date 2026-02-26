# CinemaTicket - API Documentation

## Project Assumptions / Technical Details

1. The application is written in **Golang** using the **Fiber** framework.
2. The application does **not** have an authentication system.
3. All data is stored **in-memory**.
4. All API documentation is stored in **Postgress**.
5. Requests are saved in **Postman**.
6. All documentation is saved in the project repository.

## APIs to be Developed

* **GET /api/movies** – Retrieve the list of movies.
* **GET /api/movies/{id}/screenings** – Retrieve active screenings for a specific movie.
* **GET /api/screenings/{id}/seats** – View the status of seats for a screening (available / reserved / sold).
* **POST /api/screenings/{id}/reserve** – Temporarily reserve seats for a user (based on user_id and seat_ids).
* **POST /api/screenings/{id}/confirm** – Finalize the ticket purchase atomically, including:

  * Checking if seats are already reserved
  * Moving seats from Reserved → Sold
  * Creating Tickets for each seat
  * Generating ticket details
  * Sending a message via **Bale Bot API** to the user's phone number
  * Rolling back the operation in case of message sending failure (in-memory)
* **GET /api/users/{id}/tickets** – Retrieve the list of tickets purchased by a user.
* **PUT /api/users/{id}** – Update the user's name (only the Name field is editable).
* **DELETE /api/screenings/{id}** – Delete a screening (soft delete: IsActive = false).
