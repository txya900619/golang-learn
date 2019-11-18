package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type (
	TodoSQLmodel struct {
		Describe string `json:"user_input" gorm:"column:todo_describe"`
		Status   bool   `json:"status" gorm:"column:todo_status"`
		ID       uint   `json:"item_id" gorm:"primary_key;column:todo_id"`
	}
	Status struct {
		Status bool `json:"status"`
	}
)

func main() {
	router := gin.Default()
	router.Delims("[[", "]]")
	router.LoadHTMLGlob("./templates/*.html")
	router.Static("/css", "./templates/css")
	router.Static("/js", "./templates/js")
	router.GET("/todolist", gettingtodolist)
	router.GET("/todolist/lists", getlists)
	router.POST("/todolist", newtodo)
	router.PUT("/todolist/:id", changeStatus)
	router.DELETE("/todolist/:id", deletetodo)
	router.Run(":8888")

}
func gettingtodolist(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
func getlists(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:Fuck06050@/todolist?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		panic("failed to connect database")
	}
	var todos []TodoSQLmodel
	db.Table("todo").Find(&todos)
	c.JSON(http.StatusOK, todos)
}
func newtodo(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:Fuck06050@/todolist?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		panic("failed to connect database")
	}
	var newtodo TodoSQLmodel
	c.BindJSON(&newtodo)
	if newtodo.Describe == "" {
		c.JSON(403, nil)
	} else {
		db.Create(&newtodo)
		fmt.Println(newtodo)
		c.JSON(http.StatusOK, newtodo)
	}
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
	c.JSON(http.StatusOK, nil)
}
func deletetodo(c *gin.Context) {

	db, err := gorm.Open("mysql", "root:Fuck06050@/todolist?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		panic("failed to connect database")
	}
	db.Table("todo").Where("todo_id=?", c.Param("id")).Delete(TodoSQLmodel{})
}
func (todo *TodoSQLmodel) TableName() string {
	return "todo"
}
