package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type (
	todoSQLmodel struct {
		Describe string    `json:"todothing" gorm:"column:todo_describe"`
		Deadline time.Time `json:"deadline" gorm:"column:todo_deadline"`
		Status   bool      `json:"status" gorm:"column:todo_status"`
		ID       uint      `json:"id" gorm:"column:todo_id"`
	}
	Status struct {
		Status int `json:"status"`
	}
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*.html")
	router.GET("/todolist", gettingtodolist)
	router.GET("/todolist/lists", getlists)
	router.POST("/todolist", newtodo)
	router.PUT("/todolist/lists/:id", changeStatus)
	router.Run(":8888")
}
func gettingtodolist(c *gin.Context) {
	c.HTML(200, "todolist.html", nil)
}
func getlists(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:Fuck06050@/todolist?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		panic("failed to connect database")
	}
	var todos []todoSQLmodel
	db.Table("todo").Find(&todos)
	c.JSON(http.StatusOK, todos)
}
func newtodo(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:Fuck06050@/todolist?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		panic("failed to connect database")
	}
	var newtodo todoSQLmodel
	c.BindJSON(&newtodo)
	db.Create(&newtodo)
}
func changeStatus(c *gin.Context) {
	id := c.Param("id")
	db, err := gorm.Open("mysql", "root:Fuck06050@/todolist?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		panic("failed to connect database")
	}
	var status Status
	c.BindJSON(&status)
	fmt.Println(status.Status)
	db.Table("todo").Where("todo_id=?", id).Update("todo_status", status.Status)
}
func (todo *todoSQLmodel) TableName() string {
	return "todo"
}
