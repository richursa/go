package main

import (
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var randomValue string

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Sample webpage time : " + time.Now().String() + "\n" + "containerRandomNumber = " + randomValue + "\nHAAAI GAUTHAM \nENTHARUND CHELLAKILI"))
}
func main() {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 50
	randomValue = strconv.Itoa(rand.Intn(max-min+1) + min)
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
