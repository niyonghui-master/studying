package main

import (
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleDelete(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func handlePut(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func handlePost(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	return
}
