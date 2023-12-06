
# Assestment by Tempo

Reza Wahyu Ramadhan
[![Built with GO-Gin Framework](https://img.shields.io/badge/built%20with-GoGin-ff69b4.svg?logo=cookiecutter)](https://github.com/karec/cookiecutter-flask-restful)

## Documentation

Create your own database with mysql engine.

Copy the .env-example to .env and set up the credential.

```
cp .env-example .env
```

run main program

```
go run cmd/*
```

or if you using windows run the main program without asterisk symbol

```
go run cmd/
```

The API documentation is on [Documentation](https://documenter.getpostman.com/view/11131161/2s9YeN1oGj) or access the json file Tempo Api.postman_collection.json on this repository

## QUESTION

#### How to handle the middleware for user authentication


This projects is using Json Web Token (JWT) to handle user authentication, the benefit of using JWT isbeing a streamlined and secure authentication process. JWTs provide a standardized and stateless method for representing claims between parties, making them an ideal solution for middleware. By including JWT as a middleware component, i can implement a token-based authentication mechanism, reducing the need for server-side sessions and associated storage.
