package cyoa

// Define all types used
// Make sure Story map[string]Chapter, Chapter struct, Option struct types are distinctly defined
// Remember that Arc = Chapter, Story = Paragraphs,
// Add struct tags as well

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
