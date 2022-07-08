package schema

type FluffyPackage struct {
	InlinePackage      *InlinePackage
	InlinePackageArray []InlinePackage
}

func (x *FluffyPackage) UnmarshalJSON(data []byte) error {
	x.InlinePackageArray = nil
	x.InlinePackage = nil
	var c InlinePackage
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.InlinePackageArray, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.InlinePackage = &c
	}
	return nil
}

func (x *FluffyPackage) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.InlinePackageArray != nil, x.InlinePackageArray, x.InlinePackage != nil, x.InlinePackage, false, nil, false, nil, false)
}
