package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
)

// Define all types used
// Make sure Story map[string]Chapter, Chapter struct, Option struct types are distinctly defined
// Remember that Arc = Chapter, Story = Paragraphs,
// Add struct tags as well

// REFACTOR
// Add a JsonStory function that takes an io.Reader and returns (Story, error) to story.go.
// it should decode into a new var story, handle a case for it having an error and return the story var if no error
// Edit main.go to reflect this as well

// grab the template and place it in a variable called defaultHandlerTmpl
var defaultHandlerTmpl = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      {{if .Options}}
        <ul>
        {{range .Options}}
          <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
        </ul>
      {{end}}
  </body>
</html>`

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

// Define a template variable and assign to tpl in a func init(), then globally declare var tpl
var tpl *template.Template

// Create NewHandler function that takes in a Story and returns http.Handler interface

func NewHandler(s Story) http.Handler {
	return handler{s}
}

// Create a new struct handler and a ServeHTTP method that allows it conform to the http.Handle interface
type handler struct {
	s Story
}

// Write ServeHTTP for main handler where it it takes in w and r and uses tpl.Execute() to write some the first Chapter with key "intro" to respWriter w and handle the error by panicking.
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
}

func JsonStory(r io.Reader) (Story, error) {
	var story Story
	d := json.NewDecoder(r)
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"chapter"`
}
