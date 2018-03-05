package services

import (
	"log"
	"os/exec"
)

func RestartIpsec() bool {
	cmd := exec.Command("/etc/init.d/ipsec", "restart")
	err := cmd.Run()
	if err != nil {
		log.Println("restart ipsec faild :", err)
	}
	return err == nil

}
