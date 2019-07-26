package controller

import (
	"fmt"
	"net/http"

	"github.com/codepository/go-test/database"
	"github.com/codepository/go-test/util"
)

func InsertToDB(writer http.ResponseWriter, request *http.Request) {
	opend, db := database.OpenDB()
	if opend {
		fmt.Println("open success")
	} else {
		fmt.Println("open faile:")
	}
	uid := util.GetNowtimeMD5()
	nowTimeStr := util.GetTime()
	stmt, err := db.Prepare("insert userinfo set username=?,departname=?,created=?,password=?,uid=?")
	util.CheckErr(err)
	res, err := stmt.Exec("wangbiao", "研发中心", nowTimeStr, "123456", uid)
	util.CheckErr(err)
	id, err := res.LastInsertId()
	util.CheckErr(err)
	if err != nil {
		fmt.Fprintf(writer, "插入数据失败")
	} else {
		fmt.Fprintf(writer, "插入数据成功：%d", id)
	}
}
