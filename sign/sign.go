package sign

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/JREAMLU/core/crypto"
	"github.com/JREAMLU/core/global"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

//GenerateSign 生成签名 参数key全部按键值排序     ToUpper(md5(sha1(base64(urlencode(SecretKey1Value1Key2Value2SecretTime)))))
//strtoupper( md5 ( sha1( base64_encode( urlencode( secret_key . static::serialize( request_data ) . secret_key . request_time ) ) ) ) )
func GenerateSign(requestData []byte, requestTime int64, secretKey string) string {
	var rdata map[string]interface{}
	json.Unmarshal([]byte(requestData), &rdata)
	str := Serialize(rdata)
	var serial bytes.Buffer
	serial.WriteString(secretKey)
	serial.WriteString(str.(string))
	serial.WriteString(secretKey)
	serial.WriteString(strconv.FormatInt(int64(requestTime), 10))
	urlencodeSerial := url.QueryEscape(serial.String())
	urlencodeBase64Serial := base64.StdEncoding.EncodeToString([]byte(urlencodeSerial))
	sign, _ := crypto.Sha1(urlencodeBase64Serial)
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
		var sortedKeys []string
		for _, key := range keys {
			sortedKeys = append(sortedKeys, key.Interface().(string))
		}
		sort.Strings(sortedKeys)
		for _, key := range sortedKeys {
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

// ValidSignT 签名验证
func ValidSignT(requestData []byte, sign string, timestamp int64, secretKey string) error {
	var rdata map[string]interface{}
	json.Unmarshal(requestData, &rdata)

	jsonData, err := json.Marshal(rdata)
	if err != nil {
		return err
	}

	signed := GenerateSign(jsonData, timestamp, secretKey)

	if sign != signed {
		beego.Info("sign: ", sign, "==", signed)
		return errors.New(i18n.Tr(global.Lang, "sign.INVALIDSIGNATURE"))
	}

	expire, _ := beego.AppConfig.Int64("sign.expire")
	if diff := time.Now().Unix() - timestamp; diff > expire {
		return errors.New(i18n.Tr(global.Lang, "sign.SIGNATURETIMEEXPIRED"))
	}

	return nil
}
