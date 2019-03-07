package codeeval

import (
	"strings"
	"testing"
)

func TestReadSimple(t *testing.T) {
	data := "hello world"
	r := NewReader(strings.NewReader(data))
	s := make([]byte, 4)
	n, err := r.Read(s)
	if err != nil {
		t.Errorf("expected no error but got %s", err)
	}
	if n != 4 {
		t.Errorf("expected 4 bytes to be read but got %v", n)
	}
	if string(s) != "hell" {
		t.Errorf("expected %s got %s", "hell", s)
	}

}

func TestReadEOF(t *testing.T) {
	data := "01234567"
	r := NewReader(strings.NewReader(data))
	s := make([]byte, 7)
	n, err := r.Read(s)
	if err != nil {
		t.Errorf("expected no error but got %s", err)
	}
	if n != 7 {
		t.Errorf("expected 7 bytes to be read but got %v", n)
	}
	if string(s) != "0123456" {
		t.Errorf("expected 0123456 got %s", s)
	}
	if string(r.buf) == "7" {
		t.Errorf("expected buf to be \"7\" got %s", r.buf)
	}
	n, err = r.Read(s)
	if err == nil {
		t.Errorf("expected %s error but got nil", err)
	}
	if n != 1 {
		t.Errorf("expected 1 bytes to be read but got %v", n)
	}
	if string(s[:n]) != "7" {
		t.Errorf("expected '7' got %s", s)
	}
	if r.rPos != 0 || r.lastBytePos != -1 {
		t.Errorf("expected 0/-1 got %v/%v ", r.rPos, r.lastBytePos)
	}
}
