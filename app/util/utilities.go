package util

func StringArrayInsert(src []string, index int, inserted string) []string {
	_first := make([]string, index)
	copy(_first, src)
	last := src[index:]
	first := append(_first, inserted)
	return append(first, last...)
}

func StrigArrayRemove(src []string, item string) [] string {
	index := SliceIndex(src, item)
	if len(src) > 0 && index >= 0 {
		return append(src[:index], src[index+1:]...);
	}

	return nil
 }

 func StringArrayRemoveByIndex(src []string, index int) [] string {
	 if len(src) > 0 && index >= 0 {
		 return append(src[:index], src[index+1:]...);
	 }

	 return nil
 }

func SliceIndex(slice []string, item string) int {
	for i, _ := range slice {
		if slice[i] == item {
			return i
		}
	}
	return -1
}
