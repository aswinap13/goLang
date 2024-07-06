package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hey welcome to Time")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("2006-01-02 15-04-05 Monday"))

	createdDate := time.Date(2020, time.August, 30, 23, 23, 0, 0, time.UTC)
	//Year, time.Month, date, timeHR, timeMin, timeSec, timeMilliSec, TimeZone

	fmt.Println(createdDate.Format("2006-01-02 Monday 15-04-05.00"))

	//Year :: 2006
	//Month :: 01
	//Date :: 02
	//Day :: Monday
	//TimeH :: 15
	//Minute :: 04
	//Second :: 05
}
