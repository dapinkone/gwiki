//serves up plain web pages.
// This code originally taken from https://gobyexample.com/http-servers

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func login(w http.ResponseWriter, req *http.Request) {
	// TODO: generate and store sessionids?
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
func sanitizeFormData(s string) string {
	// TODO: test this for edge cases.
	res := strings.Replace(s, ".", "", -1)
	return strings.Replace(res, "/", "", -1)
}

func serverFactory(filename string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		dat, err := ioutil.ReadFile("www" + filename)
		check(err)
		fmt.Fprintf(w, string(dat))
	}
}

func view(w http.ResponseWriter, req *http.Request) {
	v := req.Method
	if v == http.MethodGet {
		req.ParseForm()
						// TODO: sanitize pageid
		pageid := sanitizeFormData(req.Form["page"][0])

		fmt.Fprintf(w, "viewing %v<br />\n", pageid)


		// TODO: future feature: store pages in a databse.
		// TODO: read file if exists. otherwise, provide option to edit/create.
		dat, err := ioutil.ReadFile("www/pgs/" + pageid)
		if err != nil { // TODO: redirect to edit page
			fmt.Fprintf(w, "not found")

		} else {
			fmt.Fprintf(w, string(dat))
		}
	} // Else?
}

func main() {
	http.HandleFunc("/", serverFactory("/index.html"))
	http.HandleFunc("/view", view)
	http.HandleFunc("/style.css", serverFactory("/style.css"))
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8090", nil)
}
