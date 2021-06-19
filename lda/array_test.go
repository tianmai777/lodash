package lda

import (
	. "github.com/jakesally/lodash/ldc"
	"testing"
)

func TestChunk(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	dest := Chunk(ToSliceByInt(source), 10)
	if len(dest) != 1 {
		t.Errorf("TestChunk fail dest>%v", dest)
	}
	dest = Chunk(ToSliceByInt(source), 2)
	if len(dest) != 3 {
		t.Errorf("TestChunk fail dest>%v", dest)
	}
	dest1 := ToSliceInt(dest[1])
	if dest1[1] != 4 {
		t.Errorf("TestChunk fail dest>%v", dest)
	}
	source = []int{1, 2, 3, 4}
	dest = Chunk(ToSliceByInt(source), 2)
	if len(dest) != 2 {
		t.Errorf("TestChunk fail dest>%v", dest)
	}
}

func TestDifference(t *testing.T) {
	dest := Difference([]interface{}{"1", "3", "5"}, "1", "5")
	if dest[0] != "3" {
		t.Errorf("TestDifference fail dest>%v", dest)
	}
}

func TestDifferenceBy(t *testing.T) {
	array := []interface{}{&Node{Id: "1", Name: "x"}, &Node{Id: "2", Name: "y"}}
	values := []interface{}{&Node{Id: "1", Name: "x"}}
	diff := DifferenceBy(array, values, Id)
	dest := diff[0].(*Node)
	if dest.Name != "y" {
		t.Errorf("DifferenceBy fail dest>%v", dest)
	}
	diff = DifferenceBy(array, values, func(a interface{}) interface{} {
		a1 := a.(*Node)
		return a1.Id
	})
	dest = diff[0].(*Node)
	if dest.Name != "y" {
		t.Errorf("DifferenceBy fail dest>%v", dest)
	}
}

func TestDifferenceWith(t *testing.T) {
	array := []interface{}{&Node{Id: "1", Name: "x"}, &Node{Id: "2", Name: "y"}}
	values := []interface{}{&Node{Id: "1", Name: "x"}}
	diff := DifferenceWith(array, values, func(a, b interface{}) int {
		a1 := a.(*Node)
		b1 := b.(*Node)
		if a1.Id > b1.Id {
			return 1
		} else if a1.Id == b1.Id {
			return 0
		} else {
			return -1
		}
	})
	dest := diff[0].(*Node)
	if dest.Name != "y" {
		t.Errorf("DifferenceWith fail dest>%v", dest)
	}
}

func TestIntersection(t *testing.T) {
	// 1.int
	dest := Intersection([]interface{}{"1", "3", "5"}, []interface{}{"3"})
	if dest[0] != "3" {
		t.Errorf("TestIntersection fail dest>%v", dest)
	}
}

func TestIntersectionBy(t *testing.T) {
	array := []interface{}{&Node{Id: "1", Name: "x"}, &Node{Id: "2", Name: "y"}}
	values := []interface{}{&Node{Id: "1", Name: "x"}}
	inter := IntersectionBy(array, values, Id)
	dest := inter[0].(*Node)
	if dest.Name != "x" {
		t.Errorf("IntersectionBy fail dest>%v", dest)
	}
}

func TestIntersectionWith(t *testing.T) {
	array := []interface{}{&Node{Id: "1", Name: "x"}, &Node{Id: "2", Name: "y"}}
	values := []interface{}{&Node{Id: "1", Name: "x"}}
	inter := IntersectionWith(array, values, func(a, b interface{}) int {
		a1 := a.(*Node)
		b1 := b.(*Node)
		if a1.Id > b1.Id {
			return 1
		} else if a1.Id == b1.Id {
			return 0
		} else {
			return -1
		}
	})
	dest := inter[0].(*Node)
	if dest.Name != "x" {
		t.Errorf("IntersectionWith fail dest>%v", dest)
	}
}
