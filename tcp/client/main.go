package main

import (
	"log"
	"net"
	"strconv"
)

func main() {
	c, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Panic(err)
	}
	defer c.Close()

	c.Write([]byte("05GET04"))

	s := make([]byte, 2)
	c.Read(s)
	log.Printf("Reading %s bytes from server", string(s))

	size, err := strconv.ParseInt(string(s), 10, 64)
	if err != nil {
		log.Panic(err)
	}

	d := make([]byte, size-2)
	c.Read(d)
	log.Printf("Response: %s", string(d))
}
