package impl

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"fmt"
	"go_ecommerce/global"
	"go_ecommerce/internal/consts"
	"go_ecommerce/internal/database"
	"go_ecommerce/internal/model"
	"go_ecommerce/internal/utils"
	"go_ecommerce/internal/utils/auth"
	"go_ecommerce/internal/utils/crypto"
	"go_ecommerce/internal/utils/random"
	"go_ecommerce/internal/utils/sendto"
	"go_ecommerce/pkg/response"

	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"
)

type sUserLogin struct {
	r *database.Queries
}

func NewUserLoginImpl(r *database.Queries) *sUserLogin {
	return &sUserLogin{
		r: r,
	}
}

// ---- TWO FACTOR AUTHEN -----

// two-factor authentication
func (s *sUserLogin) IsTwoFactorEnabled(ctx context.Context, userId int) (codeResult int, rs bool, err error) {
	return 200, true, nil
}

func (s *sUserLogin) SetupTwoFactorAuth(ctx context.Context, in *model.SetupTwoFactorAuthInput) (codeResult int, err error) {
	// Logic
	// 1. Check isTwoFactorEnabled -> true return
	isTwoFactorAuth, err := s.r.IsTwoFactorEnabled(ctx, in.UserId)
	if err != nil {
		return response.ErrCodeTwoFactorAuthSetupFailed, err
	}
	if isTwoFactorAuth > 0 {
		return response.ErrCodeTwoFactorAuthSetupFailed, fmt.Errorf("Two-factor authentication is already enabled")
	}
	// 2. crate new type Authe
	err = s.r.EnableTwoFactorTypeEmail(ctx, database.EnableTwoFactorTypeEmailParams{
		UserID:            in.UserId,
		TwoFactorAuthType: database.PreGoAccUserTwoFactor9999TwoFactorAuthTypeEMAIL,
		TwoFactorEmail:    sql.NullString{String: in.TwoFactorEmail, Valid: true},
	})
	if err != nil {
		return response.ErrCodeTwoFactorAuthSetupFailed, err
	}

	// 3. send otp to in.TwoFactorEmail
	keyUserTwoFator := crypto.GetHash("2fa:" + strconv.Itoa(int(in.UserId)))
	go global.Rdb.Set(ctx, keyUserTwoFator, "123456", time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err()
	// if err != nil {
	// 	return response.ErrCodeTwoFactorAuthSetupFailed, err
	// }
	return response.CodeSuccess, nil
}

// Verify Two Factor Authentication
func (s *sUserLogin) VerifyTwoFactorAuth(ctx context.Context, in *model.TwoFactorVerificationInput) (codeResult int, err error) {
	// 1. Check isTwoFactorEnabled
	isTwoFatorAuth, err := s.r.IsTwoFactorEnabled(ctx, in.UserId)
	if err != nil {
		return response.ErrCodeTwoFactorAuthVerifyFailed, err
	}
	if isTwoFatorAuth > 0 {
		return response.ErrCodeTwoFactorAuthVerifyFailed, fmt.Errorf("Two-factor authentication is not enabled")
	}

	// 2. Check Otp in redis avaible
	keyUserTwoFator := crypto.GetHash("2fa:" + strconv.Itoa(int(in.UserId)))
	otpVerifyAuth, err := global.Rdb.Get(ctx, keyUserTwoFator).Result()
	if err == redis.Nil {
		return response.ErrCodeTwoFactorAuthVerifyFailed, fmt.Errorf("Key %s does not exists", keyUserTwoFator)
	} else if err != nil {
		return response.ErrCodeTwoFactorAuthVerifyFailed, err
	}
	// 3. check otp
	if otpVerifyAuth != in.TwoFactorCode {
		return response.ErrCodeTwoFactorAuthVerifyFailed, fmt.Errorf("OTP does not match")
	}

	// 4. udpoate status
	err = s.r.UpdateTwoFactorStatus(ctx, database.UpdateTwoFactorStatusParams{
		UserID:            in.UserId,
		TwoFactorAuthType: database.PreGoAccUserTwoFactor9999TwoFactorAuthTypeEMAIL,
	})
	if err != nil {
		return response.ErrCodeTwoFactorAuthVerifyFailed, err
	}
	// 5. remove otp
	_, err = global.Rdb.Del(ctx, keyUserTwoFator).Result()
	if err != nil {
		return response.ErrCodeTwoFactorAuthVerifyFailed, err
	}
	return 200, nil
}

// ---- END TWO FACTOR AUTHEN ----

func (s *sUserLogin) Login(ctx context.Context, in *model.LoginInput) (codeResult int, out model.LoginOutput, err error) {

	userBase, err := s.r.GetOneUserInfo(ctx, in.UserAccount)
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}
	// 2. check password?
	if !crypto.MatchingPassword(userBase.UserPassword, in.UserPassword, userBase.UserSalt) {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("does not match password")
	}

	// 3. check two-factor authentication

	isTwoFactorEnable, err := s.r.IsTwoFactorEnabled(ctx, uint32(userBase.UserID))
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("does not match password")
	}
	if isTwoFactorEnable > 0 {

		// otpNew := random.GenerateSixDigiOtp()
		// if __DEV__ == "TEST_USER" {
		// 	otpNew = 123456
		// }
		// sen otp to in.TwoFactorEmail
		keyUserLoginTwoFactor := crypto.GetHash("2fa:otp:" + strconv.Itoa(int(userBase.UserID)))
		err = global.Rdb.SetEx(ctx, keyUserLoginTwoFactor, "111111", time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err()
		if err != nil {
			return response.ErrCodeAuthFailed, out, fmt.Errorf("set otp redis faiuled")
		}
		// send otp via twofactorEmail
		// get email 2fa
		infoUserTwoFactor, err := s.r.GetTwoFactorMethodByIDAndType(ctx, database.GetTwoFactorMethodByIDAndTypeParams{
			UserID:            uint32(userBase.UserID),
			TwoFactorAuthType: database.PreGoAccUserTwoFactor9999TwoFactorAuthTypeEMAIL,
		})
		if err != nil {
			return response.ErrCodeAuthFailed, out, fmt.Errorf("get two factor method failed")
		}
		// go sendto.SendEmailToJavaByAPI()
		log.Println("send OTP 2FA to Email::", infoUserTwoFactor.TwoFactorEmail)
		go sendto.SendTextEmailOtp([]string{infoUserTwoFactor.TwoFactorEmail.String}, os.Getenv("SENDER_EMAIL"), "111111")

		out.Message = "send OTP 2FA to Email, pls het OTP by Email.."
		return response.CodeSuccess, out, nil
	}

	// 4. update password time
	go s.r.LoginUserBase(ctx, database.LoginUserBaseParams{
		UserLoginIp:  sql.NullString{String: "127.0.0.1", Valid: true},
		UserAccount:  in.UserAccount,
		UserPassword: in.UserPassword, // khong can
	})

	// 5. Create UUID User
	subToken := utils.GenerateCliTokenUUID(int(userBase.UserID))
	log.Println("subtoken:", subToken)
	// 6. get user_info table
	infoUser, err := s.r.GetUser(ctx, uint64(userBase.UserID))
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}
	// convert to json
	infoUserJson, err := json.Marshal(infoUser)
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("convert to json failed: %v", err)
	}
	// 7. give infoUserJson to redis with key = subToken
	err = global.Rdb.Set(ctx, subToken, infoUserJson, time.Duration(consts.TIME_2FA_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}
	// 8. create token
	out.AccessToken, err = auth.CreateToken(subToken, "30m")
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}
	// Tạo refreshToken
	refreshToken, err := auth.CreateToken(subToken, "1h")
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}
	out.RefreshToken = refreshToken

	err = global.Rdb.Set(ctx, subToken+"_refresh", refreshToken, time.Duration(consts.TIME_REFRESH_TOKEN)*time.Hour).Err()
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}
	return 200, out, nil
}

func (s *sUserLogin) RefreshToken(ctx context.Context, refreshToken string) (out model.LoginOutput, err error) {
	// Giải mã refreshToken
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.JWT.API_SECRET_KEY), nil
	})

	if err != nil || !token.Valid {
		return out, fmt.Errorf("refreshToken không hợp lệ")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["sub"] == nil {
		return out, fmt.Errorf("refreshToken không hợp lệ")
	}

	subToken := claims["sub"].(string)

	// Kiểm tra refreshToken trong Redis
	savedRefreshToken, err := global.Rdb.Get(ctx, subToken+"_refresh").Result()
	if err != nil || savedRefreshToken != refreshToken {
		return out, fmt.Errorf("refreshToken không hợp lệ hoặc đã hết hạn")
	}

	// Tạo accessToken và refreshToken mới
	newAccessToken, err := auth.CreateToken(subToken, "30m")
	if err != nil {
		return out, err
	}

	err = global.Rdb.Set(ctx, subToken, newAccessToken, time.Duration(consts.TIME_2FA_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		return out, err
	}

	newRefreshToken, err := auth.CreateToken(subToken, "1h")
	if err != nil {
		return out, err
	}

	// Lưu refreshToken mới vào Redis
	err = global.Rdb.Set(ctx, subToken+"_refresh", newRefreshToken, time.Duration(consts.TIME_REFRESH_TOKEN)*time.Hour).Err()
	if err != nil {
		return out, err
	}

	out.AccessToken = newAccessToken
	out.RefreshToken = newRefreshToken
	return out, nil
}

func (s *sUserLogin) Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error) {
	//1. hash email
	fmt.Printf("Verify key %s\n", in.VerifyKey)
	fmt.Printf("Verify type %d\n", in.VerifyType)

	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))
	fmt.Printf("hasKey %s\n", hashKey)

	//2. check if existing user in database
	userFound, err := s.r.CheckUserBaseExists(ctx, in.VerifyKey)
	if err != nil {
		return response.ErrCodeUserHassExits, err
	}
	if userFound > 0 {
		return response.ErrCodeUserHassExits, fmt.Errorf("user already registered")
	}

	//3. Create OTP
	userKey := utils.GetUserKey(hashKey)
	otpFound, err := global.Rdb.Get(ctx, userKey).Result()

	switch {
	case err == redis.Nil:
		fmt.Println("Key does not exist")
	case err != nil:
		fmt.Println("Get fail")
		return response.ErrInvalidOtp, err
	case otpFound != "":
		return response.ErrInvalidOtp, fmt.Errorf("")
	}
	//4. Generate OTP
	otpNew := random.GenerateSixDigiOtp()
	if in.VerifyPurpose == "TEST_USER" {
		otpNew = 123456
	}
	fmt.Printf("Otp is %d\n", otpNew)
	//5. Save OTP in redis with expiration time
	err = global.Rdb.SetEx(ctx, userKey, strconv.Itoa(otpNew), time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err()

	if err != nil {
		return response.ErrInvalidOtp, err
	}

	//6. Send email otp
	switch in.VerifyType {
	case consts.EMAIL:
		err = sendto.SendTextEmailOtp([]string{in.VerifyKey}, os.Getenv("SENDER_EMAIL"), strconv.Itoa(otpNew))
		if err != nil {
			return response.ErrSendEmailOtp, err
		}
		//7. Save otp to MYSQL
		result, err := s.r.InsertOTPVerify(ctx, database.InsertOTPVerifyParams{
			VerifyOtp:     strconv.Itoa(otpNew),
			VerifyKey:     in.VerifyKey,
			VerifyKeyHash: hashKey,
			VerifyType:    sql.NullInt32{Int32: 1, Valid: true},
		})
		if err != nil {
			return response.ErrSendEmailOtp, err
		}
		//8 getLastID
		lastIdVerifyUser, err := result.LastInsertId()

		if err != nil {
			return response.ErrSendEmailOtp, err
		}
		log.Println("lastIdVerifyUser", lastIdVerifyUser)

		return response.CodeSuccess, nil

	case consts.MOBILE:
		return response.CodeSuccess, nil

	}
	return response.CodeFail, nil
}
func (s *sUserLogin) VerifyOTP(ctx context.Context, in *model.VerifyInput) (out model.VerifyOTPOutput, err error) {

	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))

	//Get otp
	otpFound, err := global.Rdb.Get(ctx, utils.GetUserKey((hashKey))).Result()

	if err != nil {
		return out, err
	}

	if in.VerifyCode != otpFound {
		// Nếu sai 3 lần 1 phủt chưa xử lý
		return out, fmt.Errorf("OTP not match")
	}
	infoOTP, err := s.r.GetInfoOTP(ctx, hashKey)

	if err != nil {
		return out, err
	}

	err = s.r.UpdateUserVerificationStatus(ctx, hashKey)
	if err != nil {
		return out, err
	}

	out.Token = infoOTP.VerifyKeyHash
	out.Message = "Succes"
	return out, err
}
func (s *sUserLogin) UpdatePasswordRegister(ctx context.Context, token string, password string) (userId int, err error) {
	//Token is already verified
	infoOTP, err := s.r.GetInfoOTP(ctx, token)
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}
	// Check isVerified OK
	if infoOTP.IsVerified.Int32 == 0 {
		return response.ErrCodeUserOtpNotExists, fmt.Errorf("User OTP not verified")
	}
	// update user_base

	userBase := database.AddUserBaseParams{}
	userBase.UserAccount = infoOTP.VerifyKey
	userSalt, err := crypto.GenerateSalt(16)

	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}
	userBase.UserSalt = userSalt
	userBase.UserPassword = crypto.HashPassword(password, userSalt)

	newUserBase, err := s.r.AddUserBase(ctx, userBase)

	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}

	user_id, err := newUserBase.LastInsertId()

	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}

	// add user_id to user_info table
	newUserInfo, err := s.r.AddUserHaveUserId(ctx, database.AddUserHaveUserIdParams{
		UserID:               uint64(user_id),
		UserAccount:          infoOTP.VerifyKey,
		UserNickname:         sql.NullString{String: infoOTP.VerifyKey, Valid: true},
		UserAvatar:           sql.NullString{String: "", Valid: true},
		UserState:            1,
		UserMobile:           sql.NullString{String: "", Valid: true},
		UserGender:           sql.NullInt16{Int16: 0, Valid: true},
		UserBirthday:         sql.NullTime{Time: time.Time{}, Valid: false},
		UserEmail:            sql.NullString{String: infoOTP.VerifyKey, Valid: true},
		UserIsAuthentication: 1,
	})

	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}

	user_id, err = newUserInfo.LastInsertId()

	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}

	return int(user_id), nil
}
