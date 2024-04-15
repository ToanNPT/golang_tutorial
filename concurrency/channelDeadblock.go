package main

import "time"

func main() {
	sendEmail([]email{
		{
			body: "Are you going to make it?",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "I need a break",
			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "What were you thinking?",
			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})

	deadLocksendEmail([]email{
		{
			body: "Are you going to make it?",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "I need a break",
			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "What were you thinking?",
			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
}

type email struct {
	body string
	date time.Time
}

/*
This function will be blocked because the channel isOldMail is not buffered and the main goroutine is waiting for the channel to be empty.
*/
func deadLocksendEmail(emailList []email) {
	isOldMail := make(chan bool)

	for _, email := range emailList {
		if email.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldMail <- true
			continue
		}
		isOldMail <- false
	}

	for i := 0; i < len(emailList); i++ {
		if <-isOldMail {
			println("Email is old")
		} else {
			println("Email is new")
		}
	}

}

func sendEmail(emailList []email) {
	isOldMail := make(chan bool)

	go func(emailList []email) {
		for _, email := range emailList {
			if email.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
				isOldMail <- true
				continue
			}
			isOldMail <- false
		}
	}(emailList)

	for i := 0; i < len(emailList); i++ {
		if <-isOldMail {
			println("Email is old")
		} else {
			println("Email is new")
		}
	}
}
