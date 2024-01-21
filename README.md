# golang-projects
All Golang-based projects

# Setup and Installation of packages
```sh
# 1. Initilize project as module
go mod init github.com/sharhan-alhassan/golang-projects/fiber/official

# 2. Install your packages
go get github.com/gofiber/fiber/v2  

# 3. 
```

# Modules Setup
```sh
# 1. Say you have a package locally that you want to access from an application
go mod init example.com/greetings

# 2. In the application which will access it
go mod init example.com/hello

# 3. Since the exampl.e.com/greetings is not publich in remote, you need to make
# reference to it from the golang application that it's available on local machine
# directory somewhere
# On the Go app, do
go mod edit -replace example.com/greetings=../02-greetings

# Meaning check the 02-greetings/ directory one step back and make reference to its package

# 4. Install it
go mod tidy
```

# Build and Run
```sh
# Build the app
go build

# Run the executable
./hello-module

# Install and run without using ./
# To install, get your install directory
go list -f '{{.Target}}'
# output
/Users/sharhan/go/bin/hello-module

# Add the path to the PATH
export PATH=$PATH:/Users/sharhan/go/bin/hello-module

# Now install
go install
```