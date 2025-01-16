package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func tagsUnion(uTags []string, vd *viewData) *viewData {
	log.Println(uTags)
	rdb.ZUnionStore(rdx, "tempstore", &redis.ZStore{Keys: uTags}).Result()
	ids, _ := rdb.ZRevRange(rdx, "tempstore", 0, 15).Result()
	for _, id := range ids {
		r := report{}
		err := rdb.HGetAll(rdx, id).Scan(&r)
		if err != nil {
			log.Println(err)
		}
		vd.Media = append(vd.Media, r)
	}
	return vd
}

// exeTmpl is used to build and execute an html template.
func exeTmpl(w http.ResponseWriter, r *http.Request, view *viewData, tmpl string) {
	if view == nil {
		view = &viewData{}
	}

	err := t.ExecuteTemplate(w, tmpl, view)
	if err != nil {
		log.Println(err)
	}
}

// ajaxResponse is used to respond to ajax requests with arbitrary data in the
// format of map[string]string
func ajaxResponse(w http.ResponseWriter, res map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println(err)
	}
}

// genPostID generates a post ID
func genPostID(length int) (ID string) {
	symbols := "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i <= length; i++ {
		s := rand.Intn(len(symbols))
		ID += symbols[s : s+1]
	}
	return
}
