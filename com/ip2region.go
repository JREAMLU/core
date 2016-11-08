package com

import (
	"github.com/mohong122/ip2region/binding/golang"
)

var IP2Region *ip2region.Ip2Region

func InitIP2Region(path string) error {
	var err error
	IP2Region, err = ip2region.New(path)
	defer IP2Region.Close()
	if err != nil {
		return err
	}
	return nil
}

func Query(ipList []string, mode string) (map[string]ip2region.IpInfo, error) {
	var err error
	var ipinfo = make(map[string]ip2region.IpInfo)
	for _, ip := range ipList {
		switch mode {
		case "memory":
			ipinfo[ip], err = IP2Region.MemorySearch(ip)
			if err != nil {
				return nil, err
			}
		case "binary":
			ipinfo[ip], err = IP2Region.BinarySearch(ip)
			if err != nil {
				return nil, err
			}
		case "btree":
			ipinfo[ip], err = IP2Region.BtreeSearch(ip)
			if err != nil {
				return nil, err
			}
		}
	}
	return ipinfo, nil
}
