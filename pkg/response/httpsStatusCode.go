package response

const (
	ErrCodeSuccess =20001 // Success
	ErrParamInvalid =20003 // Email invalid
	ErrInvalidToken = 30001 // Token is Invalid
	ErrInvalidOtp = 30002
	ErrSendEmailOtp = 30003
	// Register code
	ErrCodeUserHassExits = 50001 // User has already been assigned
)
var msg = map[int]string{
	ErrCodeSuccess:			"success",
	ErrParamInvalid:		"Email is invalid", 
	ErrInvalidToken:    	"Token is invalid",
	ErrCodeUserHassExits: 	"User has already been assigned",
	ErrInvalidOtp: 			"Otp is error",
	ErrSendEmailOtp: 		"Send email failed",
}