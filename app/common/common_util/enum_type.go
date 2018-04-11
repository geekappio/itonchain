package common_util

// Define an enumeration interface
type Enum interface {
	ValueOf(value string) *EnumType
}

// Define an enumeration type
type EnumType struct {
	Value string
}

// Save all defined enumeration types.
var enumTypeMap = make(map[string]*EnumType)

// Return enumeration type object by its value
func (*EnumType) ValueOf(value string) *EnumType {
	return enumTypeMap[value]
}

// Define a new enumeration type with specified value
func DefEnumType(value string) (enumType *EnumType) {
	if enumTypeMap[value] != nil {
		return enumTypeMap[value]
	} else {
		enumType = &EnumType{Value: value}
		enumTypeMap[value] = enumType
	}
	return
}

