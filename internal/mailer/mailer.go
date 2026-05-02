package mailer

const (
	FromName = "Gopher Socials"
	MaxRetries = 3
)

type Client interface {
	Send(username, email, activationURL string) error
}