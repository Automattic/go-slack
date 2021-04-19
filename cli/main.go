package main

import (
	"fmt"
	"os"

	"github.com/Automattic/go-slack"
)

func main() {
	token := os.Getenv("TOKEN")
	if token == "" {
		fmt.Printf("Missing TOKEN environment variable\n")
		os.Exit(1)
	}

	if len(os.Args) < 4 {
		usage()
		os.Exit(2)
	}

	channel := os.Args[1]

	messages := []slack.Message{
		slack.Message{
			Title: os.Args[2],
			Text:  os.Args[3],
		},
	}

	client := slack.NewClient(token)
	err := client.Send(channel, messages)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

func usage() {
	fmt.Printf("Usage: %s <channel> <title> <message>\n", os.Args[0])
}
