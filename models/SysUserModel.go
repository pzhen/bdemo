package models

import (
	"github.com/astaxie/beego/orm"
	"encoding/json"
	"errors"
)

type SysUser struct {
	Id         int    `orm:"column(user_id);auto" description:"用户id"`
	UserName   string `orm:"column(user_name);size(64)" description:"登录名"`
	NickName   string `orm:"column(nick_name);size(64)" description:"昵称"`
	RoleId     string `orm:"column(role_id);size(64)" description:"角色id"`
	Photo      string `orm:"column(photo);size(128)" description:"头像"`
	Password   string `orm:"column(password);size(32)" description:"密码"`
	Salt       string `orm:"column(salt);size(6)" description:"密码盐值"`
	Email      string `orm:"column(email);size(64)"`
	Mobile     string `orm:"column(mobile);size(32)"`
	CreateTime uint   `orm:"column(create_time)"`
	UpdateTime int    `orm:"column(update_time);null"`
	LastTime   uint   `orm:"column(last_time)"`
	LastIp     string `orm:"column(last_ip);size(15)"`
	LoginCount uint   `orm:"column(login_count)"`
	UserType   uint8  `orm:"column(user_type)"`
	UserStatus uint8  `orm:"column(user_status)"`
}

func init() {
	orm.RegisterModel(new(SysUser))
}

func AddSysUser(m *SysUser) (id int64, err error) {
	id, err = orm.NewOrm().Insert(m)
	return
}

func GetSysUserById(id int) (v *SysUser, err error) {
	o := orm.NewOrm()
	v = &SysUser{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetUserInfoBySession(s interface{}) *SysUser {
	u := new(SysUser)
	value, ok := s.(string);
	if !ok {
		return u
	}
	json.Unmarshal([]byte(value), &u)
	return u
}

func GetSysUserByUserName(userName string) (u *SysUser, err error) {
	u = new(SysUser)
	if userName == "" {
		return u, errors.New("userName param is empty")
	}
	u.UserName = userName
	if err = orm.NewOrm().Read(u, "UserName"); err == nil {
		return u, nil
	}
	return nil, err
}
