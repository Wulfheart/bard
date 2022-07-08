package schema

// Indicates whether this package has been abandoned, it can be boolean or a package
// name/URL pointing to a recommended alternative. Defaults to false.
//
// The default style of handling dirty updates, defaults to false and can be any of true,
// false or "stash".
//
// Defaults to "php-only" which checks only the PHP version. Setting to true will also check
// the presence of required PHP extensions. If set to false, Composer will not create and
// require a platform_check.php file as part of the autoloader bootstrap.
//
// What to do after prompting for authentication, one of: true (store), false (do not store)
// or "prompt" (ask every time), defaults to prompt.
//
// When running Composer in a directory where there is no composer.json, if there is one
// present in a directory above Composer will by default ask you whether you want to use
// that directory's composer.json instead. One of: true (always use parent if needed), false
// (never ask or use it) or "prompt" (ask every time), defaults to prompt.
type Abandoned struct {
	Bool   *bool
	String *string
}

func (x *Abandoned) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, nil, nil, &x.Bool, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Abandoned) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, x.Bool, x.String, false, nil, false, nil, false, nil, false, nil, false)
}
