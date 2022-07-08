package schema

type InlinePackage struct {
	Archive     *InlinePackageArchive `json:"archive,omitempty"`     
	Authors     []Author              `json:"authors,omitempty"`     
	Autoload    *Autoload             `json:"autoload,omitempty"`    
	Bin         *Bin                  `json:"bin"`                   // A set of files, or a single file, that should be treated as binaries and symlinked into; bin-dir (from config).
	Conflict    map[string]string     `json:"conflict,omitempty"`    
	Description *string               `json:"description,omitempty"` 
	Dist        *Dist                 `json:"dist,omitempty"`        
	Extra       *Extra                `json:"extra"`                 
	Homepage    *string               `json:"homepage,omitempty"`    
	IncludePath []string              `json:"include-path,omitempty"`// DEPRECATED: A list of directories which should get added to PHP's include path. This is; only present to support legacy projects, and all new code should preferably use; autoloading.
	Keywords    []string              `json:"keywords,omitempty"`    
	License     *Comment              `json:"license"`               
	Name        string                `json:"name"`                  // Package name, including 'vendor-name/' prefix.
	Provide     map[string]string     `json:"provide,omitempty"`     
	Replace     map[string]string     `json:"replace,omitempty"`     
	Require     map[string]string     `json:"require,omitempty"`     
	RequireDev  map[string]string     `json:"require-dev,omitempty"` 
	Source      *Source               `json:"source,omitempty"`      
	Suggest     map[string]string     `json:"suggest,omitempty"`     
	TargetDir   *string               `json:"target-dir,omitempty"`  // DEPRECATED: Forces the package to be installed into the given subdirectory path. This is; used for autoloading PSR-0 packages that do not contain their full path. Use forward; slashes for cross-platform compatibility.
	Time        *string               `json:"time,omitempty"`        
	Type        *string               `json:"type,omitempty"`        
	Version     string                `json:"version"`               
}
