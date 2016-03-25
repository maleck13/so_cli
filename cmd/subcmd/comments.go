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

const LIST_COMMENTS_API = "https://api.stackexchange.com/2.2/questions/%s/comments?order=desc&sort=creation&site=stackoverflow&filter=withBody"

var test = `{"items":[{"owner":{},"edited":false,"score":0,"creation_date":1458692007,"post_id":36166791,"comment_id":59972415,"body":"No Problem pnovotak! Happy Golearning"}]}`

type CommentsResponse struct {
	Items []*Comment `json:"items"`
}

type Comment struct {
	Body  string `json:"body"`
	Score int    `json:"score"`
}

var commentTemplate = `
	{{range .Items}}
		|body  | {{.Body}}
		|score | {{.Score}}
		- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	{{end}}
`

func GetComments(context *cli.Context) {
	args := context.Args()
	if len(args) == 0 || "" == args[0] {
		log.Fatal("missing argument questionid. Usage: ", context.Command.Usage)
	}
	quid := args[0]
	t := template.New("comments")
	t, err := t.Parse(commentTemplate)
	if nil != err {
		log.Fatal(err)
	}
	url := fmt.Sprintf(LIST_COMMENTS_API, quid)
	resp, err := http.Get(url)
	if nil != err {
		log.Fatal("error getting response ", err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	comments := &CommentsResponse{}
	if err := decoder.Decode(comments); err != nil {
		log.Fatal("failed to decode response ", err)
	}
	t.Execute(os.Stdout, comments)

}
