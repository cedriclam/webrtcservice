package main

import (
	"flag"
	"os"

	"webrtcservice/server"
)

var (
	port             int
	host             string
	databasePathName string
)

const (
	defaultPort             = 8080
	defaultHost             = "0.0.0.0"
	defaultDatabasePathName = "./local.db"
)

func main() {
	flag.IntVar(&port, "p", defaultPort, "server port")
	flag.StringVar(&host, "h", defaultHost, "server listen addresse")
	flag.StringVar(&databasePathName, "db", defaultDatabasePathName, "sqlite db path name")

	flag.Parse()

	os.Exit(run())
}

func run() int {
	config := &server.Config{
		Port: port,
		Host: host,
	}

	srv := server.NewServer(config)
	defer srv.Close()
	if err := srv.Serve(); err != nil {
		return 1
	}
	return 0
}
