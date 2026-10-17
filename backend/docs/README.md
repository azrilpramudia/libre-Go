# LibreGo API Documentation

LibreGo is a library book repository backend built with Go, Gin, PostgreSQL, Redis, and the OpenLibrary API. This document describes all available endpoints for frontend integration.

## Table of Contents

- [Base URL](#base-url)
- [Response Format](#response-format)
- [Authentication](#authentication)
- [Endpoints](#endpoints)
  - [Health Check](#health-check)
  - [Auth](#auth)
  - [Books](#books)
  - [Members](#members)
  - [Loans](#loans)
- [Error Codes Reference](#error-codes-reference)
- [Environment Variables](#environment-variables)

---

## Base URL

```
http://localhost:5000
```

Adjust the port according to your local `.env` configuration (`APP_PORT`).

---

## Response Format

All endpoints return a consistent JSON envelope:

**Success:**

```json
{
  "success": true,
  "data": { ... }
}
```

**Error:**

```json
{
  "success": false,
  "error": "error message here"
}
```

---

## Authentication

LibreGo uses **JWT Bearer tokens** for protected endpoints. After logging in via `POST /auth/login`, include the returned token in the `Authorization` header for protected requests:

```
Authorization: Bearer <your-jwt-token>
```

Tokens expire after **24 hours**.

**Protected endpoint groups:**

- `/loans/*` — all loan endpoints require authentication

---

## Endpoints

### Health Check

#### `GET /health`

Checks whether the API is running.

**Auth required:** No

**Response `200`:**

```json
{
  "status": "ok"
}
```

> Note: this endpoint does not use the standard response envelope.

---

### Auth

#### `POST /auth/register`

Registers a new member account.

**Auth required:** No

**Request body:**

```json
{
  "full_name": "Budi Santoso",
  "email": "budi@example.com",
  "phone": "08123456789",
  "password": "rahasia123"
}
```

| Field       | Type   | Required | Notes                 |
| ----------- | ------ | -------- | --------------------- |
| `full_name` | string | Yes      |                       |
| `email`     | string | Yes      | Must be a valid email |
| `phone`     | string | No       |                       |
| `password`  | string | Yes      | Minimum 6 characters  |

**Response `201`:**

```json
{
  "success": true,
  "data": {
    "ID": "b1f2e3a4-...",
    "FullName": "Budi Santoso",
    "Email": "budi@example.com",
    "Phone": "08123456789",
    "CreatedAt": "2026-10-17T10:00:00Z"
  }
}
```

**Possible errors:**
| Status | Error | Cause |
|--------|--------------------------------|-------------------------------------|
| 400 | validation error message | Missing/invalid fields |
| 409 | `email already registered` | Email already taken |

---

#### `POST /auth/login`

Logs a member in and returns a JWT token.

**Auth required:** No

**Request body:**

```json
{
  "email": "budi@example.com",
  "password": "rahasia123"
}
```

**Response `200`:**

```json
{
  "success": true,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIs..."
  }
}
```

**Possible errors:**
| Status | Error | Cause |
|--------|----------------------------------|--------------------------------|
| 400 | validation error message | Missing/invalid fields |
| 401 | `invalid email or password` | Wrong credentials |

---

### Books

#### `POST /books`

Creates a new book manually (without OpenLibrary data).

**Auth required:** No

**Request body:**

```json
{
  "title": "Laskar Pelangi",
  "isbn": "9789793062792",
  "stock": 3
}
```

| Field   | Type   | Required | Notes                        |
| ------- | ------ | -------- | ---------------------------- |
| `title` | string | Yes      |                              |
| `isbn`  | string | No       |                              |
| `stock` | int32  | No       | Defaults to 1 if less than 1 |

**Response `201`:**

```json
{
  "success": true,
  "data": {
    "ID": "c3d4e5f6-...",
    "Title": "Laskar Pelangi",
    "ISBN": "9789793062792",
    "OpenLibraryID": "",
    "CoverURL": "",
    "Stock": 3,
    "Available": 3,
    "CreatedAt": "2026-10-17T10:00:00Z",
    "UpdatedAt": "2026-10-17T10:00:00Z"
  }
}
```

---

#### `GET /books`

Lists books with pagination.

**Auth required:** No

**Query params:**
| Param | Type | Default | Notes |
|----------|------|---------|-------------------|
| `limit` | int | 10 | Items per page |
| `offset` | int | 0 | Items to skip |

**Example:** `GET /books?limit=20&offset=0`

**Response `200`:**

```json
{
  "success": true,
  "data": [
    {
      "ID": "c3d4e5f6-...",
      "Title": "Laskar Pelangi",
      "ISBN": "9789793062792",
      "Stock": 3,
      "Available": 2,
      "CreatedAt": "2026-10-17T10:00:00Z",
      "UpdatedAt": "2026-10-17T10:00:00Z"
    }
  ]
}
```

---

#### `GET /books/:id`

Gets a single book by ID.

**Auth required:** No

**Path params:**
| Param | Type | Notes |
|-------|------|-----------------|
| `id` | UUID | Book ID |

**Response `200`:**

```json
{
  "success": true,
  "data": {
    "ID": "c3d4e5f6-...",
    "Title": "Laskar Pelangi",
    "Stock": 3,
    "Available": 2
  }
}
```

**Possible errors:**
| Status | Error | Cause |
|--------|---------------------|-------------------------|
| 400 | `invalid id` | Malformed UUID |
| 404 | `book not found` | No book with that ID |

---

#### `GET /books/search/openlibrary`

Searches the OpenLibrary catalog by title. Results are cached in Redis for 1 hour (case-insensitive cache key).

**Auth required:** No

**Query params:**
| Param | Type | Required | Notes |
|---------|--------|----------|-----------------|
| `title` | string | Yes | Book title to search |

**Example:** `GET /books/search/openlibrary?title=the+hobbit`

**Response `200`:**

```json
{
  "success": true,
  "data": {
    "docs": [
      {
        "title": "The Hobbit",
        "author_name": ["J.R.R. Tolkien"],
        "isbn": null,
        "cover_i": 14627509,
        "key": "/works/OL27482W"
      }
    ]
  }
}
```

> `isbn` is frequently `null` — OpenLibrary's search endpoint does not always return ISBN data at the work level.

**Possible errors:**
| Status | Error | Cause |
|--------|-----------------------------|----------------------------------|
| 400 | `title query param required`| Missing `title` query param |
| 502 | error message | OpenLibrary API unreachable/error|

---

#### `POST /books/import/openlibrary`

Creates a new book using data from an OpenLibrary search result (see above). Automatically finds or creates the author and links them to the book.

**Auth required:** No

**Request body:**

```json
{
  "title": "The Hobbit",
  "author_name": "J.R.R. Tolkien",
  "isbn": "",
  "cover_i": 14627509,
  "key": "/works/OL27482W",
  "stock": 2
}
```

| Field         | Type   | Required | Notes                                            |
| ------------- | ------ | -------- | ------------------------------------------------ |
| `title`       | string | Yes      |                                                  |
| `author_name` | string | No       | Creates the author if not already present        |
| `isbn`        | string | No       |                                                  |
| `cover_i`     | int    | No       | OpenLibrary cover ID, converted to a cover URL   |
| `key`         | string | Yes      | OpenLibrary work key, stored as `openlibrary_id` |
| `stock`       | int32  | No       | Defaults to 1 if less than 1                     |

**Response `201`:** Same shape as `POST /books`.

---

### Members

> **Deprecated:** `POST /members` has been superseded by `POST /auth/register`, which also sets a password for login. New member creation should go through `/auth/register`.

#### `GET /members`

Lists members with pagination.

**Auth required:** No

**Query params:** Same as `GET /books` (`limit`, `offset`).

**Response `200`:**

```json
{
  "success": true,
  "data": [
    {
      "ID": "b1f2e3a4-...",
      "FullName": "Budi Santoso",
      "Email": "budi@example.com",
      "Phone": "08123456789",
      "CreatedAt": "2026-10-17T10:00:00Z"
    }
  ]
}
```

---

#### `GET /members/:id`

Gets a single member by ID.

**Auth required:** No

**Path params:**
| Param | Type | Notes |
|-------|------|-------------|
| `id` | UUID | Member ID |

**Response `200`:** Same shape as a single item from `GET /members`.

**Possible errors:**
| Status | Error | Cause |
|--------|------------------------|---------------------------|
| 400 | `invalid id` | Malformed UUID |
| 404 | `member not found` | No member with that ID |

---

### Loans

All loan endpoints require authentication. Include `Authorization: Bearer <token>` in the request headers.

#### `POST /loans/borrow`

Borrows a book for a member. Atomically decrements the book's `available` count and creates a loan record with a due date 7 days from now.

**Auth required:** Yes

**Request body:**

```json
{
  "book_id": "c3d4e5f6-...",
  "member_id": "b1f2e3a4-..."
}
```

**Response `201`:**

```json
{
  "success": true,
  "data": {
    "ID": "d4e5f6a7-...",
    "BookID": "c3d4e5f6-...",
    "MemberID": "b1f2e3a4-...",
    "BorrowedAt": "2026-10-17T10:00:00Z",
    "DueAt": "2026-10-24T10:00:00Z",
    "ReturnedAt": null,
    "Status": "borrowed",
    "FineAmount": 0
  }
}
```

**Possible errors:**
| Status | Error | Cause |
|--------|---------------------------|------------------------------------------|
| 400 | validation/parse error | Missing or invalid `book_id`/`member_id` |
| 401 | auth error | Missing/invalid/expired token |
| 409 | `book not available` | No copies of the book currently available |

---

#### `PATCH /loans/:id/return`

Marks a loan as returned. Atomically increments the book's `available` count (capped at `stock`) and calculates a late fine if applicable.

**Auth required:** Yes

**Path params:**
| Param | Type | Notes |
|-------|------|-----------|
| `id` | UUID | Loan ID |

**Fine calculation:** Rp 2,000 per day late (minimum 1 day if overdue at all). `0` if returned on or before the due date.

**Response `200`:**

```json
{
  "success": true,
  "data": {
    "ID": "d4e5f6a7-...",
    "BookID": "c3d4e5f6-...",
    "MemberID": "b1f2e3a4-...",
    "BorrowedAt": "2026-10-17T10:00:00Z",
    "DueAt": "2026-10-24T10:00:00Z",
    "ReturnedAt": "2026-10-20T15:30:00Z",
    "Status": "returned",
    "FineAmount": 0
  }
}
```

**Possible errors:**
| Status | Error | Cause |
|--------|-------------------------------|---------------------------------------|
| 400 | `invalid id` | Malformed UUID |
| 401 | auth error | Missing/invalid/expired token |
| 409 | `loan already returned` | This loan was already marked returned |

---

#### `GET /loans/member/:memberId`

Lists all loans (active and returned) for a given member.

**Auth required:** Yes

**Path params:**
| Param | Type | Notes |
|------------|------|-------------|
| `memberId` | UUID | Member ID |

**Response `200`:**

```json
{
  "success": true,
  "data": [
    {
      "ID": "d4e5f6a7-...",
      "BookID": "c3d4e5f6-...",
      "MemberID": "b1f2e3a4-...",
      "BorrowedAt": "2026-10-17T10:00:00Z",
      "DueAt": "2026-10-24T10:00:00Z",
      "ReturnedAt": null,
      "Status": "borrowed",
      "FineAmount": 0
    }
  ]
}
```

**Possible errors:**
| Status | Error | Cause |
|--------|------------------------|-------------------------|
| 400 | `invalid member id` | Malformed UUID |
| 401 | auth error | Missing/invalid/expired token |

---

## Error Codes Reference

| Status | Meaning                                                                  |
| ------ | ------------------------------------------------------------------------ |
| 200    | Success                                                                  |
| 201    | Resource created successfully                                            |
| 400    | Bad request — validation failed or malformed input                       |
| 401    | Unauthorized — missing, invalid, or expired token                        |
| 404    | Resource not found                                                       |
| 409    | Conflict — e.g. book unavailable, duplicate email, already returned loan |
| 500    | Internal server error                                                    |
| 502    | Bad gateway — upstream OpenLibrary API error                             |

---

## Environment Variables

The frontend team doesn't need to set these (backend-only), but they're listed here for reference on what the backend expects to be running:

| Variable               | Example                                                          | Notes                          |
| ---------------------- | ---------------------------------------------------------------- | ------------------------------ |
| `APP_PORT`             | `5000`                                                           | Port the API listens on        |
| `DATABASE_URL`         | `postgres://user:pass@localhost:5433/librego_db?sslmode=disable` | PostgreSQL connection string   |
| `REDIS_ADDR`           | `localhost:6379`                                                 | Redis address                  |
| `OPENLIBRARY_BASE_URL` | `https://openlibrary.org`                                        | OpenLibrary API base URL       |
| `JWT_SECRET`           | (random long string)                                             | Secret used to sign JWT tokens |

---

## Notes for Frontend Integration

- All UUID fields (`ID`, `BookID`, `MemberID`, etc.) are standard UUID v4 strings.
- Timestamps are in RFC 3339 / ISO 8601 format (UTC).
- After login, store the JWT token securely (e.g. httpOnly cookie or secure storage) and attach it to every request under `/loans/*`.
- A book's `available` field reflects real-time availability; `stock` is the total number of physical copies and does not change on borrow/return.
- When importing a book from OpenLibrary, always run a search first (`GET /books/search/openlibrary`) to get the `key`, `cover_i`, and `author_name` values needed for the import request.

---

## Interactive API Documentation

The OpenAPI specification is available in [`openapi.yaml`](./openapi.yaml). To view it in Swagger UI without copying the specification into Swagger Editor, serve this directory locally:

```bash
cd backend/docs
python3 -m http.server 8080
```

Then open [http://localhost:8080/](http://localhost:8080/) in your browser. The Swagger UI page loads `openapi.yaml` automatically. An internet connection is required to load the Swagger UI assets from the CDN.
