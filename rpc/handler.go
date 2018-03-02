package rpc

import (
	"log"
	"script/ipsecMonitor/client/services"
	"time"
)

type Handler struct {
}

func (this *Handler) RestartIpsec(args *Args, reply *Status) error {
	log.Printf("server notice time : %s", time.Unix(args.Stamp, 0).Format("01-02 15:04:05"))
	succ := services.RestartIpsec()
	reply.Succ = succ
	if succ {
		log.Println("restart ipsec success")
	} else {
		log.Println("restart ipsec faild")
	}
	return nil
}
