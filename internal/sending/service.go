package sending

import (
	"context"
	"fmt"
	"net/smtp"
	"promotion/pkg/logger"
)

// ------------------- Core Abstraction for DIP -------------------

type NotificationSender interface {
	Send(ctx context.Context, message string) error
}

type EmailSender struct {
	SMTPHost string
	SMTPPort string
	Username string
	Password string
	From     string
}

func (e *EmailSender) Send(ctx context.Context, email, message string) error {
	to := []string{email}
	msg := []byte("Subject: Notification from GoTutor\n\n" + message)

	auth := smtp.PlainAuth("", e.Username, e.Password, e.SMTPHost)

	err := smtp.SendMail(
		e.SMTPHost+":"+e.SMTPPort,
		auth,
		e.From,
		to,
		msg,
	)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Println("✅ Email sent successfully!")
	return nil
}

type SMSSender struct {
	FromPhone string
	APIKey    string
}

func (s *SMSSender) Send(ctx context.Context, phone, message string) error {
	twilio.SendSMS(s.FromPhone, phone, message)
	fmt.Println("📱 Sending SMS to the phone:", message)
	return nil
}

type Service struct {
	log    *logger.Logger
	repo   *Repo
	sender NotificationSender // <- DIP: depends on abstraction
}

func NewService(log *logger.Logger, repo *Repo, sender NotificationSender) *Service {
	return &Service{log: log, repo: repo, sender: sender}
}
