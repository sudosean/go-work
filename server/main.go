package main 

import(
	"fmt"
	"encoding/gob"
	"net"
)
// encoding/gob makes it easy to encode Go values so that other go programs (or the same) can read them

func server()  {
  // listen on a port
  ln, err := net.Listen("tcp", ":9999")
  if err != nil {
	  panic(err)
  }
  for{
	  // accept a connection
	  c, err := ln.Accept()
	  if err != nil {
		  panic(err)
	  }
	  go handleServerConnection(c)
  }
}

func handleServerConnection(c net.Conn)  {
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		panic(err)
	} else{
		fmt.Println("Received", msg)
	}
   c.Close()
}

func client()  {
	// connect to the server
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		panic(err)
	}
	// send the message
	msg := "Hello, World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		panic(err)
	}
	c.Close()
}



func main()  {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}

