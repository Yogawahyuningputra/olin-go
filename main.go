package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connection Database
func DatabaseInit() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/olin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Database")
}
func GetTransaction() {
	query := `
		SELECT u.name, SUM(o.amount) AS amount
		FROM users u
		JOIN orders o ON u.id = o.user_id
		WHERE o.created_at >= '2022-01-01'
		GROUP BY u.name
		HAVING SUM(o.amount) >= 100;
	`
	// 	query2 := `
	// 		SELECT u.name, o.amount, o.created_at
	// 		FROM users u
	// 		JOIN orders o ON u.id = o.user_id;
	// `
	type Result struct {
		Name   string  `json:"name"`
		Amount float64 `json:"amount"`
	}
	var results []Result

	// Eksekusi query
	if err := DB.Raw(query).Scan(&results).Error; err != nil {
		log.Fatalf("error executing query: %v", err)
	}

	// Tampilkan hasil
	for _, result := range results {
		fmt.Println("user:", result.Name)
		fmt.Println("total amount", result.Amount)
	}

}
func main() {
	DatabaseInit()
	GetTransaction()
	fmt.Println("hello golang")
}
