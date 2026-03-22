package main

import (
	"fmt"
	"net/http"
)


func initServer(mux *http.ServeMux) {
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
		return
	}
}
