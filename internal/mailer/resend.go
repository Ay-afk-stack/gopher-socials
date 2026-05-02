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

	from := fmt.Sprintf("%s", m.fromEmail)
	subject := fmt.Sprintf("Activation link for %s", FromName)

	params := &resend.SendEmailRequest{
		From: from,
		To: []string{email},
		Subject: subject,
		Html: `
			<html>
				<body>
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
			log.Printf("Failed to send email to %v, attempt of %d of %d", email, i + 1, MaxRetries)
			log.Print(err)

			time.Sleep(time.Second * time.Duration(i + 1))
			continue
		}

		log.Printf("Email send with response id: %v", response.Id)
		return nil
	}

	return nil

}