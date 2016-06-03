package curl

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type Requests struct {
	Method string
	UrlStr string
	Header map[string]string
	Raw    string
}

//RollingCurl http请求url
func RollingCurl(r Requests) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest(
		r.Method,
		r.UrlStr,
		strings.NewReader(r.Raw),
	)

	if err != nil {
		return "", err
	}

	for hkey, hval := range r.Header {
		req.Header.Set(hkey, hval)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

//example
/*
func main() {
	res, err := RollingCurl(
		Requests{
			Method: "POST",
			UrlStr: "http://localhost/study/curl/servera.php",
			Header: map[string]string{
				"Content-Type": "application/json",
			},
			Raw: `{"name":"KII","age":24}`,
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	var result = make(map[string]interface{})
	json.Unmarshal([]byte(res), &result)
	for k, v := range result {
		fmt.Printf("%s: %v \n", k, v)
	}
}
*/
