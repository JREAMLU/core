package com

func MapMerge(ms ...map[string]interface{}) map[string]interface{} {
	var nm = make(map[string]interface{})
	for _, m := range ms {
		for k, v := range m {
			nm[k] = v
		}
	}
	return nm
}
