package iosemantic

import (
	"bytes"
	"io"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
)

var DefaultReaderFromOpts = ReaderFromOpts{BufferSize: 4096 * 100}

// ImplementsReaderFrom verifies the following properties for a reader:
//
// 1. The ReaderFrom consumes the input until an error is encountered.
// 2. io.EOF is not returned.
//
// Use ImplementsReaderFromOpts for more control over the test suite.
func ImplementsReaderFrom(t *testing.T, reader io.ReaderFrom) bool {
	return ImplementsReaderFromOpts(t, reader, DefaultReaderFromOpts)
}

type ReaderFromOpts struct {
	BufferSize int
}

func ImplementsReaderFromOpts(t *testing.T, reader io.ReaderFrom, opts ReaderFromOpts) bool {
	src := iotest.TimeoutReader(bytes.NewReader(make([]byte, opts.BufferSize)))
	f, err := reader.ReadFrom(src)
	if !assert.EqualError(t, err, iotest.ErrTimeout.Error()) {
		return false
	}
	s, err := reader.ReadFrom(src)
	return assert.NoError(t, err) && assert.Equal(t, int(s+f), opts.BufferSize)
}
