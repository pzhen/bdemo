package controllers

import (
	"strconv"

	"bdemo/models"
)

type SysMenuController struct {
	SysBaseController
}

func (c *SysMenuController) Prepare() {
	c.SysBaseController.Prepare()
}

func (c *SysMenuController) ListSysMenu() {
	menuList := models.GetSysMenuList()
	c.Data["MenuList"] = menuList
	c.TplName = "sysmenu/listSysMenu.html"
}


func (c *SysMenuController) FormSysMenu() {
	MenuId := c.Input().Get("menu_id")
	Id, _ := strconv.Atoi(MenuId)
	MenuRow := models.GetSysMenuById(Id)
	MenuList := models.GetSysMenuList()
	c.Data["MenuRow"] = MenuRow
	c.Data["MenuList"] = MenuList
	c.TplName = "sysmenu/formSysMenu.html"
}

func (c *SysMenuController) SaveSysMenu() {
	m := &models.SysMenu{}
	if err := c.ParseForm(m); err != nil {
		c.DisplayJson(0, "数据解析失败", err)
	}
	if _,err := models.SaveSysMenu(m); err != nil {
		c.DisplayJson(0, "保存失败", err)
	}
	c.DisplayJson(1, "保存成功", c.URLFor("SysMenuController.ListSysMenu"))
}

func (c *SysMenuController) ModifySysMenuStatus() {
	ids := c.Input().Get("id")
	menuStatus, _ := strconv.Atoi(c.Input().Get("status"))
	_, err := models.ModifySysMenuStatus(ids, menuStatus)
	if err != nil {
		c.DisplayJson(0, "修改失败", err.Error())
	}
	c.DisplayJson(1, "修改成功", c.URLFor("SysMenuController.ListSysMenu"))
}

func (c *SysMenuController) DeleteSysMenu() {
	ids := c.Input().Get("id")
	_, err := models.DeleteSysMenu(ids)
	if err != nil {
		c.DisplayJson(0, "修改失败", err.Error())
	}
	c.DisplayJson(1, "删除成功", c.URLFor("SysMenuController.ListSysMenu"))
}
