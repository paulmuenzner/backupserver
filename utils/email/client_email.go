package email

import "gopkg.in/gomail.v2"

type ClientConfigData struct {
	Host         string
	SmtpUsername string
	SmtpPassword string
	SmtpPort     int
}

func EmailClient(configData *ClientConfigData) (client *gomail.Dialer) {
	return gomail.NewDialer(configData.Host, configData.SmtpPort, configData.SmtpUsername, configData.SmtpPassword)
}
