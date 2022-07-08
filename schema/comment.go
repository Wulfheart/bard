package schema

// A key to store comments in
//
// License name. Or an array of license names.
//
// Occurs after the autoloader is dumped, contains one or more Class::method callables or
// shell commands.
//
// Occurs after the create-project command is executed, contains one or more Class::method
// callables or shell commands.
//
// Occurs after the install command is executed, contains one or more Class::method
// callables or shell commands.
//
// Occurs after a package is installed, contains one or more Class::method callables or
// shell commands.
//
// Occurs after a package has been uninstalled, contains one or more Class::method callables
// or shell commands.
//
// Occurs after a package is updated, contains one or more Class::method callables or shell
// commands.
//
// Occurs after the root-package is installed, contains one or more Class::method callables
// or shell commands.
//
// Occurs after the status command is executed, contains one or more Class::method callables
// or shell commands.
//
// Occurs after the update command is executed, contains one or more Class::method callables
// or shell commands.
//
// Occurs before the autoloader is dumped, contains one or more Class::method callables or
// shell commands.
//
// Occurs before the install command is executed, contains one or more Class::method
// callables or shell commands.
//
// Occurs before a package is installed, contains one or more Class::method callables or
// shell commands.
//
// Occurs before a package has been uninstalled, contains one or more Class::method
// callables or shell commands.
//
// Occurs before a package is updated, contains one or more Class::method callables or shell
// commands.
//
// Occurs before the status command is executed, contains one or more Class::method
// callables or shell commands.
//
// Occurs before the update command is executed, contains one or more Class::method
// callables or shell commands.
type Comment struct {
	AnythingArray []interface{}
	String        *string
}

func (x *Comment) UnmarshalJSON(data []byte) error {
	x.AnythingArray = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.AnythingArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Comment) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.AnythingArray != nil, x.AnythingArray, false, nil, false, nil, false, nil, false)
}
