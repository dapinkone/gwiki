//serves up plain web pages.
// This code originally taken from https://gobyexample.com/http-servers

package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"io"
)
func check(e error) {
	if e != nil { panic(e) }
}


func login(w http.ResponseWriter, req *http.Request) {
	dat, err f := openFile("www/login.html")
	defer closeFile(f)
	fmt.Fprintf(w, f.)
}

func main() {
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8090", nil)
}
