package proxy

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {
	nginxServer := newNginxServer()
	appStatusURI, createUserURI := "/app/status", "/create/user"

	httpCode, body := nginxServer.handleRequest(appStatusURI, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURI, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURI, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURI, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURI, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURI, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createUserURI, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURI, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createUserURI, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURI, httpCode, body)
}
