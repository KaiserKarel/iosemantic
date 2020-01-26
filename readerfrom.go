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
	"io"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
)

var defaultReaderFromOpts = ReaderFromOpts{BufferSize: 4096 * 100}

// ImplementsReaderFrom verifies the following properties for a reader:
//
// 1. The ReaderFrom consumes the input until an error is encountered.
// 2. io.EOF is not returned.
//
// Use ImplementsReaderFromOpts for more control over the test suite.
func ImplementsReaderFrom(t *testing.T, reader io.ReaderFrom) bool {
	return ImplementsReaderFromOpts(t, reader, defaultReaderFromOpts)
}

// ReaderFromOpts defines fine tunes controls for the ImplementsReaderFromOpts test.
type ReaderFromOpts struct {
	BufferSize int
}

// ImplementsReaderFromOpts uses providing options to perform ImplementsReaderFrom.
func ImplementsReaderFromOpts(t *testing.T, reader io.ReaderFrom, opts ReaderFromOpts) bool {
	src := iotest.TimeoutReader(bytes.NewReader(make([]byte, opts.BufferSize)))
	f, err := reader.ReadFrom(src)
	if !assert.EqualError(t, err, iotest.ErrTimeout.Error()) {
		return false
	}
	s, err := reader.ReadFrom(src)
	return assert.NoError(t, err) && assert.Equal(t, int(s+f), opts.BufferSize)
}
