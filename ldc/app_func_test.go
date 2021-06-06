package ldc

import (
	"testing"
)

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
}

func TestParseIntErr(t *testing.T) {
	i, e := ParseIntErr("666x")
	if e == nil {
		t.Errorf("TestParseIntErr fail dest>%v,%v", i, e)
	}
}
