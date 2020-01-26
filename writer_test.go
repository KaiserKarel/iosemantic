package iosemantic_test

import (
	"bytes"
	"github.com/kaiserkarel/iosemantic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImplementsWriter(t *testing.T) {
	writer := bytes.NewBuffer(make([]byte, 4096 * 100))
	assert.True(t, iosemantic.ImplementsWriter(t, writer))
}

func TestImplementsWriterOpts(t *testing.T) {
	writer := bytes.NewBuffer(make([]byte, 4096 * 100))
	assert.True(t, iosemantic.ImplementsWriterOpts(t, writer, iosemantic.WriterOpts{BufferSize: 201 * 1011}))
}