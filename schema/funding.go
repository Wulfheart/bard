package schema

type Funding struct {
	Type *string `json:"type,omitempty"`// Type of funding or platform through which funding is possible.
	URL  *string `json:"url,omitempty"` // URL to a website with details on funding and a way to fund the package.
}
