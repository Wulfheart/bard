package schema

type PurplePackage struct {
	Bool               *bool
	InlinePackage      *InlinePackage
	InlinePackageArray []InlinePackage
}

func (x *PurplePackage) UnmarshalJSON(data []byte) error {
	x.InlinePackageArray = nil
	x.InlinePackage = nil
	var c InlinePackage
	object, err := unmarshalUnion(data, nil, nil, &x.Bool, nil, true, &x.InlinePackageArray, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.InlinePackage = &c
	}
	return nil
}

func (x *PurplePackage) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, x.Bool, nil, x.InlinePackageArray != nil, x.InlinePackageArray, x.InlinePackage != nil, x.InlinePackage, false, nil, false, nil, false)
}
