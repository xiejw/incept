package main

import (
    "fmt"
    "github.com/pelletier/go-toml/v2"
)

type Service struct {
    Port          int      `toml:"port"`
    Enabled       bool     `toml:"enabled"`
    Dependencies  []string `toml:"dependencies"`
}

type Config struct {
    Params   map[string]any     `toml:"params"`
    Services map[string]Service `toml:"services"`
}

func main() {
    data := []byte(`
        [params]
        region = "us-west-1"
        env = "prod"

        [services.api]
        port = 8080
        enabled = true
        dependencies = ["db", "cache"]

        [services.db]
        port = 5432
        enabled = false
        dependencies = []

        [services.cache]
        port = 6379
        enabled = true
        dependencies = ["db"]
    `)

    var cfg Config
    if err := toml.Unmarshal(data, &cfg); err != nil {
        panic(err)
    }

    fmt.Println("Params:", cfg.Params)
    for name, svc := range cfg.Services {
        fmt.Printf("Service %s: %+v\n", name, svc)
    }

    fmt.Println("API dependencies:", cfg.Services["api"].Dependencies)
}

