package inout

import (
	"time"

	"github.com/JREAMLU/core/global"

	"github.com/beego/i18n"
)

//logic service
const (
	SUCCESS           = 0
	DATAPARAMSILLEGAL = 10000
	METAPARAMSILLEGAL = 15000
	LOGICILLEGAL      = 20000
	SYSTEMILLEGAL     = 30000
)

//httpStatus
const (
	OK                            = 200 //服务器成功处理了请求
	CREATED                       = 201 //请求执行成功，资源已创建完毕
	ACCEPTED                      = 202 //请求已接受，但服务器可能尚未处理
	NON_AUTHORITATIVE_INFORMATION = 203 //服务器返回的信息并非来自原始资源，而是来自第三方或者原始资源的子集
	NO_CONTENT                    = 204 //请求执行成功，但是想要没有内容尸体
	PARTIAL_CONTENT               = 206 //服务器已经成功处理了部分GET请求
	MOVED_PERMANENTLY             = 301 //请求的URL已移走，Respone中应该包含一个新的URI，锁门资源现在所处的位置，客户端之后的请求都应该访问新的URI
	FOUND                         = 302 //与状态码301不同的是，这里的资源移除是临时的，客户端以后的资源请求仍然使用原始的URI
	SEE_OTHER                     = 303 //用来告知客户端应该使用另一个URL来获取资源，Respone中应该包含一个另一个URI
	NOT_MODIFIED                  = 304 //客户端发送GET请求，告诉客户端资源未被修改
	TEMPORARY_REDIRECT            = 307 //类似302，临时URI
	PERMANENT_REDIRECT            = 308 //永久URI
	BAD_REQUEST                   = 400 //客户端发送另一个错误的请求
	UNAUTHORIZED                  = 401 //请求需要验证，因此客户端需要以合适的授权重新发送请求
	FORBIDDEN                     = 403 //权限错误
	NOT_FOUND                     = 404 //客户端请求有误，请求了一个不存在URI
	METHOD_NOT_ALLOWED            = 405 //有这个方法，但不被允许
	NOT_ACCEPTABLE                = 406 //header accept 必须有
	REQUESET_TIMEOUT              = 408 //请求超时
	CONFLICT                      = 409 //请求与先前请求不兼容
	GONE                          = 410 //标记服务器曾经有这个资源，现在这个资源已经不存在了
	LENGTH_REQUIRED               = 411 //header Content-Length 必须有
	PRECONDITION_FAILED           = 412 //header if-* 必须有
	PAYLOAD_TOO_LARGE             = 413 //负载太大
	URI_TOO_LONG                  = 414 //URI太长
	UNSUPPORTED_MEDIA_TYPE        = 415 //header Content-Type 不支持的类型
	RANGE_NOT_SATISFIABLE         = 416 //您的 Web 服务器认为，客户端（如您的浏览器或我们的 CheckUpDown 机器人）发送的 HTTP 数据流包含一个“范围”请求，规定了一个无法满足的字节范围 - 因为被访问的资源不覆盖这个字节范围。 例如， 如果一个图像文件资源有 1000 个字节，而被请求的范围是 500-1500 ，那就无法满足。
	EXPECTATION_FAILED            = 417 //header Expect 不一致(other)
	I_M_A_TEAPOT                  = 418 //我是茶壶(愚人节 无实际用途)
	ENHANCE_YOUR_CALM             = 420 //请求量大量大也能负载
	UNPROCESSABLE_ENTITY          = 422 //无法处理的请求实体
	UPGRADE_REQUIRED              = 426 //升级要求
	TOO_MANY_REQUESTS             = 429 //请求量大无法负载
	INTERNAL_SERVER_ERROR         = 500 //服务器请求时到内部错误
	NOT_IMPLEMENTED               = 501 //服务器不支持客户端的请求方法
	BAD_GATEWAY                   = 502 //如果服务器被设置为网关或者代理设备，但是受到了上游服务器的无效响应就会提示该错误码
	SERVICE_UNAVAILABLE           = 503 //由于临时的服务器维护或者过载，服务器当前无法处理请求。这个状况是临时的，并且将在一段时间以后恢复
	GATEWAY_TIMEOUT               = 504 //服务器（不一定是 Web 服务器）正在作为一个网关或代理来完成客户（如您的浏览器或我们的 CheckUpDown 机器人）访问所需网址的请求。 为了完成您的 HTTP 请求， 该服务器访问一个上游服务器， 但没得到及时的响应。
	HTTP_VERSION_NOT_SUPPORTED    = 505 //服务器手到的请求使用了它不支持的HTTP协议
)

type Output struct {
	Meta       MetaList    `json:"meta"`
	StatusCode int         `json:"status_code"`
	Message    interface{} `json:"message"`
	Data       interface{} `json:"data"`
}

type MetaList struct {
	RequestId string `json:"Request-Id"`
	UpdatedAt string `json:"updated_at"`
}

/*
type dataList struct {
	Total int                    `json:"total"`
	List  map[string]interface{} `json:"list"`
}
*/

/**
 *	@auther		jream.lu
 *	@intro		出参成功
 *	@logic
 *	@todo		返回值
 *	@params		params ...interface{}	切片指针
 *	@return 	?
 */
func OutputSuccess(data interface{}, requestID string) Output {
	var op Output
	op.Meta.RequestId = requestID
	op.Meta.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	op.StatusCode = SUCCESS

	op.Message = i18n.Tr(global.Lang, "outputParams.SUCCESS")

	op.Data = data

	return op
}

func OutputFail(msg interface{}, status string, requestID string) Output {
	var op Output
	op.Meta.RequestId = requestID
	op.Meta.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	switch status {
	case "SUCCESS":
		op.StatusCode = SUCCESS
	case "DATAPARAMSILLEGAL":
		op.StatusCode = DATAPARAMSILLEGAL
	case "METAPARAMSILLEGAL":
		op.StatusCode = METAPARAMSILLEGAL
	case "LOGICILLEGAL":
		op.StatusCode = LOGICILLEGAL
	case "SYSTEMILLEGAL":
		op.StatusCode = SYSTEMILLEGAL
	}

	op.Message = msg

	op.Data = make(map[string]interface{})

	return op
}
