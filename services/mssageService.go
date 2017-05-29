package services

import "gopkg.in/gomail.v2"

type MessageService struct {

}

func (this MessageService) SendEmail(from, to, title, body, contentType string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", title)
	m.SetBody(contentType, body)

	d := gomail.NewDialer("smtp.elasticemail.com", 587, "ademuanthony@gmail.com", "dda405b5-4b2b-42ec-bd42-ce791665da77")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
		return err
	}
	return nil
}
