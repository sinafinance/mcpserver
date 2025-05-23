package main

import (
	"flag"
	"log"
	tool "mcpserver/mcp/tool/forexExchangeTool"
	"os"
	"strings"

	"github.com/mark3labs/mcp-go/server"
)

const MethodStdio = "stdio"
const MethodSSE = "sse"

func main() {
	method := flag.String("m", "stdio", "部署方式：sse或者stdio，默认为stdio")
	address := flag.String("addr", "localhost:12345", "服务地址：sse模式需要，默认为 localhost:12345")
	s := server.NewMCPServer(
		"sinafinance-mcpserver",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	tool.NewForexExchangeRateTool(s)
	switch strings.ToLower(*method) {
	case MethodStdio:
		startStdio(s)
	case MethodSSE:
		startSSE(s, *address)
	default:
		log.Printf("not support deployment method:%s\n", *method)
	}
}

func startStdio(s *server.MCPServer) {
	log.SetOutput(os.Stderr) // 避免污染 stdout

	if err := server.ServeStdio(s); err != nil {
		panic(err)
	}
}

func startSSE(s *server.MCPServer, addr string) {
	sse := server.NewSSEServer(s)
	if err := sse.Start(addr); err != nil {
		panic(err)
	}
}
