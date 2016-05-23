package server

import (
	"fmt"
	"net"
)

// Config Server configuration parameters
type Config struct {
	Port int
	Host string
}

func getPort(port int) string {
	return fmt.Sprintf("%d", port)
}

func (c Config) getAddr() string {
	return net.JoinHostPort(c.Host, getPort(c.Port))
}
