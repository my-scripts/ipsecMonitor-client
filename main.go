package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"path/filepath"
	"script/ipsecMonitor/client/config"
	localrpc "script/ipsecMonitor/client/rpc"
)

func main() {
	cfg := config.Config{}
	err := cfg.Load(filepath.Join("conf", "imclient.json"))
	if err != nil {
		log.Panicln("failed to load config: %s", err)
		return
	}

	// rpc
	rpc.Register(&localrpc.Handler{})
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp4", fmt.Sprintf(":%d", cfg.Rpc.Port))
	if err != nil {
		log.Panicln("unable to listen", err)
		return
	}
	http.Serve(listener, nil)
}
