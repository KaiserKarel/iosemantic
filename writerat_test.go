package iosemantic_test

import (
	"testing"

	"github.com/djherbis/buffer"

	"github.com/kaiserkarel/iosemantic"
	"github.com/stretchr/testify/assert"
)

func TestImplementsWriterAt(t *testing.T) {
	var length int64 = 4096 * 100
	writer := buffer.New(length)
	assert.True(t, iosemantic.ImplementsWriterAt(t, writer, length))
}

func TestImplementsWriterAtOpts(t *testing.T) {
	var length int64 = 4096 * 100
	writer := buffer.New(length)
	assert.True(t, iosemantic.ImplementsWriterAtOpts(t, writer, length, iosemantic.WriterAtOpts{BufferSize: 1999}))
}
