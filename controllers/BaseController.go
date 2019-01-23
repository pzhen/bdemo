package controllers

import (
	"fmt"
	"strings"
	"time"
	"bdemo/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type SysBaseController struct {
	beego.Controller

	ControllerName string
	ActionName     string
	CurrentUrl     string
	CurrUserInfo   *models.SysUser
}

//开放地址
var openAction = []string{
	"SysUserController.LoginForm",
	"SysUserController.LoginAction",
}

//登录后不需要权限地址
var loginOpenAction = []string{
	"SysHomeController.Index",
	"SysHomeController.ServerInfo",
}

//不需要收集操作日志的地址
var logOpenAction = []string{
	"SysUserController.LoginForm",
	"SysUserController.LoginAction",
}

//返回统一的格式
type SerResJson struct {
	Code    int
	Message string
	Data    interface{}
}

//权限控制,操作日志收集
func (b *SysBaseController) Prepare() {
	b.ControllerName, b.ActionName = b.GetControllerAndAction()
	b.CurrentUrl = b.ControllerName + "." + b.ActionName
	//权限判断
	if !b.CheckAuth() {
		b.DisplayStatus(0, "对不起您没有权限", "")
	}
	//用户信息
	b.CurrUserInfo = models.GetUserInfoBySession(b.GetSession("UserSession"))
	b.Data["CurrUserInfo"] = b.CurrUserInfo

	//记录日志
	b.logCollection()
}

//收集操作日志
func (b *SysBaseController) logCollection() {
	flag := false
	for _, v := range logOpenAction {
		if v == b.CurrentUrl {
			flag = true
		}
	}
	if flag == false {
		b.Controller.Ctx.Request.ParseForm()

		formJson, _ := json.Marshal(b.Controller.Ctx.Request.Form)
		log         := models.SysLog{}

		log.Url        = fmt.Sprintf("%s", b.Controller.Ctx.Request.URL)
		log.UrlFor     = b.CurrentUrl
		log.UserId     = b.CurrUserInfo.Id
		log.UserName   = b.CurrUserInfo.UserName
		log.FormData   = string(formJson)
		log.CreateTime = uint(time.Now().Unix())

		models.AddSysLog(&log)
	}
}

//权限验证
func (b *SysBaseController) CheckAuth() bool {
	flag := false
	for _, v := range openAction {
		if v == b.CurrentUrl {
			flag = true
			return flag
		}
	}
	userInfo := models.GetUserInfoBySession(b.GetSession("UserSession"))
	//登录
	if userInfo.Id > 0 {
		for _, v := range loginOpenAction {
			if v == b.CurrentUrl {
				flag = true
				return flag
			}
		}
		if userInfo.UserType == 1 {
			flag = true
			return flag
		}
		userPowerList := models.GetUserMenuByRoleIdArr(*userInfo)
		for _, value := range userPowerList {
			if value.DefaultUrl != "" {
				sArr := strings.Split(value.DefaultUrl, ".")
				if sArr[0] == b.ControllerName {
					for _, v := range value.Operates {
						if v == b.ActionName {
							flag = true
							break
						}
					}
				}
			}
		}
	}
	return flag
}

//公共返回方法
func (b *SysBaseController) DisplayJson(code int, message string, data interface{}) {
	b.Data["json"] = &SerResJson{
		Code:    code,
		Message: message,
		Data:    data,
	}
	b.ServeJSON()
	b.StopRun()
}

//公共返回方法
func (b *SysBaseController) DisplayStatus(code int, message string, data interface{}) {
	if b.Ctx.Input.IsAjax() {
		b.DisplayJson(code, message, data)
	}
	b.Abort("403")
}
