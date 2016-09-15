# UDP

Simple command line application to send text to a specified UDP server

Usage:

    Send "text" over UDP

    Usage:
        udp text

    -host string
      the UDP host (default "localhost")
    -port int
        the UDP port (default 9099)

Example:

    udp -port 1234 red,green,blue


## Build Instructions

Go is easy.

Run from current directory:

    go run main.go

Install into go/bin directory

    go install
