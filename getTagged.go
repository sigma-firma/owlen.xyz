package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/redis/go-redis/v9"
)

type keyw [][]string

func (i keyw) MarshalBinary() ([]byte, error) { return json.Marshal(i) }

func getTagged(w http.ResponseWriter, r *http.Request) {
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
		json.Unmarshal([]byte(newR.Keywords), &newR.KeywordSlice)
		json.Unmarshal(marshaled, &newR.SplitCategories)
		// vd. = s[1 : len(s)-1]
		vd.Media = append(vd.Media, newR)
	}
	t = strings.ReplaceAll(parts[len(parts)-1], "%20", " ")
	vd.Selected = append(vd.Tags, tagger(t))
	exeTmpl(w, r, &vd, "tag.tmpl")

}
