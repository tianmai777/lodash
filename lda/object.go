package lda

// values from map only support AMapStrStr,AMapStrInf
func Values(obj interface{}) []interface{} {
	var result []interface{}
	switch obj.(type) {
	case map[string]string:
		aMap := obj.(map[string]string)
		for _, v := range aMap {
			result = append(result, v)
		}
	case map[string]interface{}:
		aMap := obj.(map[string]interface{})
		for _, v := range aMap {
			result = append(result, v)
		}
	}
	return result
}
