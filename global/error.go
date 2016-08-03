package global

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type ErrorController struct {
	beego.Controller
	i18n.Locale
}

func (r *ErrorController) Error401() {
	r.Data["json"] = 401
	r.ServeJSON()
}

func (r *ErrorController) Error402() {
	r.Data["json"] = 402
	r.ServeJSON()
}

func (r *ErrorController) Error403() {
	r.Data["json"] = 403
	r.ServeJSON()
}

func (r *ErrorController) Error404() {
	r.Data["json"] = 404
	r.ServeJSON()
}

func (r *ErrorController) Error405() {
	r.Data["json"] = 405
	r.ServeJSON()
}

func (r *ErrorController) Error500() {
	r.Data["json"] = 500
	r.ServeJSON()
}

func (r *ErrorController) Error501() {
	r.Data["json"] = 501
	r.ServeJSON()
}

func (r *ErrorController) Error502() {
	r.Data["json"] = 502
	r.ServeJSON()
}

func (r *ErrorController) Error503() {
	r.Data["json"] = 503
	r.ServeJSON()
}

func (r *ErrorController) Error504() {
	r.Data["json"] = 504
	r.ServeJSON()
}
