package curl

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Requests struct
type Requests struct {
	HTTPClient *http.Client
}

// Responses struct
type Responses struct {
	Response *http.Response
	Body     string
	Data     interface{}
}

const (
	// MaxIdleConnections maxIdleConnections
	MaxIdleConnections = 100
	// MaxConnectionIdleTime 连接池中一个连接可以idle的时长
	MaxConnectionIdleTime = 60 * time.Second
	timeout               = 3
)

// NewRequests new requests
func NewRequests() *Requests {
	return &Requests{
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost: MaxIdleConnections,
				IdleConnTimeout:     MaxConnectionIdleTime,
			},
			Timeout: time.Duration(timeout) * time.Second,
		},
	}
}

//RollingCurl http请求url
func (r *Requests) RollingCurl(Method string, URLStr string, Header map[string]string, Raw string, RetryTimes int64, data interface{}) (rp Responses, err error) {
	var i int64
	req, err := http.NewRequest(
		Method,
		URLStr,
		strings.NewReader(Raw),
	)

	if err != nil {
		return rp, err
	}

	for hkey, hval := range Header {
		req.Header.Set(hkey, hval)
	}

RELOAD:

	resp, err := r.HTTPClient.Do(req)
	if err != nil {
		i++
		if i < RetryTimes {
			goto RELOAD
		}
		return rp, err
	}
	rp.Response = resp

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return rp, err
	}

	rp.Body = string(body)

	if data != nil {
		err = json.Unmarshal(body, data)
		if err != nil {
			return rp, err
		}

		rp.Data = data
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return rp, err
	}

	rp.Data = data

	return rp, nil
}
