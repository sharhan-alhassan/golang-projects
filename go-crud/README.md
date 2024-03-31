
## Setup
```sh
mkdir -p $GOPATH/src/github.com/sharhan-alhassan/golang-projects/go-crud && cd "$_"
```

## 1. Install Go Compile Daemon
```sh
# Set up
go mod init

# 1. Install Go Compile Daemon
go get github.com/githubnemo/CompileDaemon

# Install it
go install github.com/githubnemo/CompileDaemon

## Usage
CompileDaemon -command="./go-crud" 

```

## 2. Go dotenv to manage/load .env variables
```sh
go get github.com/joho/godotenv

# Install it
go install github.com/joho/godotenv
```

## 3. Install Gin framework
```sh
go get -u github.com/gin-gonic/gin

```

## Install Go ORM library
```sh
# Install gorm library
go get -u gorm.io/gorm

# Install postgres driver for gorm
go get -u gorm.io/driver/postgres

```

## Define a Model
```sh

type Post struct {
    gorm.Model
    Title string
    Body string
}

- The gorm.Model will auto include fields in its parent's struct

- Example of such fields are: 

CreatedAt time.Time
UpdatedAt time.Time
DeletedAt gorm.DeletedAt
```

## Migrate a Database
```sh
go run migrate/migrate.go

# Output
➜  go-crud go run migrate/migrate.go

2024/03/21 02:17:03 /home/sharhan/go/1.20.4/src/github.com/sharhan-alhassan/golang-projects/go-crud/migrate/migrate.go:19 SLOW SQL >= 200ms
[729.379ms] [rows:1] SELECT count(*) FROM information_schema.tables WHERE table_schema = CURRENT_SCHEMA() AND table_name = 'posts' AND table_type = 'BASE TABLE'

2024/03/21 02:17:03 /home/sharhan/go/1.20.4/src/github.com/sharhan-alhassan/golang-projects/go-crud/migrate/migrate.go:19 SLOW SQL >= 200ms
[277.690ms] [rows:0] CREATE TABLE "posts" ("id" bigserial,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz,"title" text,"body" text,PRIMARY KEY ("id"))

2024/03/21 02:17:04 /home/sharhan/go/1.20.4/src/github.com/sharhan-alhassan/golang-projects/go-crud/migrate/migrate.go:19 SLOW SQL >= 200ms
[213.674ms] [rows:0] CREATE INDEX IF NOT EXISTS "idx_posts_deleted_at" ON "posts" ("deleted_at")

```

## Noes

- Use `Context.JSON` instead of `Context.IndentedJSON` because `Context.IndentedJSON` is CPU intensive because of the pretty printed format of the data