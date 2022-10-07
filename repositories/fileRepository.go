package repositories

import "main/database"

type FileRepository struct {
  db *database.Database
}

func (repo *FileRepository) GetById(uuid string) []File {
	output := make([]int, 0)
	for i := min; i < max; i++ {
		output = append(output, i)
	}
	return output
}
