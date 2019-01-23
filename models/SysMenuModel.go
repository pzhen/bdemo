package models

import (
	"fmt"
	"errors"
	"strconv"

	"bdemo/utils"
	"encoding/json"
	"github.com/astaxie/beego/orm"
)

type SysMenuAction struct {
	FuncId   int    `json:"func_id"`
	FuncName string `json:"func_name"`
	FuncDesc string `json:"func_desc"`
}

type SysMenu struct {
	Id         int             `orm:"column(menu_id);auto"`
	MenuRootid int             `orm:"column(menu_rootid);null" description:"上级id"`
	MenuName   string          `orm:"column(menu_name);size(60)" description:"菜单名称"`
	MenuUrl    string          `orm:"column(menu_url);size(60)" description:"所属类"`
	MenuFuncs  string          `orm:"column(menu_funcs);size(1024)" description:"所属方法"`
	MenuIcon   string          `orm:"column(menu_icon);size(50)" description:"图标"`
	MenuLock   int8            `orm:"column(menu_lock)" description:"锁定"`
	MenuStatus int8            `orm:"column(menu_status)" description:"状态"`
	MenuLevel  int8            `orm:"column(menu_level)" description:"层级"`
	MenuPath   string          `orm:"column(menu_path)" description:"路径"`

	Operates   []string        `orm:"-"`
	FuncsInfo  []SysMenuAction `orm:"-"`
}

type UserMenuIterm struct {
	MenuId     int
	MenuName   string
	MenuIcon   string
	MenuLevel  int8
	DefaultUrl string
	MenuRootid int
	Operates   []string
}

func init() {
	orm.RegisterModel(new(SysMenu))
}

func GetSysMenuList() []*SysMenu {
	data := make([]*SysMenu, 0)
	orm.NewOrm().QueryTable(Table_Sys_Menu).OrderBy("menu_path").All(&data)

	if len(data) > 0 {
		for key, value := range data {
			var SysMenuAcitonS = make([]SysMenuAction, 0)
			json.Unmarshal([]byte(value.MenuFuncs), &SysMenuAcitonS)
			data[key].FuncsInfo = SysMenuAcitonS
		}
	}
	return data
}

func AddSysMenu(m *SysMenu) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	v := SysMenu{Id: int(id)}
	if m.MenuRootid > 0 {
		v.MenuPath = "-" + strconv.Itoa(m.MenuRootid) + "-" + strconv.Itoa(int(id)) + "-"
	} else {
		v.MenuPath = "-" + strconv.Itoa(int(id)) + "-"
	}
	o.Update(&v, "MenuPath")
	return
}

func GetSysMenuById(id int) *SysMenu {
	data := new(SysMenu)
	if id > 0 {
		orm.NewOrm().QueryTable(Table_Sys_Menu).Filter("menu_id", id).One(data)
	}
	return data
}

func SaveSysMenu(m *SysMenu) (err error) {
	o := orm.NewOrm()
	v := SysMenu{Id: m.Id}

	if m.MenuRootid > 0 {
		m.MenuPath = "-" + strconv.Itoa(m.MenuRootid) + "-" + strconv.Itoa(m.Id) + "-"
	} else {
		m.MenuPath = "-" + strconv.Itoa(m.Id) + "-"
	}

	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

//删除菜单
func DeleteSysMenu(ids string) (num int64, err error) {
	s, i := utils.GetWhereInSqlByStrId(ids)
	if len(i) == 0 {
		return 0, errors.New("参数错误")
	}
	o := orm.NewOrm()
	res, err := o.Raw("DELETE FROM sys_menu WHERE menu_id in ("+s+")", i).Exec()
	num, _ = res.RowsAffected()
	return num, err
}

//修改菜单状态
func ModifySysMenuStatus(ids string, menuStatus int) (num int64, err error) {
	s, i := utils.GetWhereInSqlByStrId(ids)
	if len(i) == 0 {
		err = errors.New("参数错误")
		return 0, err
	}
	o := orm.NewOrm()
	res, err := o.Raw("UPDATE sys_menu SET menu_status = ? WHERE menu_id in ("+s+")", menuStatus, i).Exec()
	num, _ = res.RowsAffected()
	return num, err
}

//获取用户权限列表
//根据用户信息获取用户菜单以及菜单下的权限
func GetUserMenuByRoleIdArr(userInfo SysUser) map[int]*UserMenuIterm {
	roleIds := userInfo.RoleId
	userMenu := make(map[int]*UserMenuIterm)
	menuList := GetSysMenuList()
	mapList := GetSysRoleMenuActionMap(roleIds)
	for _, value := range menuList {
		if value.MenuStatus == 0 {
			continue
		}
		Flag := false
		if userInfo.UserType == 1 {
			Flag = true
		}
		tmpOper := make([]string, 0)
		for _, v := range mapList {
			if v.MenuId == value.Id {
				Flag = true
				for _, oper := range value.FuncsInfo {
					if oper.FuncId == v.ActionId {
						tmpOper = append(tmpOper, oper.FuncName)
					}
				}
			}
		}
		if Flag == true {
			userMenu[value.Id] = &UserMenuIterm{}
			userMenu[value.Id].MenuId = value.Id
			userMenu[value.Id].MenuName = value.MenuName
			userMenu[value.Id].MenuIcon = value.MenuIcon
			userMenu[value.Id].DefaultUrl = value.MenuUrl
			userMenu[value.Id].MenuRootid = value.MenuRootid
			userMenu[value.Id].MenuLevel = value.MenuLevel
			userMenu[value.Id].Operates = tmpOper
		}
	}
	return userMenu
}
