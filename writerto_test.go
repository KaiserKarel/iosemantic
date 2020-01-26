package iosemantic_test

import (
	"bytes"
	"github.com/kaiserkarel/iosemantic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImplementsWriterTo(t *testing.T) {
	writer := bytes.NewBuffer(make([]byte, 4096*100))
	assert.True(t, iosemantic.ImplementsWriterTo(t, writer))
}

func TestImplementsWriterToOpts(t *testing.T) {
	writer := bytes.NewBuffer(make([]byte, 4096*100))
	assert.True(t, iosemantic.ImplementsWriterToOpts(t, writer, iosemantic.WriterToOpts{BufferSize: 303 * 299}))
}
