# User Management Service

A simple REST API service for managing users.

## Installation

This service requires Go and Docker to be installed. To run the service, follow these steps:

1. Clone the repository to your local machine.
2. Navigate to the project directory.
3. Run `make start-build` to build and start the service with Docker.

The service will be running on `http://localhost:8080`.

## Endpoints

| Endpoint | Method | Parameters | Description |
| --- | --- | --- | --- |
| `/users` | GET | `query`: Search query to filter results. <br> `offset`: Number of records to skip. <br> `limit`: Maximum number of records to return. | Get a list of users. |
| `/users/{id}` | GET | `id`: ID of the user. | Get a specific user by ID. |
| `/users` | POST | `name`: User's name. <br> `email`: User's email. <br> `password`: User's password. <br> `country`: User's country. | Add a new user. |
| `/users/{id}` | PUT | `id`: ID of the user. <br> `name`: User's name. <br> `email`: User's email. <br> `password`: User's password. <br> `country`: User's country. | Update an existing user by ID. |
| `/users/{id}` | DELETE | `id`: ID of the user. | Delete a user by ID. |

## Running Tests

This service comes with both unit and integration tests. You can run them using the following make commands:

- `make unit-test`: Run the unit tests.
- `make integration-test`: Run the integration tests.
- `make start`: Start the service without building.
- `make start-build`: Build and start the service.

## Example Usage

### GET `/users`

Get a list of users with the email set to "xyz@gmail.com":

```bash
curl --location --request GET 'http://localhost:8080/users?query=email%3D%27xyz@gmail.com%27&offset=10&limit=5'
```

### GET `/users/{id}`

Get a user with ID `1bc298ef-5793-41d5-a450-1aff04c6d6f0`:

```bash
curl --location --request GET 'http://localhost:8080/users/1bc298ef-5793-41d5-a450-1aff04c6d6f0'
```

### POST `/users`

Add a new user:

```bash
curl --location --request POST 'http://localhost:8080/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "John Doe",
    "email": "johndoe@example.com",
    "password": "password123",
    "country": "US"
}'
```

### PUT `/users/{id}`

Update a user with ID `1bc298ef-5793-41d5-a450-1aff04c6d6f0`:

```bash
curl --location --request PUT 'http://localhost:8080/users/1bc298ef-5793-41d5-a450-1aff04c6d6f0' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Jane Doe",
    "email": "janedoe@example.com",
    "password": "password123",
    "country": "US"
}'
```

Sure! Here's the documentation for the DELETE `/users/{id}` endpoint:

### DELETE `/users/{id}`

Delete a user with the specified ID.

#### Request

```
DELETE /users/{id}
```

#### Path parameters

| Parameter | Type   | Description            |
| --------- | ------ | ---------------------- |
| id        | string | The ID of the user to delete |

#### Response

| Status code | Description                                |
| ----------- | ------------------------------------------ |
| 204         | User deleted successfully                   |
| 404         | User not found                              |
| 500         | Internal server error                       |

Here's an example curl command to delete a user with ID `1bc298ef-5793-41d5-a450-1aff04c6d6f0`:

```
curl -X DELETE http://localhost:8080/users/1bc298ef-5793-41d5-a450-1aff04c6d6f0
```