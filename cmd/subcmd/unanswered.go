package subcmd

import (
	"encoding/json"
	"fmt"
	"github.com/maleck13/so_cli/Godeps/_workspace/src/github.com/codegangsta/cli"
	"log"
	"net/http"
	"os"
	"text/template"
)

const UNANSWERED_API = "https://api.stackexchange.com/2.2/questions/unanswered?order=desc&sort=activity&tagged=%s&site=stackoverflow"

type UnansweredQuest struct {
	Tags  []string `json:"tags"`
	Link  string   `json:"link"`
	Title string   `json:"title"`
	Score int      `json:"score"`
	Id    int      `json:"question_id"`
}

type UnansweredResponse struct {
	Items []*UnansweredQuest `json:"items"`
}

var unansweredTemplate = `
	{{range .Items}}
	|title | {{.Title}}
	|tags  | {{range .Tags }}{{ . }} {{end}}
	|link  | {{.Link}}
	|score | {{.Score}}
	|id    | {{.Id}}
	 - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	{{end}}
`
var FlagTag string

func GetUnanswered(context *cli.Context) {
	t := template.New("unanswered")
	t, err := t.Parse(unansweredTemplate)
	if nil != err {
		log.Fatal(err)
		return
	}
	url := fmt.Sprintf(UNANSWERED_API, FlagTag)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	var questions UnansweredResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&questions); err != nil {
		log.Fatal("failed to decode json ", err)
	}
	t.Execute(os.Stdout, questions)
}
