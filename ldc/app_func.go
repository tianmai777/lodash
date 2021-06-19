package ldc

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	json "github.com/json-iterator/go"
	"log"
	"math"
	"strconv"
	"text/template"
	"time"
)

// format unix time second
func FormatTime(second int64, format string) string {
	var t = time.Unix(second, 0)
	return t.Format(format)
}

// str => md5
func Md5(str string) string {
	result := md5.Sum([]byte(str))
	return string(hex.EncodeToString(result[:]))
}

// to str from obj
func ToStr(obj interface{}) string {
	switch obj.(type) {
	case int:
		return strconv.Itoa(obj.(int))
	case int8:
		return strconv.Itoa(int(obj.(int8)))
	case int16:
		return strconv.Itoa(int(obj.(int16)))
	case int32:
		return strconv.Itoa(int(obj.(int32)))
	case int64:
		return strconv.Itoa(int(obj.(int64)))
	case uint:
		return strconv.Itoa(int(obj.(uint)))
	case uint8:
		return strconv.Itoa(int(obj.(uint8)))
	case uint16:
		return strconv.Itoa(int(obj.(uint16)))
	case uint32:
		return strconv.Itoa(int(obj.(uint32)))
	case nil:
		return ""
	default:
		return fmt.Sprintf("%v", obj)
	}
}

// to json silence
func ToJson(obj interface{}) string {
	result, err := ToJsonErr(&obj)
	if err != nil {
		log.Printf("ToJson err > %v", err)
	}
	return result
}

// to json with err
func ToJsonErr(obj interface{}) (string, error) {
	return json.ConfigCompatibleWithStandardLibrary.MarshalToString(&obj)
}

// ToPercent(0.123, 3, %) => 12.300%
func ToPercent(i float64, precision int, suffix string) string {
	fmtStr := "%." + ToStr(precision) + "f"
	return fmt.Sprintf(fmtStr, i*math.Pow10(2)) + suffix
}

// parse int silence
func ParseInt(str string) int {
	i, err := ParseIntErr(str)
	if err != nil {
		log.Printf("ParseInt err > %v", err)
	}
	return i
}

// parse int with err
func ParseIntErr(str string) (int, error) {
	return strconv.Atoi(str)
}

// parse object silence
func ParseObj(str string, obj interface{}) {
	err := ParseObjErr(str, obj)
	if err != nil {
		log.Printf("ParseObj err > %v", err)
	}
}

// parse object with err
func ParseObjErr(str string, obj interface{}) error {
	return json.ConfigCompatibleWithStandardLibrary.UnmarshalFromString(str, &obj)
}

// parse str map silence
func ParseStrMap(str string) map[string]string {
	var result map[string]string
	ParseObj(str, &result)
	return result
}

// parse str-obj map silence
func ParseStrObjMap(str string) map[string]interface{} {
	var result map[string]interface{}
	ParseObj(str, &result)
	return result
}

// to unix timestamp, format "2006-01-02 15:04:05"
func UnixTime(date string) int {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tp, _ := time.ParseInLocation(ADateTime, date, loc)
	return int(tp.Unix())
}

// TimeSegment, format 2006-01-02 15:04:05
func TimeSegment(start string, end string, seconds int) []string {
	st := UnixTime(start)
	ed := UnixTime(end)
	var result []string
	for st <= ed {
		result = append(result, time.Unix(int64(st), 0).Format(ADateTime))
		st += seconds
	}
	return result
}

type Node struct {
	Id       string
	Name     string
	ParentId string
	Children []*Node
	Ext      interface{}
}

// list to tree
func ToTree(aList []*Node) []*Node {
	var resultList []*Node
	for _, aNode := range aList {
		var findParentNode = false
		for _, anotherNode := range aList {
			if aNode.ParentId == anotherNode.Id {
				findParentNode = true
				anotherNode.Children = append(anotherNode.Children, aNode)
				break
			}
		}
		if !findParentNode {
			resultList = append(resultList, aNode)
		}
	}
	return resultList
}

// parse html
func ParseHtml(htmlStr string) *goquery.Document {
	doc, err := ParseHtmlErr(htmlStr)
	if err != nil {
		log.Printf("BuildDocument err > %v", err)
	}
	return doc
}

// parse html with err
func ParseHtmlErr(htmlStr string) (*goquery.Document, error) {
	bf := bytes.NewBufferString(htmlStr)
	return goquery.NewDocumentFromReader(bf)
}

// text template gen with err
func TemplateGenErr(text string, args map[string]interface{}) (string, error) {
	tmp, err := template.New("").Parse(text)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	err = tmp.Execute(buf, args)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// text template gen silence
func TemplateGen(text string, args map[string]interface{}) string {
	str, err := TemplateGenErr(text, args)
	if err != nil {
		log.Printf("BuildTemplate err > %v", err)
	}
	return str
}
