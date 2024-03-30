package main

import (
	"fmt"
	"strconv"
)

type shape interface {
	area() float64
	perimeter() float64
}

type react struct {
	width, height float64
}

func (r react) area() float64 {
	return r.width * r.height
}

func (r react) perimeter() float64 {
	return 2*r.width + 2*r.height
}

//sample 2

type message interface {
	getMessage() string
}

type sendingReport struct {
	reportNo int
	content  string
}

type birthdayWish struct {
	date     string
	fullname string
	content  string
}

func (s sendingReport) getMessage() string {
	return "Report Number " + strconv.Itoa(s.reportNo) + " : " + s.content
}

func (b birthdayWish) getMessage() string {
	return "Happy birthday " + b.fullname + " on " + b.date + " : " + b.content
}

func inquireMessage(m message) {
	fmt.Println("This is required message: ")
	fmt.Println(m.getMessage())
	fmt.Println("End of message ======================")
}

func main() {
	r := react{width: 12, height: 2}
	s := shape(r)
	fmt.Printf("Area: %f, Perimeter: %f\n", s.area(), s.perimeter())

	//sample 2
	birthDayMess := birthdayWish{date: "2020-01-01", fullname: "Toan NPT", content: "Wish you all the best"}
	inquireMessage(birthDayMess)

	reportMess := sendingReport{reportNo: 1, content: "This is the first report"}
	inquireMessage(reportMess)
}
