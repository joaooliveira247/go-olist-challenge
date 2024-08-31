package utils

import "github.com/google/uuid"

func SlicesHasSameUUIDs(a, b []uuid.UUID) bool {
	if len(a) != len(b) {
		return false
	}

	countMap := make(map[uuid.UUID]int)

	for _, v := range a {
		countMap[v]++
	}

	for _, v := range b {
		if countMap[v] == 0 {
			return false
		}
		countMap[v]--
	}
	return true
}
