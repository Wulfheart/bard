package schema

// This is an object of {"pattern": true|false} with packages which are allowed to be loaded
// as plugins, or true to allow all, false to allow none. Defaults to {} which prompts when
// an unknown plugin is added.
type AllowPlugins struct {
	Bool    *bool
	BoolMap map[string]bool
}

func (x *AllowPlugins) UnmarshalJSON(data []byte) error {
	x.BoolMap = nil
	object, err := unmarshalUnion(data, nil, nil, &x.Bool, nil, false, nil, false, nil, true, &x.BoolMap, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *AllowPlugins) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, x.Bool, nil, false, nil, false, nil, x.BoolMap != nil, x.BoolMap, false, nil, false)
}
