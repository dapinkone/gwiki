//serves up plain web pages.
// This code originally taken from https://gobyexample.com/http-servers

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//	"strings"
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
	if v := req.Method; v == http.MethodPost {
		//		fmt.Fprintf(w, "Post recieved")
		req.ParseForm()
		for name, headers := range req.Header {
			for _, h := range headers {
				fmt.Fprintf(w, "%v: %v<br/>", name, h)
			}
		}
		for k, v := range req.Form {
			fmt.Fprintf(w, "%v: %v<br/>", k, v)
		}
	}
}
func serverFactory(filename string) func(http.ResponseWriter, *http.Request) {
	// WARN NOTE: 2020-12-02: dapinkone: filename needs to be sanitized.
	return func(w http.ResponseWriter, req *http.Request) {
		dat, err := ioutil.ReadFile("www" + filename)
		check(err)
		fmt.Fprintf(w, string(dat))
	}
}

func view(w http.ResponseWriter, req *http.Request) {
	v := req.Method;
	if v == http.MethodGet {
		req.ParseForm()
		fmt.Fprintf(w, "viewing %v<br />\n", req.Form["page"])
		pageid := req.Form["page"]
		// TODO: sanitize pageid

		// TODO: check for existance of file

		// TODO: read file if exists. otherwise, provide option to edit/create.


	} // Else?
}

func main() {
	http.HandleFunc("/", serverFactory("/index.html"))
	http.HandleFunc("/view", view)
	http.HandleFunc("/style.css", serverFactory("/style.css"))
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8090", nil)
}
