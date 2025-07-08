package store

import (
	"github.com/akshaykumar-tech/school-management/model"
	"github.com/akshaykumar-tech/school-management/db"
)

type UserStore struct {
	DB *db.Postgres
}

func NewUserStore(db *db.Postgres) *UserStore {
	return &UserStore{DB: db}
}

func (s *UserStore) CreateUser(user *model.User) error {
	return s.DB.DB.Create(user).Error
}

func (s *UserStore) GetUser(id uint) (*model.User, error) {
	var user model.User
	err := s.DB.DB.First(&user, id).Error
	return &user, err
}

func (s *UserStore) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := s.DB.DB.Find(&users).Error
	return users, err
}

func (s *UserStore) UpdateUser(user *model.User) error {
	return s.DB.DB.Save(user).Error
}

func (s *UserStore) DeleteUser(id uint) error {
	return s.DB.DB.Delete(&model.User{}, id).Error
} 