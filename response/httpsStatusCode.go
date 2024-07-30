package response

const (
	ErrCodeSuccess =20001 // Success
	ErrParamInvalid =20003 // Email invalid
	ErrInvalidToken = 30001 // Token is Invalid
)
var msg = map[int]string{
	ErrCodeSuccess:		"success",
	ErrParamInvalid:	"Email is invalid", 
	ErrInvalidToken:    "Token is invalid",
}