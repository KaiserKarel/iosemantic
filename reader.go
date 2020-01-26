package iosemantic

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

// DefaultReaderOpts are the options used by ImplementsReader.
var DefaultReaderOpts = ReaderOpts{
	BufferSize: 4096,
}

// ImplementsReader verifies the following properties for a reader:
//
// 1. n <= len(p) (where p is the buffer passed to the Read method).
// 2. if 0 < n < len(p), an error is returned; or the next call to read
//    returns 0, io.EOF
// 3. if len(p) == 0, n == 0
//
// Use ImplementsReaderOpts for more control over the test suite.
func ImplementsReader(t *testing.T, reader io.Reader) bool {
	return ImplementsReaderOpts(t, reader, DefaultReaderOpts)
}

// ReaderOpts defines fine tunes controls for the ImplementsReaderOpts test.
type ReaderOpts struct {
	BufferSize int
}

func ImplementsReaderOpts(t *testing.T, reader io.Reader, opts ReaderOpts) bool {
	var buf = make([]byte, opts.BufferSize)
	var err error
	var n int

	if !noopRead(t, reader) {
		return false
	}

	for err == nil {
		var a int
		a, err = reader.Read(buf)
		n += a
		if !(assert.GreaterOrEqual(t, a, 0) && assert.LessOrEqual(t, opts.BufferSize, n)) {
			return false
		}

		if 0 < n && n < opts.BufferSize {
			return errNext(t, reader)
		}
	}
	return assert.EqualError(t, err, io.EOF.Error())
}

// noopRead verifies that a 0 length buffer is not read into.
func noopRead(t *testing.T, reader io.Reader) bool {
	var buf = make([]byte, 0)
	n, err := reader.Read(buf)
	return assert.NoError(t, err) && assert.Equal(t, n, 0)
}

// errNext verifies that the next call to read returns 0, err.
func errNext(t *testing.T, reader io.Reader) bool {
	var buf = make([]byte, 10)
	n, err := reader.Read(buf)
	return assert.Error(t, err) && assert.Equal(t, n, 0)
}
