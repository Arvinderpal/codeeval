package codeeval

import (
	"fmt"
	"os"
	"testing"
)

var _ = fmt.Printf // For debugging; delete when done.

func TestRead4k(t *testing.T) {
	cases := []struct {
		in    string
		want  []byte
		want2 int
	}{
		{"_test_data/d1.txt", []byte("sdfljslfjsldjflsjflsfd n,sngetiueorjlkfj"), BUFFER_SIZE},
		// {"_test_data/d2.txt", []byte("XXXXXXX"), 4096},
	}
	for _, c := range cases {
		f, _ := os.Open(c.in)
		got, got2 := Read4k(f)
		if got2 != c.want2 {
			t.Errorf("Length mismatch: want: %v --- got len: %v", c.want2, got2)
		}
		for i := 0; i < got2; i++ {
			if got[i] != c.want[i] {
				t.Errorf("string mismatch at index %v: want: %v --- got: %v", i, string(c.want[i]), string(got[i]))
				break
			}
		}
	}
}

func TestReadBuf(t *testing.T) {
	cases := []struct {
		in string
	}{
		{"_test_data/d1.txt"},
		{"_test_data/d2.txt"},
	}
	subcases := []struct {
		caseidx int
		in      int
		want    []byte
		want2   int
	}{
		{0, 2, []byte("sd"), 2},
		{0, 2, []byte("fl"), 2},
		{0, 2, []byte("js"), 2},
		{0, 2, []byte("lf"), 2},
		{0, 16, []byte("jsldjflsjflsfd n"), 16},
		{0, 11, []byte(",sngetiueor"), 11},
		{0, 20, []byte("jlkfjxz"), 7},
		{1, 14, []byte("<!DOCTYPE html"), 14},
		{1, 1, []byte(">"), 1},
		{1, 1, []byte("\n"), 1},
	}
	i := 0
	for cidx, c := range cases {
		// fmt.Printf("cidx: %v\n", cidx)
		fd, _ := os.Open(c.in)
		mr := MyBufReader{f: fd, start: 0, end: 0}
		for ; i < len(subcases); i++ {
			if cidx == subcases[i].caseidx {
				got, got2 := mr.ReadBuf(subcases[i].in)
				if got2 != subcases[i].want2 {
					t.Errorf("Length mismatch: want: %v --- got len: %v", subcases[i].want2, got2)
				}
				for j := 0; j < got2; j++ {
					if got[j] != subcases[i].want[j] {
						t.Errorf("string mismatch at index %v: want: %v --- got: %v", j, string(subcases[i].want[j]), string(got[j]))
						break
					}
				}
			}
			if cidx < subcases[i].caseidx {
				// assume subcases are ordered
				break
			}
		}

	}
}
