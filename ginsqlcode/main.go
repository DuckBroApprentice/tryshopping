package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//1、鏈接數據庫
	connStr := "root:123456@tcp(127.0.0.1:3306)/ginsql"

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	//db為連結上數據庫的實例

	/*
		//創建數據庫表格
		//person: id, name, age
		_, err = db.Exec("create table person(" +
			"id int auto_increment primary key," +
			"name varchar(12) not null," +
			"age int default 1" +
			");")

		if err != nil {
			log.Fatal(err.Error())
			return
		} else {
			fmt.Println("數據庫表創建成功")
		}
		執行一次後再執行就會報錯!!!
	*/

	//插入數據庫到數據庫表
	_, err = db.Exec("insert into person(name,age)"+
		"values(?,?);", "Davie", 18)

	if err != nil {
		log.Fatal(err.Error())
		return
	} else {
		fmt.Println("數據插入成功")
	}

	//查詢數據庫
	rows, err := db.Query("select id,name,age from person")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

scan:
	if rows.Next() {
		person := new(Person)
		err = rows.Scan(&person.Id, &person.Name, &person.Age)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(person.Id, person.Name, person.Age)
		goto scan
	}
	//假設rows有多筆數據
	//.Next()
	//->
	//Davie 18
	//Tom 15
	//Jack 20
	//Lily 25
	//->
}

// 與sql表單一致
type Person struct {
	Id   int
	Name string
	Age  int
}
