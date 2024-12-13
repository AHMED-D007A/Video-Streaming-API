package main

import "Video-Streaming-Platform/server"

func main() {
	server := server.NewServer(":8080", nil)

	if err := server.Start(); err != nil {
		panic(err)
	}
}
