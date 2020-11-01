package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "net/http"
	"strconv"
)

// func main() {
// 	router := gin.Default()
// 	router.Static("styles", "./styles")
// 	router.LoadHTMLGlob("templates/*.html")
// 	router.GET("/", func(ctx *gin.Context) {
// 		ctx.HTML(200, "index.html", gin.H{})
// 	})
// 	router.Run()
// }
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	db_init()
	r.GET("/", func(c *gin.Context) {
		people := get_all()
		c.HTML(200, "index.html", gin.H{
			"people": people,
		})
	})
	r.POST("/new", func(c *gin.Context) {
		name := c.PostForm("name")
		age, _ := strconv.Atoi(c.PostForm("age"))
		create(name, age)
		c.Redirect(302, "/")
	})
	r.Run()
}

type Person struct {
	gorm.Model
	Name string
	Age  int
}

func db_init() {
	db, err := sqlConnect()
	if err != nil {
		panic("failed to connect database\n")
	}

	db.AutoMigrate(&Person{})
}
func create(name string, age int) {
	db, err := sqlConnect()
	if err != nil {
		panic("failed to connect database\n")
	}
	db.Create(&Person{Name: name, Age: age})
}
func get_all() []Person {
	db, err := sqlConnect()
	if err != nil {
		panic("failed to connect database\n")
	}
	var people []Person
	db.Find(&people)
	return people

}
func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "cake"
	PASS := "sw1908su"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "cheesecake"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}
