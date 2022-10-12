package repositories

import (
	database "main/database"
	models "main/models"
)

type FileRepository struct {
	db *database.Database
}

func (repo *FileRepository) GetById(uuid string) []models.File {
	output := make([]models.File, 0)
	// for i := min; i < max; i++ {
	// 	output = append(output, i)
	// }
	return output
}
