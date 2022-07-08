package schema

// The install method Composer will prefer to use, defaults to auto and can be any of
// source, dist, auto, or an object of {"pattern": "preference"}.
type PreferredInstall struct {
	String    *string
	StringMap map[string]string
}

func (x *PreferredInstall) UnmarshalJSON(data []byte) error {
	x.StringMap = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, false, nil, true, &x.StringMap, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *PreferredInstall) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, false, nil, x.StringMap != nil, x.StringMap, false, nil, false)
}
