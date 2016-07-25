package conv

import "net/http"

func GetToMap(r *http.Request) map[string]interface{} {
	r.ParseForm()
	return SthToMap(r.Form)
}
