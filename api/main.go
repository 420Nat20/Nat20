package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/420Nat20/Nat20/nat-20/server"
	"github.com/420Nat20/Nat20/nat-20/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx, _ := context.WithCancel(context.Background())
	api := &server.Server{
		Ctx: ctx,
	}

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s dbname=%s user=%s password=%s port=%s",
			os.Getenv("PSQL_HOST"),
			os.Getenv("PSQL_DBNAME"),
			os.Getenv("PSQL_USER"),
			os.Getenv("PSQL_PASS"),
			os.Getenv("PSQL_PORT")))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB created")

	api.DB = db

	api.UserService = service.UserService{
		Ctx: ctx,
		DB:  db,
	}

	api.InitServer()
}
