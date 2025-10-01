// viewData represents the root model used to dynamically update the app
package main

import (
	"html/template"
	"log"
	"strings"

	"github.com/redis/go-redis/v9"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// rand.Seed(time.Now().UTC().UnixNano())
	funcMap := template.FuncMap{
		"Id": func(s string) string {
			return strings.ReplaceAll(s, " ", "_")
		},
	}

	t = template.Must(template.New("").Funcs(funcMap).ParseGlob("internal/*/*/*"))
}
func main() {
	if len(logFilePath) > 1 {
		f := setupLogging()
		defer f.Close()
	}

	getReports()

	ctx, srv := bolt()

	log.Println("Waiting for connections @ http://localhost" + srv.Addr)

	<-ctx.Done()
}
func getReports() {
	opts := &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  150,
	}

	topTags, err := rdb.ZRevRangeByScore(rdx, "TOPTAGS", opts).Result()
	if err != nil {
		log.Println(err)
	}

	for _, tag := range topTags {
		tags = append(tags, tagger(strings.ReplaceAll(tag, "::", " ")))
	}

	topCats, err := rdb.ZRevRangeByScore(rdx, "TOPCATS", opts).Result()
	if err != nil {
		log.Println(err)
	}

	for _, cat := range topCats {
		cats = append(cats, cat)
	}

}
