package cyoa

import (
	"encoding/json"
	"io"
)

// Define all types used
// Make sure Story map[string]Chapter, Chapter struct, Option struct types are distinctly defined
// Remember that Arc = Chapter, Story = Paragraphs,
// Add struct tags as well

// REFACTOR
// Add a JsonStory function that takes an io.Reader and returns (Story, error) to story.go.
// it should decode into a new var story, handle a case for it having an error and return the story var if no error
// Edit main.go to reflect this as well

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
	Title       string   `json:"title"`
	Parapghraps []string `json:"story"`
	Options     []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"chapter"`
}
