package iosemantic_test

import (
	"bytes"
	"testing"

	"github.com/kaiserkarel/iosemantic"
	"github.com/stretchr/testify/assert"
)

func TestImplementsReaderAt(t *testing.T) {
	length := 4096 * 100
	reader := bytes.NewReader(make([]byte, length))
	assert.True(t, iosemantic.ImplementsReaderAt(t, reader, int64(length)))
}

func TestImplementsReaderAtOpts(t *testing.T) {
	length := 4096 * 100
	reader := bytes.NewReader(make([]byte, length))
	assert.True(t, iosemantic.ImplementsReaderAtOpts(t, reader, int64(length), iosemantic.ReaderAtOpts{BufferSize: 1999}))
}
