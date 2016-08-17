package main

import (
	"fmt"
	"time"

	"github.com/zhangjunfang/dateTool"
)

func main() {

	var format = "2006-01-02 15:04:05.999999999"
	n := time.Date(2016, 8, 16, 17, 00, 00, 00, time.UTC)

	fmt.Println(time.Now().Sub(n))
	fmt.Println(dateTool.New(n).BeginningOfMinute().Format(format))
}
