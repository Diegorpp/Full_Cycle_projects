package main

import "net/http"

func main(){
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80",nil)
}

func Hello(w http.ResponseWriter, r *http.Request){
	w.write( []Byte("<h1> Welcome to full cycle</h1>"))
}

