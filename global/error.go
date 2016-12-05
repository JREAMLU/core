package global

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

// ErrorController error controller
type ErrorController struct {
	beego.Controller
	i18n.Locale
}

// Error401 401
func (r *ErrorController) Error401() {
	r.Data["json"] = 401
	r.ServeJSON()
}

// Error402 402
func (r *ErrorController) Error402() {
	r.Data["json"] = 402
	r.ServeJSON()
}

// Error403 403
func (r *ErrorController) Error403() {
	r.Data["json"] = 403
	r.ServeJSON()
}

// Error404 404
func (r *ErrorController) Error404() {
	r.Data["json"] = 404
	r.ServeJSON()
}

// Error405 405
func (r *ErrorController) Error405() {
	r.Data["json"] = 405
	r.ServeJSON()
}

// Error500 500
func (r *ErrorController) Error500() {
	r.Data["json"] = 500
	r.ServeJSON()
}

// Error501 501
func (r *ErrorController) Error501() {
	r.Data["json"] = 501
	r.ServeJSON()
}

// Error502 502
func (r *ErrorController) Error502() {
	r.Data["json"] = 502
	r.ServeJSON()
}

// Error503 503
func (r *ErrorController) Error503() {
	r.Data["json"] = 503
	r.ServeJSON()
}

// Error504 504
func (r *ErrorController) Error504() {
	r.Data["json"] = 504
	r.ServeJSON()
}
