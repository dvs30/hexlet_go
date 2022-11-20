package main

import (
	"fmt"
	"hello/greeting"
	"hello/hxlt"

	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println(greeting.Hello())
	logrus.Println("Hello, Hexlet!")
	fmt.Println(hxlt.Even([]int{1, 2, 3, 4}))
}
