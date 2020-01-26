// Copyright 2020 Karel L. Kubat
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
// documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit
// persons to whom the Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
// Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
// WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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

// WriterOpts defines fine tunes controls for the ImplementsWriterOpts test.
type WriterOpts struct {
	BufferSize int
}

// ImplementsWriterOpts uses providing options to perform ImplementsWriter.
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
