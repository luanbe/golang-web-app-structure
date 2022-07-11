package service

import (
	"github.com/luanbe/golang-web-app-structure/app/models/entity"
	"github.com/luanbe/golang-web-app-structure/app/repository"
)

type UserService interface {
	AddUser(username, email string) *entity.User
}

type UserServiceImpl struct {
	//logger      log.Logger
	baseRepo repository.BaseRepository
	userRepo repository.UserRepository
}

func NewUserService(
	//lg log.Logger,
	baseRepo repository.BaseRepository,
	userRepo repository.UserRepository,
) UserService {
	return &UserServiceImpl{baseRepo, userRepo}
}

func (s *UserServiceImpl) AddUser(username, email string) *entity.User {
	s.baseRepo.BeginTx()
	User := entity.User{
		Email:    email,
		UserName: username,
	}
	result, err := s.userRepo.AddUser(&User)
	if err != nil {
		return nil
	}
	return result
}
