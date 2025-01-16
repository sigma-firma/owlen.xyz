package main

import (
	"log"
	"net/http"
	"strings"
)

func getScience(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.ReplaceAll(r.RequestURI, "%20", " "), "/")
	if strings.Contains(parts[2], "+") {
		tags_ := strings.Split(parts[2], "+")
		vd := tagsUnion(tags_, &viewData{})
		exeTmpl(w, r, vd, "tag.tmpl")
		return
	}
	vd := viewData{}
	rp := report{}
	err := rdb.HGetAll(rdx, parts[2]).Scan(&rp)
	if err != nil {
		log.Println(err)
	}
	vd.Report = rp
	vd.Report.KeywordSlice = strings.Split(rp.Categories, " ")

	exeTmpl(w, r, &vd, "viewFull.tmpl")
}
