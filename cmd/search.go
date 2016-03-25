package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/maleck13/so_cli/Godeps/_workspace/src/github.com/codegangsta/cli"
	"log"
	"net/http"
	"net/url"
	"os"
	"text/template"
)

const SEARCH_QUESTION_API = "https://api.stackexchange.com/2.2/search/advanced?order=desc&sort=activity&q=%s&site=stackoverflow"

var (
	Flag_Search_Query string
	Flag_Search_Tag   string
)

func SearchCmd() cli.Command {
	return cli.Command{
		Name:  "search",
		Usage: "search stackoverflow. --query=\"some question\" --tags=go;nodejs tags is optional and limits to certain subjects",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "query",
				Value:       "whats the meaning of life",
				Usage:       "search --query=\"some query\"",
				Destination: &Flag_Search_Query,
			},
			cli.StringFlag{
				Name:        "tags",
				Value:       "",
				Usage:       "search --query=\"some query\" --tags=go;nodejs",
				Destination: &Flag_Search_Tag,
			},
		},
		Action: search,
	}
}

var searchTemplate = `
	{{range .Items}}
		|title   	| {{.Title}}
		|answers 	| {{.Answers}}
		|score   	| {{.Score}}
		|is answered 	| {{.IsAnswered}}
		|link		| {{.Link}}
		- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	{{end}}
`

type QuestionResults struct {
	Items []*Question
}

type Question struct {
	Title      string `json:"title"`
	Answers    int    `json:"answer_count"`
	Score      int    `json:"score"`
	IsAnswered bool   `json:"is_answered"`
	Link       string `json:"link"`
}

func search(context *cli.Context) {
	var api = fmt.Sprintf(SEARCH_QUESTION_API, url.QueryEscape(Flag_Search_Query))
	if Flag_Search_Tag != "" {
		api += "&tagged=" + url.QueryEscape(Flag_Search_Tag)
	}

	t := template.New("questions")
	t, err := t.Parse(searchTemplate)
	if nil != err {
		log.Fatal(err)
		return
	}
	log.Print(api)
	resp, err := http.Get(api)
	if nil != err {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	questions := &QuestionResults{}
	if err := decoder.Decode(questions); err != nil {
		log.Fatal("failed to decode json ", err)
	}
	if err := t.Execute(os.Stdout, questions); err != nil {
		log.Fatal("failed to output template ", err)
	}

}
