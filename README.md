# gowebsocket

A simple WebSocket echo server and client built with Go and [coder/websocket](https://github.com/coder/websocket).

## How it works

- **Server** (`server.go`): Accepts WebSocket connections on `localhost:8080` and echoes back any message it receives.
- **Client** (`client.go`): Connects to the server, sends `"hello"`, reads the echoed response, and prints it.
- **Main** (`main.go`): Starts the HTTP server and immediately runs the client as a goroutine.

## Running

```bash
go run .
```

You should see output like:

```
Server started
received hello
```

The server receives the message and echoes it back; the client prints the response.

## Dependencies

- [github.com/coder/websocket](https://github.com/coder/websocket) — modern, context-based WebSocket library for Go.
