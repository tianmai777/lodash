package lodash

import "github.com/cheekybits/genny/generic"

type TP generic.Type

//go:generate genny -in=$GOFILE -out=gen_$GOFILE gen "TP=BUILTINS"

// to slice TP
func ToSliceTP(items []interface{}) []TP {
	var result = make([]TP, 0, len(items))
	for _, item := range items {
		result = append(result, item.(TP))
	}
	return result
}

// to slice object by TP
func ToSliceObjByTP(items []TP) []interface{} {
	var result = make([]interface{}, 0, len(items))
	for _, item := range items {
		result = append(result, item)
	}
	return result
}

// if sta return tru else return fal
func IfTP(sta bool, tru TP, fal TP) TP {
	if sta {
		return tru
	} else {
		return fal
	}
}
