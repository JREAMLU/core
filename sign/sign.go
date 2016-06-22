package sign

import (
	"bytes"
	"encoding/json"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/JREAMLU/core/crypto"
)

//GenerateSign 生成签名 参数key全部按键值排序     ToUpper(md5(sha1(SecretKey1Value1Key2Value2SecretTime)))
func GenerateSign(requestData []byte, requestTime int64, secretKey string) string {
	var rdata map[string]interface{}
	json.Unmarshal([]byte(requestData), &rdata)
	str := Serialize(rdata)
	var serial bytes.Buffer
	serial.WriteString(secretKey)
	serial.WriteString(str.(string))
	serial.WriteString(secretKey)
	serial.WriteString(strconv.FormatInt(int64(requestTime), 10))
	sign, _ := crypto.Sha1(serial.String())
	sign, _ = crypto.MD5(sign)

	return strings.ToUpper(sign)
}

// Serialize 序列化 && 递归ksort
func Serialize(data interface{}) interface{} {
	var buffer bytes.Buffer
	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(data)
		for i := 0; i < s.Len(); i++ {
			serial := Serialize(s.Index(i).Interface())
			if reflect.TypeOf(serial).Kind() == reflect.Float64 {
				serial = strconv.Itoa(int(serial.(float64)))
			}
			buffer.WriteString(strconv.Itoa(i))
			buffer.WriteString(serial.(string))
		}
		return buffer.String()
	case reflect.Map:
		s := reflect.ValueOf(data)
		keys := s.MapKeys()
		//ksort
		sorted_keys := make([]string, 0)
		for _, key := range keys {
			sorted_keys = append(sorted_keys, key.Interface().(string))
		}
		sort.Strings(sorted_keys)
		for _, key := range sorted_keys {
			serial := Serialize(s.MapIndex(reflect.ValueOf(key)).Interface())
			if reflect.TypeOf(serial).Kind() == reflect.Float64 {
				serial = strconv.Itoa(int(serial.(float64)))
			}
			buffer.WriteString(key)
			buffer.WriteString(serial.(string))
		}
		return buffer.String()
	}

	return data
}
