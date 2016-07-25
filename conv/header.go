package conv

import "net/http"

func HeaderToMap(r *http.Request) map[string]interface{} {
	return SthToMap(r.Header)
}
