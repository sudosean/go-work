package main

import(
	"fmt"

	"strings"
)

func main()  {
	// contains, substring
	fmt.Println(strings.Contains("test", "es"))
	// count
	fmt.Println(strings.Count("test", "t"))
	// prefix
	fmt.Println(strings.HasPrefix("test", "te"))
	// suffix
	fmt.Println(strings.HasSuffix("test", "st"))
	// index of char
	fmt.Println(strings.Index("test", "e"))
	// join strings
	fmt.Println(strings.Join([]string{"a", "b"},"-"))


	// convert string to slice of bytes
	arr := []byte("test")
	str := string([]byte{'t', 'e', 's','t'})
	fmt.Println(arr,str)
}