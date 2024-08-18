package engage

import (
	"log"
	"fmt"
	"net/smtp"
	. "github.com/rohit123sinha456/payredapp/config"
)

// // ConfigStruct contains configuration for sending emails
// type ConfigStruct struct {
// 	From string
// 	Pass string
// }

// sendEmail sends an email using the provided configuration and message body
func SendEmail(body ConfigStruct, to []string, subject, msgBody string) {
	from := body.From
	pass := body.Pass
	fmt.Printf("%v", to)
	msg := "From: " + from + "\n" +
		"To: " + fmt.Sprintf("%v", to) + "\n" +
		"Subject: " + subject + "\n\n" +
		msgBody

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, to, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	log.Println("Successfully sent to " + fmt.Sprintf("%v", to))
}
