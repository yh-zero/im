package test

import (
	"fmt"
	"os"
	"testing"
)

func TestOS(t *testing.T) {
	var MailPassword = os.Getenv("MailPassword")
	fmt.Println("MailPassword ========> ", MailPassword)
	fmt.Println("MailPassword ========> ", MailPassword)
}
