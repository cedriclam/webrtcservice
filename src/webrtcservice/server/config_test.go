package server

import "testing"

func TestGetAddr(t *testing.T) {
	c := &Config{
		Host: "127.0.0.1",
		Port: 32000,
	}
	if res := c.getAddr(); res != "127.0.0.1:32000" {
		t.Error("wrong getAddr, current:", res)
	}
}
