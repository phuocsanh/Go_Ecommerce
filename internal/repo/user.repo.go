package repo

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
	return true
}

func NewUserRepository() IUserRepository { return &userRepository{} }
