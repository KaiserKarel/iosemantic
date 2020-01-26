package iosemantic_test

import (
	"bytes"
	"github.com/kaiserkarel/iosemantic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImplementsReader(t *testing.T) {
	reader := bytes.NewBuffer(make([]byte, 4096 * 100))
	assert.True(t, iosemantic.ImplementsReader(t, reader))
}

func TestImplementsReaderOpts(t *testing.T) {
	reader := bytes.NewBuffer(make([]byte, 4096 * 100))
	assert.True(t, iosemantic.ImplementsReaderOpts(t, reader, iosemantic.ReaderOpts{BufferSize:1999}))
}