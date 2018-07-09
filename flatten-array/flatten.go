package flatten

func Flatten(array interface{}) []interface{} {
	var result []interface{}

	switch x := array.(type) {
	case int:
		result = append(result, x)
	case []interface{}:
		for _, v := range x {
			result = append(result, Flatten(v)...)
		}
	}
	if result == nil {
		return []interface{}{}
	}
	return result
}
