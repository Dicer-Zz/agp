package main

import (
	"fmt"
	"net/http"
	"strings"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	urList := strings.Split(req.URL.Path[1:], "/")
	fmt.Println(urList[0])
	if urList[0] == "hello" {
		// fmt.Fprintf(w, "Hello, "+strings.ToUpper(urList[1]))
		fmt.Fprintf(w, "<h1>Hello, %s</h1>", urList[1])
	} else {
		fmt.Println("The first req is not HELLO")
		fmt.Fprintf(w, "<h1>Bad Request.</h1>")
	}
}
func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(HelloServer))

}
