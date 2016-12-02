package com

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

// IP2Int ip to int
func IP2Int(ip string) int64 {
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

// Int2IP int to ip
func Int2IP(ip int64) string {
	ulMask := [4]int64{0x000000FF, 0x0000FF00, 0x00FF0000, 0xFF000000}
	var result [4]string
	for i := 0; i < 4; i++ {
		result[3-i] = strconv.FormatInt((ip&ulMask[i])>>(uint(i)*8), 10)
	}
	return strings.Join(result[:], ".")
}

// ExternalIP external ip
func ExternalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}
