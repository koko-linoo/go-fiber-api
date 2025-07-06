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
