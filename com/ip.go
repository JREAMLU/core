package com

import (
	"net"
	"strconv"
	"strings"
)

const (
	LO       = "lo"
	LOOPBACK = "loopback"
	ETH0     = "eth0"
	ETH1     = "eth1"
)

func Ip2Int(ip string) int64 {
	arrary := strings.Split(ip, ".")
	if len(arrary) != 4 {
		return 0
	}
	A, err := strconv.Atoi(arrary[0])
	if err != nil {
		return 0
	}
	B, err := strconv.Atoi(arrary[1])
	if err != nil {
		return 0
	}
	C, err := strconv.Atoi(arrary[2])
	if err != nil {
		return 0
	}
	D, err := strconv.Atoi(arrary[3])
	if err != nil {
		return 0
	}
	return int64(((A*256+B)*256+C)*256 + D)
}

func Int2Ip(ip int64) string {
	ulMask := [4]int64{0x000000FF, 0x0000FF00, 0x00FF0000, 0xFF000000}
	var result [4]string
	for i := 0; i < 4; i++ {
		result[3-i] = strconv.FormatInt((ip&ulMask[i])>>(uint(i)*8), 10)
	}
	return strings.Join(result[:], ".")
}

func GetServerIP() string {
	list, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	var ipMap = make(map[string]string)
	ipMap[LO] = ""
	ipMap[LOOPBACK] = ""
	ipMap[ETH0] = ""
	ipMap[ETH1] = ""

	for _, iface := range list {
		// fmt.Printf("%d name=%s %v\n", i, iface.Name, iface)
		addrs, err := iface.Addrs()
		if err != nil {
			panic(err)
		}
		for _, addr := range addrs {
			// fmt.Printf(" %d %v\n", j, addr)
			name := strings.Split(strings.ToLower(iface.Name), " ")
			ip := strings.Split(addr.String(), "/")

			switch name[0] {
			case LO:
			case LOOPBACK:
			case ETH0:
			case ETH1:
			}
			if name[0] == LO || name[0] == LOOPBACK || name[0] == ETH0 || name[0] == ETH1 {
				ipMap[name[0]] = ip[0]
			}
		}
	}

	if ipMap[ETH0] != "" {
		return ipMap[ETH0]
	}
	if ipMap[ETH1] != "" {
		return ipMap[ETH1]
	}
	if ipMap[LO] != "" {
		return ipMap[LO]
	}
	if ipMap[LOOPBACK] != "" {
		return ipMap[LOOPBACK]
	}

	return "127.0.0.1"
}
