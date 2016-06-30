package conv

import (
	"encoding/json"
	"net/http"
)

func HeaderToJson(r *http.Request) string {
	var header = make(map[string]interface{})
	for k, v := range r.Header {
		vint64, err := AssignGetInt64(v[0])
		if err == nil {
			header[k] = vint64
			continue
		}
		vfloat64, err := AssignGetFloat64(v[0])
		if err == nil {
			header[k] = vfloat64
			continue
		}
		header[k] = v[0]
	}
	js, _ := json.Marshal(header)
	return string(js)
}
