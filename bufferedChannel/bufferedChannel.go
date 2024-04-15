package main

import "fmt"

func addEmailsToQueue(emails []string) chan string {
	emailsToSend := make(chan string, len(emails))
	for _, email := range emails {
		emailsToSend <- email
	}
	return emailsToSend
}

func sendEmails(batchSize int, ch chan string) {
	for i := 0; i < batchSize; i++ {
		email := <-ch
		fmt.Println("Sending email:", email)
	}
}

func main() {
	ch := addEmailsToQueue([]string{"email1", "email2", "email3", "email4", "email5"})

	sendEmails(1, ch)
	ch <- "email6"
	ch <- "email7"

}
