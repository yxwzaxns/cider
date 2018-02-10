package core

import "strings"

func ParsePort(p string) [2][2]string {
	// [[remote_ip:]remote_port[-remote_port]:]port[/protocol]
	// "127.0.0.1:8000:80/tcp"
	// return [[ip,port],[port]]
	res := strings.Split(p, ":")
	portInfo := [2][2]string{}
	switch len(res) {
	case 1:
		portInfo[0][0] = ""
		portInfo[0][1] = ""
		portInfo[1][0] = res[0]
		break
	case 2:
		portInfo[0][0] = "127.0.0.1"
		portInfo[0][1] = res[0]
		portInfo[1][0] = res[1]
		break
	case 3:
		portInfo[0][0] = res[0]
		portInfo[0][1] = res[1]
		portInfo[1][0] = res[2]
		break
	default:

	}
	return portInfo
}
