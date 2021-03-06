package main

import (
	"os"
	"testing"
)

func TestWriteDefaultConf(t *testing.T) {
	conf := &Config{}
	os.Unsetenv("TYK_GW_LISTENPORT")
	writeDefaultConf(conf)
	if conf.ListenPort != 8080 {
		t.Error("Expected ListenPort to be set to its default")
	}
	*conf = Config{}
	os.Setenv("TYK_GW_LISTENPORT", "9090")
	writeDefaultConf(conf)
	if conf.ListenPort != 9090 {
		t.Error("Expected ListenPort to be set to 9090")
	}
}
