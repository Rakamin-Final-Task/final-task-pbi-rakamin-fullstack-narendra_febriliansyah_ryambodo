# User

## User Registration API

Endpoint : POST /users/register

Request Body :

```json
{
  "username": "username",
  "email": "example@gmail.com",
  "password": "password"
}
```

Response Body Success :

```json
{
  "username": "username",
  "email": "example@gmail.com",
}
```

Response Body Error :

- status : 400

```json
{
  "errors": "Register failed"
}
```

## User Login API

Endpoint : GET /users/login

Request Body :

```json
{
  "email": "example@gmail.com",
  "password": "password"
}
```

Response Body Success :

```json
{
  "token": "unique-token" 
}
```

Response Body Error :

- status : 400

```json
{
  "errors": "Login failed"
}
```

## Update User API

Headers :

- Authorization : token

Endpoint : PUT /users/update

Request Body :

```json
{
  "username": "new_username",
  "email": "example@gmail.com",
}
```

Response Body Success :

```json
{
  "username": "new_username",
  "email": "example@gmail.com",
}
```

Response Body Error :

- status : 400

```json
{
  "errors": "Update failed"
}
```

## Delete User API

Headers :

- Authorization : token

Endpoint : DELETE /users/delete

Request Body :

```json
{
  "email": "example@gmail.com",
}
```

Response Body Success :

```json
{
  "email": "example@gmail.com",
}
```

Response Body Error :

- status : 400

```json
{
  "errors": "Delete failed"
}
```
