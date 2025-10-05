package server

import (
	"fmt"
	"http-server/parser"
	"io"
	"net"
	"strconv"
	"strings"
)

const (
	maxHeaderSize = 1024
	baseResponse  = "HTTP/1.1 200 OK\r\nContent-Length: 5\r\n\r\nHello"
	httpType0     = "HTTP/1.0"
	httpType1     = "HTTP/1.1"
)

// Server
// supports tcp connections/*
type Server struct {
	port     int
	listener net.Listener
	router   *Router
}

type Response struct {
	StatusCode int
	HTTPType   string
	Headers    map[string]string
	Body       []byte
}

func checkErorr(err error) {
	if err != nil {
		panic(err)
	}
}

func NewServer(port int) *Server {
	return &Server{port, net.Listener(nil), NewRouter()}
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
	//loops because keep alive is the standard so if it's not changed and there isn't a EOF
	//End Of File then it just continues
	for {
		//min size for header 20B max 60B
		message := make([]byte, maxHeaderSize)
		//n is the ammount of bytes that are relevant so only header[n:] is important!!
		n, err := conn.Read(message)
		//break at EOF because it means the clien terminated the connection
		if err == io.EOF {
			break
		} else if err != nil {
			checkErorr(err)
		}
		//read everything up to n (inklusive)
		//fmt.Println(string(header[:n]))
		req, err := parser.ParseRequest(message[:n])
		checkErorr(err)
		response := s.router.useRoute(req.Path, req)
		_, err = conn.Write(formatResponse(response))
		if err != nil {
			return
		}
		if !(req.HTTPType == "HTTP/1.0" || req.Headers["Connection"] == "close") {
			continue
		}
		defer func(conn net.Conn) {
			err := conn.Close()
			checkErorr(err)
		}(conn)
		break
	}
}

func (s *Server) sendString(html string, conn net.Conn) {
	_, err := conn.Write([]byte(html))
	checkErorr(err)
}

func (s *Server) Handle(Method, path string, handler Handler) {
	s.router.Handle(Method, path, handler)
}

// AddFileSystem
/**
this method adds a folder and all folders under it to the filesystem of the server
any requests the go from the basepath to any file in the system using a get Request can get pulled
*/
func (s *Server) AddFileSystem(folderPath string) {
	if strings.HasSuffix(folderPath, "/") {
		folderPath = folderPath[:len(folderPath)-1]
	}
	s.router.sourceFolder = folderPath
	files := listFiles(folderPath)
	for _, file := range files {
		file = strings.Replace(file, folderPath, "", 1)
		file = strings.ReplaceAll(file, "\\", "/")
		s.router.Handle("GET", file, StreamHandler)
	}
	fmt.Println(s.router.Routes)
}

func (s *Server) Close() {
	err := s.listener.Close()
	checkErorr(err)
}
