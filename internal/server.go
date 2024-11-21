package internal

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

const rootDir = "./www"

type Server struct {
	port string
}

func InitServer(port string) *Server {
	return &Server{
		port: port,
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
		go HandleConnection(c)
	}
}

func ServeFile(conn net.Conn, path string) {
	cleanPath := filepath.Clean(path)
	fullPath := filepath.Join(rootDir, cleanPath)

	if !strings.HasPrefix(fullPath, filepath.Clean(rootDir)) {
		HTTPResponse(conn, 403, "Forbidden", "text/plain")
		return
	}

	content, err := os.ReadFile(fullPath)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			HTTPResponse(conn, 404, "File Not Found", "text/plain")
		} else {
			HTTPResponse(conn, 500, "Internal Server Error", "text/plain")
		}
		return
	}
	HTTPResponse(conn, 200, string(content), "text/html")
}

func HTTPResponse(conn net.Conn, statusCode int, body string, contentType string) {
	statusText := map[int]string{
		200: "OK",
		404: "Not Found",
		500: "Internal Server Error",
	}[statusCode]

	response := fmt.Sprintf(
		"HTTP/1.1 %d %s\r\nContent-Length: %d\r\nContent-Type: %s\r\n\r\n%s",
		statusCode, statusText, len(body), contentType, body,
	)

	conn.Write([]byte(response))
	conn.Close()
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading request:", err)
		return
	}

	requestLine = strings.TrimSpace(requestLine)
	parts := strings.Split(requestLine, " ")
	if len(parts) < 3 {
		HTTPResponse(conn, 400, "Bad Request", "text/plain")
		return
	}

	method, path, protocol := parts[0], parts[1], parts[2]
	fmt.Printf("Method: %s, Path: %s, Protocol: %s\n", method, path, protocol)

	if protocol != "HTTP/1.1" {
		HTTPResponse(conn, 505, "HTTP Version Not Supported", "text/plain")
		return
	}

	if method != "GET" {
		HTTPResponse(conn, 405, "Method Not Allowed", "text/plain")
		return
	}

	if path == "/" {
		path = "/index.html"
	}
	filePath := "." + path
	ServeFile(conn, filePath)
}
