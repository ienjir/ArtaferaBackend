# ArtaferaBackend

The backend for the Artafera Website

## Startup

To start the Artafera backend you the following things:

- Go 1.23.2
- Docker

Run `go mod tidy` to install all dependency's. Then run `docker compose up -d` to start the postgres server, but only if you are not in production.

## Production

To start using this project for prod make the following steps:
- In `database.go` remove the loop that drops all tables when rebuilding the application to prevent data loss
- In `main.go` remove the `database.GenerateFakeData()` call to not have data when running the application
- In `main.go` set the right proxies
- In `docker-compose.yml` change the credentials to the database because currently they are not safe and only for testing
- Put the right credentials for the datbase in the `.env` file
- Switch the mode in the `.env` file to production

Afterwards compose docker with `docker compose up -d`. 