package main

import (
	"flag"
	"os"

	"webrtcservice/server"
)

var (
	port             int
	host             string
	indexTplFilename string
)

const (
	defaultPort             = 8080
	defaultHost             = "0.0.0.0"
	defaultIndexTplFilename = "./src/template/index.html"
)

func main() {
	flag.IntVar(&port, "p", defaultPort, "server port")
	flag.StringVar(&host, "h", defaultHost, "server listen addresse")
	flag.StringVar(&indexTplFilename, "indexTpl", defaultIndexTplFilename, "path to the index.html template")

	flag.Parse()

	os.Exit(run())
}

func run() int {
	config := &server.Config{
		Port:          port,
		Host:          host,
		IndexFileName: indexTplFilename,
	}

	srv := server.NewServer(config)
	defer srv.Close()
	if err := srv.Serve(); err != nil {
		return 1
	}
	return 0
}
