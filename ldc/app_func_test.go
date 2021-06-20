package ldc

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

// echo 'ldc/app_func_test.go:13:1: TestLocation'
func TestLocation(t *testing.T) {
	loc := Location()
	assert.Equal(t, "lodash ldc/app_func_test.go:12:1 TestLocation", loc)
}

func TestFormatTime(t *testing.T) {
	// dest = FormatTime(time.Now().Unix(), ADateTime)
	dest := FormatTime(1624077283, ADateTime)
	assert.Equal(t, "2021-06-19 12:34:43", dest)
}

func TestMd5(t *testing.T) {
	dest := Md5("1")
	if dest != "c4ca4238a0b923820dcc509a6f75849b" {
		t.Errorf("TestMd5 fail dest>%v", dest)
	}
}

func TestToStr(t *testing.T) {
	tests := []struct {
		In  interface{}
		Out string
	}{
		{int(1), "1"},
		{int8(1), "1"},
		{int16(1), "1"},
		{int32(1), "1"},
		{int64(1), "1"},
		{uint(1), "1"},
		{uint8(1), "1"},
		{uint16(1), "1"},
		{uint32(1), "1"},
		{nil, ""},
		{bool(true), "true"},
	}
	for _, te := range tests {
		if ToStr(te.In) != te.Out {
			t.Errorf("TestToStr fail dest>%v,%v", te.In, te.Out)
		}
	}
}

func TestParseInt(t *testing.T) {
	i := ParseInt("666")
	if i != 666 {
		t.Errorf("TestParseInt fail dest>%v", i)
	}
	i = ParseInt("666x")
	assert.Equal(t, 0, i)
}

func TestParseObj(t *testing.T) {
	var aArr []int
	ParseObj("[1,2]", &aArr)
	assert.Contains(t, aArr, 1)
	var bArr []int
	ParseObj("[1,2", &bArr)
}

func TestParseStrMap(t *testing.T) {
	r := ParseStrMap("{\"Id\":\"123\"}")
	if r["Id"] != "123" {
		t.Errorf("TestParseStrMap err > %v", r["id"])
	}
}

func TestParseStrObjMap(t *testing.T) {
	r := ParseStrObjMap("{\"Id\":\"123\"}")
	if r["Id"] != "123" {
		t.Errorf("TestParseStrMap err > %v", r["id"])
	}
}

func TestToJson(t *testing.T) {
	str := ToJson([]int{1, 2})
	assert.Equal(t, str, "[1,2]")
	str = ToJson(math.NaN())
	assert.Equal(t, "", str)
}

func TestToJwt(t *testing.T) {
	secret := "jake"
	jwtStr := ToJwt(secret, map[string]interface{}{
		"test": "only",
	})
	assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZXN0Ijoib25seSJ9.rQyDU5XrUPSlrD5Et_ThO9HqKt0UpZC-g6bB1Hhn5dY",
		jwtStr)
	mc, err := ParseJwt(secret, jwtStr)
	assert.NoError(t, err)
	assert.Equal(t, "only", mc["test"])
}

func TestToPercent(t *testing.T) {
	str := ToPercent(0.123, 3, "%")
	assert.Equal(t, str, "12.300%")
}

func TestParseTime(t *testing.T) {
	r := ParseTime("2021-06-12 18:06:18")
	assert.Equal(t, 1623492378, r)
}

func TestTimeSegment(t *testing.T) {
	r := TimeSegment("2021-06-12 00:00:00", "2021-06-20 00:00:00", 24*60*60)
	assert.Contains(t, r, "2021-06-12 00:00:00")
	assert.Contains(t, r, "2021-06-13 00:00:00")
	assert.Contains(t, r, "2021-06-20 00:00:00")
}

func TestToTree(t *testing.T) {
	var aList []*Node
	aList = append(aList, &Node{Id: "1", ParentId: "0", Name: "节点1"})
	aList = append(aList, &Node{Id: "2", ParentId: "0", Name: "节点2"})
	aList = append(aList, &Node{Id: "3", ParentId: "0", Name: "节点3"})
	aList = append(aList, &Node{Id: "4", ParentId: "1", Name: "节点4"})
	aList = append(aList, &Node{Id: "5", ParentId: "1", Name: "节点5"})
	aList = append(aList, &Node{Id: "6", ParentId: "1", Name: "节点6"})
	aList = append(aList, &Node{Id: "7", ParentId: "4", Name: "节点7"})
	aList = append(aList, &Node{Id: "8", ParentId: "4", Name: "节点8"})
	aList = append(aList, &Node{Id: "9", ParentId: "5", Name: "节点9"})

	aTree := ToTree(aList)
	assert.Equal(t, 3, len(aTree))
	// fmt.Println(ToJson(aTree))
}

func TestTemplateGen(t *testing.T) {
	a := map[string]interface{}{
		"Name": "jk",
		"Age":  "30",
	}
	r := TemplateGen("{{.Name}} today is {{.Age}}", a)
	assert.Equal(t, "jk today is 30", r)
	r = TemplateGen("{{.Name}} today is {{.Age}},{{", a)
	assert.Equal(t, "", r)
}

func TestParseHtml(t *testing.T) {
	htmlStr := "<html><body><a href=\"mpInfo+\">mpinfo</a><a href=\"mpInfo+\">mpinfo2</a></body></html>"
	html := ParseHtml(htmlStr)
	var items []string
	html.Find("a").Each(func(i int, selection *goquery.Selection) {
		items = append(items, selection.Text())
	})
	assert.Equal(t, 2, len(items))
	assert.Contains(t, items, "mpinfo2")
}
