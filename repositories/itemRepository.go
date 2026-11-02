package repositories

import (
	"database/sql"
	"log"
	"october26Project/entities"
)

type ItemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{db: db}
} //цель создавать обьект репозитория когда конструктор вызывают

func (itemRepository *ItemRepository) AddItem(item entities.Item) {
	_, err := itemRepository.db.Exec("insert into items(model, price, weight,invoice_price) values($1,$2,$3,$4)", item.Model, item.Price, item.Weight, item.InvoicePrice)
	if err != nil {
		log.Fatal(err)
	}
}

func (itemRepository *ItemRepository) UpdateItem(updItem entities.Item) {
	_, err := itemRepository.db.Exec("update items set model = $1, price = $2, weight = $3 where id = $4", updItem.Model, updItem.Price, updItem.Weight, updItem.ID)
	if err != nil {
		log.Fatal(err)
	}
}

func (itemRepository *ItemRepository) DeleteItem(id int64) {
	_, err := itemRepository.db.Exec("delete from items where id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
}

func (itemRepository *ItemRepository) GetAllItems() []entities.Item {
	var items []entities.Item
	rows, err := itemRepository.db.Query("select * from items")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var item entities.Item
		rows.Scan(&item.ID, &item.Model, &item.Price, &item.Weight, &item.InvoicePrice)
		items = append(items, item)
	}
	return items
}

func (itemRepository *ItemRepository) GetItemByID(id int64) entities.Item {
	var item entities.Item
	row := itemRepository.db.QueryRow("select * from items where id = $1", id)
	row.Scan(&item.ID, &item.Model, &item.Price, &item.Weight, &item.InvoicePrice)
	return item
}

func (itemRepository *ItemRepository) GetItemsByModel(model string) []entities.Item {
	var items []entities.Item
	rows, err := itemRepository.db.Query("select * from items where model = $1", model)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var item entities.Item
		rows.Scan(&item.ID, &item.Model, &item.Price, &item.Weight, &item.InvoicePrice)
		items = append(items, item)
	}
	return items
}

func (itemRepository *ItemRepository) GetItemsByPriceFilter(minPrice float64, maxPrice float64) []entities.Item {
	var items []entities.Item
	rows, err := itemRepository.db.Query("select * from items where price >= $1 and price <= $2", minPrice, maxPrice)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var item entities.Item
		rows.Scan(&item.ID, &item.Model, &item.Price, &item.Weight, &item.InvoicePrice)
		items = append(items, item)
	}
	return items
}
