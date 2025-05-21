package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type User struct {
    ID    uint   `json:"id" gorm:"primaryKey"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var DB *gorm.DB

func ConnectDB() {
    dsn := "root:123456@tcp(db:3306)/testDB?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }

    db.AutoMigrate(&User{})

    DB = db
}

func SeedUsers() {
    var count int64
    DB.Model(&User{}).Count(&count)
    if count > 0 {
        fmt.Println("Users already seeded.")
        return
    }

    users := []User{
        {Name: "Nguyen Van A", Email: "nguyenvanA@gmail.com"},
        {Name: "Nguyen Van B", Email: "nguyenvanB@gmail.com"},
        {Name: "Nguyen Van C", Email: "nguyenvanC@gmail.com"},
    }

    for _, user := range users {
        DB.Create(&user)
    }
}

// List users API
func GetUsers(c *gin.Context) {
    var users []User
    result := DB.Find(&users)

    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    fmt.Println("Danh sách người dùng:")
    for _, user := range users {
        fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
    }

    c.JSON(http.StatusOK, users)
}

func main() {
    ConnectDB()
    SeedUsers()

    r := gin.Default()
	
	r.SetTrustedProxies([]string{"127.0.0.1"})

    r.GET("/users", GetUsers)

    r.Run(":8080")
}
