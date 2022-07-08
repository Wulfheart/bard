// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    composerSchema, err := UnmarshalComposerSchema(bytes)
//    bytes, err = composerSchema.Marshal()

package schema

import "encoding/json"

func UnmarshalComposerSchema(data []byte) (Schema, error) {
	var r Schema
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Schema) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Schema struct {
	Comment             *Comment               `json:"_comment"`          // A key to store comments in
	Abandoned           *Abandoned             `json:"abandoned"`         // Indicates whether this package has been abandoned, it can be boolean or a package; name/URL pointing to a recommended alternative. Defaults to false.
	Archive             *ComposerSchemaArchive `json:"archive,omitempty"` // Options for creating package archives for distribution.
	Authors             []Author               `json:"authors,omitempty"`
	Autoload            *Autoload              `json:"autoload,omitempty"`
	AutoloadDev         *AutoloadDev           `json:"autoload-dev,omitempty"`   // Description of additional autoload rules for development purpose (eg. a test suite).
	Bin                 *Bin                   `json:"bin"`                      // A set of files, or a single file, that should be treated as binaries and symlinked into; bin-dir (from config).
	Config              *Config                `json:"config,omitempty"`         // Composer options.
	Conflict            map[string]string      `json:"conflict,omitempty"`       // This is an object of package name (keys) and version constraints (values) that conflict; with this package.
	DefaultBranch       *bool                  `json:"default-branch,omitempty"` // Internal use only, do not specify this in composer.json. Indicates whether this version; is the default branch of the linked VCS repository. Defaults to false.
	Description         *string                `json:"description,omitempty"`    // Short package description.
	Dist                *Dist                  `json:"dist,omitempty"`
	Extra               *Extra                 `json:"extra"`                  // Arbitrary extra data that can be used by plugins, for example, package of type; composer-plugin may have a 'class' key defining an installer class name.
	Funding             []Funding              `json:"funding,omitempty"`      // A list of options to fund the development and maintenance of the package.
	Homepage            *string                `json:"homepage,omitempty"`     // Homepage URL for the project.
	IncludePath         []string               `json:"include-path,omitempty"` // DEPRECATED: A list of directories which should get added to PHP's include path. This is; only present to support legacy projects, and all new code should preferably use; autoloading.
	Keywords            []string               `json:"keywords,omitempty"`
	License             *Comment               `json:"license"`                        // License name. Or an array of license names.
	MinimumStability    *MinimumStability      `json:"minimum-stability,omitempty"`    // The minimum stability the packages must have to be install-able. Possible values are:; dev, alpha, beta, RC, stable.
	Name                *string                `json:"name,omitempty"`                 // Package name, including 'vendor-name/' prefix.
	NonFeatureBranches  []string               `json:"non-feature-branches,omitempty"` // A set of string or regex patterns for non-numeric branch names that will not be handled; as feature branches.
	PreferStable        *bool                  `json:"prefer-stable,omitempty"`        // If set to true, stable packages will be preferred to dev packages when possible, even if; the minimum-stability allows unstable packages.
	Provide             map[string]string      `json:"provide,omitempty"`              // This is an object of package name (keys) and version constraints (values) that this; package provides in addition to this package's name.
	Readme              *string                `json:"readme,omitempty"`               // Relative path to the readme document.
	Replace             map[string]string      `json:"replace,omitempty"`              // This is an object of package name (keys) and version constraints (values) that can be; replaced by this package.
	Repositories        *Repositories          `json:"repositories"`                   // A set of additional repositories where packages can be found.
	Require             map[string]string      `json:"require,omitempty"`              // This is an object of package name (keys) and version constraints (values) that are; required to run this package.
	RequireDev          map[string]string      `json:"require-dev,omitempty"`          // This is an object of package name (keys) and version constraints (values) that this; package requires for developing it (testing tools and such).
	Scripts             *Scripts               `json:"scripts,omitempty"`              // Script listeners that will be executed before/after some events.
	ScriptsDescriptions map[string]string      `json:"scripts-descriptions,omitempty"` // Descriptions for custom commands, shown in console help.
	Source              *Source                `json:"source,omitempty"`
	Suggest             map[string]string      `json:"suggest,omitempty"` // This is an object of package name (keys) and descriptions (values) that this package; suggests work well with it (this will be suggested to the user during installation).
	Support             *Support               `json:"support,omitempty"`
	TargetDir           *string                `json:"target-dir,omitempty"` // DEPRECATED: Forces the package to be installed into the given subdirectory path. This is; used for autoloading PSR-0 packages that do not contain their full path. Use forward; slashes for cross-platform compatibility.
	Time                *string                `json:"time,omitempty"`       // Package release date, in 'YYYY-MM-DD', 'YYYY-MM-DD HH:MM:SS' or 'YYYY-MM-DDTHH:MM:SSZ'; format.
	Type                *string                `json:"type,omitempty"`       // Package type, either 'library' for common packages, 'composer-plugin' for plugins,; 'metapackage' for empty packages, or a custom type ([a-z0-9-]+) defined by whatever; project this package applies to.
	Version             *string                `json:"version,omitempty"`    // Package version, see https://getcomposer.org/doc/04-schema.md#version for more info on; valid schemes.
}
