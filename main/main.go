package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"october26Project/entities"
	"october26Project/repositories"
	"october26Project/services"
	"strconv"
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

func main() {
	InitDB()
	defer CloseDB()

	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(*itemRepository)

	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			idStr := r.URL.Query().Get("id")
			model := r.URL.Query().Get("model")
			minPriceStr := r.URL.Query().Get("minPrice")
			maxPriceStr := r.URL.Query().Get("maxPrice")

			if idStr != "" && model == "" && minPriceStr == "" && maxPriceStr == "" {
				id, _ := strconv.Atoi(idStr)
				itemDTO := itemService.GetItemByID(int64(id))
				json.NewEncoder(w).Encode(itemDTO)
			} else if idStr == "" && model == "" && minPriceStr == "" && maxPriceStr == "" {
				itemDTOs := itemService.GetAllItems()
				json.NewEncoder(w).Encode(itemDTOs)
			} else if idStr == "" && model != "" && minPriceStr == "" && maxPriceStr == "" {
				itemDTOs := itemService.GetItemsByModel(model)
				json.NewEncoder(w).Encode(itemDTOs)
			} else if idStr == "" && model == "" && minPriceStr != "" && maxPriceStr != "" {
				minPrice, err := strconv.ParseFloat(minPriceStr, 64)
				if err != nil {
					log.Fatal(err)
				}
				maxPrice, err := strconv.ParseFloat(maxPriceStr, 64)
				if err != nil {
					log.Fatal(err)
				}
				itemDTOs := itemService.GetItemsByPriceFilter(minPrice, maxPrice)
				json.NewEncoder(w).Encode(itemDTOs)
			}
		} else if r.Method == http.MethodPost {
			var item entities.Item
			json.NewDecoder(r.Body).Decode(&item)
			itemService.AddItem(item)
		} else if r.Method == http.MethodPut {
			var updItem entities.Item
			json.NewDecoder(r.Body).Decode(&updItem)
			itemService.UpdateItem(updItem)
		} else if r.Method == http.MethodDelete {
			idStr := r.URL.Query().Get("id")
			id, _ := strconv.Atoi(idStr)
			itemService.DeleteItem(int64(id))
		}
	})

	server := http.Server{
		Addr: "localhost:4040",
	}
	fmt.Println("Server application is starting", server.Addr)
	server.ListenAndServe()
}
