package models

import (
	"bdemo/utils"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SysRoleFormData struct {
	RoleId     int
	Intro      string
	MenuMap    string
	RoleName   string
	RoleStatus int
}

type SysRole struct {
	Id         int    `orm:"column(role_id);auto" description:"角色ID"`
	RoleName   string `orm:"column(role_name);size(60)" description:"角色名称"`
	Intro      string `orm:"column(intro);null" description:"角色介绍"`
	RoleStatus int    `orm:"column(role_status)" description:"状态"`
	CreateTime uint   `orm:"column(create_time)"`
	UpdateTime uint   `orm:"column(update_time)"`
}

type SysRoleMenuMap struct {
	Id       int `orm:"column(role_id);auto" description:"角色ID"`
	MenuId   int `orm:"column(menu_id)" description:"菜单id"`
	ActionId int `orm:"column(action_id);size(255);null" description:"操作权限"`
}

func init() {
	orm.RegisterModel(new(SysRole), new(SysRoleMenuMap))
}

//获取一条角色
func GetSysRoleById(id int) (v *SysRole, err error) {
	o := orm.NewOrm()
	v = &SysRole{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//获取角色对应菜单以及菜单下方法
func GetSysRoleMenuActionMap(roleIds string) (v []SysRoleMenuMap) {
	data := make([]SysRoleMenuMap, 0)
	roleIdArr := utils.StringsSplitToSliceInt(roleIds, ",")
	if len(roleIdArr) == 0 {
		return data
	}

	o := orm.NewOrm()
	qs := o.QueryTable(Table_Sys_Role_Menu_Map)
	qs.Filter("role_id__in", roleIdArr)
	qs.All(&data)
	return data
}

func AddSysRole(m *SysRoleFormData) (id int64, err error) {
	var roleInfo SysRole
	roleInfo.Intro = m.Intro
	roleInfo.RoleName = m.RoleName
	roleInfo.RoleStatus = m.RoleStatus

	o := orm.NewOrm()
	id, err = o.Insert(&roleInfo)

	//关系入库
	MenuMapArr := strings.Split(m.MenuMap, ",")
	for _, v := range MenuMapArr {
		if v == "" {
			continue
		}

		insert := SysRoleMenuMap{}
		insert.Id = int(id)
		if strings.Contains(v, "-") == true {
			mapArr := strings.Split(v, "-")
			insert.MenuId, _ = strconv.Atoi(mapArr[0])
			insert.ActionId, _ = strconv.Atoi(mapArr[1])
		} else {
			insert.MenuId, _ = strconv.Atoi(v)
			insert.ActionId = 0
		}

		o.Insert(&insert)
	}
	return
}

func SaveSysRole(m *SysRoleFormData) (id int64, err error) {
	var roleInfo SysRole
	roleInfo.Id = m.RoleId
	roleInfo.Intro = m.Intro
	roleInfo.RoleName = m.RoleName
	roleInfo.RoleStatus = m.RoleStatus

	o := orm.NewOrm()
	id, err = o.Update(&roleInfo)
	o.Delete(&SysRoleMenuMap{Id: m.RoleId})

	//关系入库
	MenuMapArr := strings.Split(m.MenuMap, ",")
	for _, v := range MenuMapArr {
		if v == "" {
			continue
		}
		insert := SysRoleMenuMap{}
		insert.Id = m.RoleId
		if strings.Contains(v, "-") == true {
			mapArr := strings.Split(v, "-")
			insert.MenuId, _ = strconv.Atoi(mapArr[0])
			insert.ActionId, _ = strconv.Atoi(mapArr[1])
		} else {
			insert.MenuId, _ = strconv.Atoi(v)
			insert.ActionId = 0
		}

		o.Insert(&insert)
	}
	return
}

func GetSysRoleListByPage(where map[string]string, pageNum int, rowsNum int, orderBy string) ([]*SysRole, int, error) {
	dataSql := "SELECT * FROM sys_role WHERE 1=1 "
	countSql := "SELECT count(*) AS count FROM sys_role WHERE 1=1 "

	start := (pageNum - 1) * rowsNum

	if v, ok := where["RoleName"]; ok && v != "" {
		dataSql += "AND role_name LIKE %" + v + "%"
		countSql += "AND role_name LIKE %" + v + "%"
	}

	dataSql += " ORDER BY " + orderBy + " LIMIT " + strconv.Itoa(start) + "," + strconv.Itoa(rowsNum)

	data := make([]*SysRole, 0)
	o := orm.NewOrm()
	_, err := o.Raw(dataSql).QueryRows(&data)

	d := struct {
		Count int `orm:"column(count)"`
	}{0}

	err = o.Raw(countSql).QueryRow(&d)
	return data, d.Count, err
}

func DeleteSysRole(ids string) (num int64, err error) {
	s, i := utils.GetWhereInSqlByStrId(ids)
	if len(i) == 0 {
		return 0, errors.New("参数错误")
	}
	o := orm.NewOrm()
	res, err := o.Raw("DELETE FROM sys_role WHERE role_id IN ("+s+")", i).Exec()
	o.Raw("DELETE FROM sys_role_menu_map WHERE role_id IN ("+s+")", i).Exec()
	num, _ = res.RowsAffected()
	return num, err
}

// 修改状态
func ModifySysRoleStatus(ids string, roleStatus int) (num int64, err error) {
	s, i := utils.GetWhereInSqlByStrId(ids)
	if len(i) == 0 {
		err = errors.New("参数错误")
		return 0, err
	}
	o := orm.NewOrm()
	res, err := o.Raw("UPDATE sys_role SET role_status = ? WHERE role_id IN ("+s+")", roleStatus, i).Exec()
	num, _ = res.RowsAffected()
	return num, err
}
