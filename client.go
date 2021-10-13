package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {


	a:=1
	fmt.Println(string(a))
	b := bytes.NewBufferString("a=1&b=2")

	res, err := http.Post("http://localhost:9090", "multipart/form-data", b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ioutil.ReadAll(res.Body))
}
