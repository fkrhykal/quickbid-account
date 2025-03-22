# QuickBid Account API

Welcome to the QuickBid Account API. This API enables seamless account management, facilitating interactions across all QuickBid services. 🚀

## Authentication Endpoints

### Sign Up

Registers a new user.

**Endpoint:**

```
POST /sign-up
```

**Request Body:**

```json
{
  "username": "string",
  "password": "string"
}
```

**Response:**

```json
HTTP/1.1 201 Created
Content-Type: application/json

{
  "id": "string"
}
```

#### Sign Up Error Handling

- `400 Bad Request` – Invalid or missing input
- `409 Conflict` – Username already exists
- `500 Internal Server Error` – Server-side error

### Sign In

Authenticates an existing user and returns an access token.

**Endpoint:**

```
POST /sign-in
```

**Request Body:**

```json
{
  "username": "string",
  "password": "string"
}
```

**Response:**

```json
HTTP/1.1 200 OK
Content-Type: application/json

{
  "token": "string"
}
```

#### Sign In Error Handling

- `400 Bad Request` – Invalid or missing input
- `401 Unauthorized` – Incorrect username or password
- `500 Internal Server Error` – Server-side error
