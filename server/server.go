package server

import (
	"fmt"
	"http-server/parser"
	"net"
	"strconv"
)

// Server
// supports tcp connections/*
type Server struct {
	port     int
	listener net.Listener
	parser   *parser.Parser
}

func checkErorr(err error) {
	if err != nil {
		panic(err)
	}
}

func NewServer(port int) *Server {
	return &Server{port, net.Listener(nil), parser.NewParser()}
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
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	//min size for header 20B max 60B
	header := make([]byte, 60)
	//n is the ammount of bytes that are relevant so only header[n:] is important!!
	n, err := conn.Read(header)
	if err != nil {
		fmt.Printf("Error reading header: %s\nerr: %s\n", conn, err.Error())
	}
	//read everything up to n (inklusive)
	fmt.Println(string(header[:n]))
	defer func(conn net.Conn) {
		err := conn.Close()
		checkErorr(err)
	}(conn)
}

func (s *Server) Close() {
	err := s.listener.Close()
	checkErorr(err)
}
