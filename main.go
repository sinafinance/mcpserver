package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	log.SetOutput(os.Stderr) // 避免污染 stdout

	s := server.NewMCPServer(
		"demo",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	tool := mcp.NewTool("hello_world",
		mcp.WithDescription("Say hello to someone"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Name of the person to greet"),
		),
	)

	s.AddTool(tool, helloHandler)

	log.Println("MCP server started, waiting for JSON-RPC requests...")

	if err := server.ServeStdio(s); err != nil {
		log.Printf("Server error: %v", err)
	}
}

func helloHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, ok := request.GetArguments()["name"].(string)
	if !ok {
		return nil, errors.New("name must be a string")
	}
	return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
}
