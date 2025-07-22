# ArtaferaBackend

### The backend for the <a href="https://wwww.Artafera.ch">Artafera Website</a>

## Development

To start the Artafera backend you the following things:

- Go 1.24.4
- Docker

Run `go mod tidy` to install all dependency's. Then run `docker compose up -d` to start the postgres and minio services, but only if you are not in production. <br>
Afterwards go to <a href="localhost:9000">`localhost:9000`</a>, log in with the credentials and generate an access/private token and put them in the `.env` file.

#### Database
- Database: `PostgreSQL`
- Host: `localhost`
- Database: `ArtaferaDB`
- Port: `5432` (Default)
- User: `DBAdmin`
- Password: `AVerySecurePassword`

#### MinIO
- Port: `9000`
- Admin page: `localhost:9001`
- User: `S3Admin`
- Password: `AVerySecurePassword`

#### Bruno
There is also a Bruno folder called `bruno` (duh) in which I have (hopefully someday) all requests. Because Bruno is new (and currently a bit shitty) there are some issues. So until they are so kind and fix these issues its these issues you are stuck with copy pasting your generated JWT into the auth tab of the collection. If i write docs about it it will be in the `docs` folder and not in bruno itself.

## Production

To start using this project for prod make the following steps:
- In `main.go` set the right proxies
- In `docker-compose.yml` change the credentials to the database because currently they are not safe and only for development
- Put the right credentials in the `.env` file for all services 
- Switch the mode in the `.env` file to `release`

Afterward run `docker compose up -d`. 

For any more specific documentation refer to the right `md` file in the `src/docs` folder