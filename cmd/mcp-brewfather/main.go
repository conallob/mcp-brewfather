package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/conallob/mcp-brewfather/internal/brewfather"
	"github.com/conallob/mcp-brewfather/internal/tools"
	"github.com/mark3labs/mcp-go/server"
)

var version = "dev"

func main() {
	vFlag := flag.Bool("version", false, "Print version and exit")
	flag.Parse()

	if *vFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	userID := os.Getenv("BREWFATHER_USER_ID")
	apiKey := os.Getenv("BREWFATHER_API_KEY")

	if userID == "" || apiKey == "" {
		log.Fatal("BREWFATHER_USER_ID and BREWFATHER_API_KEY environment variables are required")
	}

	client := brewfather.NewClient(userID, apiKey)

	s := server.NewMCPServer(
		"mcp-brewfather",
		version,
		server.WithToolCapabilities(false),
		server.WithRecovery(),
	)

	tools.Register(s, client)

	log.Printf("mcp-brewfather %s starting (stdio)", version)
	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
