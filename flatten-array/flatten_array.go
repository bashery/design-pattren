package flatten

func Flatten(array interface{}) []interface{} {
	slc := make([]interface{}, 0)
	for _, v := range array.([]interface{}) {
		switch v.(type) {
		case []interface{}:
			slc = append(slc, Flatten(v)...)
		case int:
			slc = append(slc, v)
		}
	}
	return slc
}
