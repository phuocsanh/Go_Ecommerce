package repo

import (
	"fmt"
	"go_ecommerce/global"
	"go_ecommerce/internal/model"
)

// type UserRepo struct {}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetInfoUser()string {
// 	return "sanh 123"
// }

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
}

// GetUserByEmail implements IUserRepository.
func (u *userRepository) GetUserByEmail(email string) bool {
	row := global.Mdb.Table(TableNameGoCrmUser).Where("user_email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	fmt.Print("row", row)
	return row != NumberNull
}

func NewUserRepository() IUserRepository { return &userRepository{} }
