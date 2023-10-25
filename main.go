package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/keighl/postmark"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	log.Println("Hello")
	SERVER_TOKEN := os.Getenv("SERVER_TOKEN")
	ACCOUNTT_TOKEN := os.Getenv("ACCOUNTT_TOKEN")
	client := postmark.NewClient(SERVER_TOKEN, ACCOUNTT_TOKEN)
	// Create a new email message.
	message := postmark.Email{
		From:     "no-reply@gricdserver.com",
		To:       "nimi@figorr.com",
		Subject:  "Hello from Postmark!",
		TextBody: "This is just a friendly 'hello' from your friends at Postmark.",
	}

	// Send the email message.
	sendResult, err := client.SendEmail(message)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error sending email:", err)

		return
	}


	jsonData, err := json.Marshal(sendResult)
	jsonText := string(jsonData)

	log.Printf("The email was sent successfully! %s\n", jsonText)
}
