package conv

import "net/http"

// GetToMap get to map
func GetToMap(r *http.Request) map[string]interface{} {
	r.ParseForm()
	return SthToMap(r.Form)
}
