package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// dsn := "host=localhost user=postgres password=postgres dbname=payslip_db port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// }

	r := gin.Default()

	// Test health route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})
	})

	r.Run(":8080")
}
