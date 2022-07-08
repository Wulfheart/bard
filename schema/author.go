package schema

// List of authors that contributed to the package. This is typically the main maintainers,
// not the full list.
type Author struct {
	Email    *string `json:"email,omitempty"`   // Email address of the author.
	Homepage *string `json:"homepage,omitempty"`// Homepage URL for the author.
	Name     string  `json:"name"`              // Full name of the author.
	Role     *string `json:"role,omitempty"`    // Author's role in the project.
}
