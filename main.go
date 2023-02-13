package main

import (
	"flag"
	"log"

	"github.com/SAMCHEE/webserver-go/server"
	"github.com/SAMCHEE/webserver-go/siteconfig"
)

var (
	FLAG_PORT int
	FLAG_DEV  bool
)

func init() {
	flag.IntVar(&FLAG_PORT, "port", 4000, "port number")
	flag.BoolVar(&FLAG_DEV, "dev", false, "true=localhost")
	flag.Parse()
	if err := siteconfig.Init(FLAG_DEV, uint32(FLAG_PORT)); err != nil {
		panic(err)
	}
}

func main() {
	s := server.NewServer(uint32(FLAG_PORT))
	log.Fatal(s.Run())
}
