package com

func MapMerge(ms ...map[string]interface{}) map[string]interface{} {
	var nm = make(map[string]interface{})
	for _, m := range ms {
		for k, v := range m {
			mu.Lock()
			nm[k] = v
			mu.Unlock()
		}
	}
	return nm
}
