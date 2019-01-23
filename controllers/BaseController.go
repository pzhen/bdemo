package controllers

import (
	"fmt"
	"os"
	"runtime"
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

//登录后开放地址
var loginOpenAction = []string{
	"SysHomeController.Index",
	"SysHomeController.ServerInfo",
}

type SerResJson struct {
	Code    int
	Message string
	Data    interface{}
}

//权限控制,操作日志收集
func (b *SysBaseController) Prepare() {
	b.ControllerName, b.ActionName = b.GetControllerAndAction()
	b.CurrentUrl = b.ControllerName + "." + b.ActionName

	// 权限判断
	if !b.CheckAuth() {
		b.DisplayStatus(0, "对不起您没有权限", "")
	}

	userInfo := models.GetUserInfoBySession(b.GetSession("UserSession"))
	b.CurrUserInfo = userInfo
	b.Data["SysUserName"] = b.CurrUserInfo.UserName

	// 系统配置信息
	b.Data["OS"] = beego.AppConfig.String("os")
	b.Data["Author"] = beego.AppConfig.String("author")
	b.Data["GOPATH"] = os.Getenv("GOPATH")
	b.Data["AppName"] = beego.AppConfig.String("appname")
	b.Data["Version"] = beego.AppConfig.String("version")
	b.Data["GOVersion"] = runtime.Version()
	b.Data["MySqlVersion"] = beego.AppConfig.String("mysqlversion")
	b.Data["UploadLimit"] = beego.AppConfig.String("uploadlimit")

	// 记录日志
	flag := false
	for _, v := range openAction {
		if v == b.CurrentUrl {
			flag = true
		}
	}
	if flag == false {
		b.Controller.Ctx.Request.ParseForm()
		formJson, _ := json.Marshal(b.Controller.Ctx.Request.Form)

		log := models.SysLog{}
		log.Url = fmt.Sprintf("%s", b.Controller.Ctx.Request.URL)
		log.UrlFor = b.CurrentUrl
		log.UserId = b.CurrUserInfo.Id
		log.UserName = b.CurrUserInfo.UserName
		log.FormData = string(formJson)
		log.CreateTime = uint(time.Now().Unix())

		models.AddSysLog(&log)
	}
}

func (b *SysBaseController) DisplayStatus(code int, message string, data interface{}) {
	if b.Ctx.Input.IsAjax() {
		b.DisplayJson(code, message, data)
	}
	b.Abort("403")
}

func (b *SysBaseController) DisplayJson(code int, message string, data interface{}) {
	b.Data["json"] = &SerResJson{
		Code:    code,
		Message: message,
		Data:    data,
	}
	b.ServeJSON()
	b.StopRun()
}

// 权限验证
func (b *SysBaseController) CheckAuth() bool {
	flag := false
	for _, v := range openAction {
		if v == b.CurrentUrl {
			flag = true
			return flag
		}
	}
	userInfo := models.GetUserInfoBySession(b.GetSession("UserSession"))
	// 登录
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
