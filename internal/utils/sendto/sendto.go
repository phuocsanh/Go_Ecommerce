package sendto

import (
	"fmt"
	"go_ecommerce/global"
	"net/smtp"
	"strings"

	"go.uber.org/zap"
)
const (
	SMTPHost = "smtp.gmail.com"
    SMTPPort = "587" 
	SMTPName = "phuocsanh61688@gmail.com"
	SMTPPassword = ""

)
type EmailAddress struct {
	Address string `json:"address"`
	Name string `json:"name"`
}
type Mail struct {
	From EmailAddress
	To []string
	Subject string
	Body string
}


 func BuilderMessage(mail Mail)string  {
	
	msg:= "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n",strings.Join(mail.To,";"))
	msg += fmt.Sprintf("Subject: %s\r\n",mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n",mail.Body)
	return msg
 }
func SendTextEmailOtp(to []string, from string, otp string) error  {
	contentEmail := Mail{
		From: EmailAddress{Address: from, Name: "test"},
		To:to,
		Subject: "Otp verification",
		Body: fmt.Sprintf("Your OTP is %s. Please enter it to  verify your account.", otp),
	}
	messageMail := BuilderMessage(contentEmail)
	
	authention := smtp.PlainAuth("",SMTPName,SMTPPassword,SMTPHost)

	err := smtp.SendMail(SMTPHost+":"+SMTPPort,authention, from,to,[]byte(messageMail))
	if err != nil {
		global.Logger.Error("Email send failed", zap.Error(err))
		return err
	}
	return nil
}