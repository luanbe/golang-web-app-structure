package repository

import "github.com/luanbe/golang-web-app-structure/app/models/entity"

type UserRepository interface {
	AddUser(user *entity.User) (*entity.User, error)
}

type UserRepositoryImpl struct {
	base BaseRepository
}

func NewUserRepository(br BaseRepository) UserRepository {
	return &UserRepositoryImpl{br}
}

func (r *UserRepositoryImpl) AddUser(User *entity.User) (*entity.User, error) {
	err := r.base.GetDB().Create(User).Error
	if err != nil {
		return nil, err
	}
	return User, nil
}
