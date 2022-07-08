package schema

// Arbitrary extra data that can be used by plugins, for example, package of type
// composer-plugin may have a 'class' key defining an installer class name.
type Extra struct {
	AnythingArray []interface{}
	AnythingMap   map[string]interface{}
}

func (x *Extra) UnmarshalJSON(data []byte) error {
	x.AnythingArray = nil
	x.AnythingMap = nil
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.AnythingArray, false, nil, true, &x.AnythingMap, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Extra) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.AnythingArray != nil, x.AnythingArray, false, nil, x.AnythingMap != nil, x.AnythingMap, false, nil, false)
}
