package conv

import (
	"encoding/json"
	"net/http"

	"github.com/JREAMLU/core/com"
)

//GetHGC header get cookies
func GetHGC(r *http.Request) string {
	header := HeaderToJson(r)
	get := GetToJson(r)
	cookies := CookiesToJson(r)

	result := com.MapMerge(header, get, cookies)

	js, _ := json.Marshal(result)
	return string(js)
}
