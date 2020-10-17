package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)
type Participant struct {
    Name string `json:"Name"`
    Email string `json:"Email"`
    RSVP string `json:"RSVP"`
}
var Articles []Participant
func allArticles(w http.ResponseWriter, r *http.Request){
	Articles = []Participant{
        Participant{Name: "Roy", Email: "roy@yahoo.com", RSVP: "Yes"},
        Participant{Name: "Sushant", Email: "sus@gmail.com", RSVP: "No"},
    }
    fmt.Println("Endpoint Hit: allArticles")
    json.NewEncoder(w).Encode(Articles)
}
func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Welcome to the HomePage!")
	fmt.Fprintln(w, "Use the link http://localhost:10000/participants to see Participants detail")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/participants", allArticles)
    http.ListenAndServe(":10000", nil)
}

func main() {
    handleRequests()
}
