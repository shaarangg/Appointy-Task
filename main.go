package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shaarangg/Appointy-Task/controllers"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint hit")
}
func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
func main() {
	handleRequests()
	controllers.Dbinit()
}
