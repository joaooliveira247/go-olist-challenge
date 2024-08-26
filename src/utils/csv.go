package utils

import (
	"encoding/csv"
	"os"
)

type AuthorCSV struct {
	Name string
}

func CSVToAuthor(path string, header bool) ([]AuthorCSV, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	if !header {
		lines = lines[1:]
	}

	var authors []AuthorCSV

	for _, line := range lines {
		author := AuthorCSV{Name: line[0]}
		authors = append(authors, author)
	}

	return authors, nil

}
