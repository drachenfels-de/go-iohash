package iohash

import (
	"fmt"
	"hash"
	"io"
)

func StringOfHash(h hash.Hash) string {
	return fmt.Sprintf("%x", h.Sum(nil))
}

type HashWriter struct {
	io.Writer
	hash.Hash
}

func (h *HashWriter) Write(p []byte) (n int, err error) {
	n, err = h.Writer.Write(p)
	if err == nil {
		// Hash.Write never returns an error (see godoc interface definition)
		h.Hash.Write(p[:n])
	}
	return
}

func NewWriter(w io.Writer, h hash.Hash) *HashWriter {
	return &HashWriter{w, h}
}

type HashReader struct {
	io.Reader
	hash.Hash
}

func (h *HashReader) Read(p []byte) (int, error) {
	n, err := h.Reader.Read(p)
	if n > 0 {
		// Hash.Write never returns an error (see godoc interface definition)
		h.Hash.Write(p[:n])
	}
	return n, err
}

func NewReader(r io.Reader, h hash.Hash) *HashReader {
	return &HashReader{r, h}
}
