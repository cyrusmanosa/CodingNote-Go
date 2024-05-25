package main

// SendGrid woodygame123456@gmail.com Cyrusmkc19941207

import (
	// "bytes"
	"fmt"
	"net/smtp"
	// "text/template"
	// "gopkg.in/gomail.v2"
)

func sendMailSimple(subject, body string, to []string) {
	// 設定 account
	auth := smtp.PlainAuth("", "studiocmkc0110@gmail.com", "iodvpvmlyvadnhfb", "smtp.gmail.com")
	msg := "Subject: " + subject + "\n" + body
	err := smtp.SendMail("smtp.gmail.com:587", auth, "studiocmkc0110@gmail.com", to, []byte(msg))
	if err != nil {
		fmt.Println(err)
	}
}

// func sendMailSimpleHTML(subject, templatePath string, to []string) {
// 	// Get HTML
// 	var body bytes.Buffer
// 	t, err := template.ParseFiles(templatePath)
// 	t.Execute(&body, struct{ Name string }{Name: "Cyrus"})

// 	// 設定 account
// 	auth := smtp.PlainAuth(
// 		"",
// 		"studiocmkc0110@gmail.com",
// 		"iodvpvmlyvadnhfb",
// 		"smtp.gmail.com",
// 	)

// 	//  大細草要跟返
// 	headers := "MIME-version: 1.0; \nContent-Type: text/html; charset=\"UTF-8\";"
// 	msg := "Subject: " + subject + "\n" + headers + "\n\n" + body.String()
// 	err = smtp.SendMail("smtp.gmail.com:587", auth, "studiocmkc0110@gmail.com", to, []byte(msg))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// // SendGo plugin
// func sendGoMail(templatePath string) {
// 	// Get HTML
// 	var body bytes.Buffer
// 	t, err := template.ParseFiles(templatePath)
// 	t.Execute(&body, struct{ Name string }{Name: "Cyrus"})
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// Send Go Mail
// 	m := gomail.NewMessage()
// 	m.SetHeader("From", "studiocmkc0110@gmail.com")
// 	m.SetHeader("To", "studiocmkc0110@gmail.com")
// 	m.SetHeader("Subject", "Hello!")
// 	m.SetBody("text/html", body.String())
// 	m.Attach("./456.jpeg")

// 	d := gomail.NewDialer("smtp.gmail.com", 587, "studiocmkc0110@gmail.com", "iodvpvmlyvadnhfb")

// 	// Send the email to Bob, Cora and Dan.
// 	if err := d.DialAndSend(m); err != nil {
// 		panic(err)
// 	}
// }

func main() {
	sendMailSimple(
		"Another subject123",
		"Another body",
		[]string{"studiocmkc0110@gmail.com"},
	)

	// sendMailSimpleHTML(
	// 	"Another subject",
	// 	"./test.html",
	// 	[]string{"woodygame123456@gmail.com"},
	// )

	// sendGoMail("./test.html")
}
