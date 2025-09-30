package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/redis/go-redis/v9"
)

func getCat(w http.ResponseWriter, r *http.Request) {
	log.Println(" test")
	opts := &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  20,
	}

	parts := strings.Split(r.RequestURI, "/")
	t := strings.ReplaceAll(parts[len(parts)-1], "%20", " ")
	reports, err := rdb.ZRangeByScore(rdx, t, opts).Result()
	if err != nil {
		log.Println(err)
	}
	vd := viewData{}
	for _, r := range reports {
		newR := report{}
		err := rdb.HGetAll(rdx, r).Scan(&newR)
		if err != nil {
			log.Println(err)
		}
		marshaled, err := json.Marshal(strings.Fields(newR.Categories))
		if err != nil {
			log.Println(err)
		}
		err = json.Unmarshal([]byte(newR.Keywords), &newR.KeywordSlice)
		if err != nil {
			log.Println(err)
		}
		err = json.Unmarshal(marshaled, &newR.SplitCategories)
		if err != nil {
			log.Println(err)
		}
		// vd. = s[1 : len(s)-1]
		vd.Media = append(vd.Media, newR)
	}
	t = strings.ReplaceAll(parts[len(parts)-1], "%20", " ")
	vd.Selected = append(vd.Tags, tagger(t))
	exeTmpl(w, r, &vd, "tag.tmpl")
}
