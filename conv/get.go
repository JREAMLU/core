package conv

import (
	"encoding/json"
	"net/http"
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
