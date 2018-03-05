package services

import (
	"io/ioutil"
	"log"
	"net"
	"regexp"
	"strings"
)

type HaService struct {
	VirtualIpAddress string
	IsMaster         bool
}

var haService *HaService

const (
	KeepAlivedCfgpath = "/etc/keepalived/keepalived.conf"
)

func GetHaService() *HaService {
	if haService != nil {
		return haService
	}
	haService = &HaService{}
	haService.VirtualIpAddress = haService.GetCurrentHaVirtualIpAddress()
	haService.GetCurrentIsMaster()
	return haService
}

func (this *HaService) GetCurrentHaVirtualIpAddress() string {
	data, err := ioutil.ReadFile(KeepAlivedCfgpath)
	if err != nil {
		log.Println(err)
	}

	reg := regexp.MustCompile(`(?m)^.*[\d]*\.[\d]*\.[\d]*\.[\d]*.*$`)
	str := reg.FindAllString(string(data), -1)
	return strings.TrimSpace(str[0])
}

func (this *HaService) GetCurrentIsMaster() bool {
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		if iface.Name == this.GetCurrentHaInterface() {
			addrs, _ := iface.Addrs()
			for _, addr := range addrs {
				if strings.Split(addr.String(), "/")[0] == this.GetCurrentHaVirtualIpAddress() {
					this.IsMaster = true
					return true
				}
			}
		}
	}
	this.IsMaster = false
	return false
}

func (this *HaService) GetCurrentHaInterface() string {
	data, err := ioutil.ReadFile(KeepAlivedCfgpath)
	if err != nil {
		log.Println(err)
	}

	reg := regexp.MustCompile(`(?m)^.*interface.*$`)
	str := reg.FindAllString(string(data), -1)
	strs := strings.Split(strings.TrimSpace(str[0]), " ")
	if err != nil {
		log.Println(err)
	}
	return strs[1]
}
