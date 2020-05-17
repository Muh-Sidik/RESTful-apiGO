package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

var db *gorm.DB
var err error

type UserTable struct {
	gorm.Model
	Name  string
	Email string
}

func InitialMigration() { //migration table function
	db, err = gorm.Open("mysql", "root:@/testUser?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed connect to Database")
	}

	defer db.Close()

	db.AutoMigrate(&UserTable{})

}

func AllUsers(c echo.Context) error {
	db, err := gorm.Open("mysql", "root:@/testUser?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Cannot connect to database")
	}
	defer db.Close()

	var users []UserTable
	db.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func NewUser(c echo.Context) error {
	db, err := gorm.Open("mysql", "root:@/testUser?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Cannot connect to database")
	}
	defer db.Close()

	name := c.Param("name")
	email := c.Param("email")

	db.Create(&UserTable{Name: name, Email: email})

	return c.String(http.StatusCreated, "New user successfully created")
}

func ShowUser(c echo.Context) error {
	db, err := gorm.Open("mysql", "root:@/testUser?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Cannot connect to database")
	}
	defer db.Close()

	name := c.Param("name")

	var user UserTable

	db.Where("name = ?", name).Find(&user)

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	db, err := gorm.Open("mysql", "root:@/testUser?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Cannot connect to database")
	}
	defer db.Close()

	name := c.Param("name")

	var user UserTable

	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	return c.String(http.StatusNoContent, "User Successfully Deleted")
}

func UpdateUser(c echo.Context) error {
	db, err := gorm.Open("mysql", "root:@/testUser?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Cannot connect to database")
	}
	defer db.Close()

	name := c.Param("name")
	email := c.Param("email")

	var user UserTable

	db.Where("name = ?", name).Find(&user)

	user.Email = email

	db.Save(&user)

	return c.String(http.StatusAccepted, "User Successfully Updated")
}
