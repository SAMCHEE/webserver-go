package main

import (
	"flag"
	"log"

	"github.com/SAMCHEE/webserver-go/server"
)

var (
	FLAG_PORT int
)

func init() {
	flag.IntVar(&FLAG_PORT, "port", 4000, "port number")
	flag.Parse()
}

func main() {
	s := server.NewServer(uint32(FLAG_PORT))
	log.Fatal(s.Run())
}
