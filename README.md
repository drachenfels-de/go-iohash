# go-iohash
Use iohash to generate a hash from the IO of an io.Reader or io.Writer.
It simply wraps an existing io.Reader or io.Writer and additionaly feeds
the io to a wrapped hash.

[See test for usage!](iohash_test.go)
