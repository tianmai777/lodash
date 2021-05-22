package lodash

import (
	"github.com/emirpasic/gods/lists/arraylist"
	. "github.com/jakesally/lodash/common"
	"testing"
)

type Node struct {
	Id       string
	Name     string
	ParentId string
	Children []*Node
}

// to slice object by Node
func ToSliceObjByNode(items []*Node) []interface{} {
	var result = make([]interface{}, 0, len(items))
	for _, item := range items {
		result = append(result, item)
	}
	return result
}

func TestFindFieldVal(t *testing.T) {
	aNode := &Node{Id: "1", Name: "x"}
	dest := findFieldVal(aNode, Id)
	if dest != "1" {
		t.Errorf("TestFindFieldVal fail >%v", dest)
	}
}

func TestMap(t *testing.T) {
	// 1.[]node
	var source = arraylist.New(Node{Id: "1", Name: "x"}, Node{Id: "2", Name: "y"})
	dest := Map(source.Values(), Id)
	if source.Size() != len(dest) {
		t.Errorf("map []node id fail >%v,%v", source, dest)
	}
	if dest[0] != "1" {
		t.Errorf("map []node id fail dest[0] >%v", dest[0])
	}

	//2.[]*node
	source = arraylist.New(&Node{Id: "1", Name: "x"}, &Node{Id: "2", Name: "y"})
	dest = Map(source.Values(), Id)
	if source.Size() != len(dest) {
		t.Errorf("map []*node id fail >%v,%v", source, dest)
	}
	if dest[0] != "1" {
		t.Errorf("map []*node id fail dest[0] >%v", dest[0])
	}

	//3.func
	dest = Map(source.Values(), func(item interface{}) interface{} {
		i := item.(*Node)
		i.Name = i.Id + i.Name
		return i.Name
	})
	if source.Size() != len(dest) {
		t.Errorf("map []*node func fail >%v,%v", source, dest)
	}
	if dest[0] != "1x" {
		t.Errorf("map []*node func fail dest[0] >%v", dest[0])
	}
}

func TestFilter(t *testing.T) {
	// 1.func
	nodes := []*Node{&Node{Id: "1", Name: "x"}, &Node{Id: "2", Name: "y"}}
	dest := Filter(ToSliceObjByNode(nodes), func(item interface{}) bool {
		i := item.(*Node)
		return i.Name == "y"
	})
	if (dest[0].(*Node)).Name != "y" {
		t.Errorf("filter []*node func fail dest[0] >%v", dest[0])
	}
	// 2.map
	dest = Filter(ToSliceObjByNode(nodes), map[interface{}]interface{}{
		Name: "x",
	})
	if (dest[0].(*Node)).Name != "x" {
		t.Errorf("filter []*node map field fail dest[0] >%v", dest[0])
	}
}
