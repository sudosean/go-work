package main

import(
	"fmt"

	"hash/crc32"

	"crypto/sha1"
)

func main()  {
	// non cryptographic hash functions can be found underneath the hash package
	// adler32, crc32, crc64 and fnv
	
	// Package crc32 implements the 32-bit //* cyclic redundancy check *//, or CRC-32,
	// checksum. See https://en.wikipedia.org/wiki/Cyclic_redundancy_check for
	// information.

	// create a hasher
	h := crc32.NewIEEE()
	//write our dat to it
	h.Write([]byte("test"))
	//calculate the crc32
	v := h.Sum32()
	fmt.Println(v)
	fmt.Println(h.BlockSize())

	// * a common use for crc32 is to compare two files, if the check sums are the same it is 
	// * highly likely they are the same (not 100% certain though), if the value is different then they are
	// * certainly not the same

	// crypto
	// SHA-1 is cryptographically broken and should not be used for secure
	// applications.

	c := sha1.New()
	c.Write([]byte("test"))
	cs := c.Sum([]byte{})
	fmt.Println("crypto ----",cs)
     
	/// sha1 is 160 bit hash
	// * there is no native type to represent a 160bit number so we use a slice of 20bytes instead
}