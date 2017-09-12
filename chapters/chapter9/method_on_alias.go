package main

import (
	"fmt"
	"time"
)

type myTime struct {
	time.Time
}

//接受指针
func (mT *myTime) first9Str() string {
	return mT.Time.String()[0:9]
}

funcmethodOnAlias() {
	mT := myTime{time.Now()}
	fmt.Printf("%s\n", mT.first9Str())
}
