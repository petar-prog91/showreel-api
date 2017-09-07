# Requirements
In order to run this API completely out of the box with all microservices running together, I suggest that you follow the docker setup I put in here. If you wish to run all of the microservices on your own, I might provide later on an explanation how to do it, but I strongly suggest following Docker Compose standard.

- Docker & Docker Compose 3+

Requirements bellow come by using Docker setup mentioned here, so you don't have to think about installing it.
- GoLang 1.7+
- MySQL

# Installation
- Install Docker & Docker Compose.
- Checkout the code from master branch.
- Navigate into the folder of the project.
- In your terminal run command `docker-compose build`.
- Once that is finished, run `docker-compose up`.
- By default API will run under `http://localhost:8080`

# Authentification
If you imported SQL file provided in this repo under Web folder, you will have user with username: `admin` and password: `admin12345`

Currently available endpoints:
- POST /api/authenticate -> authenticate user (required JWT)
- GET /api/users -> return all the users
- GET /api/users/:id -> return a single user with ID
- POST /api/users/ -> creates a new user
- DELETE /api/users/:id -> deletes a user with ID
- PUT /api/users/:id -> updated a user with ID

All endpoints except Authenticate endpoint are protected with JWT Auth. In order to work with endpoints you will need JWT token in your header for each HTTP request.

To get JWT token you need to make HTTP POST request to `/api/authenticate` with body containing JSON:
```json
{
	"username": "admin",
	"password": "admin12345"
}
```
IF authentification is successful, you will get JWT token string.

Once you get JWT token string, you need to make a custom header `AUTH_JWT_TOKEN` with token string as a value.