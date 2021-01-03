package equality

import "reflect"

func JsonObjectsEquals(obj1, obj2 interface{}) bool {
	if reflect.TypeOf(obj1) != reflect.TypeOf(obj2) {
		return false
	}

	switch obj1.(type) {
	case map[string]interface{}:
		return JsonMapsEquals(obj1.(map[string]interface{}), obj2.(map[string]interface{}))
	case []interface{}:
		return JsonArraysEquals(obj1.([]interface{}), obj2.([]interface{}))
	default:
		return obj1 == obj2
	}
}

func JsonMapsEquals(map1, map2 map[string]interface{}) bool {
	if len(map1) != len(map2) {
		return false
	}

	for k, v := range map1 {
		val2 := map2[k]

		if (v == nil) != (val2 == nil) {
			return false
		}

		if !JsonObjectsEquals(v, val2) {
			return false
		}
	}

	return true
}

func JsonArraysEquals(array1, array2 []interface{}) bool {

	if len(array1) != len(array2) {
		return false
	}

	var matches int
	flagged := make([]bool, len(array2))
	for _, v := range array1 {
		for i, v2 := range array2 {
			if JsonObjectsEquals(v, v2) && !flagged[i] {
				matches++
				flagged[i] = true

				break
			}
		}
	}

	return matches == len(array1)

}
