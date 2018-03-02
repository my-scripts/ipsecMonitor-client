package base

import (
	"os/exec"
	"regexp"
	"strings"
)

type IpsecState struct {
	Name  string
	Index string
	Left  string
	Right string
	State string
}

func parseHeader(header string) (string, string) {
	index := strings.Index(header, "\"")
	if index == -1 {
		return header, ""
	}

	h := header[index+1:]
	h = strings.Replace(h, "\"", "", -1)
	result := strings.Split(h, "/")
	if len(result) != 2 {
		return h, ""
	}
	return result[0], result[1]
}

func parseLine(line string) *IpsecState {
	index := strings.Index(line, ":")
	if index == -1 {
		return nil
	}

	var state IpsecState
	header := line[:index]
	conn, i := parseHeader(header)
	state.Name = conn
	state.Index = i

	left := line[index+1:]
	items := strings.Split(left, ";")
	if len(items) != 3 {
		return nil
	}
	state.State = strings.TrimSpace(items[1])

	items = strings.Split(items[0], "===")
	if len(items) != 3 {
		return nil
	}

	state.Left = items[0]
	state.Right = items[2]
	return &state
}

func GetIpsecConnState() *[]IpsecState {
	cmd := exec.Command("ipsec", "auto", "status")
	o, err := cmd.Output()
	if err != nil {
		return nil
	}

	reg := regexp.MustCompile(`(?m)^.*===.*$`)

	var states []IpsecState
	lines := reg.FindAllString(string(o), -1)
	for _, line := range lines {
		state := parseLine(line)
		if state == nil {
			continue
		}
		states = append(states, *state)
	}
	return &states
}

func (this *IpsecState) GetIpsecConnState() *IpsecState {
	cmd := exec.Command("ipsec", "auto", "status")
	o, err := cmd.Output()
	if err != nil {
		return nil
	}

	reg := regexp.MustCompile(`(?m)^.*===.*$`)

	lines := reg.FindAllString(string(o), -1)
	for _, line := range lines {
		state := parseLine(line)
		if state == nil {
			continue
		}
		if state.Name == this.Name {
			return state
		}

	}
	return nil
}
