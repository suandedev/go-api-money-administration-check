package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	e.GET("/", test)
	e.Logger.Fatal(e.Start(":1323"))
}

type User struct {
	gorm.Model
	Name           string
	Class_name     string
	Jurusan        string
	Administration []Administration
}

type Administration struct {
	gorm.Model
	Semester string
	Total    uint32
	Status   bool
	UserID   uint
}

type Result struct {
	Name      string
	Semester  string
	Total     uint32
	Status    bool
	ClassName string
	Jurusan   string
}

func connect() *gorm.DB {
	// Connect to the database
	db, err := gorm.Open(sqlite.Open("administrationDb.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{}, &Administration{})
	return db
}

func test(c echo.Context) error {
	// create user and administration
	// user := createUserAndAdmin()

	// delete user and administration by id user
	// user := deleteUserAndAdminByIdUser()

	// select user and administration by status atau yang belum bayar
	// results := selectUserAndAdminByStatus(false)

	// select user and administration by status true atau yang sudah bayar
	// results := selectUserAndAdminByStatus(true)

	// read all user and administration
	// results := selectAllUserAndAdmin()

	// update administration
	admin := updateAdminById()

	return c.JSON(http.StatusOK, admin)
}

func createUserAndAdmin() *User {
	db := connect()
	user := User{
		Name:       "andi4",
		Class_name: "XII RPL 1",
		Jurusan:    "RPL",
		Administration: []Administration{
			{
				Semester: "1",
				Total:    0,
				Status:   false,
			},
			{
				Semester: "2",
				Total:    0,
				Status:   false,
			},
		},
	}

	db.Create(&user)
	return &user
}

func deleteUserAndAdminByIdUser() string {
	db := connect()
	db.Delete(&User{}, 1)
	db.Where("user_id = ?", 1).Delete(&Administration{})
	return "deleted"
}

func selectUserAndAdminByStatus(bol bool) []Result {
	var results []Result
	db := connect()
	db.Table("users").Select("users.name, administrations.semester, administrations.status, administrations.total").Joins("left join administrations on administrations.user_id = users.id").Where("administrations.status = ?", bol).Scan(&results)
	return results
}

func selectAllUserAndAdmin() []Result {
	var results []Result
	db := connect()
	db.Table("users").Select("users.name, administrations.semester, administrations.status, administrations.total, users.class_name, users.jurusan").Joins("left join administrations on administrations.user_id = users.id").Scan(&results)
	return results
}

func updateAdminById() *Administration {
	db := connect()
	var administration Administration
	db.Model(&administration).Where("id = @id AND user_id = @user_id", sql.Named("id", 1), sql.Named("user_id", 1)).Update("total", 10000011)
	return &administration
}
