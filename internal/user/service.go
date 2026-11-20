package user

import (
	"errors"

	"github.com/SifatU360/GoTix-.git/internal/user/dto"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateUser(req dto.CreateRequest) (*dto.Response,error) {
	user := User{
		Name:     req.Name,
		Email:    req.Email,
	}

	user.hashPassword(req.Password)

	err := s.repo.CreateUser(&user)
	if err != nil {
		return nil, err
	}
	response := dto.Response{
		ID:		user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
	}

	return &response, nil
}


func (s *service) LoginUser(req dto.LoginRequest) (*dto.Response,error) {
	user , err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	if !user.checkPassword(req.Password) {
		return nil, errors.New("invalid password")
	}
	response := dto.Response{
		ID:		user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
	}
	return &response, nil
}