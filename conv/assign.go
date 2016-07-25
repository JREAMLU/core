package conv

import (
	"errors"
	"strconv"
)

func AssignGetInt64(str string) (int64, error) {
	stri, _ := strconv.ParseInt(str, 10, 64)
	strs := strconv.FormatInt(stri, 10)
	if str != strs {
		return stri, errors.New("not equals")
	}
	return stri, nil
}

func AssignGetFloat64(str string) (float64, error) {
	strf, _ := strconv.ParseFloat(str, 64)
	strs := strconv.FormatFloat(strf, 'f', -1, 64)
	if str != strs {
		return strf, errors.New("not equals")
	}
	return strf, nil
}

func SthToJson(sth map[string][]string) map[string]interface{} {
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
