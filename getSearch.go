package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/redis/go-redis/v9"
)

func getSearch(w http.ResponseWriter, r *http.Request) {
	search := strings.Split(r.RequestURI, "/")[2]
	if len(search) > 0 {
		opts := &redis.ZRangeBy{
			Min:    "(" + string(search[0]),
			Max:    strings.ReplaceAll("["+search+"z", "%20", " "),
			Offset: 0,
			Count:  50,
		}
		mems, err := rdb.ZRevRangeByLex(rdx, "LEXTAGS", opts).Result()
		if err != nil {
			log.Println(err)
		}
		log.Println(mems, opts)
		keystring := strings.Join(mems, ",")
		ajaxResponse(w, map[string]string{
			"success":   "true",
			"responses": keystring,
		})
	}
}
