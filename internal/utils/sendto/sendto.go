package sendto

import (
	"crypto/tls"
	"fmt"
	"go_ecommerce/global"
	"net/smtp"
	"strings"

	"go.uber.org/zap"
)

const (
	SMTPHost     = "smtp.gmail.com"
	SMTPPort     = "587"
	SMTPName     = "phuocsanh61688@gmail.com"
	SMTPPassword = "0919461688Tp$"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

func BuilderMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)
	return msg
}

func SendTextEmailOtp(to []string, from string, otp string) error {
	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: "test"},
		To:      to,
		Subject: "Otp verification",
		Body:    fmt.Sprintf("Your OTP is %s. Please enter it to verify your account.", otp),
	}
	messageMail := BuilderMessage(contentEmail)

	auth := smtp.PlainAuth("", SMTPName, SMTPPassword, SMTPHost)

	// Cấu hình TLS tùy chỉnh
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,       // Bỏ qua xác thực chứng chỉ
		ServerName:         SMTPHost,   // Đặt ServerName để khớp với chứng chỉ
	}

	// Tạo kết nối TLS tới SMTP server
	conn, err := tls.Dial("tcp", SMTPHost+":"+SMTPPort, tlsConfig)
	if err != nil {
		global.Logger.Error("Failed to connect to SMTP server", zap.Error(err))
		return err
	}

	client, err := smtp.NewClient(conn, SMTPHost)
	if err != nil {
		global.Logger.Error("Failed to create SMTP client", zap.Error(err))
		return err
	}

	// Xác thực
	if err = client.Auth(auth); err != nil {
		global.Logger.Error("SMTP authentication failed", zap.Error(err))
		return err
	}

	// Cài đặt người gửi và người nhận
	if err = client.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = client.Rcpt(addr); err != nil {
			return err
		}
	}

	// Gửi nội dung email
	w, err := client.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(messageMail))
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}

	client.Quit()

	fmt.Print("Email sent successfully")
	return nil
}
