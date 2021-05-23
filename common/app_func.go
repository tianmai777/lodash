package common

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// str => md5
func Md5(str string) string {
	result := md5.Sum([]byte(str))
	return string(hex.EncodeToString(result[:]))
}

// export func and const to gen_export,exclude (*_test,)
func export() {
	var funcList, varList []string
	root := "../"
	er := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(info.Name(), ".go") || info.Name() == "gen_export.go" ||
			strings.HasSuffix(info.Name(), "_test.go") {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		findConst := false
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lineText := scanner.Text()
			if strings.HasPrefix(lineText, "func ") {
				idx := strings.Index(lineText, "(")
				str := lineText[5:idx]
				pkg := path[len(root):]
				pidx := strings.Index(pkg, "/")
				funcList = append(funcList, pkg[:pidx]+"."+str)
			}

			if findConst {
				lineText = strings.TrimLeft(lineText, "\t")
				idx := strings.Index(lineText, " ")
				if idx == -1 {
					continue
				}
				pkg := path[len(root):]
				pidx := strings.Index(pkg, "/")
				varList = append(varList, pkg[:pidx]+"."+lineText[:idx])
			}
			if strings.HasPrefix(lineText, "const (") {
				findConst = true
			} else if strings.HasPrefix(lineText, ")") {
				findConst = false
			}
		}
		return nil
	})
	if er != nil {
		log.Fatal(er)
	}
	sort.Strings(varList)
	sort.Strings(funcList)
	exportWrite(append(varList, funcList...))
}

// exportWrite to gen_export
func exportWrite(items []string) {
	headStr := []string{
		"package lodash",
		"import \"github.com/jakesally/lodash/common\"",
		"import \"github.com/jakesally/lodash/core\"",
		"\r\n",
	}
	genStr := strings.Join(headStr, "\r\n")
	for _, item := range items {
		pkgIdx := strings.Index(item, ".")
		itemName := item[pkgIdx+1:]
		if itemName[0] >= 'A' && itemName[0] <= 'Z' {
			genStr += fmt.Sprintf("var %v = %v\r\n", itemName, item)
		}
	}
	ioutil.WriteFile("../gen_export.go", []byte(genStr), 0644)
}
