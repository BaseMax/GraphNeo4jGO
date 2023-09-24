# Graph Neo4j Go

This is a simple twitter-like api written in go as an example to use neo4j with golang.

### Technologies

- [neo4j](https://neo4j.com/)
- postgresql
- [mux](https://github.com/gorilla/mux)
- [jwt](https://github.com/golang-jwt/jwt)
- [go-playground/validator](https://github.com/go-playground/validator)


### Api documentation
A guide for endpoints.

#### User endpoints
Base path: `host://api/v1/user`

- **Register**
  - Path: `/register/`
  - Method: `POST`
  - Request body:
```json
{
    "username": "username_length_higher_than_6",
    "name": "name of user",
    "email": "valid_email@example.com",
    "password":"a_good_password",
    "gender": 1
}
  ```
  - Response:
```json
{
	"status": "Created",
	"id": 5,
	"token": "<jwt-token>"
}
```

  - Description: available genders are: Male=1, Female=2, Other=3. use token to access authenticated endpoints.

- **Login**
  - Path: `/login/`
  - Method: `POST`
  - Request body:
```json
{
	"username": "username",
	"password": "user_password"
}
```
  - Response body:
```json
{
	"status": "Found",
	"id": 5,
	"token": "<jwt-token>"
}
```

- **Info**
  - _**Auth**_ : need jwt token in headers as Bearer token.
  - Path: `/info/{username}`
  - Method: `GET`
  - Response Body:
```json
{
	"status": "Found",
	"data": {
		"id": 3,
		"username": "username",
		"name": "users name",
		"email": "valid_email@example.com",
		"gender": 1
	}
}
```

- **Delete**
  - _**Auth**_
  - Path: `/delete/`
  - Method: `DELETE`
  - Response Body:
```json
{
	"status": "Deleted",
	"id": 4
}
```

Copyright 2023, Max Base
