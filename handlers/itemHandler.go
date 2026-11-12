package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"october26Project/entities"
	"october26Project/services"
	"strconv"
	"time"
)

type ItemHandler struct {
	itemService services.ItemService
}

func NewItemHandler(itemService services.ItemService) *ItemHandler {
	return &ItemHandler{itemService: itemService}
}

func (itemHandler *ItemHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idStr := r.URL.Query().Get("id")
		model := r.URL.Query().Get("model")
		minPriceStr := r.URL.Query().Get("minPrice")
		maxPriceStr := r.URL.Query().Get("maxPrice")

		if idStr != "" && model == "" && minPriceStr == "" && maxPriceStr == "" {
			id, _ := strconv.Atoi(idStr)
			itemDTO := itemHandler.itemService.GetItemByID(int64(id))
			json.NewEncoder(w).Encode(itemDTO)
		} else if idStr == "" && model == "" && minPriceStr == "" && maxPriceStr == "" {
			ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
			defer cancel()
			itemDTOs, err := itemHandler.itemService.GetAllItems(ctx)
			if err != nil {
				log.Println("Не удалось получить товары:", err)
			}
			json.NewEncoder(w).Encode(itemDTOs)
		} else if idStr == "" && model != "" && minPriceStr == "" && maxPriceStr == "" {
			itemDTOs := itemHandler.itemService.GetItemsByModel(model)
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
			itemDTOs := itemHandler.itemService.GetItemsByPriceFilter(minPrice, maxPrice)
			json.NewEncoder(w).Encode(itemDTOs)
		}
	}
}

func (itemHandler *ItemHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var item entities.Item
		json.NewDecoder(r.Body).Decode(&item)
		itemHandler.itemService.AddItem(item)
	}
}

func (itemHandler *ItemHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		var updItem entities.Item
		json.NewDecoder(r.Body).Decode(&updItem)
		itemHandler.itemService.UpdateItem(updItem)
	}
}

func (itemHandler *ItemHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		idStr := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idStr)
		itemHandler.itemService.DeleteItem(int64(id))
	}
}
