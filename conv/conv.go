package conv

import "net/http"

//GetHGC header get cookies
func GetHGC(r *http.Request) (string, string, string) {
	header := HeaderToJson(r)
	get := GetToJson(r)
	cookies := CookiesToJson(r)

	return cookies, header, get
}
