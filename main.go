package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/get", getHandler).Methods("GET")
	router.HandleFunc("/post", postHandler).Methods("POST")
	// route declarations continue like this

	fmt.Println("Server started")
	http.ListenAndServe(":80", router)
}
func getHandler(res http.ResponseWriter, req *http.Request) {

	body := req.URL.RawQuery

	// fmt.Println(body)
	fmt.Fprintln(res, body)
	fmt.Printf("%T\n", body)
	getResponse(body)
}

func postHandler(res http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Println(err)
	}
	myString := string(body)
	fmt.Println(myString)
	// fmt.Println(res, "POST")
	fmt.Printf("%T\n", myString)

	myString = strings.ReplaceAll(myString, "\r\n", "")
	myString = strings.ReplaceAll(myString, " ", "")
	myString = strings.ReplaceAll(myString, "\"", "")
	myString = strings.ReplaceAll(myString, ":", "=")
	myString = strings.ReplaceAll(myString, "{", "")
	myString = strings.ReplaceAll(myString, "}", "")
	myString = strings.ReplaceAll(myString, ",", "&")
	getResponse(myString)
}

func getResponse(str string) {

	urlAddr := "https://phyto-pharm.ru/postback?" + str
	resp, err := http.Get(urlAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
}
