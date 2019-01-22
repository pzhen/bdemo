package controllers

import (
	"bdemo/models"
	"bdemo/utils"
	"encoding/json"
)

type SysUserController struct {
	SysBaseController
}

func (c *SysUserController) Prepare() {
	c.SysBaseController.Prepare()
}

func (c *SysUserController) LoginForm() {
	c.TplName = "sysuser/loginForm.html"
}

func (c *SysUserController) LogOut() {
	c.DelSession("UserMenu")
	c.DelSession("UserSession")
	c.DisplayJson(1,"logout success",c.URLFor("SysUserController.LoginForm"))
}

func (c *SysUserController) LoginAction() {

	type loginForm struct {
		UserName string `form:"username"`
		Password string `form:"password"`
	}
	u := loginForm{}
	c.ParseForm(&u)

	if len(u.UserName) == 0 || len(u.Password) == 0 {
		c.DisplayStatus(0, "账号密码不能为空!", "")
	}

	u.Password = utils.String2md5(u.Password)
	userInfo, err := models.GetSysUserByUserName(u.UserName)
	if userInfo != nil && err == nil {
		if userInfo.UserStatus != 1 {
			c.DisplayStatus(0, "账号禁用,请联系管理员", "")
		}

		if u.Password != userInfo.Password {
			c.DisplayStatus(0, "密码错误", "")
		}

		// 用户信息
		userSession, _ := json.Marshal(userInfo)
		c.SetSession("UserSession", string(userSession))

		// 菜单权限
		v := models.GetUserMenuByRoleIdArr(*userInfo)
		m, _ := json.Marshal(v)
		c.SetSession("UserMenu", string(m))

		c.DisplayStatus(1, "登录成功,等待跳转", "/sys_home/index")
	} else {
		c.DisplayStatus(0, "账号不存在", "")
	}
}

