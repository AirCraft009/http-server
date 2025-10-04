package server

import (
	"fmt"
	"http-server/parser"
	"net"
	"strconv"
)

const (
	maxHeaderSize = 1024
	baseResponse  = "HTTP/1.1 200 OK\r\nContent-Length: 5\r\n\r\nHello"
)

// Server
// supports tcp connections/*
type Server struct {
	port     int
	listener net.Listener
}

func checkErorr(err error) {
	if err != nil {
		panic(err)
	}
}

func NewServer(port int) *Server {
	return &Server{port, net.Listener(nil)}
}

func (s *Server) Listen() {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(s.port))
	checkErorr(err)
	s.listener = listener
}

func (s *Server) AcceptConnections() {
	if s.listener == nil {
		s.Listen()
	}
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting conn: %s\n", err.Error())
		}
		go handleConnection(conn, s)
	}
}

func handleConnection(conn net.Conn, s *Server) {
	//min size for header 20B max 60B
	header := make([]byte, maxHeaderSize)
	//n is the ammount of bytes that are relevant so only header[n:] is important!!
	n, err := conn.Read(header)
	if err != nil {
		fmt.Printf("Error reading header: %s\nerr: %s\n", conn, err.Error())
	}
	//read everything up to n (inklusive)
	//fmt.Println(string(header[:n]))
	req, _ := parser.ParseRequest(string(header[:n]))
	fmt.Println(req)
	s.sendString(baseResponse, conn)
	defer func(conn net.Conn) {
		err := conn.Close()
		checkErorr(err)
	}(conn)
}

func (s *Server) sendString(html string, conn net.Conn) {
	_, err := conn.Write([]byte(html))
	checkErorr(err)
	select {}
}

func (s *Server) Close() {
	err := s.listener.Close()
	checkErorr(err)
}
