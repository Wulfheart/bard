package schema

// A set of additional repositories where packages can be found.
type Repositories struct {
	RepositoryElementArray []RepositoryElement
	UnionMap               map[string]*RepositoryValue
}

func (x *Repositories) UnmarshalJSON(data []byte) error {
	x.RepositoryElementArray = nil
	x.UnionMap = nil
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.RepositoryElementArray, false, nil, true, &x.UnionMap, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Repositories) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.RepositoryElementArray != nil, x.RepositoryElementArray, false, nil, x.UnionMap != nil, x.UnionMap, false, nil, false)
}
