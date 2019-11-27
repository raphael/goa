#

![Goa logo](https://goa.design/img/goa-logo.svg "Goa")

Goa is a framework for building micro-services and APIs in Go using a unique
design-first approach.

---
[![Build Status](https://travis-ci.org/goadesign/goa.svg?branch=v3)](https://travis-ci.org/goadesign/goa)
[![Windows Build status](https://ci.appveyor.com/api/projects/status/vixp37loj5i6qmaf/branch/v3?svg=true)](https://ci.appveyor.com/project/RaphaelSimon/goa-oqtis/branch/v3)
[![Godoc](https://godoc.org/goa.design/goa?status.svg)](https://godoc.org/goa.design/goa)
[![Slack](https://img.shields.io/badge/slack-gophers-orange.svg?style=flat)](https://gophers.slack.com/messages/goa/)

## Overview

Goa takes a different approach to building services by making it possible to
describe the *design* of the service API using a simple Go DSL. Goa uses the
description to generate specialized service helper code, client code and
documentation. Goa is extensible via plugins, for example the
[goakit](https://github.com/goadesign/plugins/tree/v3/goakit) plugin
generates code that leverage the Go kit library.

The service design describes the transport independent layer of the services in
the form of simple methods that accept a context and a payload and return a
result and an error. The design also describes how the payloads, results and
errors are serialized in the transport (HTTP or gRPC). For example a service
method payload may be built from an HTTP request by extracting values from the
request path, headers and body. This clean separation of layers makes it
possible to expose the same service using multiple transports. It also promotes
good design where the service business logic concerns are expressed and
implemented separately from the transport logic.

The Goa DSL consists of Go functions so that it may be extended easily to avoid
repetition and promote standards. The design code itself can easily be shared
across multiple services by simply importing the corresponding Go package again
promoting reuse and standardization across services.

## Code Generation

The Goa tool accepts the Go design package import path as input and produces the
interface as well as the glue that binds the service and client code with the
underlying transport. The code is specific to the API so that for example there
is no need to cast or "bind" any data structure prior to using the request
payload or response result. The design may define validations in which case the
generated code takes care of validating the incoming request payload prior to
invoking the service method on the server, and validating the response prior to
invoking the client code.

## Installation

Assuming you have a working [Go](https://golang.org) setup, and are in a
directory where a `go.mod` file is present:

```bash
export GO111MODULE=on
go get -u goa.design/goa/v3
go get -u goa.design/goa/v3/...
```

If you don't have a `go.mod` file present, and only want to install the
Goa command globally:

```bash
go get -u goa.design/goa/v3/cmd/goa
```

Alternatively, when NOT using Go modules (this installs Goa v2, see below):

```bash
go get -u goa.design/goa/...
```

### Goa Versions and Go Module Support

Goa v2 and Goa v3 are functionally the exact same. The only addition provided by
Goa v3 is Go module support. Goa v3 requires Go v1.11 or above, it also requires
projects that use Goa to be within modules.

Projects that use Goa v3 use `goa.design/goa/v3` as root package import path
while projects that use v2 use `goa.design/goa` (projects that use v1 use
`github.com/goadesign/goa`).

Note that the Goa v3 tool is backwards compatible and can generate code for v2
designs. This means that you don't need to swap the tool to generate code for
designs using v2 or v3 (designs using v1 use a different tool altogether).

### Vendoring

Since Goa generates and compiles code vendoring tools are not able to
automatically identify all the dependencies. In particular the `generator`
package is only used by the generated code. To alleviate this issue simply add
`goa.design/goa/codegen/generator` as a required package to the vendor manifest.
For example if you are using `dep` add the following line to `Gopkg.toml`:

``` toml
required = ["goa.design/goa/codegen/generator"]
```

This only applies to Goa v2 as vendoring is not used together with Go modules.

### Stable Versions

Goa follows [Semantic Versioning](http://semver.org/) which is a fancy way of
saying it publishes releases with version numbers of the form `vX.Y.Z` and makes
sure that your code can upgrade to new versions with the same `X` component
without having to make changes.

Releases are tagged with the corresponding version number. There is also a
branch for each major version (`v1`, `v2` and `v3`).

Current Release: `v3.0.8`

## Teaser

Note: the instructions below assume Goa v3.

### 1. Design

Create a new Goa project:

```bash
mkdir -p calcsvc/design
cd calcsvc
go mod init calcsvc
```

Create the file `design.go` in the `design` directory with the following
content:

```go
package design

import . "goa.design/goa/v3/dsl"

// API describes the global properties of the API server.
var _ = API("calc", func() {
        Title("Calculator Service")
        Description("HTTP service for adding numbers, a goa teaser")
        Server("calc", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
        })
})

// Service describes a service
var _ = Service("calc", func() {
        Description("The calc service performs operations on numbers")
        // Method describes a service method (endpoint)
        Method("add", func() {
                // Payload describes the method payload
                // Here the payload is an object that consists of two fields
                Payload(func() {
                        // Attribute describes an object field
                        Attribute("a", Int, "Left operand")
                        Attribute("b", Int, "Right operand")
                        // Both attributes must be provided when invoking "add"
                        Required("a", "b")
                })
                // Result describes the method result
                // Here the result is a simple integer value
                Result(Int)
                // HTTP describes the HTTP transport mapping
                HTTP(func() {
                        // Requests to the service consist of HTTP GET requests
                        // The payload fields are encoded as path parameters
                        GET("/add/{a}/{b}")
                        // Responses use a "200 OK" HTTP status
                        // The result is encoded in the response body
                        Response(StatusOK)
                })
        })
})
```

This file contains the design for a `calc` service which accepts HTTP GET
requests to `/add/{a}/{b}` where `{a}` and `{b}` are placeholders for integer
values. The API returns the sum of `a` and `b` in the HTTP response body.

### 2. Implement

Now that the design is done, let's run `goa` on the design package.
In the `calcsvc` directory run:

``` bash
goa gen calcsvc/design
```

This produces a `gen` directory with the following directory structure:

``` text
gen
├── calc
│   ├── client.go
│   ├── endpoints.go
│   └── service.go
└── http
    ├── calc
    │   ├── client
    │   │   ├── cli.go
    │   │   ├── client.go
    │   │   ├── encode_decode.go
    │   │   ├── paths.go
    │   │   └── types.go
    │   └── server
    │       ├── encode_decode.go
    │       ├── paths.go
    │       ├── server.go
    │       └── types.go
    ├── cli
    │   └── calc
    │       └── cli.go
    ├── openapi.json
    └── openapi.yaml

7 directories, 15 files
```

* `calc` contains the service endpoints and interface as well as a service
  client.
* `http` contains the HTTP transport layer. This layer maps the service
  endpoints to HTTP handlers server side and HTTP client methods client side.
  The `http` directory also contains a complete
  [OpenAPI 2.0](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md)
  spec for the service.

The `goa` tool can also generate example implementations for both the service
and client. These examples provide a good starting point:

``` text
goa example calcsvc/design

calc.go
cmd/calc-cli/http.go
cmd/calc-cli/main.go
cmd/calc/http.go
cmd/calc/main.go
```

The tool generated the `main` functions for two commands: one that runs the
server and one the client. The tool also generated a dummy service
implementation that prints a log message. Again note that the `example` command
is intended to generate just that: an *example*, in particular it is not
intended to be re-run each time the design changes (as opposed to the `gen`
command which should be re-run each time the design changes).

Let's implement our service by providing a proper implementation for the `add`
method. Goa generated a payload struct for the `add` method that contains both
fields. Goa also generated the transport layer that takes care of decoding the
request so all we have to do is to perform the actual sum. Edit the file
`calc.go` and change the code of the `add` function as follows:

```go
// Add returns the sum of attributes a and b of p.
func (s *calcsrvc) Add(ctx context.Context, p *calc.AddPayload) (res int, err error) {
        return p.A + p.B, nil
}
```

That's it! we have now a full-fledged HTTP service with a corresponding OpenAPI
specification and a client tool.

### 3. Run

Now let's compile and run the service:

```bash
cd cmd/calc
go build
./calc
[calcapi] 16:10:47 HTTP "Add" mounted on GET /add/{a}/{b}
[calcapi] 16:10:47 HTTP server listening on "localhost:8088"
```

Open a new console and compile the generated CLI tool:

```bash
cd calcsvc/cmd/calc-cli
go build
```

and run it:

```bash
./calc-cli calc add -a 1 -b 2
3
```

The tool includes contextual help:

``` bash
./calc-cli --help
```

Help is also available on each command:

``` bash
./calc-cli calc add --help
```

Now let's see how robust our code is and try to use non integer values:

``` bash
./calc-cli calc add -a 1 -b foo
invalid value for b, must be INT
run './calccli --help' for detailed usage.
```

The generated code validates the command line arguments against the types
defined in the design. The server also validates the types when decoding
incoming requests so that your code only has to deal with the business logic.

### 4. Document

The `http` directory contains the OpenAPI 2.0 specification in both YAML and
JSON format.

The specification can easily be served from the service itself using a file
server. The [Files](http://godoc.org/goa.design/goa/dsl/http.go#Files) DSL
function makes it possible to server static file. Edit the file
`design/design.go` and add:

```go
var _ = Service("openapi", func() {
        // Serve the file with relative path ../../gen/http/openapi.json for
        // requests sent to /swagger.json.
        Files("/swagger.json", "../../gen/http/openapi.json")
})
```

Re-run `goa gen calcsvc/design` and note the new directory `gen/openapi` and
`gen/http/openapi` which contain the implementation for a HTTP handler that
serves the `openapi.json` file.

All we need to do is mount the handler on the service mux. Add the corresponding
import statement to `cmd/calc/http.go`:

```go
import openapisvr "calcsvc/gen/http/openapi/server"
```

and mount the handler by adding the following line in the same file and after
the mux creation (e.g. one the line after the `// Configure the mux.` comment):

```go
openapisvr.Mount(mux)
```

That's it! we now have a self-documenting service. Stop the running service
with CTRL-C. Rebuild and re-run it then make requests to the newly added
`/swagger.json` endpoint:

``` bash
^C[calcapi] 16:17:37 exiting (interrupt)
[calcapi] 16:17:37 shutting down HTTP server at "localhost:8088"
[calcapi] 16:17:37 exited
go build
./calc
```

In a different console:

``` bash
curl localhost:8088/swagger.json
{"swagger":"2.0","info":{"title":"Calculator Service","description":...
```

## Resources

Consult the following resources to learn more about Goa.

### Docs

See the [goa.design](https://goa.design) website.

### Examples

The [examples](https://github.com/goadesign/examples) directory
contains simple examples illustrating basic concepts.

## Contributing

See [CONTRIBUTING](https://github.com/goadesign/goa/blob/v3/CONTRIBUTING.md).
