package inout

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/JREAMLU/core/global"
	"github.com/JREAMLU/core/sign"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/validation"
	"github.com/beego/i18n"
	"github.com/pquerna/ffjson/ffjson"
)

type MetaHeader struct {
	Source      []string `json:"Source" valid:"Required"`
	Version     []string `json:"Version" valid:"Required"`
	SecretKey   []string `json:"Secret-Key" valid:"Required"`
	RequestID   []string `json:"Request-ID" valid:"Required"`
	ContentType []string `json:"Content-Type" valid:"Required"`
	Accept      []string `json:"Accept" valid:"Required"`
	Token       []string `json:"Token" valid:"Required"`
	IP          []string `json:"Ip" valid:"Required"`
}

type Result struct {
	MetaCheckResult map[string]string
	RequestID       string
	Message         string
}

func InputParams(r *context.BeegoInput) (http.Header, []byte) {
	rawMetaHeader := r.Context.Request.Header
	rawDataBody := r.RequestBody

	js, _ := json.Marshal(rawMetaHeader)

	//记录参数日志
	beego.Trace("input params header :" + string(js))
	beego.Trace("input params body :" + string(rawDataBody))

	return rawMetaHeader, rawDataBody
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
func InputParamsCheck(meta map[string][]string, rawDataBody []byte, data ...interface{}) (result Result, err error) {
	//MetaHeader check
	metaCheckResult, err := MetaHeaderCheck(meta)
	if err != nil {
		return metaCheckResult, err
	}

	//DataParams check
	valid := validation.Validation{}

	for _, val := range data {
		is, err := valid.Valid(val)

		//日志

		//检查参数
		if err != nil {
			// handle error
			beego.Trace(i18n.Tr(global.Lang, "outputParams.SYSTEMILLEGAL") + err.Error())
		}

		if !is {
			for _, err := range valid.Errors {
				beego.Trace("input params body check : " + i18n.Tr(global.Lang, "outputParams.DATAPARAMSILLEGAL") + err.Key + ":" + err.Message)
				result.MetaCheckResult = nil
				result.RequestID = metaCheckResult.MetaCheckResult["request-id"]
				result.Message = i18n.Tr(global.Lang, "outputParams.DATAPARAMSILLEGAL") + " " + err.Key + ":" + err.Message
				return result, errors.New(i18n.Tr(global.Lang, "outputParams.DATAPARAMSILLEGAL"))
			}
		}
	}

	//sign check
	err = sign.ValidSign(rawDataBody, beego.AppConfig.String("sign.secretKey"))
	if err != nil {
		result.MetaCheckResult = nil
		result.RequestID = metaCheckResult.MetaCheckResult["request-id"]
		result.Message = err.Error()
		// return result, err
	}

	return metaCheckResult, nil
}

/**
 * meta参数验证
 * 1.map转json
 * 2.json转slice
 * 3.解析到struct
 * 4.将header 放入map 返回
 *
 * @meta 	meta  map[string][]string 	header信息 map格式
 */
func MetaHeaderCheck(meta map[string][]string) (result Result, err error) {
	rawMetaHeader, _ := ffjson.Marshal(meta)
	var metaHeader MetaHeader
	ffjson.Unmarshal(rawMetaHeader, &metaHeader)

	//日志
	fmt.Println("meta json解析:", metaHeader)
	for key, val := range meta {
		beego.Trace("meta analysis : " + key + ":" + val[0])
	}

	valid := validation.Validation{}

	is, err := valid.Valid(&metaHeader)

	//日志

	//Content-Type
	if val, ok := meta["Content-Type"]; ok {
		if val[0] != beego.AppConfig.String("Content-Type") {
			result.MetaCheckResult = nil
			result.Message = i18n.Tr(global.Lang, "outputParams.CONTENTTYPEILLEGAL")
			if val, ok := meta["request-id"]; ok {
				result.RequestID = val[0]
			}
			return result, errors.New(i18n.Tr(global.Lang, "outputParams.CONTENTTYPEILLEGAL "))
		}
	}

	//Accept
	if val, ok := meta["Accept"]; ok {
		if val[0] != beego.AppConfig.String("Accept") {
			result.MetaCheckResult = nil
			result.Message = i18n.Tr(global.Lang, "outputParams.ACCEPTILLEGAL")
			if val, ok := meta["request-id"]; ok {
				result.RequestID = val[0]
			}
			return result, errors.New(i18n.Tr(global.Lang, "outputParams.ACCEPTILLEGAL "))
		}
	}

	//检查参数
	if err != nil {
		// handle error
		beego.Trace(i18n.Tr(global.Lang, "outputParams.SYSTEMILLEGAL") + err.Error())
	}

	if !is {
		for _, err := range valid.Errors {
			beego.Trace(i18n.Tr(global.Lang, "outputParams.METAPARAMSILLEGAL") + err.Key + ":" + err.Message)
			result.MetaCheckResult = nil
			result.Message = i18n.Tr(global.Lang, "outputParams.METAPARAMSILLEGAL") + " " + err.Key + ":" + err.Message
			if val, ok := meta["request-id"]; ok {
				result.RequestID = val[0]
			}
			return result, errors.New(i18n.Tr(global.Lang, "outputParams.METAPARAMSILLEGAL "))
		}
	}

	//把meta参数放入新的struct 返回
	var metaMap = make(map[string]string)
	for key, val := range meta {
		metaMap[key] = val[0]
	}

	//日志
	if len(metaMap["request-id"]) == 0 {
		metaMap["request-id"] = getRequestID()
	}

	result.MetaCheckResult = metaMap
	result.Message = ""
	result.RequestID = ""

	return result, nil
}

//request id增加
func getRequestID() string {
	return "RRRRRRRRRRRRRRRRRRRR"
}
