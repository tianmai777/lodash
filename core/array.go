package lodash

import (
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/sets/treeset"
	"github.com/emirpasic/gods/utils"
	"strings"
)

// chunk array to list array with size
func Chunk(array []interface{}, size int) [][]interface{} {
	totalCount := len(array)
	if totalCount <= size {
		return [][]interface{}{array[0:totalCount]}
	}
	var result [][]interface{}
	for i := 0; i < totalCount; i += size {
		if i+size >= totalCount {
			result = append(result, array[i:totalCount])
		} else {
			result = append(result, array[i:i+size])
		}
	}
	return result
}

// remove difference values from array, complex object use DifferenceBy or DifferenceWith
func Difference(array []interface{}, values ...interface{}) []interface{} {
	hset := hashset.New(array...)
	hset.Remove(values...)
	return hset.Values()
}

// remove difference values from array by field name or func
func DifferenceBy(array []interface{}, values []interface{}, iteratee interface{}) []interface{} {
	var op func(interface{}) interface{}
	switch iteratee.(type) {
	case string:
		op = func(i interface{}) interface{} {
			var fieldName = strings.Title(iteratee.(string))
			return findFieldVal(i, fieldName)
		}
	case func(interface{}) interface{}:
		op = iteratee.(func(interface{}) interface{})
	}

	var newArray = Map(array, op)
	var newValues = Map(values, op)

	diff := Difference(newArray, newValues...)
	diffSet := hashset.New(diff...)
	result := Filter(array, func(item interface{}) bool {
		return diffSet.Contains(op(item))
	})
	return result
}

// remove difference values from array by comparator func
func DifferenceWith(array []interface{}, values []interface{}, comparator utils.Comparator) []interface{} {
	var aSet = treeset.NewWith(comparator)
	aSet.Add(array...)
	aSet.Remove(values...)
	return aSet.Values()
}

// find intersection from arrays, complex object use IntersectionBy or IntersectionWith
func Intersection(arrays ...[]interface{}) []interface{} {
	hset := hashset.New()
	for i, array := range arrays {
		if i == 0 {
			hset.Add(array...)
		} else {
			intersectionSet := hashset.New()
			for _, a := range array {
				if hset.Contains(a) {
					intersectionSet.Add(a)
				}
			}
			hset = intersectionSet
		}
	}
	return hset.Values()
}

// find intersection by field name or func
func IntersectionBy(array []interface{}, values []interface{}, iteratee interface{}) []interface{} {
	diffLeft := DifferenceBy(array, values, iteratee)
	result := DifferenceBy(array, diffLeft, iteratee)
	return result
}

// find intersection by by comparator func
func IntersectionWith(array []interface{}, values []interface{}, comparator utils.Comparator) []interface{} {
	diffLeft := DifferenceWith(array, values, comparator)
	result := DifferenceWith(array, diffLeft, comparator)
	return result
}
