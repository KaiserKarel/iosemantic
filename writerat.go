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
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
)

var defaultWriterAtOpts = WriterAtOpts{
	BufferSize: 4096,
}

// ImplementsWriterAt verifies the following properties for a writer:
//
// 1. 0 <= n <= len(p) where p is the buffer being written from.
// 2. if n < len(p), err != nil.
// 3. No error is returned during parallel WriteAt calls on the same destination if the ranges do not overlap.
//
// Use ImplementsWriterAtOpts for more control over the test suite.
func ImplementsWriterAt(t *testing.T, writer io.WriterAt, length int64) bool {
	return ImplementsWriterAtOpts(t, writer, length, defaultWriterAtOpts)
}

// WriterAtOpts defines fine tunes controls for the ImplementsWriterAtOpts test.
type WriterAtOpts struct {
	BufferSize int
}

// ImplementsWriterAtOpts uses providing options to perform ImplementsWriterAt.
func ImplementsWriterAtOpts(t *testing.T, writer io.WriterAt, length int64, opts WriterAtOpts) bool {
	if !ImplementsWriterOpts(t,
		toWriter(writer, 0), WriterOpts(opts)) {
		return false
	}

	grp, _ := errgroup.WithContext(context.Background())
	for i := int64(0); i < length && i < 50; i++ {
		i := i
		grp.Go(func() error {
			var buf = make([]byte, 1)
			_, err := writer.WriteAt(buf, i)
			assert.NoError(t, err)
			return err
		})
	}

	return assert.NoError(t, grp.Wait())
}

type writer struct {
	at io.WriterAt
	i  int64
}

func (w *writer) Write(p []byte) (int, error) {
	n, err := w.at.WriteAt(p, w.i)
	w.i += int64(n)
	return n, err
}

func toWriter(at io.WriterAt, offset int64) io.Writer {
	return &writer{
		at: at,
		i:  offset,
	}
}
