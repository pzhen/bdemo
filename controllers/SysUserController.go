package controllers

import (
	"bdemo/models"
	"bdemo/utils"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
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

// URLMapping ...
func (c *SysUserController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create SysUser
// @Param	body		body 	models.SysUser	true		"body for SysUser content"
// @Success 201 {int} models.SysUser
// @Failure 403 body is empty
// @router / [post]
func (c *SysUserController) Post() {
	var v models.SysUser
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddSysUser(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get SysUser by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SysUser
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SysUserController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSysUserById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get SysUser
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SysUser
// @Failure 403
// @router / [get]
func (c *SysUserController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllSysUser(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the SysUser
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SysUser	true		"body for SysUser content"
// @Success 200 {object} models.SysUser
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SysUserController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.SysUser{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateSysUserById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the SysUser
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SysUserController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSysUser(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
