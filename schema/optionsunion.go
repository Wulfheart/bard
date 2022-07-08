package schema

type OptionsUnion struct {
	Bool         *bool
	OptionsClass *OptionsClass
}

func (x *OptionsUnion) UnmarshalJSON(data []byte) error {
	x.OptionsClass = nil
	var c OptionsClass
	object, err := unmarshalUnion(data, nil, nil, &x.Bool, nil, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.OptionsClass = &c
	}
	return nil
}

func (x *OptionsUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, x.Bool, nil, false, nil, x.OptionsClass != nil, x.OptionsClass, false, nil, false, nil, false)
}
