package schema

type GitlabTokenValue struct {
	GitlabTokenClass *GitlabTokenClass
	String           *string
}

func (x *GitlabTokenValue) UnmarshalJSON(data []byte) error {
	x.GitlabTokenClass = nil
	var c GitlabTokenClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.GitlabTokenClass = &c
	}
	return nil
}

func (x *GitlabTokenValue) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.GitlabTokenClass != nil, x.GitlabTokenClass, false, nil, false, nil, false)
}
