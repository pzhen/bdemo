package controllers

import (
	"bdemo/models"
)

type SysLogController struct {
	SysBaseController
}

func (c *SysLogController) Prepare() {
	c.SysBaseController.Prepare()
}

func (c *SysLogController) GetSysLogListByPage() {
	order, by := "log_id", "desc"
	where := make(map[string]string)
	where["user_name"] 	= c.Input().Get("user_name")
	where["start_time"] = c.Input().Get("start_time")
	where["end_time"] 	= c.Input().Get("end_time")

	dataList, count := models.GetSysLogListByPage(where, PageNum, RowsNum, order, by)

	c.Data["where"] 	= where
	c.Data["DataList"] 	= dataList
	c.Data["DataCount"] = count
	c.TplName = "syslog/listSysLog.html"
}

func (c *SysLogController) DeleteSysLog() {
	ids := c.Input().Get("log_ids")
	_, err := models.DeleteSysLog(ids)
	if err != nil {
		c.DisplayJson(0, "修改失败", err.Error())
	}
	c.DisplayJson(1, "删除成功", c.URLFor("SysLogController.GetSysLogListByPage"))
}
