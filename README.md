# go-env
    
[![Go Reference](https://pkg.go.dev/badge/github.com/dmitrymomot/go-env.svg)](https://pkg.go.dev/github.com/dmitrymomot/go-env)
[![Go Report Card](https://goreportcard.com/badge/github.com/dmitrymomot/go-env)](https://goreportcard.com/report/github.com/dmitrymomot/go-env)
[![Coverage Status](https://coveralls.io/repos/github/dmitrymomot/go-env/badge.svg?branch=master)](https://coveralls.io/github/dmitrymomot/go-env?branch=master)
[![Tests](https://github.com/dmitrymomot/go-env/actions/workflows/tests.yml/badge.svg)](https://github.com/dmitrymomot/go-env/actions/workflows/tests.yml)

Helpers to work with environment variables.


## Usage Example

```go
package main

import (
    "fmt"
    "net/http"
    "github.com/dmitrymomot/go-env"
    "github.com/go-chi/chi"
)

var (
    httpPort = env.MustInt("HTTP_PORT") // required env variable, throws fatal error if it empty
    healthEndpoint = env.GetString("HEALTH_ENDPOINT", "/health") // optional env variable with fallback value
)

func main() {
    r := chi.NewRouter()
    r.Get(healthEndpoint, func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusNoContent)
    })
    http.ListenAndServe(fmt.Sprintf(":%d", httpPort), r)
}
```