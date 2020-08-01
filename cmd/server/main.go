package main

import (
    "github.com/jyotishp/go-orders/pkg/http"
    "sync"
    "time"
)

const (
    JwtSecret = "some_super_secret_token"
    JwtTtl = 5*time.Minute
)

func main() {
    go http.StartGRPC(JwtSecret, JwtTtl)
    go http.StartHTTP()

    var wg sync.WaitGroup
    wg.Add(1)
    wg.Wait()
}