package response

const (
	ErrCodeSuccess =20001 // Success
	ErrParamInvalid =20003 // Email invalid
)
var msg = map[int]string{
	ErrCodeSuccess:		"success",
	ErrParamInvalid:	"Email is invalid", 
}