package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Welcome to the Time package learning")

	presentTime :=time.Now();
	fmt.Println("current time= ", presentTime)

	fmt.Println("time formate =", presentTime.Format("01-02-2006"))  // for get current day in this formate.
	fmt.Println("time formate =", presentTime.Format("01-02-2006 Monday"))  // for get current day in this formate with current week day. this is syntax
	fmt.Println("time formate =", presentTime.Format("01-02-2006 15:04:05 Monday"))  // for get current day in this formate with current week day with current time. this is syntax

	createdDate := time.Date(2020, time.October,10,23,23,0,0, time.Local)
	fmt.Println("created date= ", createdDate);
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}