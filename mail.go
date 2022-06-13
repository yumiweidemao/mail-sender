package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
)

func main() {

	// Sender data.
	from := "csh1234_1@126.com"
	password := "guess_me"				// enter password here

	// Check command line arguments
	if len(os.Args) == 1 {
		fmt.Println("No valid email address provided. Abort.")
		return
	}

	// Receiver email address.
	to := os.Args[1:]

	// Read email subject & text
	textFile, err := ioutil.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}

	// search for the first newline character
	var i int
	for i = 0; textFile[i] != byte('\n'); i++ {
	}

	subject := string(textFile[:i])
	text := string(textFile[i+1:])

	fmt.Println("subject:", subject)
	fmt.Println("text:\n", text)

	// smtp server configuration.
	smtpHost := "smtp.126.com"
	smtpPort := "25"

	for _, to_addr := range to {
		// Message.
		message := []byte("From: csh1234_1@126.com\r\n" +
			"To: " + to_addr + "\r\n" +
			"Subject: " + subject + "\r\n\r\n" +
			text + "\r\n")

		// Authentication.
		auth := smtp.PlainAuth("", from, password, smtpHost)

		// Sending email.
		err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Email Sent Successfully to " + to_addr + "!")
	}
}
