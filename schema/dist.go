package schema

type Dist struct {
	Mirrors   []interface{} `json:"mirrors,omitempty"`  
	Reference *string       `json:"reference,omitempty"`
	Shasum    *string       `json:"shasum,omitempty"`   
	Type      string        `json:"type"`               
	URL       string        `json:"url"`                
}
