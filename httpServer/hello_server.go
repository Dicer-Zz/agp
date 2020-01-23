package main

import (
	"fmt"
	"net/http"
	"strings"
)

type Hello struct {
	Name   string
	Age    string
	Gender string
}

func (hello Hello) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	urList := strings.Split(req.URL.Path[1:], "/")
	// fmt.Print(urList[0])
	hello.Name, hello.Age, hello.Gender = urList[0], urList[1], urList[2]
	fmt.Fprintf(w, "Hello,"+hello.Name)
	fmt.Fprint(w, hello)
}
func main() {
	var hello Hello
	// http.Handle("/", hello)
	http.ListenAndServe(":8080", hello)
}
