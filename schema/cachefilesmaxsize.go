package schema

// The cache max size for the files cache, defaults to "300MiB".
type CacheFilesMaxsize struct {
	Integer *int64
	String  *string
}

func (x *CacheFilesMaxsize) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, &x.Integer, nil, nil, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *CacheFilesMaxsize) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, nil, nil, x.String, false, nil, false, nil, false, nil, false, nil, false)
}
