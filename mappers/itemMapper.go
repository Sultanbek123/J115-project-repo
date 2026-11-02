package mappers

import (
	"october26Project/dtos"
	"october26Project/entities"
)

func MapToDTO(item entities.Item) dtos.ItemDTO {
	var itemDTO dtos.ItemDTO
	itemDTO.ID = item.ID
	itemDTO.Model = item.Model
	itemDTO.Price = item.Price
	itemDTO.Weight = item.Weight
	return itemDTO
}

func MapToDTOs(items []entities.Item) []dtos.ItemDTO {
	var itemDTOs []dtos.ItemDTO
	for i := 0; i < len(items); i++ {
		itemDTO := MapToDTO(items[i])
		itemDTOs = append(itemDTOs, itemDTO)
	}
	return itemDTOs
}
