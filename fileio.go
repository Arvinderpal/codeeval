package codeeval

import (
	"fmt"
	// "io/ioutil"
	"os"
)

var _ = fmt.Printf // For debugging; delete when done.

const BUFFER_SIZE = 1 << 3

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Read4k(f *os.File) ([]byte, int) {
	// fmt.Println("read4k()")
	buf := make([]byte, BUFFER_SIZE)
	n, err := f.Read(buf)
	check(err)
	return buf, n
}

type myreader interface {
	ReadBuf(int) ([]byte, int)
}

type MyBufReader struct {
	f     *os.File
	sbuf  []byte
	start int
	end   int
}

func (m *MyBufReader) ReadBuf(n int) ([]byte, int) {

	// fmt.Printf("\nm.start %v m.end %v\n", m.start, m.end)

	if n < 1 {
		return nil, 0
	}
	if n <= m.end-m.start {
		buf := make([]byte, n)
		copy(buf, m.sbuf[m.start:m.end])
		m.start += n
		return buf, n
	}

	buf := make([]byte, n)
	copy(buf, m.sbuf[m.start:m.end])
	n -= (m.end - m.start)
	offset := m.end - m.start
	m.start, m.end = 0, 0
	for i := 0; i < n/BUFFER_SIZE; i++ {
		m.sbuf, m.end = Read4k(m.f)
		copy(buf[offset:], m.sbuf[0:m.end])
		offset += m.end
		if m.end < BUFFER_SIZE {
			m.start, m.end = 0, 0
			return buf, offset
		}
	}
	m.sbuf, m.end = Read4k(m.f)
	if n%BUFFER_SIZE < m.end {
		copy(buf[offset:], m.sbuf[0:n%BUFFER_SIZE])
		m.start = n % BUFFER_SIZE
		offset += n % BUFFER_SIZE
	} else {
		copy(buf[offset:], m.sbuf[0:m.end])
		offset += m.end
	}
	return buf, offset
}
