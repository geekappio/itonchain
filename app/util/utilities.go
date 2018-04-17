package util

type ArrayInterface interface {
	Insert(array []interface{}, index int, inserted interface{}) []interface{}
}

func StringArrayInsert(src []string, index int64, inserted string) []string{
	last := src[index:]
	first := append(src[0:index], inserted)
	return append(first, last...)
}