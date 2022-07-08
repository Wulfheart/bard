package schema

type Source struct {
	Mirrors   []interface{} `json:"mirrors,omitempty"`
	Reference string        `json:"reference"`        
	Type      string        `json:"type"`             
	URL       string        `json:"url"`              
}
