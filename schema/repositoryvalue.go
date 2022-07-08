package schema

type RepositoryValue struct {
	Bool       *bool
	Repository *Repository
}

func (x *RepositoryValue) UnmarshalJSON(data []byte) error {
	x.Repository = nil
	var c Repository
	object, err := unmarshalUnion(data, nil, nil, &x.Bool, nil, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.Repository = &c
	}
	return nil
}

func (x *RepositoryValue) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, x.Bool, nil, false, nil, x.Repository != nil, x.Repository, false, nil, false, nil, false)
}
