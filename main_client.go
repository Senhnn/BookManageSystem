package main

import (
	"fmt"
)

func main() {
	TestSlice()
	TestMap()
	TestChan()
}

func reverseLeftWords(s string, n int) string {
	r := []byte(s)
	res := append(r[n:len(r)], r[0:n]...)
	return string(res)
}

func TestSlice() {
	s := make([]int, 0, 10)
	s = append(s, 10)
	fmt.Println(s)
	SliceApp(s)
	fmt.Println(s)
}

func SliceApp(s []int) {
	// 此处slice的len + 1，但是外部的slice的len依然没变
	s = append(s, 20)
}

func TestMap() {
	s := map[string]string{}
	s["123"] = "123"
	s["456"] = "456"
	s["789"] = "789"
	fmt.Println(s)
	MapDelAdd(s)
	fmt.Println(s)
}

func MapDelAdd(s map[string]string) {
	s["000"] = "000"
	delete(s, "123")
}

func TestChan() {
	ch := make(chan int, 10)
	ch <- 10
	ch <- 20
	fmt.Println(len(ch))
	ChanOut(ch)
	fmt.Println(len(ch))
}

func ChanOut(ch chan int) {
	<-ch
	ch <- 30
	ch <- 50
}
