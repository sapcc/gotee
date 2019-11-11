package main

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"io"
	"net"
)

var (
	listenAddr = kingpin.Flag("listen", "Listen port.").Short('l').Default(":8000").String()
	outAddr1   = kingpin.Flag("output port1", "Output port 1.").Short('1').Default(":8001").String()
	outAddr2   = kingpin.Flag("output port2", "Output port 2.").Short('2').Default(":8002").String()
	debug      = kingpin.Flag("debug", "Debug.").Short('d').Default("false").Bool()
)

func init() {
	kingpin.Parse()
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
	})
	//log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	listener, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("listening on ", *listenAddr)
		defer listener.Close()
	}

	// handle incoming connections one by one
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Error(err)
		} else {
			// first out connection is required
			c1, err := net.Dial("tcp", *outAddr1)
			if err != nil {
				log.Error(err)
				continue
			}
			// second out connection is optional
			c2, err := net.Dial("tcp", *outAddr2)
			if err != nil {
				log.Warn(err)
			}
			if c2 == nil {
				go handleConnection(conn, c1)
			} else {
				go handleConnection2(conn, c1, c2)
			}
		}
	}
}

func handleConnection(conn, out net.Conn) {
	defer func() {
		conn.Close()
		out.Close()
	}()

	s := bufio.NewScanner(conn)
	for s.Scan() {
		_, err := out.Write(append(s.Bytes(), '\n'))
		if err != nil {
			log.Info(err)
			return
		}
	}
}

func handleConnection2(in, o1, o2 net.Conn) {
	defer func() {
		in.Close()
		o1.Close()
		o2.Close()
	}()

	dw := io.MultiWriter(o1, o2)

	s := bufio.NewScanner(in)
	for s.Scan() {
		_, err := dw.Write(append(s.Bytes(), '\n'))
		if err != nil {
			log.Info(err)
			return
		}
	}
}
