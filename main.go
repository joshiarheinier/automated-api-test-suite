package main

import (
	"fmt"
	"strconv"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	// "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//MySQL struct
type MySQL struct {
	Hostname     string
	Username     string
	Password     string
	MaxOpenConns string
	MaxIdleConns string
	Schema       string
}

type new_table struct {
	Id			int
	Name		string
}

var (
	engine                  = &xorm.Engine{}
	ErrFailedToConnectToSQL = "Failed to connect to mysql %v\n"
)

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		"root", "asdasdasd", "localhost", "new_schema")
	var err error
	engine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		err = fmt.Errorf(ErrFailedToConnectToSQL, err)
		panic(err.Error())
	}
	// maxCon, err := strconv.Atoi(conf.Database.MaxOpenConns)
	// if err != nil {
	// 	panic(err)
	// }
	// maxIdleCon, err := strconv.Atoi(conf.Database.MaxIdleConns)
	// engine.SetMaxOpenConns(maxCon)
	// engine.SetMaxIdleConns(maxIdleCon)
	engine.SetConnMaxLifetime(-1)

	//this will be sync the struct into the table and the table will be same with struct
	engine.Sync()
}



func setupRouter(r *gin.Engine) {
	// Disable Console Color
	// gin.DisableConsoleColor()
	var data new_table

	// Ping test
	r.POST("/insert", func(c *gin.Context) {
		data.Id, _ = strconv.Atoi(c.PostForm("id"))
		data.Name = c.PostForm("name")
		_, err := engine.Insert(&data)
		if err != nil {
			panic("Fail to insert")
		}
		c.String(http.StatusOK, "success")
	})

	r.GET("/get/", func(c *gin.Context) {
		//var sliceOfStructs []new_table
		rows, err := engine.Rows(&new_table{})
		if err != nil {

		}
		// SELECT * FROM user
		defer rows.Close()
		bean := new(new_table)
		list := []gin.H{}
		for rows.Next() {
			err = rows.Scan(bean)
			json := gin.H{
				"id": bean.Id,
				"name": bean.Name,
			}
			list = append(list, json)
		}
		c.JSON(http.StatusOK, gin.H{"data": list})
	})

	r.GET("/update/:id/:name", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Params.ByName("id"))
		data.Id = id
		data.Name = c.Params.ByName("name")
		_, err := engine.Update(&data, &new_table{Id:id})
		if err != nil {
			panic("Fail to update")
		}
		c.String(http.StatusOK, "success")
	})

	r.POST("/delete", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.PostForm("id"))
		// data.Id = id
		_, err := engine.Where("id = ?", id).Delete(&data)
		if err != nil {
			panic(err.Error())
		}
		c.String(http.StatusOK, "success")
	})

	// Get user value
	// r.GET("/user/:name", func(c *gin.Context) {
	// 	user := c.Params.ByName("name")
	// 	value, ok := db[user]
	// 	if ok {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	// 	} else {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	// 	}
	// })

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	// authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	// 	"foo":  "bar", // user:foo password:bar
	// 	"manu": "123", // user:manu password:123
	// }))

	// authorized.POST("admin", func(c *gin.Context) {
	// 	user := c.MustGet(gin.AuthUserKey).(string)

	// 	// Parse JSON
	// 	var json struct {
	// 		Value string `json:"value" binding:"required"`
	// 	}

	// 	if c.Bind(&json) == nil {
	// 		db[user] = json.Value
	// 		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	// 	}
	// })

}

func main() {
	r := gin.Default()
	setupRouter(r)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
