package iohash

import (
	"bufio"
	"errors"
	"fmt"
	"hash"
	"io"
	"os"
	"strings"
)

var ErrMismatch = errors.New("checksum mismatch")
var ErrNotFound = errors.New("no checksum found")

func PrintHash(w io.Writer, h hash.Hash, filename string) (int, error) {
	return fmt.Fprintf(w, "%x %s\n", h.Sum(nil), filename)
}

// See description of iohash.Check.
func CheckFile(checksumsFile string, h hash.Hash, filename string) error {
	f, err := os.Open(checksumsFile)
	if err != nil {
		return fmt.Errorf("failed to open checksums file %s: %s", checksumsFile, err)
	}
	defer f.Close()
	return Check(f, h, filename)
}

// Check verifies the checksum of targetFile.
// The hash from the first line matching the targetFile is checked.
// Any succeeding lines are ignored.
// The line format is the output from #PrintHash, which is
// similar to the unix coreutils commands e.g `sha1` `sha512` ...
func Check(rdr io.Reader, h hash.Hash, targetFile string) error {
	sc := bufio.NewScanner(rdr)
	hashS := fmt.Sprintf("%x", h.Sum(nil))
	for sc.Scan() {
		line := sc.Text()
		fileName := strings.TrimSpace(line[len(hashS):])
		checksum := line[:len(hashS)]
		if targetFile == fileName {
			if hashS == checksum {
				return nil
			}
			return ErrMismatch
		}
	}
	return ErrNotFound
}
