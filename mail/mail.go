package mail

import (
	"DashboardBackend/utils"
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

// Test is func to test smtp mails is reachable or not
func Test(addr string) {
	send("Test Sender", "Test  Subject", "info@fredliang.cn", "info@fredliang.cn", "Testmail")
}

// send func play a key role in mail pkg
func send(FromName, Subject, ToName, To, Content string) {
	m := gomail.NewMessage()
	m.SetHeader("From", FromName+" <"+utils.AppConfig.Mail.Username+">")
	m.SetHeader("Subject", Subject)
	m.SetAddressHeader("Cc", To, ToName)
	m.SetBody("text/html", Content)

	d := gomail.NewDialer(utils.AppConfig.Mail.Server, utils.AppConfig.Mail.Port, utils.AppConfig.Mail.Username, utils.AppConfig.Mail.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the test mail
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
