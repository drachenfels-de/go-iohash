package iohash

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha512"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
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
	require.Equal(t, content, string(readContent))
	require.Equal(t, contentSHA1Sum, StringOfHash(r.Hash))
}

func TestWriter(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	w := NewWriter(buf, sha1.New())
	_, err := w.Write([]byte(content))
	require.NoError(t, err)
	require.Equal(t, content, buf.String())
	require.Equal(t, contentSHA1Sum, StringOfHash(w.Hash))
}

func TestCheck(t *testing.T) {
	h1 := sha512.New()
	h1.Write([]byte("hello world"))
	h2 := sha512.New()
	h2.Write([]byte("foobar"))
	h3 := sha512.New()
	h2.Write([]byte("foobar"))
	checksums, err := ioutil.TempFile("", "")
	require.NoError(t, err)
	PrintHash(checksums, h1, "file1")
	PrintHash(checksums, h2, "file2")
	PrintHash(checksums, h3, "file3")
	checksums.Close()
	defer os.Remove(checksums.Name())

	// read SHA512 sums from the FILEs and check them
	require.NoError(t, CheckFile(checksums.Name(), h1, "file1"))
	require.NoError(t, CheckFile(checksums.Name(), h2, "file2"))
	require.NoError(t, CheckFile(checksums.Name(), h3, "file3"))
	require.Equal(t, ErrMismatch, CheckFile(checksums.Name(), sha512.New(), "file2"))
	require.Equal(t, ErrNotFound, CheckFile(checksums.Name(), h3, "file4"))
	require.Error(t, CheckFile("doesnotexist", h3, "file3"))
}
