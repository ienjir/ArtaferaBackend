package utils

import (
	"fmt"
	"net/smtp"
	"os"
	"strconv"
	"strings"
)

func SendContactNotification(fromEmail, message string) error {
	host := os.Getenv("SMTP_HOST")
	portStr := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASSWORD")
	to := os.Getenv("SMTP_TO")

	if host == "" || portStr == "" || user == "" || password == "" || to == "" {
		return fmt.Errorf("SMTP configuration incomplete, skipping email notification")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return fmt.Errorf("invalid SMTP_PORT: %w", err)
	}

	subject := "Neue Kontaktanfrage von " + fromEmail
	body := buildContactEmailBody(fromEmail, message)

	msg := strings.Join([]string{
		"From: " + user,
		"To: " + to,
		"Subject: " + subject,
		"MIME-Version: 1.0",
		"Content-Type: text/plain; charset=UTF-8",
		"",
		body,
	}, "\r\n")

	addr := fmt.Sprintf("%s:%d", host, port)
	auth := smtp.PlainAuth("", user, password, host)

	return smtp.SendMail(addr, auth, user, []string{to}, []byte(msg))
}

func buildContactEmailBody(fromEmail, message string) string {
	return fmt.Sprintf(
		"Eine neue Nachricht wurde über das Kontaktformular auf ArtaFera gesendet.\r\n\r\nVon: %s\r\n\r\nNachricht:\r\n%s",
		fromEmail,
		message,
	)
}
