package conv

import "net/http"

func HeaderToJson(r *http.Request) map[string]interface{} {
	return SthToJson(r.Header)
}
