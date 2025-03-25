# QuickBid Account API

Welcome to the QuickBid Account API. This API enables seamless account management, facilitating interactions across all QuickBid services. 🚀

## Authentication Endpoints

### Sign Up

Registers a new user.

#### Request

```
POST /api/v1/sign-up
```

```
Content-Type: application/json
```

```json
{
  "username": "string",
  "password": "string"
}
```

#### Success Response

```
HTTP/1.1 201 Created
Content-Type: application/json
```

```json
{
  "id": "string"
}
```

#### Error Response

##### `400 Bad Request`

Invalid or missing input

```json
{
  "code": 400,
  "error": {
    "username": "string",
    "password": "string"
  }
}
```

##### `409 Conflict`

Username already used

```json
{
  "code": 409,
  "error": "username already used"
}
```

##### `500 Internal Server Error`

Unexpected server error

```json
{
  "code": 500,
  "error": "Internal Server Error"
}
```
