package curl

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Requests struct {
	Method     string
	UrlStr     string
	Header     map[string]string
	Raw        string
	RetryTimes int64
	Timeout    int64
}

type Responses struct {
	Response *http.Response
	Body     string
}

//RollingCurl http请求url
func RollingCurl(r Requests) (rp Responses, err error) {
	var i int64 = 0
RELOAD:
	client := &http.Client{
		Timeout: time.Duration(r.Timeout) * time.Second,
	}

	req, err := http.NewRequest(
		r.Method,
		r.UrlStr,
		strings.NewReader(r.Raw),
	)

	if err != nil {
		return rp, err
	}

	for hkey, hval := range r.Header {
		req.Header.Set(hkey, hval)
	}

	resp, err := client.Do(req)
	rp.Response = resp
	if err != nil {
		i++
		if i < r.RetryTimes {
			goto RELOAD
		}
		return rp, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return rp, err
	}

	rp.Body = string(body)
	return rp, nil
}
