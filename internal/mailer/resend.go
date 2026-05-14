package mailer

import (
	"fmt"
	"log"
	"time"

	"github.com/resend/resend-go/v3"
)

type ResendMailer struct {
	fromEmail string
	apiKey string
	client *resend.Client
}

func NewResendMailer(apiKey, fromEmail string) *ResendMailer {
	client := resend.NewClient(apiKey)

	return &ResendMailer{
		fromEmail: fromEmail,
		apiKey: apiKey,
		client: client,
	}
}

func (m *ResendMailer) Send(username, email, activationURL string) error {

	from := m.fromEmail
	subject := fmt.Sprintf("Activation link for %s", FromName)

	params := &resend.SendEmailRequest{
		From: from,
		To: []string{email},
		Subject: subject,
		Html: `
			<html>
				<body>
				<h3>Hi ` + username + `</h3>
				<p>Please proceed to following URL to verify your account:</p><br>
				<a href = "`+ activationURL +`">
			`+ activationURL + 
			`
				</a></body>
			</html>
		`,
	}

	for i := 0; i < MaxRetries; i++ {
		response, err := m.client.Emails.Send(params)
		if err != nil {
			time.Sleep(time.Second * time.Duration(i + 1))
			continue
		}

		log.Printf("Email send with response id: %v", response.Id)
		return nil
	}

	return nil

}