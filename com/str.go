package com

import "bytes"

// StringJoin string join
func StringJoin(content ...string) string {
	var str bytes.Buffer
	for _, cnt := range content {
		str.WriteString(cnt)
	}

	return str.String()
}
