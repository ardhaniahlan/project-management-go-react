package utils

import (
	"project-management-be/models"

	"github.com/google/uuid"
)

func SortingListsByPosition(lists []models.List, order []uuid.UUID) []models.List {
	ordered := make([]models.List, 0, len(order))

	listMap := make(map[uuid.UUID]models.List)

	for _, list := range lists {
		listMap[list.PublicID] = list
	}

	for _, publicID := range order {
		if list, exists := listMap[publicID]; exists {
			ordered = append(ordered, list)
		}
	}
	return ordered
}