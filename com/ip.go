package com

import (
	"strconv"
	"strings"
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
