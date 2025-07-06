# go-fiber-api

A simple RESTful API built with Go and Fiber.

## Getting Started

To get started, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/koko-linoo/go-fiber-api.git
```

2. Change into the project directory:

```bash
cd go-fiber-api
```

3. Install the dependencies:

```bash
go mod download
```

4. Run the application:

```bash
air
```

## Usage

### Create a User

To create a new user, send a POST request to the `/users` endpoint with the following JSON body:

```json
{
  "fullName": "John Doe",
  "username": "johndoe",
  "email": "johndoe@example.com",
  "phone": "1234567890",
  "password": "password"
}
```

### Update a User

To update a user, send a PUT request to the `/users/:id` endpoint with the following JSON body:

```json
{
  "fullName": "Jane Doe",
  "username": "janedoe",
  "email": "janedoe@example.com",
  "phone": "9876543210",
  "password": "password"
}
```

### Get a User

To get a user, send a GET request to the `/users/:username` endpoint.

### Get All Users

To get all users, send a GET request to the `/users` endpoint.

### Delete a User

To delete a user, send a DELETE request to the `/users/:id` endpoint.

## License

This project is licensed under the MIT License.

## Author

[Koko Linoo](https://github.com/koko-linoo)
