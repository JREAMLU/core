package conv

import (
	"encoding/json"
	"net/http"
)

func CookiesToJson(r *http.Request) string {
	var cookies = make(map[string]interface{})
	for _, v := range r.Cookies() {
		vint64, err := AssignGetInt64(v.Value)
		if err == nil {
			cookies[v.Name] = vint64
			continue
		}
		vfloat64, err := AssignGetFloat64(v.Value)
		if err == nil {
			cookies[v.Name] = vfloat64
			continue
		}
		cookies[v.Name] = v.Value
	}
	js, _ := json.Marshal(cookies)
	return string(js)
}
