package test

import (
	"DashboardBackend/mail"
	"testing"
	"fmt"
)

func TestMails(t *testing.T) {
	fmt.Println("======================sending======================")
	mail.Test("info@fredliang.cn")
	fmt.Println("========================done=======================")
}

