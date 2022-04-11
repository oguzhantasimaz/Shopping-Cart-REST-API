# Shopping-Cart-REST-API
Shopping Cart Rest API with Golang

## Introduction
This is E-commerce API app implementations and written in Golang using gin web framework.
This is not a finished project by any means, but it is a starting point for learning and understanding the concepts.
I tried to implement clean architecture and code in a way that it is easy to understand and maintain.

## Features
- Authentication / Authorization
- JWT middleware for authentication
- SHA512 hashing and base64 encoding for password
- GORM ORM for database
- CRUD operations on products, carts, orders, categories, users
- Orders, users may place an order
- Graceful shutdown

## Project structure:
- CMD: Main application file
- Config: Configuration files
- Repositories: CRUD operations handling
- Models: Mvc, it is our domain data.
- Controllers: This is the Model-View-Controller, they receive the request from the user, they ask the services to perform an action for them on the database.
- Services: contains some business logic for each model, and for authorization
- Middlewares: it contains middlewares(golang functions) that are triggered before the controller action, for example, a middleware which
  reads the request looking for the Jwt token and trying to authenticate the user before forwarding the request to the corresponding controller
  action
- PKG: it contains the packages that are used by the application


## TODO:
- [ ] Fix known bugs
- [ ] Heroku deployment
- [ ] GithubActions (CI/CD) for deployment
- [ ] Swagger Integration
- [ ] Unit Tests
- [ ] Integration Tests
- [ ] Create Bulk Category with CSV
- [ ] Pagination
- [ ] Add query params to get all products' endpoint for filtering
