package main

import (
	"os"
	"time"

	clockface "github.com/qinchenfeng/HelloLearnGoWithTest/1_Go_fundamentals/17_Math"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
