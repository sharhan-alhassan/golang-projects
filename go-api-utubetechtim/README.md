## Tutotial Source: https://www.youtube.com/watch?v=bj77B59nkTQ

## Set up
```sh
cd /home/sharhan/go/1.20.4/src/github.com/sharhan-alhassan/golang-projects/go-api-utubetechtim

go mod init
```

## Install dependencies
```sh
go get -u github.com/gin-gonic/gin

```

## Running with CURL utitily
```sh
# Send GET request to query all books
curl localhost:8080/books

# Send POST request to create a book using the payload in body.json file
curl localhost:8080/books --include --header "Content-Type: applicatoin/json" -d @body.json --request "POST"

# Send PATCH request to checkout book
curl localhost:8080/checkout?id=1 --include --header "Content-Type: applicatoin/json" --request "PATCH"

```