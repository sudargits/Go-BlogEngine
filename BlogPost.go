package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"github.com/kataras/iris"
)



const BLOGPOST  = "blog_post"

type BlogModel struct {
	id int `json:"id"`
	title string `json:"title"`
	article string `json:"article"`
}
func connectDB() *sql.DB{
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/blog")
	checkErr(err)
	return db
}
func checkErr(err error) {
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
}
func allPost(ctx *iris.Context){
	var db = connectDB()
	defer db.Close()

	rows, err := db.Query("SELECT id,article,title FROM blog_post")
	checkErr(err)
	defer rows.Close()

	var result []BlogModel
	for rows.Next() {
		var model = BlogModel{}
		var err = rows.Scan(&model.id,&model.article,&model.title)
		checkErr(err)
		result = append(result,model)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	ctx.JSON(iris.StatusOK,BlogModel{id:1,article:"aaa",title:"bbbb"})
}