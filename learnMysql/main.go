package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Fuck06050@/test?charset=utf8")
	checkErr(err)
	stmt, err := db.Prepare("INSERT userinfo SET username=?, department=?, created=?")
	checkErr(err)
	res, err := stmt.Exec("wayne", "ntut", "2019.9.9")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
	stmt, err = db.Prepare("update userinfo set username=? where uid=? ")
	checkErr(err)
	res, err = stmt.Exec("wayneupdate", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
		stmt, err = db.Prepare("delete from userinfo where uid=?")
		checkErr(err)

		res, err = stmt.Exec(id)
		checkErr(err)

		affect, err = res.RowsAffected()
		checkErr(err)

		fmt.Println(affect)

		defer db.Close()
	}
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
