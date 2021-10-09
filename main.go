package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/shaarangg/Appointy-Task/controllers"
)

func handleFunc(response http.ResponseWriter, request *http.Request) {
	url := request.URL.Path
	if strings.Contains(url, "/users/posts") {
		controllers.AllPostRequests(response, request)
		return
	} else if strings.Contains(url, "/users") {
		controllers.UserRequests(response, request)
		return
	} else if strings.Contains(url, "/posts") {
		controllers.PostRequests(response, request)
		return
	} else {
		fmt.Fprintf(response, "url")
	}
}
func handleRequests() {
	http.HandleFunc("/", handleFunc)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	controllers.Dbinit()
	handleRequests()
}
