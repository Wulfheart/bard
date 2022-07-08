package schema

// Description of additional autoload rules for development purpose (eg. a test suite).
type AutoloadDev struct {
	Classmap []interface{}   `json:"classmap,omitempty"`// This is an array of paths that contain classes to be included in the class-map generation; process.
	Files    []interface{}   `json:"files,omitempty"`   // This is an array of files that are always required on every request.
	Psr0     map[string]*Bin `json:"psr-0,omitempty"`   // This is an object of namespaces (keys) and the directories they can be found into; (values, can be arrays of paths) by the autoloader.
	Psr4     map[string]*Bin `json:"psr-4,omitempty"`   // This is an object of namespaces (keys) and the PSR-4 directories they can map to (values,; can be arrays of paths) by the autoloader.
}
