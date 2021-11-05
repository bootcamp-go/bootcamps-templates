package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// db, _ := sql.Open("mysql", "root:root@/meli")
	router := gin.Default()

	// Codigo de ayuda
	// warehouseRepository := warehouse.NewRepository(db)
	// warehouseService := warehouse.NewService(warehouseRepository)
	// warehouseHandler := handler.NewWarehouse(warehouseService)
	// warehousesRoutes := router.Group("/api/v1/warehouses")
	// {
	// 	warehousesRoutes.GET("/", warehouseHandler.GetAll())
	// 	warehousesRoutes.GET("/:id", warehouseHandler.Get())
	// 	warehousesRoutes.POST("/", warehouseHandler.Store())
	// 	warehousesRoutes.PATCH("/:id", warehouseHandler.Update())
	// 	warehousesRoutes.DELETE("/:id", warehouseHandler.Delete())
	// }

	if err := router.Run(); err != nil {
		panic(err)
	}
}
