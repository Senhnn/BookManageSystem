package main

import (
	"bookmanagesystem/proto/base"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:10000")
	if err != nil {
		fmt.Println(err)
	}
	st := &base.Aaaa{}
	res, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(res)
	proto.Unmarshal(res, st)
	fmt.Println(st)
	fmt.Println(resp.Header)
}
