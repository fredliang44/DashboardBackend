package mail

import (
	"DashboardBackend/utils"
	"bytes"
	"fmt"
	"net/mail"
	"sync"

	"github.com/go-mailer/send"
)

// Test is func to test smtp mails is reachable or not
func Test(addr string) {
	from := mail.Address{Name: "Test Mail", Address: utils.AppConfig.Mail.Username}
	sender, err := send.NewSmtpSender(utils.AppConfig.Mail.Server+":"+string(utils.AppConfig.Mail.Port), from, utils.AppConfig.Mail.Password)
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	msg := &send.Message{
		Subject: "异步发送邮件测试",
		Content: bytes.NewBufferString("<h1>你好，异步测试邮件内容</h1>"),
		To:      []string{addr},
	}
	err = sender.AsyncSend(msg, false, func(err error) {
		defer wg.Done()
		if err != nil {
			fmt.Println("发送邮件出现错误：", err)
		}
	})
	if err != nil {
		panic(err)
	}
	wg.Wait()
	fmt.Println("邮件发送完成")
}
