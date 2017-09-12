package main

import (
  "fmt"
  "time"
)

const (
  // DATE_FORMAT 自定义时间格式化
  DATE_FORMAT = "2006-01-02 15:04:05"
)

func timeDemo()  {
  var timeNow = time.Now();
  fmt.Println(timeNow)
  //time.ANSIC : "Mon Jan _2 15:04:05 2006"
  fmt.Println(timeNow.Format(time.ANSIC))
  // RFC3339 : 2006-01-02T15:04:05Z07:00
  fmt.Println(timeNow.Format(time.RFC3339))
  //
  fmt.Println(timeNow.Format(DATE_FORMAT))

  fmt.Println()

  timeFormat := timeNow.Format(DATE_FORMAT)
	fmt.Println("timeNow(time format):", timeFormat)

	//string转化为时间，layout必须为 "2006-01-02 15:04:05"
	t, _ := time.Parse(DATE_FORMAT, "2017-06-22 08:37:18")
	fmt.Println("t(time format)", t)

	//某个时间点 前后判断
	trueOrFalse := t.After(timeNow)
	if trueOrFalse {
		fmt.Printf("t:%s 在timeNow之后!\n",t.Format(DATE_FORMAT))
	} else {
		fmt.Printf("t:%s 在timeNow之前!\n",t.Format(DATE_FORMAT))
	}
	fmt.Println("Over!")
}
