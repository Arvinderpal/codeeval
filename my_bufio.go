package codeeval

import (
	"errors"
	"io"
)

const defaultBufSize = 4

var ErrInvalidSlice = errors.New("buffered reader: input slice is empty")

// Reader is a buffered reader. Given an object that implements
// the io.Reader interface, this Reader will perform buffered IO.
type Reader struct {
	rd          io.Reader
	buf         []byte
	rPos        int   // position in the buffer
	lastBytePos int   // index of last byte in buffer
	err         error // we need to keep track of any errors/EOF
}

func NewReader(rd io.Reader) *Reader {
	r := &Reader{
		rd:          rd,
		buf:         make([]byte, defaultBufSize),
		rPos:        0,
		lastBytePos: -1,
	}
	return r
}

// Read will attempt to read as much as len(s) bytes into the slice s.
// It will return the actual number of bytes read or an error.
func (r *Reader) Read(s []byte) (int, error) {
	var n int
	var err error
	// enough bytes in r.buf to satisfy s
	byteInBuf := (r.lastBytePos + 1) - r.rPos
	if len(s) < byteInBuf {
		n = len(s)
		upTo := r.rPos + n
		copy(s, r.buf[r.rPos:upTo])
		r.rPos += n
		return n, r.err
	}

	// first, we copy the bytes (if any) in buf to s
	if byteInBuf > 0 {
		copy(s, r.buf[r.rPos:r.lastBytePos+1])
		n = byteInBuf
		r.rPos = 0
		r.lastBytePos = -1
	}

	// next, we read as many blocks of defaultBufSize as we can
	// directly into s; however, for the last read, we may have
	// < defaultBufSize room left in s; in that case, we read
	// defaultBufSize into r.buf and then copy some amount
	// in s; keeping the rest in buf for next call.
	bytesLeft := len(s[n:]) % defaultBufSize
	bytesToRead := len(s[n:]) - bytesLeft
	var nx int
	if bytesToRead > 0 {
		nx, err = r.rd.Read(s[n : n+bytesToRead])
		if err != nil {
			return n + nx, err
		}
	}
	var nb int
	nb, r.err = r.rd.Read(r.buf)
	if nb > bytesLeft {
		copy(s[n+nx:], r.buf[:bytesLeft])
		r.rPos = bytesLeft
		r.lastBytePos = nb - 1
		return n + nx + bytesLeft, r.err
	}
	copy(s[n+nx:], r.buf[:nb])
	return n + nx + nb, r.err
}
