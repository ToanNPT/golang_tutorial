package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Email struct {
	from string
	to   string
	body string
}

func logMessage(sms string) {
	fmt.Printf("Message log: %s\n", sms)
}

func logEmail(email Email) {
	fmt.Printf("Email log: %s\n", email.body)
}

func logCommonMessage(smsCh chan string, emailCh chan Email) {
	for {
		select {
		case sms, ok := <-smsCh:
			if !ok {
				return
			} else {
				logMessage(sms)
			}
		case email, ok := <-emailCh:
			if !ok {
				return
			} else {
				logEmail(email)
			}
		}
	}
}

func main() {
	testFunc2(
		[]string{
			"Hello, Golang!",
			"Heelo, World!",
		},
		[]Email{
			{"me", "you", "Hello, Golang!"},
			{"me", "you", "Hello, World!"},
		})
}

func testFunc2(sms []string, emails []Email) {
	fmt.Println("Starting...")

	chSms, chEmails := sendToLogger(sms, emails)

	logCommonMessage(chSms, chEmails)
	fmt.Println("===============================")
}

func sendToLogger(sms []string, emails []Email) (chSms chan string, chEmails chan Email) {
	chSms = make(chan string)
	chEmails = make(chan Email)
	go func() {
		for i := 0; i < len(sms) && i < len(emails); i++ {
			done := make(chan struct{})
			s := sms[i]
			e := emails[i]
			t1 := time.Millisecond * time.Duration(rand.Intn(1000))
			t2 := time.Millisecond * time.Duration(rand.Intn(1000))
			go func() {
				time.Sleep(t1)
				chSms <- s
				done <- struct{}{}
			}()
			go func() {
				time.Sleep(t2)
				chEmails <- e
				done <- struct{}{}
			}()
			<-done
			<-done
			time.Sleep(10 * time.Millisecond)
		}
		close(chSms)
		close(chEmails)
	}()
	return chSms, chEmails
}
