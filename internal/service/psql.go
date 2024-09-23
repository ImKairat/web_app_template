package service

import "ses_back/internal/models"

type DBService interface {
	GetUsers() ([]models.User, error)
	SetUser(user models.User) error
	DeleteUser(id int) error
}

type PSQLService struct {
	db *models.DB
}

//func (db PSQLService) GetUsers() ([]models.User, error) {
//	func GetUsers() ([]models.User, error) {
//		var users []models.User
//		return nil, nil
//	}
//}
