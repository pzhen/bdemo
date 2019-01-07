package models

import (
	"github.com/astaxie/beego/orm"
	"os"
	"fmt"
	"encoding/csv"
	"io"
	"strings"
)

type OrderFlow struct {
	Id     int    `orm:"column(id);auto"`
	Field1 string `orm:"column(field1);" description:"交易时间"`
	Field2 string `orm:"column(field2);" description:"公众账号ID"`
	Field3 string `orm:"column(field3);" description:"商户号"`
	Field4 string `orm:"column(field4);" description:"特约商户号"`
	Field5 string `orm:"column(field5);" description:"设备号"`

	Field6 string `orm:"column(field6);" description:"微信订单号"`
	Field7 string `orm:"column(field7);" description:"商户订单号"`
	Field8 string `orm:"column(field8);" description:"用户标识"`
	Field9 string `orm:"column(field9);" description:"交易类型"`
	Field10 string `orm:"column(field10);" description:"交易状态"`

	Field11 string `orm:"column(field11);" description:"付款银行"`
	Field12 string `orm:"column(field12);" description:"货币种类"`
	Field13 string `orm:"column(field13);" description:"应结订单金额"`
	Field14 string `orm:"column(field14);" description:"代金券金额"`
	Field15 string `orm:"column(field15);" description:"微信退款单号"`

	Field16 string `orm:"column(field16);" description:"商户退款单号"`
	Field17 string `orm:"column(field17);" description:"退款金额"`
	Field18 string `orm:"column(field18);" description:"充值券退款金额"`
	Field19 string `orm:"column(field19);" description:"退款类型"`
	Field20 string `orm:"column(field20);" description:"退款状态"`
	Field21 string `orm:"column(field21);" description:"商品名称"`
	Field22 string `orm:"column(field22);" description:"商户数据包"`
	Field23 string `orm:"column(field23);" description:"手续费"`
	Field24 string `orm:"column(field24);" description:"费率"`
	Field25 string `orm:"column(field25);" description:"订单金额"`

	Field26 string `orm:"column(field26);" description:"申请退款金额"`
	Field27 string `orm:"column(field27);" description:"费率备注"`
}

func (t *OrderFlow) TableName() string {
	return "order_flow"
}

func init() {
	orm.RegisterModel(new(OrderFlow))
}

func AddOrderFlow(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 这个方法体执行完成后，关闭文件
	defer file.Close()

	reader := csv.NewReader(file)

	flowRow := &OrderFlow{}
	o := orm.NewOrm()

	for {
		// Read返回的是一个数组，它已经帮我们分割了，
		record, err := reader.Read()
		// 如果读到文件的结尾，EOF的优先级居然比nil还高！
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("记录集错误:", err)
			return
		}

		flowRow.Id += 1
		flowRow.Field1 = strings.Replace(record[0], "`", "", -1 )
		flowRow.Field2 = strings.Replace(record[1], "`", "", -1 )
		flowRow.Field3 = strings.Replace(record[2], "`", "", -1 )
		flowRow.Field4 = strings.Replace(record[3], "`", "", -1 )
		flowRow.Field5 = strings.Replace(record[4], "`", "", -1 )

		flowRow.Field6 = strings.Replace(record[5], "`", "", -1 )
		flowRow.Field7 = strings.Replace(record[6], "`", "", -1 )
		flowRow.Field8 = strings.Replace(record[7], "`", "", -1 )
		flowRow.Field9 = strings.Replace(record[8], "`", "", -1 )
		flowRow.Field10 = strings.Replace(record[9], "`", "", -1 )

		flowRow.Field11 = strings.Replace(record[10], "`", "", -1 )
		flowRow.Field12 = strings.Replace(record[11], "`", "", -1 )
		flowRow.Field13 = strings.Replace(record[12], "`", "", -1 )
		flowRow.Field14 = strings.Replace(record[13], "`", "", -1 )
		flowRow.Field15 = strings.Replace(record[14], "`", "", -1 )

		flowRow.Field16 = strings.Replace(record[15], "`", "", -1 )
		flowRow.Field17 = strings.Replace(record[16], "`", "", -1 )
		flowRow.Field18 = strings.Replace(record[17], "`", "", -1 )
		flowRow.Field19 = strings.Replace(record[18], "`", "", -1 )
		flowRow.Field20 = strings.Replace(record[19], "`", "", -1 )

		flowRow.Field21 = strings.Replace(record[20], "`", "", -1 )
		flowRow.Field22 = strings.Replace(record[21], "`", "", -1 )
		flowRow.Field23 = strings.Replace(record[22], "`", "", -1 )
		flowRow.Field24 = strings.Replace(record[23], "`", "", -1 )
		flowRow.Field25 = strings.Replace(record[24], "`", "", -1 )
		flowRow.Field26 = strings.Replace(record[25], "`", "", -1 )
		flowRow.Field27 = strings.Replace(record[26], "`", "", -1 )


		//for i := 0; i < len(record); i++ {
		//	fmt.Print(record[i] + "******")
		//}
		//fmt.Print("\n")


		o.Insert(flowRow)
	}

}
