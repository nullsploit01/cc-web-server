# CCWS - Command-Line Web Server

This project is a custom implementation of a basic web server built with Go. It was developed as part of a coding challenge [here](https://codingchallenges.fyi/challenges/challenge-webserver/). The server demonstrates how to handle HTTP requests and serve static files efficiently, providing a foundation for learning or extending to more complex web server features.

## Features

- Serve static files from a specified directory.
- Simple and lightweight implementation.
- Handles HTTP GET requests securely, with protections against directory traversal attacks.
- Configurable server port via command-line flags.

## Getting Started

These instructions will help you set up and run the project on your local machine for development and testing purposes.

### Prerequisites

- You need to have Go installed on your machine (Go 1.18 or later is recommended).
- You can download and install Go from [https://golang.org/dl/](https://golang.org/dl/).

### Installing

Clone the repository to your local machine:

```bash
git clone https://github.com/nullsploit01/cc-web-server.git
cd cc-web-server
```

### Building

Compile the project using:

```bash
go build -o ccws
```

### Usage

To run the server, execute the compiled binary. You can specify the port using the --port (or -p) flag.

#### Start the server on the default port (8080):

```bash
./ccws
```

#### Start the server on a custom port (e.g., 9090):

```bash
./ccws --port 9090
```

### Examples of Accessing the Server

    1.	Access the root endpoint to serve index.html:

```bash
curl http://localhost:8080/
```

    2.	Request a specific file (e.g., hello.html):

```bash
curl http://localhost:8080/hello.html
```

    3.	Attempt to access a non-existent file:

```bash
curl http://localhost:8080/nonexistent.html
```

Expected response: 404 Not Found

    4.	Attempt directory traversal (protected):

```bash
curl http://localhost:8080/../../../../etc/passwd
```

Expected response: 403 Forbidden

### Running the Tests

```bash
go test ./...
```
