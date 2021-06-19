package ldc

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/dgrijalva/jwt-go"
	json "github.com/json-iterator/go"
	"log"
	"math"
	"runtime"
	"strconv"
	"strings"
	"text/template"
	"time"
)

// location for project,file,line,method
func Location() string {
	pc, file, line, _ := runtime.Caller(1)
	methodPath := runtime.FuncForPC(pc).Name()
	fileIdx := strings.Index(file, "src/") + 4
	fileSubStr := file[fileIdx:]
	fileIdx = strings.Index(fileSubStr, "/") + 1
	methodIdx := strings.LastIndex(methodPath, ".") + 1
	return fmt.Sprintf("%v %v:%v:1 %v", fileSubStr[:fileIdx-1], fileSubStr[fileIdx:], line, methodPath[methodIdx:])
}

// format unix time second, now => time.Now().Unix()
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

// to jwt silence
func ToJwt(secret string, args map[string]interface{}) string {
	// jwt.StandardClaims{}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(args))
	str, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Printf("%v err > %v", Location(), err)
	}
	return str
}

// to json silence
func ToJson(obj interface{}) string {
	result, err := ToJsonErr(&obj)
	if err != nil {
		log.Printf("%v err > %v", Location(), err)
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
		log.Printf("%v err > %v", Location(), err)
	}
	return i
}

// parse int with err
func ParseIntErr(str string) (int, error) {
	return strconv.Atoi(str)
}

// parse jwt
func ParseJwt(secret string, jwtStr string) (map[string]interface{}, error) {
	token, err := jwt.Parse(jwtStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method err > %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("jwt not valid")
	}
}

// parse object silence
func ParseObj(str string, obj interface{}) {
	err := ParseObjErr(str, obj)
	if err != nil {
		log.Printf("%v err > %v", Location(), err)
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
		log.Printf("%v err > %v", Location(), err)
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
		log.Printf("%v err > %v", Location(), err)
	}
	return str
}
