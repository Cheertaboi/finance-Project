package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/harsh/finance-project/backend/internal/db"
	"github.com/harsh/finance-project/backend/internal/expense"
)

func main() {
	database := db.InitDB("expense.db")

	repo := expense.NewRepository(database)
	handler := expense.NewHandler(repo)

	r := gin.Default()
	r.POST("/expenses", handler.CreateExpense)
	r.GET("/expenses", handler.ListExpenses)

	log.Println("Server running on :8080")
	r.Run(":8080")
}
