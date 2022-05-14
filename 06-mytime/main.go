package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println("Present time: ", presentTime)

	fmt.Println("\nPresent time formated: ", presentTime.Format("01-02-2006"))                 //this is the default format
	fmt.Println("\nPresent time formated: ", presentTime.Format("01-02-2006 15:04:05 Monday")) //this is the default format

	createDate := time.Date(2022, time.May, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println("\n\nCreateDate: ", createDate)
	fmt.Println("CreateDate: ", createDate.Format("01-02-2006 Monday"))

}
