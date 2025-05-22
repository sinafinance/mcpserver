package main

import (
	"log"
	tool "mcpserver/mcp/tool/forexExchangeTool"
	"os"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	log.SetOutput(os.Stderr) // 避免污染 stdout

	s := server.NewMCPServer(
		"sinafinance-mcpserver",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	tool.NewForexExchangeRateTool(s)

	log.Println("MCP server started, waiting for JSON-RPC requests...")

	if err := server.ServeStdio(s); err != nil {
		log.Printf("Server error: %v", err)
	}
}
