package utils

import (
	"github.com/google/uuid"
	"github.com/rizqizyd/project-management-be/models"
)

func SortListsByPosition(lists []models.List, order []uuid.UUID) []models.List {
	ordered := make([]models.List, 0, len(order))

	listMap := make(map[uuid.UUID]models.List)
	for _, list := range lists {
		listMap[list.PublicID] = list
	}

	for _, id := range order {
		if list, exists := listMap[id]; exists {
			ordered = append(ordered, list)
		}
	}

	return ordered
}
