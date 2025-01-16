package main

import (
	"context"
	// _ "embed"
	"html/template"
	"os"

	"github.com/redis/go-redis/v9"
)

// ckey/ctxkey is used as the key for the HTML context and is how we retrieve
// token information and pass it around to handlers
type ckey int

const ctxkey ckey = iota

var (
	t           *template.Template
	servicePort                    = ":" + os.Getenv("servicePort")
	logFilePath                    = os.Getenv("logFilePath")
	templates   *template.Template = template.New("main")
	companyName string             = "BoltApp"
	PDFs        []string
	////go:embed tags.txt
	//byteTags []byte
	tags    []tagger
	cats    []string
	reports map[string]report   = make(map[string]report)
	catalog map[string][]string = make(map[string][]string)
	rdb                         = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       9,  // use default DB
	})
	rdx = context.Background()
)
