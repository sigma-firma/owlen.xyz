package main

import (
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	v := viewData{}
	v.Tags = append(v.Tags, tags...)
	v.Categories = append(v.Categories, cats...)
	exeTmpl(w, r, &v, "main.tmpl")
}
