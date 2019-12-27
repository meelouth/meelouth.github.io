package main

import (
	"os"

	"github.com/kurenkoff/bstu_backend/repository"
	"github.com/kurenkoff/bstu_backend/server"

	_ "github.com/go-sql-driver/mysql"
)

const DSN = "root:1234@tcp(localhost:3338)/games?collation=utf8mb4_unicode_ci&parseTime=true"

func main() {
	DSN := os.Getenv("DSN")
	sql := repository.New(DSN)

	srv := server.New(sql)
	srv.Run()
}
