package iosemantic

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
)

var DefaultWriterAtOpts = WriterAtOpts{
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
	return ImplementsWriterAtOpts(t, writer, length, DefaultWriterAtOpts)
}

type WriterAtOpts struct {
	BufferSize int
}

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
