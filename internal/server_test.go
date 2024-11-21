package internal

import (
	"bytes"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"
)

const testRootDir = "./test_www"

func setup() {
	os.MkdirAll(testRootDir, os.ModePerm)
	os.WriteFile(filepath.Join(testRootDir, "index.html"), []byte("<html><body>Test Index</body></html>"), 0644)
	os.WriteFile(filepath.Join(testRootDir, "forbidden.html"), []byte("<html><body>Forbidden File</body></html>"), 0644)
}

func teardown() {
	os.RemoveAll(testRootDir)
}

type mockConn struct {
	bytes.Buffer
	closed bool
}

func (m *mockConn) Close() error {
	m.closed = true
	return nil
}

func (m *mockConn) RemoteAddr() net.Addr {
	return nil
}

func (m *mockConn) LocalAddr() net.Addr {
	return nil
}

func (m *mockConn) SetDeadline(t time.Time) error {
	return nil
}

func (m *mockConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (m *mockConn) SetWriteDeadline(t time.Time) error {
	return nil
}

func TestServeFile_Success(t *testing.T) {
	setup()
	defer teardown()

	conn := &mockConn{}
	rootDir = testRootDir // Point to the test root directory

	ServeFile(conn, "/index.html")

	if !conn.closed {
		t.Errorf("Expected connection to be closed")
	}

	if !bytes.Contains(conn.Bytes(), []byte("200 OK")) {
		t.Errorf("Expected status 200, got response: %s", conn.String())
	}

	if !bytes.Contains(conn.Bytes(), []byte("Test Index")) {
		t.Errorf("Expected body to contain 'Test Index', got response: %s", conn.String())
	}
}

func TestServeFile_NotFound(t *testing.T) {
	setup()
	defer teardown()

	conn := &mockConn{}
	rootDir = testRootDir

	ServeFile(conn, "/nonexistent.html")

	if !conn.closed {
		t.Errorf("Expected connection to be closed")
	}

	if !bytes.Contains(conn.Bytes(), []byte("404 Not Found")) {
		t.Errorf("Expected status 404, got response: %s", conn.String())
	}
}

func TestHTTPResponse(t *testing.T) {
	conn := &mockConn{}
	HTTPResponse(conn, 200, "Hello, World!", "text/plain")

	if !conn.closed {
		t.Errorf("Expected connection to be closed")
	}

	if !bytes.Contains(conn.Bytes(), []byte("200 OK")) {
		t.Errorf("Expected status 200, got response: %s", conn.String())
	}

	if !bytes.Contains(conn.Bytes(), []byte("Hello, World!")) {
		t.Errorf("Expected body to contain 'Hello, World!', got response: %s", conn.String())
	}
}
