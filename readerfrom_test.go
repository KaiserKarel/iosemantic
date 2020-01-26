package iosemantic_test

import (
	"bytes"
	"testing"

	"github.com/kaiserkarel/iosemantic"
	"github.com/stretchr/testify/assert"
)

func TestImplementsReaderFrom(t *testing.T) {
	reader := bytes.NewBuffer(make([]byte, 4096*100))
	assert.True(t, iosemantic.ImplementsReaderFrom(t, reader))
}

func TestImplementsReaderFromOpts(t *testing.T) {
	reader := bytes.NewBuffer(make([]byte, 4096*100))
	assert.True(t, iosemantic.ImplementsReaderFromOpts(t, reader, iosemantic.ReaderFromOpts{BufferSize: 303 * 299}))
}
