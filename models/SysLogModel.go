package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"errors"
	"bdemo/utils"
)

type SysLog struct {
	Id         int    `orm:"column(log_id);auto"`
	Url        string `orm:"column(url);size(512)" description:"操作地址"`
	UrlFor     string `orm:"column(urlfor);size(255)" description:"urlFor"`
	UserId     int    `orm:"column(user_id); description:"用户ID"`
	UserName   string `orm:"column(user_name);size(64)" description:"用户名称"`
	FormData   string `orm:"column(form_data);" description:"操作数据"`
	CreateTime uint   `orm:"column(create_time); description:"操作时间"`
}

func init() {
	orm.RegisterModel(new(SysLog))
}

func AddSysLog(l *SysLog) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(l)
	return
}

func GetSysLogListByPage(where map[string]string, pageNum int, rowsNum int, orderBy string) ([]*SysLog, int, error) {
	dataSql := "SELECT * FROM sys_log WHERE 1=1 "
	countSql := "SELECT count(*) AS count FROM sys_log WHERE 1=1 "

	start := (pageNum - 1) * rowsNum

	if v, ok := where["UserName"]; ok && v != "" {
		dataSql += "AND user_name LIKE %" + v + "%"
		countSql += "AND user_name LIKE %" + v + "%"
	}

	dataSql += " ORDER BY " + orderBy + " LIMIT " + strconv.Itoa(start) + "," + strconv.Itoa(rowsNum)

	data := make([]*SysLog, 0)
	o := orm.NewOrm()
	_, err := o.Raw(dataSql).QueryRows(&data)

	d := struct {
		Count int `orm:"column(count)"`
	}{0}

	err = o.Raw(countSql).QueryRow(&d)
	return data, d.Count, err
}

func DeleteSysLog(ids string) (num int64, err error) {
	s, i := utils.GetWhereInSqlByStrId(ids)
	if len(i) == 0 {
		return 0, errors.New("参数错误!")
	}
	o := orm.NewOrm()
	res, err := o.Raw("DELETE FROM sys_log WHERE log_id IN ("+s+")", i).Exec()
	num, _ = res.RowsAffected()
	return num, err
}
