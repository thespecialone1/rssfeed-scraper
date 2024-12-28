package main

import (
	"net/http"
)


func errHandler(w http.ResponseWriter, r *http.Request){
	respondWithErrors(w, 400, "something went wrong")
}
 