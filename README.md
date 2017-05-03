# GOAPI

## Table of Contents
* [Summary](#summary)
* [Demo](#demo)
* [Technologies Used](#technologies-used)
* [Building](#building)
* [Running](#running)
* [Library Usage](#library-usage)
* [Extending the API](#extending-the-api)

## Summary
This golang library is a framework for implementing an API server.  The
statically-linked all-in-one CLI binary built by this project may run as either
the client or the server, depending on the command line arguments.

The API server responds to both GRPC and JSON REST requests.  The framework
also includes an `/auth` endpoint which responds with JWT auth tokens,
middleware that enforces authentication, a user database, SSL encryption, and
other standard requirements for any API server.

The example client and server implementation demonstrates a key/value store,
which is only available to authenticated users.

## Demo

![Animated Image of Terminal](https://github.com/dcwangmit01/goapi/raw/master/demo/demo.gif)

## Technologies Used

The goapi library uses [GRPC](http://www.grpc.io/) and
[grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway), which allows us
to code-generate API service and client libraries.  Alternative golang API frameworks to
GRPC/GW include API toolkits such as [Gin](https://github.com/gin-gonic/gin) or
[Gorilla](http://www.gorillatoolkit.org/).  GRPC is nice to use because of
code-generation.

This project provides:

* An API server implementation which responds to requests over
  [GRPC](http://www.grpc.io/) [protocol
  buffers](https://developers.google.com/protocol-buffers/) as well as JSON
  Rest.  All JSON Rest requests are proxied over
  [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) to the GRPC
  server.
* Code-generation capabilities for API service and client libraries, as
  described by protocol buffers `.proto` service description file.
* A [cobra](https://github.com/spf13/cobra) CLI tool with subcommands that can
  launch the API server as well as drive the API as a CLI client.  The CLI tool
  stores its settings in [viper](https://github.com/spf13/viper)
  configuration.
* A Makefile which generates code resources including:
    * GRPC code files including protobuf, grpc, grpc-gateway, swagger
    * SSL certs using [cfssl](https://github.com/cloudflare/cfssl)
    * Swagger asset files with embedded swagger-ui, to serve the swagger file
    * Cross-compiling statically linked Linux binaries for amd64 and armv7.
* A test framework using [ginkgo](https://github.com/onsi/ginkgo) and [gomega
  matchers](https://github.com/onsi/gomega)
* Integration with a semantic version package vendoring system using
  [glide](https://github.com/Masterminds/glide)
* Provides an `/auth` endpoint which returns a JWT token when presented with
  valid credentials.  This authentication system is based on
  [JWT](github.com/dgrijalva/jwt-go).
* An [Role Based Access Control](https://github.com/mikespook/gorbac)
  authorization system.
* A structured logging system using
  [logrus](https://github.com/sirupsen/logrus)
* The tools to code-generate the API framework from services and endpoints
* A basic user database
* An example key/value API service for storing values to sqlite.

## Building

* To print Makefile usage, run `make`
* To build everyting run `make all`

```bash
$ make
all                            run all targets
cert_gen                       generate go-bindata cert files
check                          checks
clean                          delete all non-repo files
code_gen                       generate grpc go files from proto spec
compile                        build the binaries for amd64
compilex                       build the binaries for all platforms
demo                           run and record the demo-magic script
deps                           install host dependencies
resource_gen                   generate go-bindata swagger files
vendor                         install/build all 3rd party vendor libs and bins
help                           Print list of Makefile targets
```

## Running

```bash
# Show Usage
goapi

# Start the GRPC+JSON server on port 10080
goapi serve

### Use the GRPC CLI Client

# Obtain an Auth Token
goapi auth login

# Create a Key
goapi keyval create mykey myval

# Read a Key
goapi keyval read mykey

# Update a Key
goapi keyval update mykey myval2

# Delete a Key
goapi keyval delete mykey

### Use Curl

# Obtain an Auth Token
TOKEN=$(curl -fsSL -X POST -k https://localhost:10080/v1/auth -H "Content-Type: text/plain" -d '{"grant_type": "password", "username": "admin", "password": "password"}' | jq --raw-output '.access_token')

# Create a Key
curl -k -X PUT    https://localhost:10080/v1/keyval/mykey -H "Content-Type: text/plain" -H "Authorization: Bearer $TOKEN" -d '{"value": "myval1"}'

# Read a Key
curl -k -X GET    https://localhost:10080/v1/keyval/mykey -H "Content-Type: text/plain" -H "Authorization: Bearer $TOKEN"

# Update a Key
curl -k -X POST   https://localhost:10080/v1/keyval/mykey -H "Content-Type: text/plain" -H "Authorization: Bearer $TOKEN" -d '{"value": "myval2"}'

# Delete a Key
curl -k -X DELETE https://localhost:10080/v1/keyval/mykey -H "Content-Type: text/plain" -H "Authorization: Bearer $TOKEN"

### View the Swagger UI

# View the swagger-ui
curl -k https://localhost:10080/swagger-ui/

# View the swagger file that has been generated from GRPC
curl -k https://localhost:10080/swagger.json
```

## Library Usage

To use this project as a library copy the contents of `<goapi_project>/example`
to your own project `<your_project>`.  Rewrite the imports with `sed -i
's@dcwangmit01/example@your_project`.  Then copy over and edit the Makefile

## Extending the API

To extend the api, edit any the `.proto` files to add additional Protobuf
messages as well as services.  There is a common proto file in the library
`pb/goapi.proto` as well as a proto file in the example `example/pb/app.proto`.

Here's how to extend the examples API

```bash

# Edit the .proto files to add messages as well as services
#   As example, consider you create a new Tasklist service
<your_fav_editor> example/pb/app.proto

# Update the code-generated files (*.pb.go, *pb.gw.go)
make code_gen

# Create an implementation of the new TasklistServer interface that you've
#   defined in the .proto file, which has been auto-generated for you in the
#   *.pb.go file.  Follow the existing keyval example.
<your_fav_editor> example/service/tasklist.go

# Add a subcommand to the CLI tool that can drive the new Tasklist service.
#   This subcommand will call the TasklistClient that has been auto-generated
#   for you in the *.pb.go file.  Follow the existing keyval example.
<your_fav_editor> example/cmd/tasklist.go
```

