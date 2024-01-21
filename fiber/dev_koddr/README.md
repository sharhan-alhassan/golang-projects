
## [Tutorial Source](https://dev.to/koddr/welcome-to-fiber-an-express-js-styled-fastest-web-framework-written-with-on-golang-497)


## Setup
```sh
# Create dedicated docker network
docker network create -d bridge dev-network

# Create PostgreSQL container
docker run --rm -d \
    --name dev-postgres \
    --network dev-network \
    -e POSTGRES_USER=postgres \
    -e POSTGRES_PASSWORD=password \
    -e POSTGRES_DB=postgres \
    -v ${HOME}/dev-postgres/data/:/var/lib/postgresql/data \
    -p 5432:5432 \
    postgres

# Migrate db 
migrate \
    -path $(PWD)/platform/migrations \
    -database "postgres://postgres:password@localhost/postgres?sslmode=disable" \
    up
    
# Build Fiber image
docker build -t fiber:0.0.1 .

# Start container
docker run --rm -d \
    --name dev-fiber \
    --network dev-network \
    -p 5000:5000 \
    fiber:0.0.1
```