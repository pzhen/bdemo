package controllers

import (
	"encoding/json"
	"time"

	"bdemo/models"
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
	c.TplName = "syshome/serverinfo.html"
}
