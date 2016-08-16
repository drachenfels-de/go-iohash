package main

import (
	"bytes"
	"crypto/sha1"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

const (
	content        = "Hello World"
	contentSHA1Sum = "0a4d55a8d778e5022fab701977c5d840bbc486d0"
)

func TestReader(t *testing.T) {
	in := strings.NewReader(content)
	r := NewReader(in, sha1.New())
	readContent, _ := ioutil.ReadAll(r)
	assert.Equal(t, content, string(readContent))
	assert.Equal(t, contentSHA1Sum, StringOfHash(r.Hash))
}

func TestWriter(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	w := NewWriter(buf, sha1.New())
	_, err := w.Write([]byte(content))
	assert.NoError(t, err)
	assert.Equal(t, content, buf.String())
	assert.Equal(t, contentSHA1Sum, StringOfHash(w.Hash))
}
