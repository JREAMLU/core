package com

// MapMerge map merge
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

// EqualMapInt equal map int
func EqualMapInt(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

// EqualMapInt64 equal map int64
func EqualMapInt64(x, y map[string]int64) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

// EqualMapString equal map string
func EqualMapString(x, y map[string]string) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

// EqualMapFloat32 equal map float32
func EqualMapFloat32(x, y map[string]float32) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

// EqualMapFloat64 equal map float64
func EqualMapFloat64(x, y map[string]float64) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

// EqualMapInterface equal map interface
func EqualMapInterface(x, y map[string]interface{}) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
