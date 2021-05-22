package lodash

import "testing"

func TestMd5(t *testing.T) {
	dest := Md5("1")
	if dest != "c4ca4238a0b923820dcc509a6f75849b" {
		t.Errorf("TestMd5 fail dest>%v", dest)
	}
}
