package main

import (
	"fmt"
	"strings"
)

func main() {
	// To search for a smaller string in a bigger string, use the `Contains` function:
	fmt.Println(strings.Contains("test", "es")) // true

	// To count the number of times a smaller string occurs in a bigger string, use the `Count`
	// function:
	fmt.Println(strings.Count("test", "t")) // 2

	// To determine if a bigger string starts with a smaller string, use the `HasPrefix`
	// function:
	fmt.Println(strings.HasPrefix("test", "te")) // true

	// To determine if a bigger string ends with a smaller string, use the `HasSuffix`
	// function:
	fmt.Println(strings.HasSuffix("test", "st")) // true

	// To find the position of a smaller string in a bigger string, use the `Index` function
	// (it returns `-1 if not found):
	fmt.Println(strings.Index("test", "e")) // 1

	// To take a list of strings and join them together in a single string separated by another
	// string (e.g., a comma), use the `Join function:
	fmt.Println(strings.Join([]string{"a", "b"}, "-")) // "a-b"

	// To repeat a string, use the `Repeat` function:
	fmt.Println(strings.Repeat("a", 5)) // "aaaaa"

	// To replace a smaller string in a bigger string with some other string, use the `Replace`
	// function.
	fmt.Println(strings.Replace("aaaa", "a", "b", 2)) // "bbaa"

	// To split a string into a list of strings by a separating string (e.g., a comma), use the
	// `Split` function (`Split` is the reverse of `Join`):
	fmt.Println(strings.Split("a-b-c-d-e", "-")) // []string{"a", "b", "c", "d", "e"}

	// To convert a string to all lowercase letters, use the `ToLower` function:
	fmt.Println(strings.ToLower("TEST")) // "test"

	// To convert a string to all uppercase letters, use the `ToUpper` function:
	fmt.Println(strings.ToUpper("test")) // "TEST"

	// Sometimes we need to work with strings as binary data. To convert a string to a slice
	// of bytes (and vice versa), do this:
	arr := []byte("test")
	fmt.Println(arr)
	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(str)
}
