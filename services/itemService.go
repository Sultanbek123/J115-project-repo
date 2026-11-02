package services

import (
	"october26Project/dtos"
	"october26Project/entities"
	"october26Project/mappers"
	"october26Project/repositories"
)

type ItemService struct {
	itemRepository repositories.ItemRepository
}

func NewItemService(itemRepository repositories.ItemRepository) *ItemService {
	return &ItemService{itemRepository: itemRepository}
}

func (itemService *ItemService) AddItem(item entities.Item) {
	if item.Model != "" && item.Price >= 0 && item.Weight > 0 {
		itemService.itemRepository.AddItem(item)
	}
}

func (itemService *ItemService) UpdateItem(updItem entities.Item) {
	itemService.itemRepository.UpdateItem(updItem)
}

func (itemService *ItemService) DeleteItem(id int64) {
	itemService.itemRepository.DeleteItem(id)
}

func (itemService *ItemService) GetItemByID(id int64) dtos.ItemDTO {
	item := itemService.itemRepository.GetItemByID(id)
	itemDTO := mappers.MapToDTO(item)
	return itemDTO
}

func (itemService *ItemService) GetAllItems() []dtos.ItemDTO {
	items := itemService.itemRepository.GetAllItems()
	itemDTOs := mappers.MapToDTOs(items)
	return itemDTOs
}

func (itemService *ItemService) GetItemsByModel(model string) []dtos.ItemDTO {
	items := itemService.itemRepository.GetItemsByModel(model)
	itemDTOs := mappers.MapToDTOs(items)
	return itemDTOs
}

func (itemService *ItemService) GetItemsByPriceFilter(minPrice float64, maxPrice float64) []dtos.ItemDTO {
	items := itemService.itemRepository.GetItemsByPriceFilter(minPrice, maxPrice)
	itemDTOs := mappers.MapToDTOs(items)
	return itemDTOs
}
