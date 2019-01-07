package controllers

import (

	"bdemo/models"
)

type OrderController struct {
	SysBaseController
}

func (c *OrderController) Prepare() {
	c.SysBaseController.Prepare()
}

func (c *OrderController) ListOrderFlow() {
	c.TplName = "order/ListOrderFlow.html"
}

func (c *OrderController)FormAddOrderFlow() {
	c.TplName = "order/formAddOrderFlow.html"
}

func (c *OrderController)UploadOrderFlow() {
	f, h, _ := c.GetFile("filename")
	path := "./upload/file/" + h.Filename
	f.Close()
	c.SaveToFile("filename", path)
	models.AddOrderFlow(path)
	c.DisplayJson(1,"ok","")
}

func (c *OrderController)OrderFlowTongJiTpl() {
	c.TplName = "order/OrderFlowTongJiTpl.html"
}

func (c *OrderController)OrderFlowTongJi() {
	s := make([]string,5)
	s[0] = "111"
	s[1] = "222"
	s[2] = "333"
	s[3] = "444"
	s[4] = "555"
	m := make([]float64,5)

	m[0] = 10.50
	m[1] = 12.00
	m[2] = 15.00
	m[3] = 20.00
	m[4] = 5.00

	type tmp struct {
		Jiner []float64
		Name []string
	}

	c.Data["json"] = &tmp{m,s}
	c.ServeJSON()
}


func (c *OrderController)AddOrderFlow() {

}

func (c *OrderController) DeleteOrderFlow() {
	ids := c.Input().Get("menu_ids")
	_, err := models.DeleteSysMenu(ids)
	if err != nil {
		c.DisplayJson(0, "修改失败", err.Error())
	}
	c.DisplayJson(1, "删除成功", c.URLFor("SysMenuController.ListSysMenu"))
}
