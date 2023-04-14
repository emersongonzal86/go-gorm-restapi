package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		w.Write([]byte("Hello Word!"))
	})

	http.ListenAndServe(":3000",r)
}