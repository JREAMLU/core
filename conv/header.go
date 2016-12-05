package conv

import "net/http"

// HeaderToMap header to map
func HeaderToMap(r *http.Request) map[string]interface{} {
	return SthToMap(r.Header)
}
