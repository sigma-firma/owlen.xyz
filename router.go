package main

import "net/http"

// registerRoutes registers the routes with the provided *http.ServeMux
func registerRoutes(mux *http.ServeMux) {
        mux.HandleFunc("/search/", getSearch)
        mux.HandleFunc("/science/", getScience)
        mux.HandleFunc("/cat/", getCat)
	mux.HandleFunc("/tag/", getTagged)
	mux.HandleFunc("/submitData", submitData)
	mux.HandleFunc("/", home)
}