package ldc

import "testing"

func TestToSliceTP(t *testing.T) {
	source := []interface{}{1, 2, 3, 4, 5}
	dest := ToSliceTP(source)
	if dest[0] != TP(1) {
		t.Errorf("TestToSliceTP fail dest>%v", dest)
	}
}

func TestToSliceObjByTP(t *testing.T) {
	source := []interface{}{1, 2, 3, 4, 5}
	dest := ToSliceByTP(ToSliceTP(source))
	if dest[0] != 1 {
		t.Errorf("TestToSliceObjByTP fail dest>%v", dest)
	}
}

func TestIfTP(t *testing.T) {
	dest := IfTP(true, TP(1), TP(2))
	if dest != TP(1) {
		t.Errorf("TestIfTP fail dest>%v", dest)
	}
	dest = IfTP(false, TP(1), TP(2))
	if dest != TP(2) {
		t.Errorf("TestIfTP fail dest>%v", dest)
	}
}
