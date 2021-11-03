package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/usuario/repositorio/internal/domain"
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
	type response struct {
		Data domain.Warehouse `json:"data"`
	}

	return func(c *gin.Context) {

	}
}

func (w *Warehouse) GetAll() gin.HandlerFunc {
	type response struct {
		Data []domain.Warehouse `json:"data"`
	}

	return func(c *gin.Context) {

	}
}

func (w *Warehouse) Store() gin.HandlerFunc {
	type request struct {
		Address            string `json:"address"`
		Telephone          string `json:"telephone"`
		WarehouseCode      string `json:"warehouse_code"`
		MinimunCapacity    int    `json:"minimun_capacity"`
		MinimunTemperature int    `json:"minimun_temperature"`
		SectionNumber      int    `json:"section_number"`
	}

	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {

	}
}

func (w *Warehouse) Update() gin.HandlerFunc {
	type request struct {
		Address            string `json:"address"`
		Telephone          string `json:"telephone"`
		WarehouseCode      string `json:"warehouse_code"`
		MinimunCapacity    int    `json:"minimun_capacity"`
		MinimunTemperature int    `json:"minimun_temperature"`
		SectionNumber      int    `json:"section_number"`
	}

	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {

	}
}

func (w *Warehouse) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
