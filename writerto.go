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

type WriterToOpts struct {
	BufferSize int
}

func ImplementsWriterToOpts(t *testing.T, writer io.WriterTo, opts WriterToOpts) bool {
	src := &TimeoutWriter{
		bytes.NewBuffer(make([]byte, opts.BufferSize)),
		true}
	n, err := writer.WriteTo(src)
	assert.EqualError(t, err, iotest.ErrTimeout.Error())
	assert.Zero(t, n)

	n, err = writer.WriteTo(src)
	return assert.NoError(t, err) &&
		assert.NotZero(t, opts.BufferSize, int(n), "expected at least a write gte zero, but actually wrote: %d", n)
}

// TimeOutWriter resembles iotest.TimeoutReader.
type TimeoutWriter struct {
	writer io.Writer
	err bool
}

// Write buffer p to underlying writer. If this is the first write, return ErrTimeout.
func (t *TimeoutWriter) Write(p []byte) (int,error)  {
	if t.err {
		t.err = false
		return 0, iotest.ErrTimeout
	}
	return t.writer.Write(p)
}