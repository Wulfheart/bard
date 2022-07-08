package schema

// Options for creating package archives for distribution.
type ComposerSchemaArchive struct {
	Exclude []interface{} `json:"exclude,omitempty"`// A list of patterns for paths to exclude or include if prefixed with an exclamation mark.
	Name    *string       `json:"name,omitempty"`   // A base name for archive.
}
