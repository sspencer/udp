package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

const (
	defaultPort = 9099
	defaultHost = "localhost"
)

var (
	port int
	host string
)

func init() {
	flag.Usage = usage

	flag.IntVar(&port, "port", defaultPort, "the UDP port")
	flag.StringVar(&host, "host", defaultHost, "the UDP host")

	flag.Parse()
}

func main() {
	if flag.NArg() != 1 {
		usage()
		os.Exit(1)
	}

	text := flag.Args()[0]
	send(host, port, text)
}

func usage() {
	fmt.Fprintf(os.Stderr, "Send \"text\" over UDP\n\nUsage:\n\t%s text\n\n", os.Args[0])
	flag.PrintDefaults()
}

func send(host string, port int, msg string) {
	address := fmt.Sprintf("%s:%v", host, port)
	conn, err := net.Dial("udp", address)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	cnt, err := conn.Write([]byte(msg))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Sent %d bytes to %s\n", cnt, address)
}
