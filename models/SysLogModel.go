package models

import (
	"github.com/astaxie/beego/orm"
	"errors"
	"bdemo/utils"
	"math"
	"strconv"
	"strings"
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
	return orm.NewOrm().Insert(l)
}

func GetSysLogListByPage(where map[string]string, pageNum int, rowsNum int, order string, by string) ([]*SysLog, int) {
	data := make([]*SysLog, 0)
	start := (int(math.Abs(float64(pageNum))) - 1) * rowsNum

	var sql = "1=1"
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*")
	qb.From(Table_Sys_Log)

	if v, ok := where["user_name"]; ok && v != "" {
		keywords := utils.TrimString(v)
		sql += " AND user_name like \"%" + keywords + "%\""
	}

	if v, ok := where["start_time"]; ok && v != "" {
		startTime := utils.GetTimestamp(v)
		sql += " AND create_time >= " + strconv.Itoa(int(startTime))
	}

	if v, ok := where["end_time"]; ok && v != "" {
		endTime := utils.GetTimestamp(v)
		sql += " AND create_time <= " + strconv.Itoa(int(endTime))
	}

	qb.Where(sql)

	if strings.ToLower(by) == "desc" {
		qb.OrderBy(order).Desc()
	} else {
		qb.OrderBy(order).Asc()
	}

	qb.Limit(rowsNum).Offset(start)
	orm.NewOrm().Raw(qb.String()).QueryRows(&data)
	num := GetSysLogCount(where)
	return data, num
}

func GetSysLogCount(where map[string]string) int {
	c := struct {
		Count int `orm:"column(count)"`
	}{}
	var sql = "1=1"
	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("count(*) as count")
	qb.From(Table_Sys_Log)

	if v, ok := where["user_name"]; ok && v != "" {
		keywords := utils.TrimString(v)
		sql += " AND user_name like \"%" + keywords + "%\""
	}

	if v, ok := where["start_time"]; ok && v != "" {
		startTime := utils.GetTimestamp(v)
		sql += " AND create_time >= " + strconv.Itoa(int(startTime))
	}

	if v, ok := where["end_time"]; ok && v != "" {
		endTime := utils.GetTimestamp(v)
		sql += " AND create_time <= " + strconv.Itoa(int(endTime))
	}

	qb.Where(sql)
	orm.NewOrm().Raw(qb.String()).QueryRow(&c)
	return c.Count
}


func DeleteSysLog(ids string) (num int64, err error) {
	s, i := utils.GetWhereInSqlByStrId(ids)
	if len(i) == 0 {
		return 0, errors.New("参数错误")
	}
	o := orm.NewOrm()
	res, err := o.Raw("DELETE FROM sys_log WHERE log_id IN ("+s+")", i).Exec()
	num, _ = res.RowsAffected()
	return num, err
}
