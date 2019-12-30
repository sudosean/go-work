package main

// the io package consists of two important interfaces:
// Reader
// Writer
// Many functions in Go take Readers or Writers as arguments:
// fucn Copy(dst Writer, src Reader)(written int64, err error)

// to read or write to []byte or string you can use the Buffer struct found in the bytes package
// var buf bytes.Buffer
// buf.Write([]byte("test"))

// A buffer does not have to be initilaized and it supports both reader and writter interfaces.
// you can convert it into a []byte by calling buf.Bytes(). if you only need to read froma string, 
// you can also use strings.NewReader function which is more efficient

import(
	"fmt"
	"os"
	"io/ioutil"
)

func main(){
 /// this is the long way to read a file
  file, err := os.Open("test.txt")
  if err != nil {
	  panic(err)
  }
  defer file.Close()

  // get file size
  stat, err := file.Stat()
  if err != nil {
	  panic(err)
  }
  // read file
  bs := make([]byte, stat.Size())
  _, err = file.Read(bs)
  if err != nil {
	  panic(err)
  }

  str := string(bs)
  fmt.Println(str)

  // This is the fast way
	bs2, err := ioutil.ReadFile("test2.txt")
	if err != nil {
		panic(err)
	}
	str2 := string(bs2)
	fmt.Println(str2)

	// os.Create to create a file
	// os.Open to get contents of a directory, then call Readdir(-1) -1 is to get all contents
	// path/filepath package: use .Walk to recrursively go through program
}