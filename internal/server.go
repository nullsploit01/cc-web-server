package internal

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

type Server struct {
	port  string
	store map[string]string
	mu    sync.RWMutex
}

func InitServer(port string) *Server {
	return &Server{
		port:  port,
		store: make(map[string]string),
	}
}

func (s *Server) StartServer() error {
	l, err := net.Listen("tcp", ":"+s.port)
	fmt.Printf("Listening on port %s\n", s.port)
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go s.handleConnection(c)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	fmt.Printf("Accepted connection from %s\n", conn.RemoteAddr().String())

	for {
		command, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading from client:", err)
			}
			break
		}

		log.Println(command)
	}
}
