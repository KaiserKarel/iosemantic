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
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
	"testing/iotest"
)

var DefaultWriterToOpts = WriterToOpts{
	BufferSize: 4096,
}

// ImplementsWriterTo verifies the following properties for a reader:
//
// 1. The WriterTo writes to the writer until it is finished, or an error is encountered.
// 2. Any error is returned.
//
// Use ImplementsWriterToOpts for more control over the test suite.
func ImplementsWriterTo(t *testing.T, writer io.WriterTo) bool {
	return ImplementsWriterToOpts(t, writer, DefaultWriterToOpts)
}

// WriterToOpts defines fine tunes controls for the ImplementsWriterToOpts test.
type WriterToOpts struct {
	BufferSize int
}

// ImplementsWriterToOpts uses providing options to perform ImplementsWriterTo.
func ImplementsWriterToOpts(t *testing.T, writer io.WriterTo, opts WriterToOpts) bool {
	src := &timoutWriter{
		bytes.NewBuffer(make([]byte, opts.BufferSize)),
		true}
	n, err := writer.WriteTo(src)
	assert.EqualError(t, err, iotest.ErrTimeout.Error())
	assert.Zero(t, n)

	n, err = writer.WriteTo(src)
	return assert.NoError(t, err) &&
		assert.NotZero(t, opts.BufferSize, int(n), "expected at least a write gte zero, but actually wrote: %d", n)
}

// timeOutWriter resembles iotest.TimeoutReader.
type timoutWriter struct {
	writer io.Writer
	err    bool
}

// Write buffer p to underlying writer. If this is the first write, return ErrTimeout.
func (t *timoutWriter) Write(p []byte) (int, error) {
	if t.err {
		t.err = false
		return 0, iotest.ErrTimeout
	}
	return t.writer.Write(p)
}
