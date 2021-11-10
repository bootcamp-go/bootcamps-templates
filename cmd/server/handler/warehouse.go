package handler

import (
	"github.com/gin-gonic/gin"
)

type Warehouse struct {
	// warehouseService warehouse.Service
}

func NewWarehouse() *Warehouse {
	return &Warehouse{
		// warehouseService: w,
	}
}

func (w *Warehouse) Get() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (w *Warehouse) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (w *Warehouse) Create() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (w *Warehouse) Update() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (w *Warehouse) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
