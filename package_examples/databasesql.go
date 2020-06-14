package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	//"reflect"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Site ...
type Site struct {
	ID    int64  `field:"id"`
	Key   string `field:"key"`
	Value string `field:"value"`
	State int64  `field:"state"`
}

func main() {
	dsn := "root:root@/goadmin?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	checkErr(err)

	err = db.Ping()
	checkErr(err)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(10)
	// 设置最大链接数
	db.SetMaxOpenConns(20)

	Query(db)

}

//Create
func Create(db sql.DB) {
	stmt, err := db.Prepare("insert into user set UserName=?,Password=?")
	checkErr(err)
	res, err := stmt.Exec("peter", "panlei")
	checkErr(err)
	lastInsertId, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(lastInsertId)
}

//Read
func Query(db *sql.DB) {
	rows, err := db.Query("select * from goadmin_site where 1")
	checkErr(err)

	cols, err := rows.Columns()
	fmt.Println(len(cols))
	vals := make([]interface{}, len(cols))

	for i, _ := range cols {
		vals[i] = new(sql.RawBytes)
	}
	//log.Println(vals)

	for rows.Next() {
		err := rows.Scan(vals...)
		checkErr(err)

		fmt.Println(vals...)

	}

}

//Update
func Update(db sql.DB) {

}

//Delete
func Delete(db sql.DB) {

}
