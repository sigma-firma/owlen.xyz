package main

import "strings"

type report struct {
	ID              string   `json:"id,omitempty"                    redis:"id"`
	Submitter       string   `json:"submitter,omitempty"             redis:"submitter"`
	Authors         string   `json:"authors,omitempty"               redis:"authors"`
	Title           string   `json:"title,omitempty"                 redis:"title"`
	Comments        string   `json:"comments,omitempty"              redis:"comments"`
	JournalRef      string   `json:"journal-ref,omitempty"           redis:"journal-ref"`
	Doi             string   `json:"doi,omitempty"                   redis:"doi"`
	ReportNo        string   `json:"report-no,omitempty"             redis:"report-no"`
	Categories      string   `json:"categories"                      redis:"categories"`
	SplitCategories []string `json:"splitCategories,omitempty"`
	// License    interface{} `json:"license,omitempty"               redis:"license"`
	Abstract      string   `json:"abstract,omitempty"              redis:"abstract"`
	Keywords      string   `json:"keywords"                        redis:"keywords"`
	KeywordSlice  []string `json:"keywordSlice"`
	UpdateDate    string   `json:"update_date"                     redis:"update_date"`
	AuthorsParsed keyw     `json:"authors_parsed"                  redis:"authors_parsed"`
}
type tagger string

func (t *tagger) Id(s tagger) tagger {
	return tagger(strings.ReplaceAll(string(s), " ", "_"))
}

type viewData struct {
	PageTitle   string
	CompanyName string
	Media       []report
	Tags        []tagger
	Categories  []string
	Selected    []tagger
	Report      report
	Id          func(string) string
}
