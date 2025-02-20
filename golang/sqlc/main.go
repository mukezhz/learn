package main

import (
    "github.com/mukezhz/learn/tree/main/golang/sqlc/bootstrap"

    "github.com/joho/godotenv"
)

func main() {
    _ = godotenv.Load()
    _ = bootstrap.RootApp.Execute()
}
