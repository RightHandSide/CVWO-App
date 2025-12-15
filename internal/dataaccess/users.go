package dataaccess

import (
	"github.com/RightHandSide/CVWO-App/internal/database"
	"github.com/RightHandSide/CVWO-App/internal/models"
)

func List(db *database.Database) ([]models.User, error) {
	users := []models.User{
		{
			ID:   1,
			Name: "CVWO",
		},
	}
	return users, nil
}
