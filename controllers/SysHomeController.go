package controllers

import (
	"encoding/json"
	"time"
	"bdemo/models"
	"github.com/astaxie/beego"
	"os"
	"runtime"
)

type SysHomeController struct {
	SysBaseController
}

func (c *SysHomeController) Prepare() {
	c.SysBaseController.Prepare()
}

func (c *SysHomeController) Index() {
	um := make(map[int]*models.UserMenuIterm)
	json.Unmarshal([]byte(c.GetSession("UserMenu").(string)), &um)
	c.Data["UserMenu"] = um
	c.TplName = "syshome/index.html"
}

func (c *SysHomeController) ServerInfo() {
	CurrTime := time.Now().Format("2006-01-02 03:04:05 PM")
	c.Data["CurrTime"] = CurrTime

	// 系统配置信息
	c.Data["OS"] 			= beego.AppConfig.String("os")
	c.Data["Author"] 		= beego.AppConfig.String("author")
	c.Data["GOPATH"] 		= os.Getenv("GOPATH")
	c.Data["AppName"] 		= beego.AppConfig.String("appname")
	c.Data["Version"] 		= beego.AppConfig.String("version")
	c.Data["GOVersion"] 	= runtime.Version()
	c.Data["UploadLimit"] 	= beego.AppConfig.String("uploadlimit")
	c.Data["MySqlVersion"] 	= beego.AppConfig.String("mysqlversion")

	c.TplName = "syshome/serverinfo.html"
}
