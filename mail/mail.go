package mail

import (
	"DashboardBackend/utils"
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

// Test is func to test smtp mails is reachable or not
func Test(addr string) {

	m := gomail.NewMessage()
	m.SetHeader("From", "Test Mail Sender <"+utils.AppConfig.Mail.Username+">")
	m.SetHeader("Subject", "Test Mail")
	m.SetAddressHeader("Cc", addr, "Test Mail Receiver")
	m.SetBody("text/html", "Test Mail!")

	d := gomail.NewDialer(utils.AppConfig.Mail.Server, utils.AppConfig.Mail.Port, utils.AppConfig.Mail.Username, utils.AppConfig.Mail.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the test mail
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
