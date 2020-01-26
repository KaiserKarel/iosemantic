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

	"golang.org/x/sync/errgroup"

	"github.com/stretchr/testify/assert"
)

var defaultReaderAtOpts = ReaderAtOpts{
	BufferSize: 4096,
}

// ImplementsReaderAt verifies the following properties for a io.ReaderAt:
//
// 1. n <= len(p) (where p is the buffer passed to the Read method).
// 2. if 0 < n < len(p), an error is returned;
// 3. if len(p) == 0, n == 0
// 4. Parallel ReadAt calls do not result in errors.
//
// ImplementsReaderAt is a more strict version of ImplementsReader, just like the semantics of io.Reader and io.ReaderAt.
// Use ImplementsReaderAtOpts for more control over the test suite.
func ImplementsReaderAt(t *testing.T, reader io.ReaderAt, length int64) bool {
	return ImplementsReaderAtOpts(t, reader, length, defaultReaderAtOpts)
}

// ReaderAtOpts defines fine tunes controls for the ImplementsReaderAtOpts test.
type ReaderAtOpts struct {
	BufferSize int
}

// ImplementsReaderAtOpts uses providing options to perform ImplementsReaderAt.
func ImplementsReaderAtOpts(t *testing.T, reader io.ReaderAt, length int64, opts ReaderAtOpts) bool {
	var buf = make([]byte, opts.BufferSize)
	var err error
	var n int64

	if !noopRead(t, toReader(reader, 0)) {
		return false
	}

	for err == nil {
		var a int
		a, err = reader.ReadAt(buf, n)
		n += int64(a)
		if !(assert.GreaterOrEqual(t, a, 0) && assert.LessOrEqual(t, int64(opts.BufferSize), n)) {
			return false
		}

		if 0 < n && n < int64(opts.BufferSize) {
			return assert.Error(t, err)
		}
	}

	grp, _ := errgroup.WithContext(context.Background())
	for i := int64(0); i < length && i < 50; i++ {
		i := i
		grp.Go(func() error {
			var buf = make([]byte, opts.BufferSize)
			_, err := reader.ReadAt(buf, i)
			assert.NoError(t, err)
			return err
		})
	}
	err2 := grp.Wait()

	return assert.EqualError(t, err, io.EOF.Error()) && assert.NoError(t, err2)
}

type reader struct {
	at io.ReaderAt
	i  int64
}

func (r *reader) Read(p []byte) (int, error) {
	n, err := r.at.ReadAt(p, r.i)
	r.i += int64(n)
	return n, err
}

func toReader(at io.ReaderAt, offset int64) io.Reader {
	return &reader{
		at: at,
		i:  offset,
	}
}
