package conv

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func GetToJson(r *http.Request) string {
	var get = make(map[string]interface{})
	r.ParseForm()
	for k, v := range r.Form {
		vint64, err := AssignGetInt64(v[0])
		if err == nil {
			get[k] = vint64
			continue
		}
		vfloat64, err := AssignGetFloat64(v[0])
		if err == nil {
			get[k] = vfloat64
			continue
		}
		get[k] = v[0]
	}
	js, _ := json.Marshal(get)
	return string(js)
}

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
