# PoCoGo

This is a simple project with API and a javascript frontend client.

https://golang.org/pkg/

## Requiriments

- golang 1.10.1

## API

Above, the proposed entrypoints to API system.

### GET /api

Endpoint to API area

### GET /api/users

Endpoint to get users

### POST /api/users

Creates an user

### GET /api/users/:id

Fetches an user

### DELETE /api/users/:id

Removes an user

### PUT /api/users/:id

Updates an user

## Entities

### User

```
{
    "name": String,
    "email": String,
    "createdAt": Date,
    "updatedAt": Date
}
```
