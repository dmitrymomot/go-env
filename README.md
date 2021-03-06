# go-env
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