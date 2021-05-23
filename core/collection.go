package core

import (
	"reflect"
)

// find item field value by field name
func findFieldVal(item interface{}, fieldName string) interface{} {
	return reflect.Indirect(reflect.ValueOf(item)).FieldByName(fieldName).Interface()
}

// map collection use field name or func
func Map(collection []interface{}, iteratee interface{}) []interface{} {
	var op func(interface{}) interface{}
	switch iteratee.(type) {
	case string:
		op = func(i interface{}) interface{} {
			return findFieldVal(i, iteratee.(string))
		}
	case func(interface{}) interface{}:
		op = iteratee.(func(interface{}) interface{})
	}
	var result []interface{}
	for _, i := range collection {
		result = append(result, op(i))
	}
	return result
}

// filter collection use map field or func
func Filter(collection []interface{}, predicate interface{}) []interface{} {
	var op func(interface{}) bool
	switch predicate.(type) {
	case map[interface{}]interface{}:
		conditionMap := predicate.(map[interface{}]interface{})
		op = func(i interface{}) bool {
			every := true
			for k, v := range conditionMap {
				every = every && (v == findFieldVal(i, k.(string)))
				if !every {
					break
				}
			}
			return every
		}
	case func(interface{}) bool:
		op = predicate.(func(interface{}) bool)
	}
	var result []interface{}
	for _, i := range collection {
		if op(i) {
			result = append(result, i)
		}
	}
	return result
}
