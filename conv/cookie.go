package conv

import "net/http"

// CookiesToMap cookies to map
func CookiesToMap(r *http.Request) map[string]interface{} {
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
	return cookies
}
