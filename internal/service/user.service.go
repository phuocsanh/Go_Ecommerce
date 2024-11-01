package service

import (
	"fmt"
	"go_ecommerce/internal/repo"
	"go_ecommerce/internal/utils/cryto"
	"go_ecommerce/internal/utils/random"
	"go_ecommerce/internal/utils/sendto"
	"go_ecommerce/pkg/response"
	"os"
	"strconv"
	"time"
)



type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository // Dùng interface không cần dùng *(con trỏ) vì nó đã là con trỏ rồi
	userAuthRepo repo.IUserAuthRepository
}


func NewUserService(userRepo repo.IUserRepository, userAuthRepo repo.IUserAuthRepository) IUserService {
	return &userService{
		userRepo: userRepo,
		userAuthRepo: userAuthRepo,
		}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// 0. hash email
	hashEmail := cryto.GetHash(email)
	fmt.Println("HashEmail %d", hashEmail)
	// 1. check otp availability
	// 2. User spam 
	// 3. Check email exist on db
	if us.userRepo.GetUserByEmail(email){
		return response.ErrCodeUserHassExits
	}
	// 4. New otp ->
	otp:= random.GenerateSixDigiOtp() 

	if( purpose == "TEST_USER"){
		otp = 123456
	}
	fmt.Printf("Otp is %d\n", otp)

	// 5. Save otp in redis with expiration	time 
	err:= us.userAuthRepo.AddOtp(hashEmail, otp, int64(10 * time.Minute) )
	if( err != nil){
		return response.ErrInvalidOtp
	}
	// 6. Send email otp
	err = sendto.SendTextEmailOtp([]string{email},os.Getenv("SENDER_EMAIL"),strconv.Itoa(otp))
	
	if err != nil{
		return response.ErrSendEmailOtp
	}
	return response.ErrCodeSuccess
}