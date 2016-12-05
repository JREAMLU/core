package conv

import (
	"errors"
	"strconv"
)

// AssignGetInt64 assgin int64
func AssignGetInt64(str string) (int64, error) {
	stri, _ := strconv.ParseInt(str, 10, 64)
	strs := strconv.FormatInt(stri, 10)
	if str != strs {
		return stri, errors.New("not equals")
	}
	return stri, nil
}

// AssignGetFloat64 assgin float64
func AssignGetFloat64(str string) (float64, error) {
	strf, _ := strconv.ParseFloat(str, 64)
	strs := strconv.FormatFloat(strf, 'f', -1, 64)
	if str != strs {
		return strf, errors.New("not equals")
	}
	return strf, nil
}

// SthToMap map to map
func SthToMap(sth map[string][]string) map[string]interface{} {
	var sths = make(map[string]interface{})
	for k, v := range sth {
		vint64, err := AssignGetInt64(v[0])
		if err == nil {
			sths[k] = vint64
			continue
		}
		vfloat64, err := AssignGetFloat64(v[0])
		if err == nil {
			sths[k] = vfloat64
			continue
		}
		sths[k] = v[0]
	}
	return sths
}

// SthsToInt64 map to map
func SthsToInt64(sth map[string]interface{}) map[string]interface{} {
	var sths = make(map[string]interface{})
	for k, v := range sth {
		switch v.(type) {
		case float64:
			v = int64(v.(float64))
		}
		sths[k] = v
	}
	return sths
}
