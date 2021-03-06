package services

import (
	"log"
	"path/filepath"
	"script/ipsecMonitor/client/config"
	"time"

	"script/ipsecMonitor/client/base"
)

type Monitor struct {
	State bool
}

func (this *Monitor) Run() {
	cfg := config.Config{}
	err := cfg.Load(filepath.Join("conf", "imclient.json"))
	if err != nil {
		log.Panicln("failed to load config: %s", err)
		return
	}
	timeout := cfg.Monitor.Timeout
	d := time.Duration(5) * time.Second
	for {
		if !this.MonitorLinkStatus() {
			if timeout > 0 {
				ics := base.GetIpsecConnState()
				if ics == nil {
					time.Sleep(d)
					continue
				}
				for _, v := range *ics {
					if !(v.State == "erouted") {
						log.Printf("%+v\n", v)
					}
				}
				timeout = timeout - 5
			}
			this.State = false
		} else {
			if !this.State {
				log.Println("ipsec is all erouted")
				this.State = true
			}
			timeout = cfg.Monitor.Timeout
		}
		time.Sleep(d)
	}
}
func (this *Monitor) MonitorLinkStatus() bool {
	ics := base.GetIpsecConnState()
	if ics == nil {
		return false
	}
	for _, v := range *ics {
		if !(v.State == "erouted") {
			return false
		}
	}
	return true
}
