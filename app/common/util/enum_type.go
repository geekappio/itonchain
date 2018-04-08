package util

type Enum interface {
	Value() string
	ValueOf(value string) *EnumType
}

type EnumType struct {
	Value string
}

var enumTypeMap map[string]*EnumType = make(map[string]*EnumType)

func (*EnumType) ValueOf(value string) *EnumType {
	return enumTypeMap[value]
}

func DefEnumType(value string) (enumType *EnumType) {
	enumType = &EnumType{Value: value}
	enumTypeMap[value] = enumType
	return
}

