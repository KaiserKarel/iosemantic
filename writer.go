package iosemantic

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

// DefaultWriterOpts are the options used by ImplementsWriter.
var DefaultWriterOpts = WriterOpts{
	BufferSize: 4096 * 100,
}

// ImplementsWriter verifies the following properties for a writer:
//
// 1. 0 <= n <= len(p) where p is the buffer being written from.
// 2. if n < len(p), err != nil.
//
// Use ImplementsWriterOpts for more control over the test suite.
func ImplementsWriter(t *testing.T, writer io.Writer) bool {
	return ImplementsWriterOpts(t, writer, DefaultWriterOpts)
}

type WriterOpts struct {
	BufferSize int
}

func ImplementsWriterOpts(t *testing.T, writer io.Writer, opts WriterOpts) bool {
	var buf = make([]byte, opts.BufferSize)
	var n int
	var err error

	for err == nil && n < opts.BufferSize {
		var a int
		chunk := buf[n:]
		a, err = writer.Write(chunk)
		n += a

		if !(assert.Greater(t, n, 0) && assert.LessOrEqual(t, n, opts.BufferSize)) {
			return false
		}

		if n < len(chunk) {
			return assert.Error(t, err)
		}
	}
	return assert.NoError(t, err) && assert.Equal(t, n, opts.BufferSize)
}
