package schema

type Exclude struct {
	Bool        *bool
	StringArray []string
}

func (x *Exclude) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	object, err := unmarshalUnion(data, nil, nil, &x.Bool, nil, true, &x.StringArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Exclude) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, x.Bool, nil, x.StringArray != nil, x.StringArray, false, nil, false, nil, false, nil, false)
}
