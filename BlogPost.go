package main

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/kataras/iris"
	"strings"
)

const BLOG_POST string  = "blog_post"

type BlogModel struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Article string `json:"article"`
}

func allPost(ctx *iris.Context){

	var db = connectDB()
	defer db.Close()

	rows, err := db.Query("SELECT id,article,title FROM "+BLOG_POST)
	checkErr(err)
	defer rows.Close()

	var result []BlogModel
	for rows.Next() {
		var model = BlogModel{}
		var err = rows.Scan(&model.Id,&model.Title,&model.Article)
		checkErr(err)
		result = append(result,model)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	ctx.JSON(iris.StatusOK, iris.Map{"status":true, "message":BaseMessage{Devel:"success", Prod:"success"},"count":len(result), "content":result})
}

func detailPost(ctx *iris.Context) {
	id,err := ctx.ParamInt("idblog")

	if err != nil {
		ctx.Redirect("/blog/all",iris.StatusOK)
	}else {
		var db = connectDB()
		defer db.Close()

		var result BlogModel

		var err = db.QueryRow("SELECT id,article,title FROM "+BLOG_POST+" where id = ?",id).Scan(&result.Id, &result.Title,&result.Article)
		if err != nil {
			fmt.Println(err.Error())
			if strings.Contains(err.Error(),"no rows") {
				ctx.JSON(iris.StatusNotFound, iris.Map{"status":true, "message":BaseMessage{Devel:err.Error(), Prod:"Data tidak ditemukan"}})
			}else{
				ctx.JSON(iris.StatusInternalServerError, iris.Map{"status":true, "message":BaseMessage{Devel:err.Error(), Prod:"Data tidak ditemukan"}})
			}
		}else{
			ctx.JSON(iris.StatusOK, iris.Map{"status":true, "message":BaseMessage{Devel:"success", Prod:"success"}, "content":result})
		}
	}
}