package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"october26Project/handlers"
	"october26Project/repositories"
	"october26Project/services"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var db *sql.DB //object для взаим с базой

func InitDB() {
	connection := "user=postgres password=postgres dbname=october26ProjectDB host=localhost port=5432 sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Success connection to Postgres")
}

func CloseDB() {
	var err error
	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func migrateUp() {
	dsn := "postgres://postgres:postgres@localhost:5432/october26ProjectDB?sslmode=disable"
	m, err := migrate.New("file://db/migrations", dsn)
	if err != nil {
		panic(err)
	}
	m.Up()
	log.Println("Success migration")
}

func migrateDown() {

}

func main() {
	InitDB()
	defer CloseDB()

	migrateUp()

	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(*itemRepository)
	itemHandler := handlers.NewItemHandler(*itemService)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /items", itemHandler.GetItems)
	mux.HandleFunc("POST /items", itemHandler.AddItem)
	mux.HandleFunc("PUT /items", itemHandler.UpdateItem)
	mux.HandleFunc("DELETE /items", itemHandler.DeleteItem)

	server := http.Server{
		Addr:    "localhost:4040",
		Handler: mux,
	}

	go func() {
		fmt.Println("Server application is starting", server.Addr)
		server.ListenAndServe()
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	log.Println("Получен сигнал остановки сервера, жду пока оставшиеся запросы доработают")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := server.Shutdown(shutdownCtx)
	if err != nil {
		log.Println("Не успели завершить запросы", err)
	}
	log.Println("Сервер остановлен")
}
