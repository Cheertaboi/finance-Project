package main

import (
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/Cheertaboi/finance-Project/backend/internal/db"
	"github.com/Cheertaboi/finance-Project/backend/internal/expense"
)

func main() {
	database := db.InitDB("expense.db")

	repo := expense.NewRepository(database)
	handler := expense.NewHandler(repo)

	r := gin.Default()

	// -------- API ROUTES --------
	r.POST("/expenses", handler.CreateExpense)
	r.GET("/expenses", handler.ListExpenses)

	// -------- FRONTEND --------
	// Resolve static directory correctly
	staticDir, _ := filepath.Abs("../../static")

	// Serve JS/CSS
	r.Static("/static", staticDir)

	// Serve index.html
	r.GET("/", func(c *gin.Context) {
		c.File(staticDir + "/index.html")
	})

	// Fallback (refresh-safe)
	r.NoRoute(func(c *gin.Context) {
		c.File(staticDir + "/index.html")
	})

	r.Run(":8080")
}
