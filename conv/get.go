package conv

import "net/http"

func GetToJson(r *http.Request) map[string]interface{} {
	r.ParseForm()
	return SthToJson(r.Form)
}
