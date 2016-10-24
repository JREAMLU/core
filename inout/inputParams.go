package inout

import (
	"bytes"
	jcontext "context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/JREAMLU/core/com"
	"github.com/JREAMLU/core/global"
	"github.com/JREAMLU/core/guid"
	"github.com/JREAMLU/core/sign"
	"github.com/beego/i18n"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/validation"
	"github.com/pquerna/ffjson/ffjson"
)

type Header struct {
	Source      []string `json:"Source" valid:"Required"`
	Version     []string `json:"Version" `
	SecretKey   []string `json:"Secret-Key" `
	RequestID   []string `json:"Request-Id" valid:"Required"`
	ContentType []string `json:"Content-Type" valid:"Required"`
	Accept      []string `json:"Accept" valid:"Required"`
	Token       []string `json:"Token" valid:"Required"`
	Timestamp   []string `json:"Timestamp" valid:"Required"`
	IP          []string `json:"Ip" valid:"Required"`
}

type Result struct {
	CheckRes  map[string]string
	RequestID string
	Message   string
}

// var Jctx jcontext.Context

func InputParams(r *context.Context) (map[string]interface{}, jcontext.Context) {
	r.Request.ParseForm()

	headerMap := r.Request.Header
	if _, ok := headerMap["Request-Id"]; !ok {
		rid := GetRequestID()
		headerMap["Request-Id"] = []string{rid}
	}
	jctx := jcontext.WithValue(jcontext.Background(), "requestID", headerMap["Request-Id"][0])
	header, _ := json.Marshal(headerMap)
	body := r.Input.RequestBody
	cookiesSlice := r.Request.Cookies()
	cookies, _ := json.Marshal(cookiesSlice)
	querystrMap := r.Request.Form
	querystrJson, _ := json.Marshal(querystrMap)
	querystring := r.Request.RequestURI

	beego.Trace(jctx.Value("requestID").(string), ":", "input params header", string(header))
	beego.Trace(jctx.Value("requestID").(string), ":", "input params body", string(body))
	beego.Trace(jctx.Value("requestID").(string), ":", "input params cookies", string(cookies))
	beego.Trace(jctx.Value("requestID").(string), ":", "input params querystrJson", string(querystrJson))
	beego.Trace(jctx.Value("requestID").(string), ":", "input params querystring", string(querystring))

	data := make(map[string]interface{})
	mu.Lock()
	data["header"] = header
	data["body"] = body
	data["cookies"] = string(cookies)
	data["querystrjson"] = string(querystrJson)
	data["headermap"] = headerMap
	data["cookiesslice"] = cookiesSlice
	data["querystrmap"] = querystrMap
	data["querystring"] = querystring
	mu.Unlock()

	return data, jctx
}

/**
 *	@auther		jream.lu
 *	@intro		入参验证
 *	@logic
 *	@todo		返回值
 *	@meta		meta map[string][]string	   rawMetaHeader
 *	@data		data []byte 					rawDataBody 签名验证
 *	@data		data ...interface{}	切片指针	rawDataBody
 *	@return 	返回 true, metaMap, error
 */
func InputParamsCheck(jctx jcontext.Context, data map[string]interface{}, stdata ...interface{}) (result Result, err error) {
	headerRes, err := HeaderCheck(jctx, data)
	timestamp, _ := strconv.ParseInt(headerRes.CheckRes["Timestamp"], 10, 64)
	token := headerRes.CheckRes["Token"]
	if err != nil {
		return headerRes, err
	}

	result.CheckRes = nil
	result.Message = ""
	result.RequestID = headerRes.RequestID

	valid := validation.Validation{}

	for _, val := range stdata {
		is, err := valid.Valid(val)
		if err != nil {
			beego.Trace(
				jctx.Value("requestID").(string), ":",
				i18n.Tr(global.Lang, "outputParams.SYSTEMILLEGAL"),
				err.Error(),
			)
			result.Message = i18n.Tr(global.Lang, "outputParams.SYSTEMILLEGAL")
			return result, err
		}

		if !is {
			for _, err := range valid.Errors {
				beego.Trace(
					jctx.Value("requestID").(string), ":",
					i18n.Tr(global.Lang, "outputParams.DATAPARAMSILLEGAL"),
					err.Key, ":", err.Message,
				)
				result.Message = com.StringJoin(
					i18n.Tr(global.Lang, "outputParams.DATAPARAMSILLEGAL"),
					" ", err.Key, ":", err.Message,
				)
				return result, errors.New(i18n.Tr(global.Lang, "outputParams.DATAPARAMSILLEGAL"))
			}
		}
	}

	if is, _ := beego.AppConfig.Bool("sign.onOff"); is {
		err = sign.ValidSignT(data["body"].([]byte), token, timestamp, beego.AppConfig.String("sign.secretKey"))
		if err != nil {
			result.Message = err.Error()
			return result, err
		}
	}

	return headerRes, nil
}

/**
 * header参数验证
 * 将header 放入map 返回
 *
 * @meta 	meta  map[string][]string 	header信息 map格式
 */
func HeaderCheck(jctx jcontext.Context, data map[string]interface{}) (result Result, err error) {
	var h Header
	ffjson.Unmarshal(data["header"].([]byte), &h)

	result.CheckRes = nil
	result.Message = ""
	result.RequestID = jctx.Value("requestID").(string)

	ct, err := HeaderParamCheck(h.ContentType, "Content-Type")
	if err != nil {
		ct.RequestID = jctx.Value("requestID").(string)
		return ct, err
	}

	at, err := HeaderParamCheck(h.Accept, "Accept")
	if err != nil {
		at.RequestID = jctx.Value("requestID").(string)
		return at, err
	}

	valid := validation.Validation{}

	is, err := valid.Valid(&h)

	if err != nil {
		beego.Trace(
			jctx.Value("requestID").(string), ":",
			i18n.Tr(global.Lang, "outputParams.SYSTEMILLEGAL"),
			err.Error(),
		)
		result.Message = i18n.Tr(global.Lang, "outputParams.SYSTEMILLEGAL")

		return result, err
	}

	if !is {
		for _, err := range valid.Errors {
			beego.Trace(
				jctx.Value("requestID").(string), ":",
				i18n.Tr(global.Lang, "outputParams.METAPARAMSILLEGAL"),
				err.Key, ":", err.Message,
			)

			result.Message = com.StringJoin(
				i18n.Tr(global.Lang, "outputParams.METAPARAMSILLEGAL"),
				" ", err.Key, ":", err.Message,
			)

			return result, errors.New(i18n.Tr(global.Lang, "outputParams.METAPARAMSILLEGAL"))
		}
	}

	var headerMap = make(map[string]string)
	for key, val := range data["headermap"].(http.Header) {
		headerMap[key] = val[0]
	}
	headerMap["request-id"] = jctx.Value("requestID").(string)
	result.CheckRes = headerMap

	return result, nil
}

//HeaderParamCheck 验证header固定信息
func HeaderParamCheck(h []string, k string) (result Result, err error) {
	if strings.ToLower(h[0]) != strings.ToLower(beego.AppConfig.String(k)) {
		message := ""
		switch k {
		case "Content-Type":
			message = i18n.Tr(global.Lang, "outputParams.CONTENTTYPEILLEGAL")
		case "Accept":
			message = i18n.Tr(global.Lang, "outputParams.ACCEPTILLEGAL")
		}

		result.CheckRes = nil
		result.Message = message
		return result, errors.New(message)
	}

	return result, nil
}

//request id增加
func GetRequestID() string {
	var requestID bytes.Buffer
	requestID.WriteString(beego.AppConfig.String("appname"))
	requestID.WriteString("-")
	requestID.WriteString(guid.NewObjectId().Hex())
	return requestID.String()
}
