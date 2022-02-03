package main

import (
	"net/http"
	"os"
	"test_task/handler"
)



func main() {
	storage_type := os.Args[1:][0]
	handler.Stor_type(storage_type)
	http.HandleFunc("/", handler.MainHandler)
	http.ListenAndServe(":8080", nil)
}
