package main

import (
	"errors"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

func handle(c net.Conn) error {
	defer c.Close()
	log.Printf("Serving connection: %s\n", c.RemoteAddr().String())

	s := make([]byte, 2)
	c.Read(s)
	log.Printf("Reading %s bytes from client", string(s))
	dataSize, err := strconv.ParseInt(string(s), 10, 64)
	if err != nil {
		return err
	}

	d := make([]byte, dataSize)
	c.Read(d)
	cmd := string(d)
	r, err := handleCommand(cmd)

	resp := strconv.Itoa(len(r)+2) + r

	c.Write([]byte(resp))
	log.Printf("Responded data %s to client", resp)
	return nil
}

func handleCommand(cmd string) (string, error) {
	if strings.HasPrefix(cmd, "GET") {
		return handleGet(cmd)
	}
	return "", errors.New("unknown command")
}

func handleGet(cmd string) (string, error) {
	n, err := strconv.ParseInt(cmd[len(cmd)-2:], 10, 64)
	if err != nil {
		return "", err
	}
	var r []string
	for i := 0; int64(i) < n; i++ {
		r = append(r, strconv.Itoa(rand.Intn(100)))
	}
	res := strings.Join(r, " ")
	return res, nil
}

func main() {
	port := ":8000"
	l, err := net.Listen("tcp4", port)
	if err != nil {
		log.Panic(err)
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	log.Printf("Server started: waiting for connection at %s", port)

	for {
		c, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handle(c)
	}

}
