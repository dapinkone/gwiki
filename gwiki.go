//serves up plain web pages.
// This code originally taken from https://gobyexample.com/http-servers

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func login(w http.ResponseWriter, req *http.Request) {
	dat, err := ioutil.ReadFile("www/login.html")
	check(err)
	fmt.Fprintf(w, string(dat))
}

func main() {
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8090", nil)
}
