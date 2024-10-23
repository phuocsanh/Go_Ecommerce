package service

import (
	"go_ecommerce/internal/repo"
	"go_ecommerce/pkg/response"
)



type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository // Dùng interface không cần dùng *(con trỏ) vì nó đã là con trỏ rồi
}


func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
		}
}

// Register implements IUserService.
func (u *userService) Register(email string, purpose string) int {
	// Check email exist
	if(u.userRepo.GetUserByEmail(email)){
		return response.ErrCodeUserHassExits
	}
	return response.ErrCodeSuccess
}