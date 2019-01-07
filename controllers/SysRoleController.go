package controllers

import (
	"bdemo/models"
	"strconv"
)

type SysRoleController struct {
	SysBaseController
}

func (c *SysRoleController) Prepare() {
	c.SysBaseController.Prepare()
}

func (c *SysRoleController) GetSysRoleListByPage() {
	var where map[string]string
	//where["RoleName"] = c.Input().Get("role_name")
	pageNum, _ := strconv.Atoi(c.Input().Get("page_num"))
	if pageNum <= 0 {
		pageNum = 1
	}
	roleList, count, _ := models.GetSysRoleListByPage(where, pageNum, 10, "role_id desc")
	c.Data["RoleList"] = roleList
	c.Data["RoleCount"] = count
	c.Data["PageNum"] = pageNum
	c.TplName = "sysrole/listSysRole.html"
}

func (c *SysRoleController) FormAddSysRole() {
	MenuList, _ := models.GetSysMenuList()
	c.Data["MenuList"] = MenuList
	c.TplName = "sysrole/formAddSysRole.html"
}

func (c *SysRoleController) AddSysRole() {
	r := &models.SysRoleFormData{}
	if err := c.ParseForm(r); err != nil {
		c.DisplayJson(0, "数据解析失败", err)
	}

	if _, err := models.AddSysRole(r); err != nil {
		c.DisplayJson(0, "保存失败", err)
	}

	c.DisplayJson(1, "保存成功", c.URLFor("SysRoleController.ListSysRole"))
}

func (c *SysRoleController) FormModifySysRole() {
	RoleId := c.Input().Get("role_id")
	Id, _ := strconv.Atoi(RoleId)
	RoleRow, _ := models.GetSysRoleById(Id)

	//获取权限
	PowerList := models.GetSysRoleMenuActionMap(RoleId)
	//fmt.Println(PowerList)

	// 所有菜单
	MenuList, _ := models.GetSysMenuList()

	c.Data["RoleRow"] = RoleRow
	c.Data["MenuList"] = MenuList
	c.Data["PowerList"] = PowerList
	c.TplName = "sysrole/formModifySysRole.html"
}

func (c *SysRoleController) SaveSysRole() {
	r := &models.SysRoleFormData{}
	if err := c.ParseForm(r); err != nil {
		c.DisplayJson(0, "数据解析失败", err)
	}

	if _, err := models.SaveSysRole(r); err != nil {
		c.DisplayJson(0, "保存失败", err)
	}

	c.DisplayJson(1, "保存成功", c.URLFor("SysRoleController.ListSysRole"))
}

func (c *SysRoleController) ModifySysRoleStatus() {
	ids := c.Input().Get("role_ids")
	roleStatus, _ := strconv.Atoi(c.Input().Get("role_status"))
	_, err := models.ModifySysRoleStatus(ids, roleStatus)
	if err != nil {
		c.DisplayJson(0, "修改失败", err.Error())
	}
	c.DisplayJson(1, "修改成功", c.URLFor("SysRoleController.ListSysRole"))
}

func (c *SysRoleController) DeleteSysRole() {
	ids := c.Input().Get("role_ids")
	_, err := models.DeleteSysRole(ids)
	if err != nil {
		c.DisplayJson(0, "修改失败", err.Error())
	}
	c.DisplayJson(1, "删除成功", c.URLFor("SysRoleController.ListSysRole"))
}
