package main

import (
	"bytes"
	"fmt"
)

func main() {
	// The `io` package consists of a few functions, but mostly interfaces used in other
	// packages. The two main interfaces are `Reader` and `Writer`. `Reader`s support reading
	// via the `Read` method. `Writer`s support writing via the `Write` method.

	// To read or write to a `[]byte` or a `string`, you can use the `Buffer` struct found in
	// the `bytes package:
	var buf bytes.Buffer // A `Buffer` doesn't have to be initialized.
	buf.Write([]byte("test"))
	fmt.Println(string(buf.Bytes()))
}
